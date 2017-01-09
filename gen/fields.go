package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
)

var mysqlTypes = []struct {
	mysqlType, goNamePrefix, baseClass, goType string
}{
	{"DATE", "Date", "EMPTY", "string"},
	{"TIME", "Time", "EMPTY", "string"},
	{"TIMESTAMP", "TimeStamp", "EMPTY", "string"},
	{"DATETIME", "DateTime", "EMPTY", "string"},
	{"YEAR", "Year", "EMPTY", "string"},
	{"TINYBLOB", "TinyBlob", "EMPTY", "[]byte"},
	{"BLOB", "Blob", "EMPTY", "[]byte"},
	{"MEDIUMBLOB", "MediumBlob", "EMPTY", "[]byte"},
	{"LONGBLOB", "LongBlob", "EMPTY", "[]byte"},
	{"BOOLEAN", "Boolean", "EMPTY", "bool"},
	{"TINYINT", "TinyInt", "INT", "int8"},
	{"SMALLINT", "SmallInt", "INT", "int16"},
	{"MEDIUMINT", "MediumInt", "INT", "int32"},
	{"INT", "Int", "INT", "int32"},
	{"BIGINT", "BigInt", "INT", "int64"},
	{"TINYINT", "TinyUint", "UINT", "uint8"},
	{"SMALLINT", "SmallUint", "UINT", "uint16"},
	{"MEDIUMINT", "MediumUint", "UINT", "uint32"},
	{"INT", "Uint", "UINT", "int32"},
	{"BIGINT", "BigUint", "UINT", "uint64"},
	{"REAL", "Real", "FLOAT", "float64"},
	{"FLOAT", "Float", "FLOAT", "float64"},
	{"DECIMAL", "Decimal", "FLOAT", "string"},
	{"NUMERIC", "Numeric", "FLOAT", "string"},
	{"BIT", "Bit", "BINARY", "string"},
	{"BINARY", "Binary", "BINARY", "[]byte"},
	{"VARBINARY", "VarBinary", "BINARY", "[]byte"},
	{"CHAR", "Char", "CHAR", "string"},
	{"VARCHAR", "VarChar", "CHAR", "string"},
	{"TINYTEXT", "TinyText", "TEXT", "string"},
	{"TEXT", "Text", "TEXT", "string"},
	{"MEDIUMTEXT", "MediumText", "TEXT", "string"},
	{"LONGTEXT", "LongText", "TEXT", "string"},
	// ToDo:
	//{"ENUM", "Enum", "ENUM", ""},
	//{"SET", "Set", "ENUM", ""},
}

type IBaseClass interface {
	Fields() []IField
}

type IField interface {
	Name() string
	GoType() string
}

/*
type EmptyClass struct{}

func (EmptyClass) Fields() []IField { return []IField{} }

type IntClass struct{}

func (EmptyClass) Fields() []IField { return []IField{} }
*/

var baseClasses = map[string]struct {
	fields []string
}{
	"EMPTY": {
		fields: []string{},
	},
	"INT": {
		fields: []string{"Length", "Zerofill"},
	},
	"UINT": {
		fields: []string{"Length", "Zerofill"},
	},
	"FLOAT": {
		fields: []string{"Length", "Decimals", "Zerofill"},
	},
	"BINARY": {
		fields: []string{"Length"},
	},
	"CHAR": {
		fields: []string{"Length", "Charset", "Collate"},
	},
	"TEXT": {
		fields: []string{"Length", "Charset", "Collate", "Binary"},
	},
}

func main() {
	filename := flag.String("filename", "", "Output filename")
	flag.Parse()

	if *filename == "" {
		fmt.Println("No filename defined")
		os.Exit(1)
	}

	file, err := os.Create(*filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	buf := &bytes.Buffer{}

	buf.WriteString("// The file was generated, do not change it manually\n\n")
	buf.WriteString("package mysql")

	for _, mysqlType := range mysqlTypes {
		buf.WriteString("\n\ntype ")
		buf.WriteString(mysqlType.goNamePrefix)
		buf.WriteString("Field struct{\n")
		buf.WriteString("Id string\n")
		buf.WriteString("Caption string\n")

		for _, field := range baseClasses[mysqlType.baseClass].fields {
			buf.WriteString(field)
			buf.WriteString(" bool\n")
		}

		buf.WriteString("NotNull bool\n")
		buf.WriteString("Default *" + mysqlType.goType + "\n")
		buf.WriteString("CheckFunc func(" + mysqlType.goType + ") error\n")
		buf.WriteString("}\n")
	}

	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	file.Write(source)
}
