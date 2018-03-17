// The file was generated, do not change by hands

package mysql

import (
	"context"
	"reflect"
	"strconv"

	"github.com/go-qbit/model"
	"github.com/go-qbit/rbac"
)

type DateField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *DateField) GetId() string      { return f.Id }
func (f *DateField) GetCaption() string { return f.Caption }
func (f *DateField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *DateField) GetStorageType() string {
	res := "DATE"

	return res
}
func (f *DateField) IsDerivable() bool                   { return false }
func (f *DateField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *DateField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *DateField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *DateField) GetDependsOn() []string              { return nil }
func (f *DateField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *DateField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *DateField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *DateField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &DateField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *DateField) IsAutoIncremented() bool { return false }
func (f *DateField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("DATE")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TimeField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *TimeField) GetId() string      { return f.Id }
func (f *TimeField) GetCaption() string { return f.Caption }
func (f *TimeField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *TimeField) GetStorageType() string {
	res := "TIME"

	return res
}
func (f *TimeField) IsDerivable() bool                   { return false }
func (f *TimeField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *TimeField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TimeField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TimeField) GetDependsOn() []string              { return nil }
func (f *TimeField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TimeField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *TimeField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *TimeField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TimeField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TimeField) IsAutoIncremented() bool { return false }
func (f *TimeField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TIME")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TimeStampField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *TimeStampField) GetId() string      { return f.Id }
func (f *TimeStampField) GetCaption() string { return f.Caption }
func (f *TimeStampField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *TimeStampField) GetStorageType() string {
	res := "TIMESTAMP"

	return res
}
func (f *TimeStampField) IsDerivable() bool { return false }
func (f *TimeStampField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *TimeStampField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TimeStampField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TimeStampField) GetDependsOn() []string              { return nil }
func (f *TimeStampField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TimeStampField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *TimeStampField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *TimeStampField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TimeStampField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TimeStampField) IsAutoIncremented() bool { return false }
func (f *TimeStampField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TIMESTAMP")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type DateTimeField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *DateTimeField) GetId() string      { return f.Id }
func (f *DateTimeField) GetCaption() string { return f.Caption }
func (f *DateTimeField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *DateTimeField) GetStorageType() string {
	res := "DATETIME"

	return res
}
func (f *DateTimeField) IsDerivable() bool { return false }
func (f *DateTimeField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *DateTimeField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *DateTimeField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *DateTimeField) GetDependsOn() []string              { return nil }
func (f *DateTimeField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *DateTimeField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *DateTimeField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *DateTimeField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &DateTimeField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *DateTimeField) IsAutoIncremented() bool { return false }
func (f *DateTimeField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("DATETIME")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type YearField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *YearField) GetId() string      { return f.Id }
func (f *YearField) GetCaption() string { return f.Caption }
func (f *YearField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *YearField) GetStorageType() string {
	res := "YEAR"

	return res
}
func (f *YearField) IsDerivable() bool                   { return false }
func (f *YearField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *YearField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *YearField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *YearField) GetDependsOn() []string              { return nil }
func (f *YearField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *YearField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *YearField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *YearField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &YearField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *YearField) IsAutoIncremented() bool { return false }
func (f *YearField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("YEAR")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TinyBlobField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *TinyBlobField) GetId() string      { return f.Id }
func (f *TinyBlobField) GetCaption() string { return f.Caption }
func (f *TinyBlobField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *TinyBlobField) GetStorageType() string {
	res := "TINYBLOB"

	return res
}
func (f *TinyBlobField) IsDerivable() bool { return false }
func (f *TinyBlobField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *TinyBlobField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TinyBlobField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TinyBlobField) GetDependsOn() []string              { return nil }
func (f *TinyBlobField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TinyBlobField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *TinyBlobField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *TinyBlobField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TinyBlobField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TinyBlobField) IsAutoIncremented() bool { return false }
func (f *TinyBlobField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TINYBLOB")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BlobField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *BlobField) GetId() string      { return f.Id }
func (f *BlobField) GetCaption() string { return f.Caption }
func (f *BlobField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *BlobField) GetStorageType() string {
	res := "BLOB"

	return res
}
func (f *BlobField) IsDerivable() bool                   { return false }
func (f *BlobField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *BlobField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *BlobField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *BlobField) GetDependsOn() []string              { return nil }
func (f *BlobField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *BlobField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *BlobField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *BlobField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BlobField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BlobField) IsAutoIncremented() bool { return false }
func (f *BlobField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BLOB")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type MediumBlobField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *MediumBlobField) GetId() string      { return f.Id }
func (f *MediumBlobField) GetCaption() string { return f.Caption }
func (f *MediumBlobField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *MediumBlobField) GetStorageType() string {
	res := "MEDIUMBLOB"

	return res
}
func (f *MediumBlobField) IsDerivable() bool { return false }
func (f *MediumBlobField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *MediumBlobField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *MediumBlobField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *MediumBlobField) GetDependsOn() []string              { return nil }
func (f *MediumBlobField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *MediumBlobField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *MediumBlobField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *MediumBlobField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &MediumBlobField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *MediumBlobField) IsAutoIncremented() bool { return false }
func (f *MediumBlobField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("MEDIUMBLOB")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type LongBlobField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *LongBlobField) GetId() string      { return f.Id }
func (f *LongBlobField) GetCaption() string { return f.Caption }
func (f *LongBlobField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *LongBlobField) GetStorageType() string {
	res := "LONGBLOB"

	return res
}
func (f *LongBlobField) IsDerivable() bool { return false }
func (f *LongBlobField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *LongBlobField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *LongBlobField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *LongBlobField) GetDependsOn() []string              { return nil }
func (f *LongBlobField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *LongBlobField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *LongBlobField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *LongBlobField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &LongBlobField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *LongBlobField) IsAutoIncremented() bool { return false }
func (f *LongBlobField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("LONGBLOB")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BooleanField struct {
	Id             string
	Caption        string
	NotNull        bool
	Default        *bool
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value bool) error
	CleanFunc      func(ctx context.Context, value bool) (bool, error)
}

func (f *BooleanField) GetId() string      { return f.Id }
func (f *BooleanField) GetCaption() string { return f.Caption }
func (f *BooleanField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(bool(false))
	} else {
		return reflect.PtrTo(reflect.TypeOf(bool(false)))
	}
}
func (f *BooleanField) GetStorageType() string {
	res := "BOOLEAN"

	return res
}
func (f *BooleanField) IsDerivable() bool { return false }
func (f *BooleanField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *BooleanField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *BooleanField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *BooleanField) GetDependsOn() []string              { return nil }
func (f *BooleanField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *BooleanField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(bool))
	} else {
		if v.(*bool) != nil {
			return f.CheckFunc(ctx, *v.(*bool))
		}
	}
	return nil
}
func (f *BooleanField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(bool))
	} else {
		if v.(*bool) != nil {
			return f.CleanFunc(ctx, *v.(*bool))
		}
	}
	return v, nil
}
func (f *BooleanField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BooleanField{id, caption, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BooleanField) IsAutoIncremented() bool { return false }
func (f *BooleanField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BOOLEAN")

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TinyIntField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *int8
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value int8) error
	CleanFunc      func(ctx context.Context, value int8) (int8, error)
}

