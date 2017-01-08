package mysql_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/storage-mysql"
	"github.com/go-qbit/storage-mysql/test"

	"github.com/stretchr/testify/suite"
)

var (
	user    string
	pass    string
	prot    string
	addr    string
	dbname  string
	dsn     string
	netAddr string
)

var (
	_ = mysql.IMysqlFieldDefinition(&mysql.TinyIntField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.SmallIntField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.MediumIntField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.IntField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.BigIntField{})

	_ = mysql.IMysqlFieldDefinition(&mysql.VarCharField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.CharField{})
)

func init() {
	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}
	user = env("MYSQL_TEST_USER", "root")
	pass = env("MYSQL_TEST_PASS", "")
	prot = env("MYSQL_TEST_PROT", "tcp")
	addr = env("MYSQL_TEST_ADDR", "localhost:3306")
	dbname = env("MYSQL_TEST_DBNAME", "gotest")
	netAddr = fmt.Sprintf("%s(%s)", prot, addr)
	dsn = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, "mysql")
}

type DBTestSuite struct {
	suite.Suite
	storage *mysql.MySQL
	user    *test.User
	phone   *test.Phone
	message *test.Message
	address *test.Address
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}

func (s *DBTestSuite) SetupTest() {
	s.storage = mysql.NewMySQL()

	s.user = test.NewUser(s.storage)
	s.phone = test.NewPhone(s.storage)
	s.message = test.NewMessage(s.storage)
	s.address = test.NewAddress(s.storage)

	model.AddOneToOneRelation(s.phone, s.user, false)
	model.AddManyToOneRelation(s.message, s.user)
	model.AddManyToManyRelation(s.user, s.address, s.storage)

	if !s.NoError(s.storage.Connect(dsn)) {
		return
	}

	_, err := s.storage.Exec(context.Background(), "DROP DATABASE IF EXISTS "+dbname)
	if !s.NoError(err) {
		return
	}

	_, err = s.storage.Exec(context.Background(), "CREATE DATABASE "+dbname)
	if !s.NoError(err) {
		return
	}

	_, err = s.storage.Exec(context.Background(), "USE "+dbname)
	if !s.NoError(err) {
		return
	}

	if !s.NoError(s.storage.InitDB(context.Background())) {
		return
	}
}

