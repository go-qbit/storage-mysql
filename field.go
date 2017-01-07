package mysql

import (
	"github.com/go-qbit/model"
)

type IMysqlFieldDefinition interface {
	model.IFieldDefinition
	WriteSQL(buf *SqlBuffer)
}