func (f *TinyIntField) GetId() string      { return f.Id }
func (f *TinyIntField) GetCaption() string { return f.Caption }
func (f *TinyIntField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(int8(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(int8(0)))
	}
}
func (f *TinyIntField) GetStorageType() string {
	res := "TINYINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *TinyIntField) IsDerivable() bool { return false }
func (f *TinyIntField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *TinyIntField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TinyIntField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TinyIntField) GetDependsOn() []string              { return nil }
func (f *TinyIntField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TinyIntField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(int8))
	} else {
		if v.(*int8) != nil {
			return f.CheckFunc(ctx, *v.(*int8))
		}
	}
	return nil
}
func (f *TinyIntField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(int8))
	} else {
		if v.(*int8) != nil {
			return f.CleanFunc(ctx, *v.(*int8))
		}
	}
	return v, nil
}
func (f *TinyIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TinyIntField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TinyIntField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *TinyIntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TINYINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
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

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type SmallIntField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *int16
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value int16) error
	CleanFunc      func(ctx context.Context, value int16) (int16, error)
}

func (f *SmallIntField) GetId() string      { return f.Id }
func (f *SmallIntField) GetCaption() string { return f.Caption }
func (f *SmallIntField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(int16(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(int16(0)))
	}
}
func (f *SmallIntField) GetStorageType() string {
	res := "SMALLINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *SmallIntField) IsDerivable() bool { return false }
func (f *SmallIntField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *SmallIntField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *SmallIntField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *SmallIntField) GetDependsOn() []string              { return nil }
func (f *SmallIntField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *SmallIntField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(int16))
	} else {
		if v.(*int16) != nil {
			return f.CheckFunc(ctx, *v.(*int16))
		}
	}
	return nil
}
func (f *SmallIntField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(int16))
	} else {
		if v.(*int16) != nil {
			return f.CleanFunc(ctx, *v.(*int16))
		}
	}
	return v, nil
}
func (f *SmallIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &SmallIntField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *SmallIntField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *SmallIntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("SMALLINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
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

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type MediumIntField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *int32
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value int32) error
	CleanFunc      func(ctx context.Context, value int32) (int32, error)
}

