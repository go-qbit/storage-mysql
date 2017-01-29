package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
)

var mysqlTypes = []struct {
	mysqlType    string
	goNamePrefix string
	baseClass    IBaseClass
	goType       string
}{
	{"DATE", "Date", EmptyClass{}, "string"},
	{"TIME", "Time", EmptyClass{}, "string"},
	{"TIMESTAMP", "TimeStamp", EmptyClass{}, "string"},
	{"DATETIME", "DateTime", EmptyClass{}, "string"},
	{"YEAR", "Year", EmptyClass{}, "string"},
	{"TINYBLOB", "TinyBlob", EmptyClass{}, "[]byte"},
	{"BLOB", "Blob", EmptyClass{}, "[]byte"},
	{"MEDIUMBLOB", "MediumBlob", EmptyClass{}, "[]byte"},
	{"LONGBLOB", "LongBlob", EmptyClass{}, "[]byte"},
	{"BOOLEAN", "Boolean", EmptyClass{}, "bool"},
	{"TINYINT", "TinyInt", IntClass{}, "int8"},
	{"SMALLINT", "SmallInt", IntClass{}, "int16"},
	{"MEDIUMINT", "MediumInt", IntClass{}, "int32"},
	{"INT", "Int", IntClass{}, "int32"},
	{"BIGINT", "BigInt", IntClass{}, "int64"},
	{"TINYINT", "TinyUint", UintClass{}, "uint8"},
	{"SMALLINT", "SmallUint", UintClass{}, "uint16"},
	{"MEDIUMINT", "MediumUint", UintClass{}, "uint32"},
	{"INT", "Uint", UintClass{}, "uint32"},
	{"BIGINT", "BigUint", UintClass{}, "uint64"},
	{"REAL", "Real", FloatClass{}, "float64"},
	{"FLOAT", "Float", FloatClass{}, "float64"},
	{"DECIMAL", "Decimal", FloatClass{}, "string"},
	{"NUMERIC", "Numeric", FloatClass{}, "string"},
	{"BIT", "Bit", BinaryClass{}, "string"},
	{"BINARY", "Binary", BinaryClass{}, "[]byte"},
	{"VARBINARY", "VarBinary", BinaryClass{}, "[]byte"},
	{"CHAR", "Char", CharClass{}, "string"},
	{"VARCHAR", "VarChar", CharClass{}, "string"},
	{"TINYTEXT", "TinyText", TextClass{}, "string"},
	{"TEXT", "Text", TextClass{}, "string"},
	{"MEDIUMTEXT", "MediumText", TextClass{}, "string"},
	{"LONGTEXT", "LongText", TextClass{}, "string"},
	// ToDo:
	//{"ENUM", "Enum", "ENUM", ""},
	//{"SET", "Set", "ENUM", ""},
}

type IBaseClass interface {
	Fields() []IField
	IsUnsigned() bool
}

type IField interface {
	Name() string
	GoType() string
}

// Base classes
type EmptyClass struct{}

func (EmptyClass) Fields() []IField { return []IField{} }
func (EmptyClass) IsUnsigned() bool { return false }

type IntClass struct{}

func (IntClass) Fields() []IField {
	return []IField{LengthField{}, AutoIncrementField{}, ZerofillField{}}
}
func (IntClass) IsUnsigned() bool { return false }

type UintClass struct{}

func (UintClass) Fields() []IField {
	return []IField{LengthField{}, AutoIncrementField{}, ZerofillField{}}
}
func (UintClass) IsUnsigned() bool { return true }

type FloatClass struct{}

func (FloatClass) Fields() []IField { return []IField{LengthField{}, DecimalsField{}, ZerofillField{}} }
func (FloatClass) IsUnsigned() bool { return false }

type BinaryClass struct{}

func (BinaryClass) Fields() []IField { return []IField{LengthField{}} }
func (BinaryClass) IsUnsigned() bool { return false }

type CharClass struct{}

func (CharClass) Fields() []IField { return []IField{LengthField{}, CharsetField{}, CollateField{}} }
func (CharClass) IsUnsigned() bool { return false }

type TextClass struct{}

