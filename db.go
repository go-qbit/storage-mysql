package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/go-qbit/model"
	"github.com/go-qbit/qerror"
	"github.com/go-qbit/timelog"

	drvMysql "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db        *sql.DB
	models    map[string]model.IModel
	modelsMtx sync.RWMutex
}

func NewMySQL() *MySQL {
	s := &MySQL{
		models: make(map[string]model.IModel),
	}

	return s
}

func (s *MySQL) Connect(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *MySQL) Disonnect() error {
	return s.db.Close()
}

func (s *MySQL) NewModel(id string, fields []model.IFieldDefinition, pk []string) model.IModel {
	mysqlFields := make([]IMysqlFieldDefinition, len(fields))
	for i, _ := range fields {
		mysqlFields[i] = fields[i].(IMysqlFieldDefinition)
	}

	return NewBaseModel(s, id, mysqlFields, nil, pk, nil)
}

func (s *MySQL) RegisterModel(m model.IModel) error {
	s.modelsMtx.Lock()
	defer s.modelsMtx.Unlock()

	if _, exists := s.models[m.GetId()]; exists {
		return qerror.New("Model '%s' is already exists", m.GetId())
	}

	s.models[m.GetId()] = m

	return nil
}

func (s *MySQL) WriteCreateSQL(sqlBuf *SqlBuffer) {
	modelLevels := s.getModelsLevels()

	sort.Sort(modelLevels)

	s.modelsMtx.RLock()
	defer s.modelsMtx.RUnlock()

	for _, modelLevel := range modelLevels {
		s.models[modelLevel.name].(*BaseModel).WriteCreateSQL(sqlBuf)
		sqlBuf.WriteString(";\n")
	}
}

func (s *MySQL) InitDB(ctx context.Context) error {
	modelLevels := s.getModelsLevels()
	sort.Sort(modelLevels)

	sqlBuf := NewSqlBuffer()

	s.modelsMtx.RLock()
	defer s.modelsMtx.RUnlock()

	for _, modelLevel := range modelLevels {
		sqlBuf.Reset()
		s.models[modelLevel.name].(*BaseModel).WriteCreateSQL(sqlBuf)
		if _, err := s.Exec(ctx, sqlBuf.GetSQL(), sqlBuf.GetArgs()...); err != nil {
			return err
		}
	}

	return nil
}

func (s *MySQL) Exec(ctx context.Context, sql string, a ...interface{}) (driver.Result, error) {
	ctx = timelog.Start(ctx, &SqlBuffer{Buffer: bytes.NewBufferString(sql), args: a})
	defer timelog.Finish(ctx)

	ct := ctx.Value(s.transactionKey())

	var (
		res driver.Result
		err error
	)

	if ct == nil {
		res, err = s.db.Exec(sql, a...)
	} else {
		res, err = ct.(*transaction).tx.Exec(sql, a...)
	}

	if err != nil {
		if warning, ok := err.(drvMysql.MySQLWarnings); ok {
			println(warning.Error())
			err = nil
		}
	}

	return res, err
}

func (s *MySQL) RawQuery(ctx context.Context, query string, a ...interface{}) (*sql.Rows, error) {
	ct := ctx.Value(s.transactionKey())

	var (
		res *sql.Rows
		err error
	)

	if ct == nil {
		res, err = s.db.Query(query, a...)
	} else {
		res, err = ct.(*transaction).tx.Query(query, a...)
	}

	if err != nil {
		if warning, ok := err.(drvMysql.MySQLWarnings); ok {
			println(warning.Error())
			err = nil
		}
	}

	return res, err
}