func (f *MediumIntField) GetId() string      { return f.Id }
func (f *MediumIntField) GetCaption() string { return f.Caption }
func (f *MediumIntField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(int32(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(int32(0)))
	}
}
func (f *MediumIntField) GetStorageType() string {
	res := "MEDIUMINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *MediumIntField) IsDerivable() bool { return false }
func (f *MediumIntField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *MediumIntField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *MediumIntField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *MediumIntField) GetDependsOn() []string              { return nil }
func (f *MediumIntField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *MediumIntField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(int32))
	} else {
		if v.(*int32) != nil {
			return f.CheckFunc(ctx, *v.(*int32))
		}
	}
	return nil
}
func (f *MediumIntField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(int32))
	} else {
		if v.(*int32) != nil {
			return f.CleanFunc(ctx, *v.(*int32))
		}
	}
	return v, nil
}
func (f *MediumIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &MediumIntField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *MediumIntField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *MediumIntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("MEDIUMINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
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

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type IntField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *int32
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value int32) error
	CleanFunc      func(ctx context.Context, value int32) (int32, error)
}

func (f *IntField) GetId() string      { return f.Id }
func (f *IntField) GetCaption() string { return f.Caption }
func (f *IntField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(int32(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(int32(0)))
	}
}
func (f *IntField) GetStorageType() string {
	res := "INT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *IntField) IsDerivable() bool                                                 { return false }
func (f *IntField) IsRequired() bool                                                  { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *IntField) GetViewPermission() *rbac.Permission                               { return f.ViewPermission }
func (f *IntField) GetEditPermission() *rbac.Permission                               { return f.EditPermission }
func (f *IntField) GetDependsOn() []string                                            { return nil }
func (f *IntField) Calc(context.Context, map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *IntField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(int32))
	} else {
		if v.(*int32) != nil {
			return f.CheckFunc(ctx, *v.(*int32))
		}
	}
	return nil
}
func (f *IntField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(int32))
	} else {
		if v.(*int32) != nil {
			return f.CleanFunc(ctx, *v.(*int32))
		}
	}
	return v, nil
}
func (f *IntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &IntField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *IntField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *IntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("INT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
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

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BigIntField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *int64
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value int64) error
	CleanFunc      func(ctx context.Context, value int64) (int64, error)
}

func (f *BigIntField) GetId() string      { return f.Id }
func (f *BigIntField) GetCaption() string { return f.Caption }
func (f *BigIntField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(int64(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(int64(0)))
	}
}
func (f *BigIntField) GetStorageType() string {
	res := "BIGINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *BigIntField) IsDerivable() bool { return false }
func (f *BigIntField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *BigIntField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *BigIntField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *BigIntField) GetDependsOn() []string              { return nil }
func (f *BigIntField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *BigIntField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(int64))
	} else {
		if v.(*int64) != nil {
			return f.CheckFunc(ctx, *v.(*int64))
		}
	}
	return nil
}
func (f *BigIntField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(int64))
	} else {
		if v.(*int64) != nil {
			return f.CleanFunc(ctx, *v.(*int64))
		}
	}
	return v, nil
}
func (f *BigIntField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BigIntField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BigIntField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *BigIntField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BIGINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
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

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TinyUintField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *uint8
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value uint8) error
	CleanFunc      func(ctx context.Context, value uint8) (uint8, error)
}

