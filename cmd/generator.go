package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/didi/gendry/scanner"
)

var mysql *sql.DB

const (
	ADMIN = iota
	STRUCT
)

const daoTableFunc = `
	// TableName get table name
	func (b {{UpperTableName}}) TableName() string {
		return "{{TableName}}"
	}`

type columns struct {
	ColumnName string `ddb:"COLUMN_NAME"`
	DataType   string `ddb:"DATA_TYPE"`
	Remark     string `ddb:"COLUMN_COMMENT"`
	Default    string `ddb:"COLUMN_DEFAULT"`
	Key        string `ddb:"COLUMN_KEY"`
}

func genDAO(tableName string) (dom string) {
	c, err := generateStruct(tableName, STRUCT)
	if err != nil {
		return
	}
	tName := toCamel(tableName)
	c += daoTableFunc
	return strings.Replace(strings.Replace(c, "{{UpperTableName}}", tName, -1), "{{TableName}}", tableName, -1)
}

func generateStruct(tableName string, typeID int) (creates string, err error) {
	rows, err := mysql.Query(fmt.Sprintf(`
		SELECT * FROM information_schema.columns
		WHERE TABLE_SCHEMA = '%s'
		AND TABLE_NAME = '%s' 
	    ORDER BY ORDINAL_POSITION;
		`, dbname, tableName))
	if err != nil {
		return
	}
	var cs []columns
	err = scanner.ScanClose(rows, &cs)
	if err != nil {
		return
	}
	if len(cs) == 0 {
		err = errors.New("找不到数据")
		return
	}
	str := ""
	switch typeID {
	case STRUCT:
		str, err = goHandler(cs)
		if err != nil {
			return "", err
		}
	case ADMIN:
		str, err = apiHandler(cs)
		if err != nil {
			return "", err
		}
	}
	return str, err
}

func goHandler(cs []columns) (str string, err error) {
	col := make(map[string]string)
	for _, c := range cs {
		ct, ok := sql2goType[c.DataType]
		if !ok {
			fmt.Println(c)
			err = errors.New("暂不支持" + c.DataType)
			return
		}
		if _, ok = existModel[c.ColumnName]; ok {
			col["dao.Model"] = "dao.Model `gorm:\"embedded\"`"
		} else {
			defaultValue := "'%s'"
			if isNumber(ct.TransferType) {
				defaultValue = "%s"
			}
			col[c.ColumnName] = fmt.Sprintf("%s %s `json:\"%s\" gorm:\"default:"+defaultValue+"\"`",
				toCamel(c.ColumnName), ct.TransferType, c.ColumnName, c.Default)
			if c.Remark != "" {
				col[c.ColumnName] += " // " + c.Remark
			}
		}
	}
	var cols []string
	cols = append(cols, col["dao.Model"])
	for k, v := range col {
		if k == "dao.Model" {
			continue
		}
		cols = append(cols, v)
	}
	return fmt.Sprintf("package %sd\n\ntype {{UpperTableName}} struct{\n%s\n}",
		tags, strings.Join(cols, "\n")), err
}

func apiHandler(cs []columns) (str string, err error) {
	var col []string
	for _, c := range cs {
		ts, ok := sql2tsType[c.DataType]
		if !ok {
			//fmt.Println(c)
			err = errors.New("暂不支持" + c.DataType)
			return
		}
		fs := fmt.Sprintf("  %s: %s;", c.ColumnName, ts.TransferType)
		if c.Remark != "" {
			fs += " // " + c.Remark
		}
		col = append(col, fs)
	}
	return fmt.Sprintf("export interface {{UpperTableName}} {\n%s\n}\n", strings.Join(col, "\n")), err
}

func genAPIMethod(tablename string) (str string) {
	str = `import { defHttp } from '/@/utils/http/axios';
import type { RespList } from '/#/axios';

import { {{UpperTableName}} } from './model/{{TableName}}';

export function {{TableName}}s(params: any) {
  return defHttp.get<RespList<{{UpperTableName}}>>({
    url: '{{TableName}}s',
    params,
  });
}

export function {{TableName}}(id: string | number) {
  return defHttp.get<{{UpperTableName}}>({
    url: '{{TableName}}/' + id,
  });
}

export function {{TableName}}save(params: {{UpperTableName}}) {
  return defHttp.post<{{UpperTableName}}>({
    url: '{{TableName}}',
	params,
  });
}
`
	str = strings.Replace(str, "{{TableName}}", tablename, -1)
	str = strings.Replace(str, "{{UpperTableName}}", toCamel(tablename), -1)
	return
}

func genAPI(tableName string) (dom string) {
	c, err := generateStruct(tableName, ADMIN)
	if err != nil {
		return
	}
	tName := toCamel(tableName)
	return strings.Replace(c, "{{UpperTableName}}", tName, -1)
}
