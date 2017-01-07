package mysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/go-qbit/model"
	"github.com/go-qbit/qerror"

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
		if _, err := s.Exec(ctx, sqlBuf.String(), sqlBuf.GetArgs()...); err != nil {
			return err
		}
	}

	return nil
}

func (s *MySQL) Exec(ctx context.Context, sql string, a ...interface{}) (driver.Result, error) {
	res, err := s.db.Exec(sql, a...)
	if err != nil {
		if warning, ok := err.(drvMysql.MySQLWarnings); ok {
			println(warning.Error())
			err = nil
		}
	}

	return res, err
}

func (s *MySQL) Add(ctx context.Context, m model.IModel, data []map[string]interface{}) ([]interface{}, error) {
	uniqKeys := make(map[string]struct{})
	for _, row := range data {
		for k, _ := range row {
			uniqKeys[k] = struct{}{}
		}
	}
	insertKeys := make([]string, 0, len(uniqKeys))
	for k, _ := range uniqKeys {
		insertKeys = append(insertKeys, k)
	}
	sort.Strings(insertKeys)

	sqlBuf := NewSqlBuffer()

	sqlBuf.WriteString("INSERT INTO ")
	sqlBuf.WriteIdentifier(m.GetId())

	sqlBuf.WriteByte('(')
	sqlBuf.WriteIdentifiersList(insertKeys)
	sqlBuf.WriteByte(')')

	sqlBuf.WriteString("VALUES")

	for i, row := range data {
		if i != 0 {
			sqlBuf.WriteByte(',')
		}
		sqlBuf.WriteByte('(')
		for j, name := range insertKeys {
			if j != 0 {
				sqlBuf.WriteByte(',')
			}
			sqlBuf.WriteValue(row[name])
		}

		sqlBuf.WriteByte(')')
	}

	_, err := s.db.Exec(sqlBuf.String(), sqlBuf.GetArgs()...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *MySQL) Query(ctx context.Context, m model.IModel, fieldsNames []string, options model.QueryOptions) ([]map[string]interface{}, error) {
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

	rows, err := s.db.Query(sqlBuf.String(), sqlBuf.GetArgs()...)
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

func Quote(value string) string {
	// FixMe: Replace termporary solution
	return "'" + strings.Replace(value, "'", "''", -1) + "'"
}
