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
	"github.com/go-qbit/timelog"

	"github.com/stretchr/testify/suite"
)

var (
	user      string
	pass      string
	prot      string
	addr      string
	dbname    string
	mysqlDsn  string
	gotestDsn string
	netAddr   string
)

var (
	_ = mysql.IMysqlFieldDefinition(&mysql.DateField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.TimeField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.DateTimeField{})
	_ = mysql.IMysqlFieldDefinition(&mysql.BooleanField{})

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
	mysqlDsn = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true&", user, pass, netAddr, "mysql")
	gotestDsn = fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true&", user, pass, netAddr, dbname)
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
	model.AddManyToOneRelation(s.message, s.user, false, "", "")
	model.AddManyToManyRelation(s.user, s.address, s.storage)

	if !s.NoError(s.storage.Connect(mysqlDsn)) {
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

	if !s.NoError(s.storage.Connect(gotestDsn)) {
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
		")ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8';\n"+

		"CREATE TABLE `user` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`name` VARCHAR(255) NOT NULL,"+
		"`lastname` VARCHAR(255) NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"INDEX `user__name`(`name`),"+
		"INDEX `user__lastname_name`(`lastname`,`name`)"+
		")ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8';\n"+

		"CREATE TABLE `_junction__user__address` ("+
		"`fk_user_id` INT UNSIGNED NOT NULL,"+
		"`fk_address_id` INT UNSIGNED NOT NULL,"+
		"PRIMARY KEY (`fk_user_id`,`fk_address_id`),"+
		"FOREIGN KEY `fk__junction__user__address__fk_address_id___address__id`(`fk_address_id`)"+
		"REFERENCES `address`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT,"+
		"FOREIGN KEY `fk__junction__user__address__fk_user_id___user__id`(`fk_user_id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT"+
		")ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8';\n"+

		"CREATE TABLE `message` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`text` VARCHAR(255) NOT NULL,"+
		"`fk_user_id` INT UNSIGNED,"+
		"PRIMARY KEY (`id`),"+
		"FOREIGN KEY `fk_message__fk_user_id___user__id`(`fk_user_id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT"+
		")ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8';\n"+

		"CREATE TABLE `phone` ("+
		"`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,"+
		"`country_code` INT UNSIGNED NOT NULL,"+
		"`code` INT UNSIGNED NOT NULL,"+
		"`number` VARCHAR(10) NOT NULL,"+
		"PRIMARY KEY (`id`),"+
		"UNIQUE INDEX `uniq_phone__country_code_code_number`(`country_code`,`code`,`number`),"+
		"FOREIGN KEY `fk_phone__id___user__id`(`id`)"+
		"REFERENCES `user`(`id`)ON UPDATE RESTRICT ON DELETE RESTRICT"+
		")ENGINE='InnoDB' DEFAULT CHARACTER SET 'UTF8';\n", sqlBuf.String(),
	)
}

func (s *DBTestSuite) TestModel_Add() {
	ctx := timelog.Start(context.Background(), "Add data")

	ctx, err := s.storage.StartTransaction(ctx)
	s.NoError(err)

	pks, err := s.user.AddFromStructs(ctx, []struct {
		Name     string
		Lastname string
	}{
		{Name: "Ivan", Lastname: "Sidorov"},
		{Name: "Petr", Lastname: "Ivanov"},
		{Name: "James", Lastname: "Bond"},
		{Name: "John", Lastname: "Connor"},
		{Name: "Sara", Lastname: "Connor"},
	}, model.AddOptions{})
	s.NoError(err)

	s.Equal(model.NewData([]string{"id"}, [][]interface{}{
		[]interface{}{uint32(1)},
		[]interface{}{uint32(2)},
		[]interface{}{uint32(3)},
		[]interface{}{uint32(4)},
		[]interface{}{uint32(5)},
	}), pks)

	phones := []struct {
		Id          int
		CountryCode int
		Code        int
		Number      int
	}{
		{Id: 1, CountryCode: 1, Code: 111, Number: 1111111},
		{Id: 3, CountryCode: 3, Code: 333, Number: 3333333},
	}
	_, err = s.phone.AddFromStructs(ctx, phones, model.AddOptions{})
	s.NoError(err)

	_, err = s.phone.AddFromStructs(ctx, phones, model.AddOptions{Replace: true})
	s.NoError(err)

	_, err = s.message.AddFromStructs(ctx, []struct {
		Id       int
		Text     string
		FkUserId int `field:"fk_user_id"`
	}{
		{Id: 10, Text: "Message 1", FkUserId: 1},
		{Id: 20, Text: "Message 2", FkUserId: 1},
		{Id: 30, Text: "Message 3", FkUserId: 1},
		{Id: 40, Text: "Message 4", FkUserId: 2},
	}, model.AddOptions{})
	s.NoError(err)

	_, err = s.address.AddFromStructs(ctx, []struct {
		Id      int
		Country string
		City    string
		Address string
	}{
		{Id: 100, Country: "USA", City: "Arlington", Address: "1022 Bridges Dr"},
		{Id: 200, Country: "USA", City: "Fort Worth", Address: "7105 Plover Circle"},
		{Id: 300, Country: "USA", City: "Crowley", Address: "524 Pecan Street"},
		{Id: 400, Country: "USA", City: "Arlington", Address: "1023 Bridges Dr"},
		{Id: 500, Country: "USA", City: "Louisville", Address: "1246 Everett Avenue"},
	}, model.AddOptions{})
	s.NoError(err)

	s.NoError(s.user.Link(ctx, s.address, []model.ModelLink{
		{[]interface{}{1}, [][]interface{}{{100}, {200}}},
		{[]interface{}{2}, [][]interface{}{{200}, {300}}},
		{[]interface{}{3}, [][]interface{}{{300}}},
		{[]interface{}{4}, [][]interface{}{{400}}},
		{[]interface{}{5}, [][]interface{}{{500}}},
	}))

	ctx, err = s.storage.Commit(ctx)
	s.NoError(err)

	timelog.Finish(ctx)
	//println(timelog.Get(ctx).Analyze().String())
}

func (s *DBTestSuite) TestModel_Query() {
	s.TestModel_Add()

	var totalRows uint64

	data, err := s.storage.Query(context.Background(), s.user, []string{"id", "name"}, model.GetAllOptions{
		OrderBy: []model.Order{
			{"id", false},
			{"name", true},
		},
		Limit:       3,
		Offset:      2,
		RowsWoLimit: &totalRows,
		ForUpdate:   true,
	})
	if !s.NoError(err) {
		return
	}

	s.Equal(uint64(5), totalRows)

	s.Equal(model.NewData([]string{"id", "name"}, [][]interface{}{
		{uint32(3), "James"},
		{uint32(4), "John"},
		{uint32(5), "Sara"},
	}), data)
}

func (s *DBTestSuite) TestModel_GetAllToStruct() {
	s.TestModel_Add()

	ctx := timelog.Start(context.Background(), "Get all to struct")

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
		ctx, &res, model.GetAllOptions{
			Filter: expr.Or(
				expr.Lt(s.user.FieldExpr("id"), expr.Value(4)),
				expr.Any(s.user, s.message, expr.Eq(s.message.FieldExpr("id"), expr.Value(10))),
				expr.Any(s.user, s.address, expr.Eq(s.address.FieldExpr("id"), expr.Value(100))),
				expr.Eq(s.user.FieldExpr("id"), nil),
			),
			OrderBy: []model.Order{
				{"id", false},
			},
			Limit: 3,
		},
	))

	timelog.Finish(ctx)
	//println(timelog.Get(ctx).Analyze().String())

	s.EqualValues(
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
					{City: "Fort Worth", Address: "7105 Plover Circle"},
					{City: "Crowley", Address: "524 Pecan Street"},
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

func (s *DBTestSuite) TestModel_Transaction() {
	_, err := s.storage.Exec(context.Background(), "CREATE TABLE number (n INT)")
	if !s.NoError(err) {
		return
	}

	ctx, err := s.storage.StartTransaction(context.Background())
	if !s.NoError(err) {
		return
	}

	_, err = s.storage.Exec(ctx, "INSERT INTO number VALUES(100)")
	if !s.NoError(err) {
		return
	}

	ctx, err = s.storage.StartTransaction(ctx)
	if !s.NoError(err) {
		return
	}

	_, err = s.storage.Exec(ctx, "UPDATE number SET n = 1000")
	if !s.NoError(err) {
		return
	}

	ctx, err = s.storage.Rollback(ctx)
	if !s.NoError(err) {
		return
	}

	ctx, err = s.storage.Commit(ctx)
	if !s.NoError(err) {
		return
	}

	rows, err := s.storage.RawQuery(ctx, "SELECT * FROM number")
	if !s.NoError(err) {
		return
	}
	s.True(rows.Next())
	var n int
	rows.Scan(&n)
	s.Equal(100, n)
}

func (s *DBTestSuite) TestBaseModel_Edit() {
	s.TestModel_Add()

	err := s.user.Edit(context.Background(), expr.Eq(s.user.FieldExpr("id"), expr.Value(3)), map[string]interface{}{
		"lastname": "NewName",
	})
	s.NoError(err)

	data, err := s.user.GetAll(
		context.Background(),
		[]string{"id", "lastname"},
		model.GetAllOptions{
			Filter: expr.Lt(expr.ModelField(s.user, "id"), expr.Value(4)),
		},
	)
	s.NoError(err)

	s.Equal([]map[string]interface{}{
		{
			"id":       uint32(1),
			"lastname": "Sidorov",
		},
		{
			"id":       uint32(2),
			"lastname": "Ivanov",
		},
		{
			"id":       uint32(3),
			"lastname": "NewName",
		},
	}, data.Maps())
}

func (s *DBTestSuite) TestBaseModel_Delete() {
	_, err := s.user.AddFromStructs(context.Background(), []struct {
		Name     string
		Lastname string
	}{
		{Name: "Ivan", Lastname: "Sidorov"},
		{Name: "Petr", Lastname: "Ivanov"},
		{Name: "James", Lastname: "Bond"},
		{Name: "John", Lastname: "Connor"},
		{Name: "Sara", Lastname: "Connor"},
	}, model.AddOptions{})
	s.NoError(err)

	s.NoError(s.user.Delete(context.Background(), expr.Eq(s.user.FieldExpr("id"), expr.Value(3))))

	data, err := s.user.GetAll(context.Background(), []string{"id"}, model.GetAllOptions{
		OrderBy: []model.Order{{"id", false}},
	})
	s.NoError(err)

	s.Equal([]map[string]interface{}{
		{"id": uint32(1)},
		{"id": uint32(2)},
		{"id": uint32(4)},
		{"id": uint32(5)},
	}, data.Maps())
}
