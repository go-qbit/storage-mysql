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
				&mysql.UintField{
					Id:            "id",
					Caption:       "ID",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.VarCharField{
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
