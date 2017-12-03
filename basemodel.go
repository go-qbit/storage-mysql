package mysql

import (
	"fmt"
	"strings"

	"github.com/go-qbit/model"
)

type BaseModel struct {
	*model.BaseModel
	indexes []Index
}

type BaseModelOpts struct {
	model.BaseModelOpts
	Indexes []Index
}

type Index struct {
	FieldNames []string
	Unique     bool
}

func NewBaseModel(db *MySQL, id string, dbFields []IMysqlFieldDefinition, derivableFields []model.IFieldDefinition, opts BaseModelOpts) *BaseModel {
	allFields := make([]model.IFieldDefinition, 0, len(dbFields)+len(derivableFields))

	for _, field := range dbFields {
		allFields = append(allFields, field)
	}

	for _, field := range derivableFields {
		if !field.IsDerivable() {
			panic(fmt.Sprintf("The field '%s' for the model '%s' is not derivable", field.GetId(), id))
		}
		allFields = append(allFields, field)
	}

	m := &BaseModel{
		BaseModel: model.NewBaseModel(id, allFields, db, opts.BaseModelOpts),
		indexes:   opts.Indexes,
	}

	db.modelsMtx.Lock()
	defer db.modelsMtx.Unlock()
	db.models[id] = m

	return m
}

func (m *BaseModel) AddField(field model.IFieldDefinition) {
	if !field.IsDerivable() {
		if _, ok := field.(IMysqlFieldDefinition); !ok {
			panic(fmt.Sprintf("The field '%s' is not derivable and not implements IMysqlFieldDefinition", field.GetId()))
		}
	}

	m.BaseModel.AddField(field)
}

func (m *BaseModel) WriteCreateSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteString("CREATE TABLE ")
	sqlBuf.WriteIdentifier(m.GetId())
	sqlBuf.WriteString(" (")

	first := true
	for _, fieldName := range m.GetFieldsNames() {
		field := m.GetFieldDefinition(fieldName)
		if field.IsDerivable() {
			continue
		}

		if first {
			first = false
		} else {
			sqlBuf.WriteByte(',')
		}

		field.(IMysqlFieldDefinition).WriteSQL(sqlBuf)
	}

	if pk := m.GetPKFieldsNames(); len(pk) > 0 {
		sqlBuf.WriteString(",PRIMARY KEY (")
		sqlBuf.WriteIdentifiersList(pk)
		sqlBuf.WriteByte(')')
	}

	for _, index := range m.indexes {
		sqlBuf.WriteByte(',')
		if index.Unique {
			sqlBuf.WriteString("UNIQUE ")
		}
		sqlBuf.WriteString("INDEX ")

		indexNameArr := []string{}
		if index.Unique {
			indexNameArr = append(indexNameArr, "uniq")
		}
		indexNameArr = append(indexNameArr, m.GetId(), "")
		indexNameArr = append(indexNameArr, index.FieldNames...)
		indexName := strings.Join(indexNameArr, "_")
		if len(indexName) > 64 {
			indexName = indexName[0:64]
		}

		sqlBuf.WriteIdentifier(indexName)
		sqlBuf.WriteByte('(')
		sqlBuf.WriteIdentifiersList(index.FieldNames)
		sqlBuf.WriteByte(')')
	}

	for _, extModel := range m.GetRelations() {
		relation := m.GetRelation(extModel)
		if relation.IsBack {
			continue
		}

		if relation.JunctionModel == nil {
			sqlBuf.WriteString(",FOREIGN KEY ")
			fkNameArr := []string{"fk", m.GetId(), ""}
			fkNameArr = append(fkNameArr, relation.LocalFieldsNames...)
			fkNameArr = append(fkNameArr, "_", relation.ExtModel.GetId(), "")
			fkNameArr = append(fkNameArr, relation.FkFieldsNames...)
			fkName := strings.Join(fkNameArr, "_")
			if len(fkName) > 64 {
				fkName = fkName[0:64]
			}
			sqlBuf.WriteIdentifier(fkName)
			sqlBuf.WriteByte('(')
			sqlBuf.WriteIdentifiersList(relation.LocalFieldsNames)
			sqlBuf.WriteString(")REFERENCES ")
			sqlBuf.WriteIdentifier(relation.ExtModel.GetId())
			sqlBuf.WriteByte('(')
			sqlBuf.WriteIdentifiersList(relation.FkFieldsNames)
			sqlBuf.WriteString(")ON UPDATE RESTRICT ON DELETE RESTRICT")
		}
	}

	sqlBuf.WriteByte(')')
	sqlBuf.WriteString("ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8'")
}