func (TextClass) Fields() []IField {
	return []IField{LengthField{}, CharsetField{}, CollateField{}, &BinaryField{}}
}
func (TextClass) IsUnsigned() bool { return false }

// Fields
type AutoIncrementField struct{}

func (AutoIncrementField) Name() string   { return "AutoIncrement" }
func (AutoIncrementField) GoType() string { return "bool" }

type LengthField struct{}

func (LengthField) Name() string   { return "Length" }
func (LengthField) GoType() string { return "int" }

type ZerofillField struct{}

func (ZerofillField) Name() string   { return "Zerofill" }
func (ZerofillField) GoType() string { return "bool" }

type BinaryField struct{}

func (BinaryField) Name() string   { return "Binary" }
func (BinaryField) GoType() string { return "bool" }

type DecimalsField struct{}

func (DecimalsField) Name() string   { return "Decimals" }
func (DecimalsField) GoType() string { return "int" }

type CharsetField struct{}

func (CharsetField) Name() string   { return "Charset" }
func (CharsetField) GoType() string { return "string" }

type CollateField struct{}

func (CollateField) Name() string   { return "Collate" }
func (CollateField) GoType() string { return "string" }

//
var typeOf = map[string]string{
	"string":  `string("")`,
	"[]byte":  "[]byte{}",
	"int8":    "int8(0)",
	"int16":   "int16(0)",
	"int32":   "int32(0)",
	"int64":   "int64(0)",
	"uint8":   "uint8(0)",
	"uint16":  "uint16(0)",
	"uint32":  "uint32(0)",
	"uint64":  "uint64(0)",
	"float64": "float64(0)",
	"bool":    "bool(false)",
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

	buf.WriteString("// The file was generated, do not change by hands\n\n")
	buf.WriteString("package mysql\n")

	buf.WriteString("import (\n")
	buf.WriteString(`"reflect"` + "\n")
	buf.WriteString(`"strconv"` + "\n")
	buf.WriteByte('\n')
	buf.WriteString(`"github.com/go-qbit/model"` + "\n")
	buf.WriteString(")\n")

	for _, mysqlType := range mysqlTypes {
		typeName := mysqlType.goNamePrefix + "Field"
		buf.WriteString("\n\ntype " + typeName + " struct{\n")
		buf.WriteString("Id string\n")
		buf.WriteString("Caption string\n")

		for _, field := range mysqlType.baseClass.Fields() {
			buf.WriteString(field.Name() + " " + field.GoType() + "\n")
		}

		buf.WriteString("NotNull bool\n")
		buf.WriteString("Default *" + mysqlType.goType + "\n")
		buf.WriteString("CheckFunc func(" + mysqlType.goType + ") error\n")
		buf.WriteString("CleanFunc func(" + mysqlType.goType + ") (" + mysqlType.goType + ", error)\n")
		buf.WriteString("}\n")

		buf.WriteString("func (f *" + typeName + ") GetId() string { return f.Id }\n")
		buf.WriteString("func (f *" + typeName + ") GetCaption() string  { return f.Caption }\n")
		buf.WriteString("func (f *" + typeName + ") GetType() reflect.Type { return reflect.TypeOf(" + typeOf[mysqlType.goType] + ") }\n")
		buf.WriteString("func (f *" + typeName + ") IsDerivable() bool { return false }\n")
		buf.WriteString("func (f *" + typeName + ") IsRequired() bool { return f.NotNull && f.Default == nil && !f.IsAutoIncremented() }\n")
		buf.WriteString("func (f *" + typeName + ") GetDependsOn() []string { return nil }\n")
		buf.WriteString("func (f *" + typeName + ") Calc(map[string]interface{}) (interface{}, error) { return nil, nil }\n")
		buf.WriteString("func (f *" + typeName + ") Check(v interface{}) error {\n" +
			"	if f.CheckFunc != nil {\n" +
			"		return f.CheckFunc(v.(" + mysqlType.goType + "))\n" +
			"	} else {\n" +
			"		return nil\n" +
			"	}\n" +
			"}\n")
		buf.WriteString("func (f *" + typeName + ") Clean(v interface{}) (interface{}, error) {\n" +
			"	if f.CleanFunc != nil {\n" +
			"		return f.CleanFunc(v.(" + mysqlType.goType + "))\n" +
			"	} else {\n" +
			"		return v, nil\n" +
			"	}\n" +
			"}\n")

		var cloneFields string
		for _, field := range mysqlType.baseClass.Fields() {
			if field.Name() == "AutoIncrement" {
				cloneFields += "false,"
			} else {
				cloneFields += "f." + field.Name() + ","
			}
		}

		buf.WriteString("func (f *" + typeName + ") CloneForFK(id string, caption string, required bool) model.IFieldDefinition {\n" +
			"return &" + typeName + "{id, caption, " + cloneFields + " required, f.Default, f.CheckFunc, f.CleanFunc}\n" +
			"}\n")

		typeFields := make(map[string]struct{})
		for _, ftype := range mysqlType.baseClass.Fields() {
			typeFields[ftype.Name()] = struct{}{}
		}

		isAutoincremented := "false"
		if _, exists := typeFields["AutoIncrement"]; exists {
			isAutoincremented = "true"
		}
		buf.WriteString("func (f *" + typeName + ") IsAutoIncremented() bool { return " + isAutoincremented + " }\n")

		buf.WriteString("" +
			"func (f *" + typeName + ") WriteSQL(sqlBuf *SqlBuffer) {\n" +
			"	sqlBuf.WriteIdentifier(f.Id)\n\n" +
			"	sqlBuf.WriteByte(' ')\n" +
			"	sqlBuf.WriteString(\"" + mysqlType.mysqlType + "\")\n\n")

		if _, exists := typeFields["Length"]; exists {
			buf.WriteString("" +
				"	if f.Length != 0 {\n" +
				"		sqlBuf.WriteByte('(')\n" +
				"		sqlBuf.WriteString(strconv.Itoa(f.Length))\n")
			if _, exists := typeFields["Decimals"]; exists {
				buf.WriteString("" +
					"	if f.Decimals != 0 {\n" +
					"		sqlBuf.WriteByte(',')\n" +
					"		sqlBuf.WriteString(strconv.Itoa(f.Decimals))\n" +
					"	}\n\n")
			}
			buf.WriteString("		sqlBuf.WriteByte(')')\n" +
				"	}\n\n")
		}

		if mysqlType.baseClass.IsUnsigned() {
			buf.WriteString("	sqlBuf.WriteString(\" UNSIGNED\")\n\n")
		}

		if _, exists := typeFields["Zerofill"]; exists {
			buf.WriteString("" +
				"	if f.Zerofill {\n" +
				"		sqlBuf.WriteString(\" ZEROFILL\")\n" +
				"	}\n\n")
		}

		if _, exists := typeFields["Binary"]; exists {
			buf.WriteString("" +
				"	if f.Binary {\n" +
				"		sqlBuf.WriteString(\" BINARY\")\n" +
				"	}\n\n")
		}

		if _, exists := typeFields["Charset"]; exists {
			buf.WriteString("" +
				"	if f.Charset != \"\" {\n" +
				"		sqlBuf.WriteString(\" CHARACTER SET \")\n" +
				"		sqlBuf.WriteValue(f.Charset)\n" +
				"	}\n\n")
		}

		if _, exists := typeFields["Collate"]; exists {
			buf.WriteString("" +
				"	if f.Collate != \"\" {\n" +
				"		sqlBuf.WriteString(\" COLLATE \")\n" +
				"		sqlBuf.WriteValue(f.Collate)\n" +
				"	}\n\n")
		}

		buf.WriteString("" +
			"	if f.NotNull {\n" +
			"		sqlBuf.WriteString(\" NOT NULL\")\n" +
			"	}\n\n")

		if _, exists := typeFields["AutoIncrement"]; exists {
			buf.WriteString("" +
				"	if f.AutoIncrement {\n" +
				"		sqlBuf.WriteString(\" AUTO_INCREMENT\")\n" +
				"	}\n\n")
		}

		buf.WriteString("" +
			"	if f.Default != nil {\n" +
			"		sqlBuf.WriteString(\" DEFAULT \")\n" +
			"		sqlBuf.WriteValue(*f.Default)\n" +
			"	}\n\n")

		buf.WriteString("}\n")
	}

	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(buf.String())
		os.Exit(1)
	}

	file.Write(source)
}