func (s *DBTestSuite) TestModel_CreateSQL() {
	sqlBuf := mysql.NewSqlBuffer()
	s.storage.WriteCreateSQL(sqlBuf)

	s.Equal("CREATE TABLE `address` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`country` VARCHAR(64) NOT NULL,"+
		"`city` VARCHAR(64) NOT NULL,"+
		"`address` VARCHAR(255) NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE INDEX `uniq_address__country_city_address`(`country`,`city`,`address`)"+
		");\n"+

		"CREATE TABLE `user` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`name` VARCHAR(255) NOT NULL,"+
		"`lastname` VARCHAR(255) NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"INDEX `user__name`(`name`),"+
		"INDEX `user__lastname_name`(`lastname`,`name`)"+
		");\n"+

		"CREATE TABLE `_junction__user__address` ("+
		"`fk__user__id` INT UNSIGNED NOT NULL,"+
		"`fk__address__id` INT UNSIGNED NOT NULL,"+
		"PRIMARY KEY (`fk__user__id`,`fk__address__id`),"+
		"FOREIGN KEY `fk__junction__user__address__fk__user__id___user__id`(`fk__user__id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT,"+
		"FOREIGN KEY `fk__junction__user__address__fk__address__id___address__id`(`fk__address__id`)"+
		"REFERENCES `address`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT);\n"+

		"CREATE TABLE `message` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`text` VARCHAR(255) NOT NULL,"+
		"`fk__user__id` INT UNSIGNED NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"FOREIGN KEY `fk_message__fk__user__id___user__id`(`fk__user__id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT"+
		");\n"+

		"CREATE TABLE `phone` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`country_code` INT UNSIGNED NOT NULL,"+
		"`code` INT UNSIGNED NOT NULL,"+
		"`number` VARCHAR(10) NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE INDEX `uniq_phone__country_code_code_number`(`country_code`,`code`,`number`),"+
		"FOREIGN KEY `fk_phone__id___user__id`(`id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT"+
		");\n", sqlBuf.String(),
	)
}

func (s *DBTestSuite) TestModel_Add() {
	_, err := s.storage.Add(context.Background(), s.user, []map[string]interface{}{
		{"id": 1, "name": "Ivan", "lastname": "Sidorov"},
		{"id": 2, "name": "Petr", "lastname": "Ivanov"},
		{"id": 3, "name": "James", "lastname": "Bond"},
		{"id": 4, "name": "John", "lastname": "Connor"},
		{"id": 5, "name": "Sara", "lastname": "Connor"},
	})
	s.NoError(err)

	_, err = s.phone.AddMulti(context.Background(), []map[string]interface{}{
		{"id": 1, "country_code": 1, "code": 111, "number": 1111111},
		{"id": 3, "country_code": 3, "code": 333, "number": 3333333},
	})
	s.NoError(err)

	_, err = s.message.AddMulti(context.Background(), []map[string]interface{}{
		{"id": 10, "text": "Message 1", "fk__user__id": 1},
		{"id": 20, "text": "Message 2", "fk__user__id": 1},
		{"id": 30, "text": "Message 3", "fk__user__id": 1},
		{"id": 40, "text": "Message 4", "fk__user__id": 2},
	})
	s.NoError(err)

	_, err = s.address.AddMulti(context.Background(), []map[string]interface{}{
		{"id": 100, "country": "USA", "city": "Arlington", "address": "1022 Bridges Dr"},
		{"id": 200, "country": "USA", "city": "Fort Worth", "address": "7105 Plover Circle"},
		{"id": 300, "country": "USA", "city": "Crowley", "address": "524 Pecan Street"},
		{"id": 400, "country": "USA", "city": "Arlington", "address": "1023 Bridges Dr"},
		{"id": 500, "country": "USA", "city": "Louisville", "address": "1246 Everett Avenue"},
	})
	s.NoError(err)

	_, err = s.user.GetRelation("address").JunctionModel.AddMulti(context.Background(), []map[string]interface{}{
		{"fk__user__id": 1, "fk__address__id": 100},
		{"fk__user__id": 1, "fk__address__id": 200},
		{"fk__user__id": 2, "fk__address__id": 200},
		{"fk__user__id": 2, "fk__address__id": 300},
		{"fk__user__id": 3, "fk__address__id": 300},
		{"fk__user__id": 4, "fk__address__id": 400},
		{"fk__user__id": 5, "fk__address__id": 500},
	})
	s.NoError(err)
}

func (s *DBTestSuite) TestModel_Query() {
	s.TestModel_Add()

	data, err := s.storage.Query(context.Background(), s.user, []string{"id", "name"}, model.QueryOptions{
		OrderBy: []model.Order{
			{"id", false},
			{"name", true},
		},
		Limit:  3,
		Offset: 2,
	})
	if !s.NoError(err) {
		return
	}

	s.Equal([]map[string]interface{}{
		{"id": uint32(3), "name": "James"},
		{"id": uint32(4), "name": "John"},
		{"id": uint32(5), "name": "Sara"},
	}, data)
}

func (s *DBTestSuite) TestModel_GetAllToStruc() {
	s.TestModel_Add()

	type PhoneType struct {
		FormattedNumber string `field:"formated_number"`
	}

	type MessageType struct {
		Text string `field:"text"`
	}

	type AddressTYpe struct {
		City    string `field:"city"`
		Address string `field:"address"`
	}

	type UserType struct {
		Id        uint32        `field:"id"`
		Lastname  string        `field:"lastname"`
		Fullname  string        `field:"fullname"`
		PhonePtr  *PhoneType    `field:"phone"`
		Messages  []MessageType `field:"message"`
		Addresses []AddressTYpe `field:"address"`
	}

	var res []UserType

	s.NoError(s.user.GetAllToStruct(
		context.Background(),
		&res,
		model.GetAllOptions{
			Filter: expr.Lt(expr.ModelField("id"), expr.Value(4)),
			OrderBy: []model.Order{
				{"id", false},
			},
			Limit: 3,
		},
	))

	s.Equal(
		[]UserType{
			{
				Id:       1,
				Lastname: "Sidorov",
				Fullname: "Ivan Sidorov",
				PhonePtr: &PhoneType{
					FormattedNumber: "+1 (111) 1111111",
				},

				Messages: []MessageType{
					{Text: "Message 1"},
					{Text: "Message 2"},
					{Text: "Message 3"},
				},
				Addresses: []AddressTYpe{
					{City: "Arlington", Address: "1022 Bridges Dr"},
					{City: "Fort Worth", Address: "7105 Plover Circle"},
				},
			},
			{
				Id:       2,
				Lastname: "Ivanov",
				Fullname: "Petr Ivanov",

				Messages: []MessageType{
					{Text: "Message 4"},
				},
				Addresses: []AddressTYpe{
					{City: "Crowley", Address: "524 Pecan Street"},
					{City: "Fort Worth", Address: "7105 Plover Circle"},
				},
			},
			{
				Id:       3,
				Lastname: "Bond",
				Fullname: "James Bond",
				PhonePtr: &PhoneType{
					FormattedNumber: "+3 (333) 3333333",
				},

				Addresses: []AddressTYpe{
					{City: "Crowley", Address: "524 Pecan Street"},
				},
			},
		},
		res,
	)
}
