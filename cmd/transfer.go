package cmd

type ColumnType struct {
	TransferType   string
	TransferInsert func(string) string
}

var sql2goType = map[string]ColumnType{
	"tinyint": {
		TransferType: "int",
	},
	"smallint": {
		TransferType: "int",
	},
	"mediumint": {
		TransferType: "int",
	},
	"int": {
		TransferType: "int",
	},
	"integer": {
		TransferType: "int",
	},
	"bigint": {
		TransferType: "int64",
	},
	"float": {
		TransferType: "float64",
	},
	"double": {
		TransferType: "float64",
	},
	"decimal": {
		TransferType: "float64",
	},
	"date": {
		TransferType: "time.Time",
	},
	"time": {
		TransferType: "string",
	},
	"year": {
		TransferType: "int",
	},
	"datetime": {
		TransferType: "time.Time",
	},
	"timestamp": {
		TransferType: "int",
	},
	"datetimeoffset": {
		TransferType: "datetime",
	},
	"char": {
		TransferType: "string",
	},
	"varchar": {
		TransferType: "string",
	},
	"tinyblob": {
		TransferType: "string",
	},
	"tinytext": {
		TransferType: "string",
	},
	"blob": {
		TransferType: "string",
	},
	"text": {
		TransferType: "string",
	},
	"mediumblob": {
		TransferType: "string",
	},
	"mediumtext": {
		TransferType: "string",
	},
	"longblob": {
		TransferType: "string",
	},
	"longtext": {
		TransferType: "string",
	},
}

var sql2tsType = map[string]ColumnType{
	"tinyint": {
		TransferType: "number",
	},
	"smallint": {
		TransferType: "number",
	},
	"mediumint": {
		TransferType: "number",
	},
	"int": {
		TransferType: "number",
	},
	"integer": {
		TransferType: "number",
	},
	"bigint": {
		TransferType: "number",
	},
	"float": {
		TransferType: "number",
	},
	"double": {
		TransferType: "number",
	},
	"decimal": {
		TransferType: "number",
	},
	"date": {
		TransferType: "string",
	},
	"time": {
		TransferType: "string",
	},
	"year": {
		TransferType: "number",
	},
	"datetime": {
		TransferType: "string",
	},
	"timestamp": {
		TransferType: "number",
	},
	"datetimeoffset": {
		TransferType: "string",
	},
	"char": {
		TransferType: "string",
	},
	"varchar": {
		TransferType: "string",
	},
	"tinyblob": {
		TransferType: "string",
	},
	"tinytext": {
		TransferType: "string",
	},
	"blob": {
		TransferType: "string",
	},
	"text": {
		TransferType: "string",
	},
	"mediumblob": {
		TransferType: "string",
	},
	"mediumtext": {
		TransferType: "string",
	},
	"longblob": {
		TransferType: "string",
	},
	"longtext": {
		TransferType: "string",
	},
}

var existModel = map[string]struct{}{
	"id":         {},
	"created_at": {},
	"updated_at": {},
	"deleted_at": {},
	"is_delete":  {},
	"version":    {},
}

func isNumber(typeName string) bool {
	switch typeName {
	case "int", "int64", "float64":
		return true
	}
	return false
}

var defaultWorkModel = map[string]struct{}{
	"id":          {},
	"is_delete":   {},
	"create_by":   {},
	"modify_by":   {},
	"version":     {},
	"create_time": {},
	"modify_time": {},
}
