package test

import (
	"github.com/go-qbit/model"
	"github.com/go-qbit/storage-mysql"
)

type User struct {
	*mysql.BaseModel
}

func NewUser(storage *mysql.MySQL) *User {
	return &User{
		BaseModel: mysql.NewBaseModel(
			storage,
			"user",
			[]mysql.IMysqlFieldDefinition{
				&mysql.UintField{
					Id:            "id",
					Caption:       "ID",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.VarCharField{
					Id:      "name",
					Caption: "Name",
					Length:  255,
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "lastname",
					Caption: "Lastame",
					Length:  255,
					NotNull: true,
				},
			},
			[]model.IFieldDefinition{
				&model.DerivableField{
					Id:        "fullname",
					Caption:   "Full name",
					DependsOn: []string{"name", "lastname"},
					Get: func(row map[string]interface{}) (interface{}, error) {
						return row["name"].(string) + " " + row["lastname"].(string), nil
					},
				},
			},
			mysql.BaseModelOpts{
				BaseModelOpts: model.BaseModelOpts{
					PkFieldsNames: []string{"id"},
				},
				Indexes: []mysql.Index{
					{[]string{"name"}, false},
					{[]string{"lastname", "name"}, false},
				},
			},
		),
	}
}
