package test

import (
	"github.com/go-qbit/storage-mysql"
)

type Message struct {
	*mysql.BaseModel
}

func NewMessage(storage *mysql.MySQL) *Message {
	return &Message{
		mysql.NewBaseModel(
			storage,
			"message",
			[]mysql.IMysqlFieldDefinition{
				&mysql.IntField{
					Id:            "id",
					Caption:       "ID",
					Unsigned:      true,
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.VarcharField{
					Id:      "text",
					Caption: "Message text",
					Length:  255,
					NotNull: true,
				},
			},
			nil,
			[]string{"id"},
			nil,
		),
	}
}
