package mysql

import (
	"reflect"
	"strconv"

	"github.com/go-qbit/model"
)

type baseIntField struct {
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

func (f *baseIntField) GetId() string      { return f.Id }
func (f *baseIntField) GetCaption() string { return f.Caption }
func (f *baseIntField) IsRequired() bool   { return f.NotNull && f.Default == 0 }
func (f *baseIntField) Check(v interface{}) error {
	if f.CheckFunc != nil {
		return f.CheckFunc(v)
	} else {
		return nil
	}
}
func (f *baseIntField) WriteSQL(ftype string, sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString(ftype)

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

// TINYINT
type TinyIntField baseIntField

func (f *TinyIntField) GetId() string      { return (*baseIntField)(f).GetId() }
func (f *TinyIntField) GetCaption() string { return (*baseIntField)(f).GetCaption() }
func (f *TinyIntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint8(0))
	} else {
		return reflect.TypeOf(int8(0))
	}
}
func (f *TinyIntField) IsDerivable() bool                                { return false }
func (f *TinyIntField) IsRequired() bool                                 { return (*baseIntField)(f).IsRequired() }
func (f *TinyIntField) GetDependsOn() []string                           { return nil }
func (f *TinyIntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *TinyIntField) Check(v interface{}) error                        { return (*baseIntField)(f).Check(v) }
func (f *TinyIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TinyIntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *TinyIntField) WriteSQL(sqlBuf *SqlBuffer) { (*baseIntField)(f).WriteSQL("TINYINT", sqlBuf) }

// SMALLINT
type SmallIntField baseIntField

func (f *SmallIntField) GetId() string      { return (*baseIntField)(f).GetId() }
func (f *SmallIntField) GetCaption() string { return (*baseIntField)(f).GetCaption() }
func (f *SmallIntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint16(0))
	} else {
		return reflect.TypeOf(int16(0))
	}
}
func (f *SmallIntField) IsDerivable() bool                                { return false }
func (f *SmallIntField) IsRequired() bool                                 { return (*baseIntField)(f).IsRequired() }
func (f *SmallIntField) GetDependsOn() []string                           { return nil }
func (f *SmallIntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *SmallIntField) Check(v interface{}) error                        { return (*baseIntField)(f).Check(v) }
func (f *SmallIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &SmallIntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *SmallIntField) WriteSQL(sqlBuf *SqlBuffer) { (*baseIntField)(f).WriteSQL("SMALLINT", sqlBuf) }

// MEDIUMINT
type MediumIntField baseIntField

func (f *MediumIntField) GetId() string      { return (*baseIntField)(f).GetId() }
func (f *MediumIntField) GetCaption() string { return (*baseIntField)(f).GetCaption() }
func (f *MediumIntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint32(0))
	} else {
		return reflect.TypeOf(int32(0))
	}
}
func (f *MediumIntField) IsDerivable() bool                                { return false }
func (f *MediumIntField) IsRequired() bool                                 { return (*baseIntField)(f).IsRequired() }
func (f *MediumIntField) GetDependsOn() []string                           { return nil }
func (f *MediumIntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *MediumIntField) Check(v interface{}) error                        { return (*baseIntField)(f).Check(v) }
func (f *MediumIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &MediumIntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *MediumIntField) WriteSQL(sqlBuf *SqlBuffer) { (*baseIntField)(f).WriteSQL("TINYINT", sqlBuf) }

// INT
type IntField baseIntField

func (f *IntField) GetId() string      { return (*baseIntField)(f).GetId() }
func (f *IntField) GetCaption() string { return (*baseIntField)(f).GetCaption() }
func (f *IntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint32(0))
	} else {
		return reflect.TypeOf(int32(0))
	}
}
func (f *IntField) IsDerivable() bool                                { return false }
func (f *IntField) IsRequired() bool                                 { return (*baseIntField)(f).IsRequired() }
func (f *IntField) GetDependsOn() []string                           { return nil }
func (f *IntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *IntField) Check(v interface{}) error                        { return (*baseIntField)(f).Check(v) }
func (f *IntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &IntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *IntField) WriteSQL(sqlBuf *SqlBuffer) { (*baseIntField)(f).WriteSQL("INT", sqlBuf) }

// BIGINT
type BigIntField baseIntField

func (f *BigIntField) GetId() string      { return (*baseIntField)(f).GetId() }
func (f *BigIntField) GetCaption() string { return (*baseIntField)(f).GetCaption() }
func (f *BigIntField) GetType() reflect.Type {
	if f.Unsigned {
		return reflect.TypeOf(uint64(0))
	} else {
		return reflect.TypeOf(int64(0))
	}
}
func (f *BigIntField) IsDerivable() bool                                { return false }
func (f *BigIntField) IsRequired() bool                                 { return (*baseIntField)(f).IsRequired() }
func (f *BigIntField) GetDependsOn() []string                           { return nil }
func (f *BigIntField) Calc(map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *BigIntField) Check(v interface{}) error                        { return (*baseIntField)(f).Check(v) }
func (f *BigIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BigIntField{id, caption, f.Length, f.Unsigned, f.Zerofill, required, false, f.Default, f.CheckFunc}
}
func (f *BigIntField) WriteSQL(sqlBuf *SqlBuffer) { (*baseIntField)(f).WriteSQL("BIGINT", sqlBuf) }
