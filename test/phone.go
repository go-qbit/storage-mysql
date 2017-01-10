package test

import (
	"fmt"
	"github.com/go-qbit/model"
	"github.com/go-qbit/storage-mysql"
)

type Phone struct {
	*mysql.BaseModel
}

func NewPhone(storage *mysql.MySQL) *Phone {
	return &Phone{
		BaseModel: mysql.NewBaseModel(
			storage,
			"phone",
			[]mysql.IMysqlFieldDefinition{
				&mysql.UintField{
					Id:            "id",
					Caption:       "ID",
					NotNull:       true,
					AutoIncrement: true,
				},

				&mysql.UintField{
					Id:      "country_code",
					Caption: "Country code",
					NotNull: true,
				},

				&mysql.UintField{
					Id:      "code",
					Caption: "Code",
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "number",
					Caption: "Number",
					Length:  10,
					NotNull: true,
				},
			},
			[]model.IFieldDefinition{
				&model.DerivableField{
					Id:        "formated_number",
					Caption:   "Formated number",
					DependsOn: []string{"country_code", "code", "number"},
					Get: func(row map[string]interface{}) (interface{}, error) {
						return fmt.Sprintf("+%d (%d) %s", row["country_code"], row["code"], row["number"]), nil
					},
				},
			},
			[]string{"id"},
			[]mysql.Index{
				{[]string{"country_code", "code", "number"}, true},
			},
		),
	}
}
