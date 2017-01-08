package mysql

import (
	"reflect"

	"github.com/go-qbit/model"
)

type baseSimpleField struct {
	Id        string
	Caption   string
	NotNull   bool
	Default   string
	CheckFunc func(interface{}) error
}

func (f *baseSimpleField) GetId() string      { return f.Id }
func (f *baseSimpleField) GetCaption() string { return f.Caption }
func (f *baseSimpleField) IsRequired() bool   { return f.NotNull && f.Default == "" }
func (f *baseSimpleField) Check(v interface{}) error {
	if f.CheckFunc != nil {
		return f.CheckFunc(v)
	} else {
		return nil
	}
}
func (f *baseSimpleField) WriteSQL(ftype string, sqlBuf *SqlBuffer) {
	sqlBuf.WriteString(QuoteIdentifier(f.Id))

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString(ftype)

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != "" {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(f.Default)
	}
}

//DATE
type DateField baseSimpleField

func (f *DateField) GetId() string                                    { return (*baseSimpleField)(f).GetId() }
func (f *DateField) GetCaption() string                               { return (*baseSimpleField)(f).GetCaption() }
func (f *DateField) GetType() reflect.Type                            { return reflect.TypeOf(string("")) }
func (f *DateField) IsDerivable() bool                                { return false }
func (f *DateField) IsRequired() bool                                 { return (*baseSimpleField)(f).IsRequired() }
func (f *DateField) GetDependsOn() []string                           { return nil }
func (f *DateField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *DateField) Check(v interface{}) error                        { return (*baseSimpleField)(f).Check(v) }
func (f *DateField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &DateField{id, caption, required, f.Default, f.CheckFunc}
}
func (f *DateField) WriteSQL(sqlBuf *SqlBuffer) { (*baseSimpleField)(f).WriteSQL("DATE", sqlBuf) }

//TIME
type TimeField baseSimpleField

func (f *TimeField) GetId() string                                    { return (*baseSimpleField)(f).GetId() }
func (f *TimeField) GetCaption() string                               { return (*baseSimpleField)(f).GetCaption() }
func (f *TimeField) GetType() reflect.Type                            { return reflect.TypeOf(string("")) }
func (f *TimeField) IsDerivable() bool                                { return false }
func (f *TimeField) IsRequired() bool                                 { return (*baseSimpleField)(f).IsRequired() }
func (f *TimeField) GetDependsOn() []string                           { return nil }
func (f *TimeField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *TimeField) Check(v interface{}) error                        { return (*baseSimpleField)(f).Check(v) }
func (f *TimeField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TimeField{id, caption, required, f.Default, f.CheckFunc}
}
func (f *TimeField) WriteSQL(sqlBuf *SqlBuffer) { (*baseSimpleField)(f).WriteSQL("TIME", sqlBuf) }

//DATETIME
type DateTimeField baseSimpleField

func (f *DateTimeField) GetId() string                                    { return (*baseSimpleField)(f).GetId() }
func (f *DateTimeField) GetCaption() string                               { return (*baseSimpleField)(f).GetCaption() }
func (f *DateTimeField) GetType() reflect.Type                            { return reflect.TypeOf(string("")) }
func (f *DateTimeField) IsDerivable() bool                                { return false }
func (f *DateTimeField) IsRequired() bool                                 { return (*baseSimpleField)(f).IsRequired() }
func (f *DateTimeField) GetDependsOn() []string                           { return nil }
func (f *DateTimeField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *DateTimeField) Check(v interface{}) error                        { return (*baseSimpleField)(f).Check(v) }
func (f *DateTimeField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &DateTimeField{id, caption, required, f.Default, f.CheckFunc}
}
func (f *DateTimeField) WriteSQL(sqlBuf *SqlBuffer) {
	(*baseSimpleField)(f).WriteSQL("DATETIME", sqlBuf)
}

//BOOLEAN
type BooleanField baseSimpleField

func (f *BooleanField) GetId() string                                    { return (*baseSimpleField)(f).GetId() }
func (f *BooleanField) GetCaption() string                               { return (*baseSimpleField)(f).GetCaption() }
func (f *BooleanField) GetType() reflect.Type                            { return reflect.TypeOf(bool(false)) }
func (f *BooleanField) IsDerivable() bool                                { return false }
func (f *BooleanField) IsRequired() bool                                 { return (*baseSimpleField)(f).IsRequired() }
func (f *BooleanField) GetDependsOn() []string                           { return nil }
func (f *BooleanField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *BooleanField) Check(v interface{}) error                        { return (*baseSimpleField)(f).Check(v) }
func (f *BooleanField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BooleanField{id, caption, required, f.Default, f.CheckFunc}
}
func (f *BooleanField) WriteSQL(sqlBuf *SqlBuffer) { (*baseSimpleField)(f).WriteSQL("BOOLEAN", sqlBuf) }
