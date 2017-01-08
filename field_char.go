package mysql

import (
	"reflect"
	"strconv"

	"github.com/go-qbit/model"
)

type baseCharField struct {
	Id        string
	Caption   string
	Length    int
	Charset   string
	Collate   string
	NotNull   bool
	Default   string
	CheckFunc func(interface{}) error
}

func (f *baseCharField) GetId() string         { return f.Id }
func (f *baseCharField) GetCaption() string    { return f.Caption }
func (f *baseCharField) GetType() reflect.Type { return reflect.TypeOf(string("")) }
func (f *baseCharField) IsRequired() bool      { return f.NotNull && f.Default == "" }
func (f *baseCharField) Check(v interface{}) error {
	if f.CheckFunc != nil {
		return f.CheckFunc(v)
	} else {
		return nil
	}
}
func (f *baseCharField) WriteSQL(ftype string, sqlBuf *SqlBuffer) {
	sqlBuf.WriteString(QuoteIdentifier(f.Id))

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString(ftype)

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

type VarCharField baseCharField

func (f *VarCharField) GetId() string                                    { return (*baseCharField)(f).GetId() }
func (f *VarCharField) GetCaption() string                               { return (*baseCharField)(f).GetCaption() }
func (f *VarCharField) GetType() reflect.Type                            { return (*baseCharField)(f).GetType() }
func (f *VarCharField) IsDerivable() bool                                { return false }
func (f *VarCharField) IsRequired() bool                                 { return (*baseCharField)(f).IsRequired() }
func (f *VarCharField) GetDependsOn() []string                           { return nil }
func (f *VarCharField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *VarCharField) Check(v interface{}) error                        { return (*baseCharField)(f).Check(v) }
func (f *VarCharField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &VarCharField{id, caption, f.Length, f.Charset, f.Collate, required, f.Default, f.CheckFunc}
}
func (f *VarCharField) WriteSQL(sqlBuf *SqlBuffer) { (*baseCharField)(f).WriteSQL("VARCHAR", sqlBuf) }

type CharField baseCharField

func (f *CharField) GetId() string                                    { return (*baseCharField)(f).GetId() }
func (f *CharField) GetCaption() string                               { return (*baseCharField)(f).GetCaption() }
func (f *CharField) GetType() reflect.Type                            { return (*baseCharField)(f).GetType() }
func (f *CharField) IsDerivable() bool                                { return false }
func (f *CharField) IsRequired() bool                                 { return (*baseCharField)(f).IsRequired() }
func (f *CharField) GetDependsOn() []string                           { return nil }
func (f *CharField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *CharField) Check(v interface{}) error                        { return (*baseCharField)(f).Check(v) }
func (f *CharField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &CharField{id, caption, f.Length, f.Charset, f.Collate, required, f.Default, f.CheckFunc}
}
func (f *CharField) WriteSQL(sqlBuf *SqlBuffer) { (*baseCharField)(f).WriteSQL("CHAR", sqlBuf) }
