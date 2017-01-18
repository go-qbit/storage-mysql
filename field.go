//go:generate go run ./gen/fields.go -filename fields_definitions.go

package mysql

import (
	"github.com/go-qbit/model"
)

type IMysqlFieldDefinition interface {
	model.IFieldDefinition
	IsAutoIncremented() bool
	WriteSQL(buf *SqlBuffer)
}
