package mysql

import (
	"reflect"
	"strconv"

	"github.com/go-qbit/model"
)

type VarcharField struct {
	Id        string
	Caption   string
	Length    int
	Charset   string
	Collate   string
	NotNull   bool
	Default   string
	CheckFunc func(interface{}) error
}

func (f *VarcharField) GetId() string                                    { return f.Id }
func (f *VarcharField) GetCaption() string                               { return f.Caption }
func (f *VarcharField) GetType() reflect.Type                            { return reflect.TypeOf(string("")) }
func (f *VarcharField) IsDerivable() bool                                { return false }
func (f *VarcharField) IsRequired() bool                                 { return f.NotNull && f.Default == "" }
func (f *VarcharField) GetDependsOn() []string                           { return nil }
func (f *VarcharField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *VarcharField) Check(v interface{}) error {
	if f.CheckFunc != nil {
		return f.CheckFunc(v)
	} else {
		return nil
	}
}
func (f *VarcharField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &VarcharField{id, caption, f.Length, f.Charset, f.Collate, required, f.Default, f.CheckFunc}
}
func (f *VarcharField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteString(QuoteIdentifier(f.Id))

	sqlBuf.WriteString(" VARCHAR")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteString(Quote(f.Charset))
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteString(Quote(f.Collate))
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != "" {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(f.Default)
	}
}
