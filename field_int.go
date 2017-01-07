package mysql

import (
	"reflect"
	"strconv"

	"github.com/go-qbit/model"
)

type IntField struct {
	Id            string
	Caption       string
	Length        int
	Unsigned      bool
	Zerofill      bool
	NotNull       bool
	AutoIncrement bool
	Default       int
	CheckFunc     func(interface{}) error
}

func (f *IntField) GetId() string      { return f.Id }
func (f *IntField) GetCaption() string { return f.Caption }
func (f *IntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint32(0))
	} else {
		return reflect.TypeOf(int32(0))
	}
}
func (f *IntField) IsDerivable() bool                                { return false }
func (f *IntField) IsRequired() bool                                 { return f.NotNull && f.Default == 0 }
func (f *IntField) GetDependsOn() []string                           { return nil }
func (f *IntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *IntField) Check(v interface{}) error {
	if f.CheckFunc != nil {
		return f.CheckFunc(v)
	} else {
		return nil
	}
}
func (f *IntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &IntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *IntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteString(" INT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Unsigned {
		sqlBuf.WriteString(" UNSIGNED")
	}

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != 0 {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(f.Default)
	}
}