func (f *TinyUintField) GetId() string      { return f.Id }
func (f *TinyUintField) GetCaption() string { return f.Caption }
func (f *TinyUintField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(uint8(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(uint8(0)))
	}
}
func (f *TinyUintField) GetStorageType() string {
	res := "TINYINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	res += " UNSIGNED"

	return res
}
func (f *TinyUintField) IsDerivable() bool { return false }
func (f *TinyUintField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *TinyUintField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TinyUintField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TinyUintField) GetDependsOn() []string              { return nil }
func (f *TinyUintField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TinyUintField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(uint8))
	} else {
		if v.(*uint8) != nil {
			return f.CheckFunc(ctx, *v.(*uint8))
		}
	}
	return nil
}
func (f *TinyUintField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(uint8))
	} else {
		if v.(*uint8) != nil {
			return f.CleanFunc(ctx, *v.(*uint8))
		}
	}
	return v, nil
}
func (f *TinyUintField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TinyUintField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TinyUintField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *TinyUintField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TINYINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	sqlBuf.WriteString(" UNSIGNED")

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type SmallUintField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *uint16
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value uint16) error
	CleanFunc      func(ctx context.Context, value uint16) (uint16, error)
}

func (f *SmallUintField) GetId() string      { return f.Id }
func (f *SmallUintField) GetCaption() string { return f.Caption }
func (f *SmallUintField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(uint16(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(uint16(0)))
	}
}
func (f *SmallUintField) GetStorageType() string {
	res := "SMALLINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	res += " UNSIGNED"

	return res
}
func (f *SmallUintField) IsDerivable() bool { return false }
func (f *SmallUintField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *SmallUintField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *SmallUintField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *SmallUintField) GetDependsOn() []string              { return nil }
func (f *SmallUintField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *SmallUintField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(uint16))
	} else {
		if v.(*uint16) != nil {
			return f.CheckFunc(ctx, *v.(*uint16))
		}
	}
	return nil
}
func (f *SmallUintField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(uint16))
	} else {
		if v.(*uint16) != nil {
			return f.CleanFunc(ctx, *v.(*uint16))
		}
	}
	return v, nil
}
func (f *SmallUintField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &SmallUintField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *SmallUintField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *SmallUintField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("SMALLINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	sqlBuf.WriteString(" UNSIGNED")

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type MediumUintField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *uint32
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value uint32) error
	CleanFunc      func(ctx context.Context, value uint32) (uint32, error)
}

func (f *MediumUintField) GetId() string      { return f.Id }
func (f *MediumUintField) GetCaption() string { return f.Caption }
func (f *MediumUintField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(uint32(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(uint32(0)))
	}
}
func (f *MediumUintField) GetStorageType() string {
	res := "MEDIUMINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	res += " UNSIGNED"

	return res
}
func (f *MediumUintField) IsDerivable() bool { return false }
func (f *MediumUintField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *MediumUintField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *MediumUintField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *MediumUintField) GetDependsOn() []string              { return nil }
func (f *MediumUintField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *MediumUintField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(uint32))
	} else {
		if v.(*uint32) != nil {
			return f.CheckFunc(ctx, *v.(*uint32))
		}
	}
	return nil
}
func (f *MediumUintField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(uint32))
	} else {
		if v.(*uint32) != nil {
			return f.CleanFunc(ctx, *v.(*uint32))
		}
	}
	return v, nil
}
func (f *MediumUintField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &MediumUintField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *MediumUintField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *MediumUintField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("MEDIUMINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	sqlBuf.WriteString(" UNSIGNED")

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type UintField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *uint32
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value uint32) error
	CleanFunc      func(ctx context.Context, value uint32) (uint32, error)
}