func (s *MySQL) Add(ctx context.Context, m model.IModel, fieldsNames []string, data [][]interface{}, opts model.AddOptions) ([]interface{}, error) {
	sqlBuf := NewSqlBuffer()

	if opts.Replace {
		sqlBuf.WriteString("REPLACE INTO ")
	} else {
		sqlBuf.WriteString("INSERT INTO ")
	}
	sqlBuf.WriteIdentifier(m.GetId())

	sqlBuf.WriteByte('(')
	sqlBuf.WriteIdentifiersList(fieldsNames)
	sqlBuf.WriteByte(')')

	sqlBuf.WriteString("VALUES")

	for i, row := range data {
		if i != 0 {
			sqlBuf.WriteByte(',')
		}
		sqlBuf.WriteByte('(')
		sqlBuf.WriteValuesList(row)
		sqlBuf.WriteByte(')')
	}

	ctx = timelog.Start(ctx, sqlBuf)
	execRes, err := s.db.Exec(sqlBuf.GetSQL(), sqlBuf.GetArgs()...)
	ctx = timelog.Finish(ctx)
	if err != nil {
		return nil, err
	}

	pKFieldsNames := m.GetPKFieldsNames()
	fieldsPos := make(map[string]int)
	for i, fieldName := range fieldsNames {
		fieldsPos[fieldName] = i
	}

	lastInsertId := int64(0)
	for _, fieldName := range pKFieldsNames {
		if _, exists := fieldsPos[fieldName]; !exists {
			if m.GetFieldDefinition(fieldName).(IMysqlFieldDefinition).IsAutoIncremented() {
				lastInsertId, _ = execRes.LastInsertId()
				break
			}
		}
	}

	res := make([]interface{}, len(data))
	for i, row := range data {
		rowRes := make([]interface{}, len(pKFieldsNames))
		for j, fieldName := range pKFieldsNames {
			fp, exists := fieldsPos[fieldName]
			if exists {
				rowRes[j] = row[fp]
			}
			if (!exists || rowRes[j] == nil) && m.GetFieldDefinition(fieldName).(IMysqlFieldDefinition).IsAutoIncremented() {
				switch m.GetFieldDefinition(fieldName).GetType().Kind() {
				case reflect.Int8:
					rowRes[j] = int8(lastInsertId)
				case reflect.Int16:
					rowRes[j] = int16(lastInsertId)
				case reflect.Int32:
					rowRes[j] = int32(lastInsertId)
				case reflect.Int64:
					rowRes[j] = int64(lastInsertId)
				case reflect.Uint8:
					rowRes[j] = uint8(lastInsertId)
				case reflect.Uint16:
					rowRes[j] = uint16(lastInsertId)
				case reflect.Uint32:
					rowRes[j] = uint32(lastInsertId)
				case reflect.Uint64:
					rowRes[j] = uint64(lastInsertId)
				default:
					panic("Not implemented")
				}
				lastInsertId++
			}
		}
		res[i] = rowRes
	}

	return res, nil
}

func (s *MySQL) Query(ctx context.Context, m model.IModel, fieldsNames []string, options model.GetAllOptions) ([]map[string]interface{}, error) {
	sqlBuf := NewSqlBuffer()

	sqlBuf.WriteString("SELECT ")
	sqlBuf.WriteIdentifiersList(fieldsNames)
	sqlBuf.WriteString(" FROM ")
	sqlBuf.WriteIdentifier(m.GetId())

	if options.Filter != nil {
		sqlBuf.WriteString(" WHERE ")
		options.Filter.GetProcessor(exprProcessor).(WriteFunc)(sqlBuf)
	}

	if len(options.OrderBy) > 0 {
		sqlBuf.WriteString(" ORDER BY ")
		for i, order := range options.OrderBy {
			if i > 0 {
				sqlBuf.WriteString(",")
			}
			sqlBuf.WriteIdentifier(order.FieldName)
			if order.Desc {
				sqlBuf.WriteString(" DESC")
			}
		}
	}

	if options.Limit > 0 {
		sqlBuf.WriteString(" LIMIT ")
		sqlBuf.WriteString(strconv.FormatUint(options.Limit, 10))
		if options.Offset > 0 {
			sqlBuf.WriteString(" OFFSET ")
			sqlBuf.WriteString(strconv.FormatUint(options.Offset, 10))
		}
	}

	//fmt.Println(sqlBuf.String())

	ctx = timelog.Start(ctx, sqlBuf)
	rows, err := s.db.Query(sqlBuf.GetSQL(), sqlBuf.GetArgs()...)
	ctx = timelog.Finish(ctx)
	if err != nil {
		return nil, err
	}

	var res []map[string]interface{}
	columnsNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rawRow := make([]interface{}, len(columnsNames))
		for i, name := range columnsNames {
			field := m.GetFieldDefinition(name)

			rawRow[i] = reflect.New(field.GetType()).Interface()
		}

		err := rows.Scan(rawRow...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, name := range columnsNames {
			row[name] = reflect.ValueOf(rawRow[i]).Elem().Interface()
		}
		res = append(res, row)
	}

	return res, nil
}

func QuoteIdentifier(identifier string) string {
	return "`" + strings.Replace(identifier, "`", "``", -1) + "`"
}

func Quote(value interface{}) string {
	// FixMe: Replace termporary solution
	var v string

	switch value := value.(type) {
	case string:
		v = strings.Replace(value, "'", "''", -1)
	case int:
		v = strconv.FormatInt(int64(value), 10)
	case int8:
		v = strconv.FormatInt(int64(value), 10)
	case int16:
		v = strconv.FormatInt(int64(value), 10)
	case int32:
		v = strconv.FormatInt(int64(value), 10)
	case int64:
		v = strconv.FormatInt(value, 10)
	case uint:
		v = strconv.FormatUint(uint64(value), 10)
	case uint8:
		v = strconv.FormatUint(uint64(value), 10)
	case uint16:
		v = strconv.FormatUint(uint64(value), 10)
	case uint32:
		v = strconv.FormatUint(uint64(value), 10)
	case uint64:
		v = strconv.FormatUint(value, 10)
	case float32:
		v = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		v = strconv.FormatFloat(value, 'f', -1, 64)
	default:
		panic(fmt.Sprintf("%T is not implemented", value))
	}

	return "'" + v + "'"
}
