package test

import (
	"github.com/go-qbit/storage-mysql"
)

type Address struct {
	*mysql.BaseModel
}

func NewAddress(storage *mysql.MySQL) *Address {
	return &Address{
		mysql.NewBaseModel(
			storage,
			"address",
			[]mysql.IMysqlFieldDefinition{
				&mysql.IntField{
					Id:            "id",
					Caption:       "ID",
					NotNull:       true,
					Unsigned:      true,
					AutoIncrement: true,
				},

				&mysql.VarCharField{
					Id:      "country",
					Caption: "Country",
					Length:  64,
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "city",
					Caption: "City",
					Length:  64,
					NotNull: true,
				},

				&mysql.VarCharField{
					Id:      "address",
					Caption: "Address",
					Length:  255,
					NotNull: true,
				},
			},
			nil,
			[]string{"id"},
			[]mysql.Index{
				{FieldNames: []string{"country", "city", "address"}, Unique: true},
			},
		),
	}
}