func (f *UintField) GetId() string      { return f.Id }
func (f *UintField) GetCaption() string { return f.Caption }
func (f *UintField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(uint32(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(uint32(0)))
	}
}
func (f *UintField) GetStorageType() string {
	res := "INT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	res += " UNSIGNED"

	return res
}
func (f *UintField) IsDerivable() bool                   { return false }
func (f *UintField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *UintField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *UintField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *UintField) GetDependsOn() []string              { return nil }
func (f *UintField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *UintField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(uint32))
	} else {
		if v.(*uint32) != nil {
			return f.CheckFunc(ctx, *v.(*uint32))
		}
	}
	return nil
}
func (f *UintField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(uint32))
	} else {
		if v.(*uint32) != nil {
			return f.CleanFunc(ctx, *v.(*uint32))
		}
	}
	return v, nil
}
func (f *UintField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &UintField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *UintField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *UintField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("INT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	sqlBuf.WriteString(" UNSIGNED")

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BigUintField struct {
	Id             string
	Caption        string
	Length         int
	AutoIncrement  bool
	Zerofill       bool
	NotNull        bool
	Default        *uint64
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value uint64) error
	CleanFunc      func(ctx context.Context, value uint64) (uint64, error)
}

func (f *BigUintField) GetId() string      { return f.Id }
func (f *BigUintField) GetCaption() string { return f.Caption }
func (f *BigUintField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(uint64(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(uint64(0)))
	}
}
func (f *BigUintField) GetStorageType() string {
	res := "BIGINT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	res += " UNSIGNED"

	return res
}
func (f *BigUintField) IsDerivable() bool { return false }
func (f *BigUintField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *BigUintField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *BigUintField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *BigUintField) GetDependsOn() []string              { return nil }
func (f *BigUintField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *BigUintField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(uint64))
	} else {
		if v.(*uint64) != nil {
			return f.CheckFunc(ctx, *v.(*uint64))
		}
	}
	return nil
}
func (f *BigUintField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(uint64))
	} else {
		if v.(*uint64) != nil {
			return f.CleanFunc(ctx, *v.(*uint64))
		}
	}
	return v, nil
}
func (f *BigUintField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BigUintField{id, caption, f.Length, false, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BigUintField) IsAutoIncremented() bool { return f.AutoIncrement }
func (f *BigUintField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BIGINT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	sqlBuf.WriteString(" UNSIGNED")

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.AutoIncrement {
		sqlBuf.WriteString(" AUTO_INCREMENT")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type RealField struct {
	Id             string
	Caption        string
	Length         int
	Decimals       int
	Zerofill       bool
	NotNull        bool
	Default        *float64
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value float64) error
	CleanFunc      func(ctx context.Context, value float64) (float64, error)
}

func (f *RealField) GetId() string      { return f.Id }
func (f *RealField) GetCaption() string { return f.Caption }
func (f *RealField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(float64(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(float64(0)))
	}
}
func (f *RealField) GetStorageType() string {
	res := "REAL"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *RealField) IsDerivable() bool                   { return false }
func (f *RealField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *RealField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *RealField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *RealField) GetDependsOn() []string              { return nil }
func (f *RealField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *RealField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(float64))
	} else {
		if v.(*float64) != nil {
			return f.CheckFunc(ctx, *v.(*float64))
		}
	}
	return nil
}
func (f *RealField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(float64))
	} else {
		if v.(*float64) != nil {
			return f.CleanFunc(ctx, *v.(*float64))
		}
	}
	return v, nil
}
func (f *RealField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &RealField{id, caption, f.Length, f.Decimals, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *RealField) IsAutoIncremented() bool { return false }
func (f *RealField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("REAL")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		if f.Decimals != 0 {
			sqlBuf.WriteByte(',')
			sqlBuf.WriteString(strconv.Itoa(f.Decimals))
		}

		sqlBuf.WriteByte(')')
	}

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type FloatField struct {
	Id             string
	Caption        string
	Length         int
	Decimals       int
	Zerofill       bool
	NotNull        bool
	Default        *float64
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value float64) error
	CleanFunc      func(ctx context.Context, value float64) (float64, error)
}

func (f *FloatField) GetId() string      { return f.Id }
func (f *FloatField) GetCaption() string { return f.Caption }
func (f *FloatField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(float64(0))
	} else {
		return reflect.PtrTo(reflect.TypeOf(float64(0)))
	}
}
func (f *FloatField) GetStorageType() string {
	res := "FLOAT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *FloatField) IsDerivable() bool                   { return false }
func (f *FloatField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *FloatField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *FloatField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *FloatField) GetDependsOn() []string              { return nil }
func (f *FloatField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *FloatField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(float64))
	} else {
		if v.(*float64) != nil {
			return f.CheckFunc(ctx, *v.(*float64))
		}
	}
	return nil
}
func (f *FloatField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(float64))
	} else {
		if v.(*float64) != nil {
			return f.CleanFunc(ctx, *v.(*float64))
		}
	}
	return v, nil
}
func (f *FloatField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &FloatField{id, caption, f.Length, f.Decimals, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *FloatField) IsAutoIncremented() bool { return false }
func (f *FloatField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("FLOAT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		if f.Decimals != 0 {
			sqlBuf.WriteByte(',')
			sqlBuf.WriteString(strconv.Itoa(f.Decimals))
		}

		sqlBuf.WriteByte(')')
	}

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type DecimalField struct {
	Id             string
	Caption        string
	Length         int
	Decimals       int
	Zerofill       bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *DecimalField) GetId() string      { return f.Id }
func (f *DecimalField) GetCaption() string { return f.Caption }
func (f *DecimalField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *DecimalField) GetStorageType() string {
	res := "DECIMAL"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *DecimalField) IsDerivable() bool { return false }
func (f *DecimalField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *DecimalField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *DecimalField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *DecimalField) GetDependsOn() []string              { return nil }
func (f *DecimalField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *DecimalField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *DecimalField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *DecimalField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &DecimalField{id, caption, f.Length, f.Decimals, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *DecimalField) IsAutoIncremented() bool { return false }
func (f *DecimalField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("DECIMAL")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		if f.Decimals != 0 {
			sqlBuf.WriteByte(',')
			sqlBuf.WriteString(strconv.Itoa(f.Decimals))
		}

		sqlBuf.WriteByte(')')
	}

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type NumericField struct {
	Id             string
	Caption        string
	Length         int
	Decimals       int
	Zerofill       bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *NumericField) GetId() string      { return f.Id }
func (f *NumericField) GetCaption() string { return f.Caption }
func (f *NumericField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *NumericField) GetStorageType() string {
	res := "NUMERIC"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *NumericField) IsDerivable() bool { return false }
func (f *NumericField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *NumericField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *NumericField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *NumericField) GetDependsOn() []string              { return nil }
func (f *NumericField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *NumericField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *NumericField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *NumericField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &NumericField{id, caption, f.Length, f.Decimals, f.Zerofill, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *NumericField) IsAutoIncremented() bool { return false }
func (f *NumericField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("NUMERIC")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		if f.Decimals != 0 {
			sqlBuf.WriteByte(',')
			sqlBuf.WriteString(strconv.Itoa(f.Decimals))
		}

		sqlBuf.WriteByte(')')
	}

	if f.Zerofill {
		sqlBuf.WriteString(" ZEROFILL")
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BitField struct {
	Id             string
	Caption        string
	Length         int
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *BitField) GetId() string      { return f.Id }
func (f *BitField) GetCaption() string { return f.Caption }
func (f *BitField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *BitField) GetStorageType() string {
	res := "BIT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *BitField) IsDerivable() bool                                                 { return false }
func (f *BitField) IsRequired() bool                                                  { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *BitField) GetViewPermission() *rbac.Permission                               { return f.ViewPermission }
func (f *BitField) GetEditPermission() *rbac.Permission                               { return f.EditPermission }
func (f *BitField) GetDependsOn() []string                                            { return nil }
func (f *BitField) Calc(context.Context, map[string]interface{}) (interface{}, error) { return nil, nil }
func (f *BitField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *BitField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *BitField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BitField{id, caption, f.Length, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BitField) IsAutoIncremented() bool { return false }
func (f *BitField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BIT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type BinaryField struct {
	Id             string
	Caption        string
	Length         int
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *BinaryField) GetId() string      { return f.Id }
func (f *BinaryField) GetCaption() string { return f.Caption }
func (f *BinaryField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *BinaryField) GetStorageType() string {
	res := "BINARY"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *BinaryField) IsDerivable() bool { return false }
func (f *BinaryField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *BinaryField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *BinaryField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *BinaryField) GetDependsOn() []string              { return nil }
func (f *BinaryField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *BinaryField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *BinaryField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *BinaryField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &BinaryField{id, caption, f.Length, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *BinaryField) IsAutoIncremented() bool { return false }
func (f *BinaryField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("BINARY")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type VarBinaryField struct {
	Id             string
	Caption        string
	Length         int
	NotNull        bool
	Default        *[]byte
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value []byte) error
	CleanFunc      func(ctx context.Context, value []byte) ([]byte, error)
}

func (f *VarBinaryField) GetId() string      { return f.Id }
func (f *VarBinaryField) GetCaption() string { return f.Caption }
func (f *VarBinaryField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf([]byte{})
	} else {
		return reflect.PtrTo(reflect.TypeOf([]byte{}))
	}
}
func (f *VarBinaryField) GetStorageType() string {
	res := "VARBINARY"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *VarBinaryField) IsDerivable() bool { return false }
func (f *VarBinaryField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *VarBinaryField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *VarBinaryField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *VarBinaryField) GetDependsOn() []string              { return nil }
func (f *VarBinaryField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *VarBinaryField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CheckFunc(ctx, *v.(*[]byte))
		}
	}
	return nil
}
func (f *VarBinaryField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.([]byte))
	} else {
		if v.(*[]byte) != nil {
			return f.CleanFunc(ctx, *v.(*[]byte))
		}
	}
	return v, nil
}
func (f *VarBinaryField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &VarBinaryField{id, caption, f.Length, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *VarBinaryField) IsAutoIncremented() bool { return false }
func (f *VarBinaryField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("VARBINARY")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type CharField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *CharField) GetId() string      { return f.Id }
func (f *CharField) GetCaption() string { return f.Caption }
func (f *CharField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *CharField) GetStorageType() string {
	res := "CHAR"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *CharField) IsDerivable() bool                   { return false }
func (f *CharField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *CharField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *CharField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *CharField) GetDependsOn() []string              { return nil }
func (f *CharField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *CharField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *CharField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *CharField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &CharField{id, caption, f.Length, f.Charset, f.Collate, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *CharField) IsAutoIncremented() bool { return false }
func (f *CharField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("CHAR")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type VarCharField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *VarCharField) GetId() string      { return f.Id }
func (f *VarCharField) GetCaption() string { return f.Caption }
func (f *VarCharField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *VarCharField) GetStorageType() string {
	res := "VARCHAR"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *VarCharField) IsDerivable() bool { return false }
func (f *VarCharField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *VarCharField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *VarCharField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *VarCharField) GetDependsOn() []string              { return nil }
func (f *VarCharField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *VarCharField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *VarCharField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *VarCharField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &VarCharField{id, caption, f.Length, f.Charset, f.Collate, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *VarCharField) IsAutoIncremented() bool { return false }
func (f *VarCharField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("VARCHAR")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TinyTextField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	Binary         bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *TinyTextField) GetId() string      { return f.Id }
func (f *TinyTextField) GetCaption() string { return f.Caption }
func (f *TinyTextField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *TinyTextField) GetStorageType() string {
	res := "TINYTEXT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *TinyTextField) IsDerivable() bool { return false }
func (f *TinyTextField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *TinyTextField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TinyTextField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TinyTextField) GetDependsOn() []string              { return nil }
func (f *TinyTextField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TinyTextField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *TinyTextField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *TinyTextField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TinyTextField{id, caption, f.Length, f.Charset, f.Collate, f.Binary, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TinyTextField) IsAutoIncremented() bool { return false }
func (f *TinyTextField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TINYTEXT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Binary {
		sqlBuf.WriteString(" BINARY")
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type TextField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	Binary         bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *TextField) GetId() string      { return f.Id }
func (f *TextField) GetCaption() string { return f.Caption }
func (f *TextField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *TextField) GetStorageType() string {
	res := "TEXT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *TextField) IsDerivable() bool                   { return false }
func (f *TextField) IsRequired() bool                    { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }
func (f *TextField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *TextField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *TextField) GetDependsOn() []string              { return nil }
func (f *TextField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *TextField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *TextField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *TextField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &TextField{id, caption, f.Length, f.Charset, f.Collate, f.Binary, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *TextField) IsAutoIncremented() bool { return false }
func (f *TextField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("TEXT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Binary {
		sqlBuf.WriteString(" BINARY")
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type MediumTextField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	Binary         bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *MediumTextField) GetId() string      { return f.Id }
func (f *MediumTextField) GetCaption() string { return f.Caption }
func (f *MediumTextField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *MediumTextField) GetStorageType() string {
	res := "MEDIUMTEXT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *MediumTextField) IsDerivable() bool { return false }
func (f *MediumTextField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *MediumTextField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *MediumTextField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *MediumTextField) GetDependsOn() []string              { return nil }
func (f *MediumTextField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *MediumTextField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *MediumTextField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *MediumTextField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &MediumTextField{id, caption, f.Length, f.Charset, f.Collate, f.Binary, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *MediumTextField) IsAutoIncremented() bool { return false }
func (f *MediumTextField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("MEDIUMTEXT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Binary {
		sqlBuf.WriteString(" BINARY")
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}

type LongTextField struct {
	Id             string
	Caption        string
	Length         int
	Charset        string
	Collate        string
	Binary         bool
	NotNull        bool
	Default        *string
	ViewPermission *rbac.Permission
	EditPermission *rbac.Permission
	CheckFunc      func(ctx context.Context, value string) error
	CleanFunc      func(ctx context.Context, value string) (string, error)
}

func (f *LongTextField) GetId() string      { return f.Id }
func (f *LongTextField) GetCaption() string { return f.Caption }
func (f *LongTextField) GetType() reflect.Type {
	if f.NotNull {
		return reflect.TypeOf(string(""))
	} else {
		return reflect.PtrTo(reflect.TypeOf(string("")))
	}
}
func (f *LongTextField) GetStorageType() string {
	res := "LONGTEXT"

	if f.Length != 0 {
		res += "(" + strconv.Itoa(f.Length) + ")"
	}

	return res
}
func (f *LongTextField) IsDerivable() bool { return false }
func (f *LongTextField) IsRequired() bool {
	return f.NotNull && f.Default == nil && !f.IsAutoIncremented()
}
func (f *LongTextField) GetViewPermission() *rbac.Permission { return f.ViewPermission }
func (f *LongTextField) GetEditPermission() *rbac.Permission { return f.EditPermission }
func (f *LongTextField) GetDependsOn() []string              { return nil }
func (f *LongTextField) Calc(context.Context, map[string]interface{}) (interface{}, error) {
	return nil, nil
}
func (f *LongTextField) Check(ctx context.Context, v interface{}) error {
	if f.CheckFunc == nil {
		return nil
	}
	if f.NotNull {
		return f.CheckFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CheckFunc(ctx, *v.(*string))
		}
	}
	return nil
}
func (f *LongTextField) Clean(ctx context.Context, v interface{}) (interface{}, error) {
	if f.CleanFunc == nil {
		return v, nil
	}
	if f.NotNull {
		return f.CleanFunc(ctx, v.(string))
	} else {
		if v.(*string) != nil {
			return f.CleanFunc(ctx, *v.(*string))
		}
	}
	return v, nil
}
func (f *LongTextField) CloneForFK(id string, caption string, required bool) model.IFieldDefinition {
	return &LongTextField{id, caption, f.Length, f.Charset, f.Collate, f.Binary, required, f.Default, f.ViewPermission, f.EditPermission, f.CheckFunc, f.CleanFunc}
}
func (f *LongTextField) IsAutoIncremented() bool { return false }
func (f *LongTextField) WriteSQL(sqlBuf *SqlBuffer) {
	sqlBuf.WriteIdentifier(f.Id)

	sqlBuf.WriteByte(' ')
	sqlBuf.WriteString("LONGTEXT")

	if f.Length != 0 {
		sqlBuf.WriteByte('(')
		sqlBuf.WriteString(strconv.Itoa(f.Length))
		sqlBuf.WriteByte(')')
	}

	if f.Binary {
		sqlBuf.WriteString(" BINARY")
	}

	if f.Charset != "" {
		sqlBuf.WriteString(" CHARACTER SET ")
		sqlBuf.WriteValue(f.Charset)
	}

	if f.Collate != "" {
		sqlBuf.WriteString(" COLLATE ")
		sqlBuf.WriteValue(f.Collate)
	}

	if f.NotNull {
		sqlBuf.WriteString(" NOT NULL")
	}

	if f.Default != nil {
		sqlBuf.WriteString(" DEFAULT ")
		sqlBuf.WriteValue(*f.Default)
	}

}
