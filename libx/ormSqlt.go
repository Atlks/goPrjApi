package libx

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// save 函数
func Save(session map[string]interface{}, dbFileName string) {
	tblx := "表格1"

	// 打印调试信息
	fmt.Printf("Entering function Save with table: %s, dbFileName: %s\n", tblx, dbFileName)
	fmt.Printf("Session data: %v\n", session)

	// 构建SQL语句
	sqlStmt := fmt.Sprintf("REPLACE INTO %s (%s) VALUES (%s)", tblx, keysToString(session), valueSqlFmt(session))
	fmt.Printf("SQL Statement: %s\n", sqlStmt)

	// 打开数据库连接
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// 执行SQL命令
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Println("Error executing SQL:", err)
		return
	}

	fmt.Println("Data saved successfully")
}

// 将 map 的键转换为字符串，格式为 "key1, key2, ..."
func keysToString(session map[string]interface{}) string {
	keys := make([]string, 0, len(session))
	for k := range session {
		keys = append(keys, k)
	}
	// 将字符串数组转换为用逗号分隔的字符串
	result := strings.Join(keys, ",")
	//sprintf := fmt.Sprintf("%s", keys)
	return result
}

func valueSqlFmt(sortedMap map[string]interface{}) string {
	var escapedValues []string

	for _, value := range sortedMap {
		if value == nil {
			escapedValues = append(escapedValues, "null")
		} else {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.String:
				escapedValues = append(escapedValues, fmt.Sprintf("'%s'", value))
			default:
				escapedValues = append(escapedValues, fmt.Sprintf("'%v'", value))
			}
		}
	}

	return strings.Join(escapedValues, ",")
}

// 将 map 的值转换为字符串，格式为 ":value1, :value2, ..."
func valuesToString(session map[string]interface{}) string {
	values := make([]string, 0, len(session))
	for _, v := range session {
		values = append(values, fmt.Sprintf(":%v", v))
	}
	// 将字符串数组转换为用逗号分隔的字符串
	result := strings.Join(values, ",")
	return result
	//return fmt.Sprintf("%s", values)
}
func crtTable(tabl string, sortedList1 map[string]interface{}, dbFileName string) {
	// Debug information
	fmt.Printf("Entering function crtTable with table: %s, dbFileName: %s\n", tabl, dbFileName)
	fmt.Printf("Session data: %v\n", sortedList1)

	// Open database connection
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create table
	sqlStmt := fmt.Sprintf("CREATE TABLE %s (id TEXT PRIMARY KEY)", tabl)
	executeNonQuery(db, sqlStmt)

	// Type map
	typeMapPHP2sqlt := map[string]string{
		"integer": "int",
		"string":  "text",
	}

	for k, v := range sortedList1 {
		if k == "id" {
			continue
		}

		sqltType, ok := typeMapPHP2sqlt[fmt.Sprintf("%T", v)]
		if !ok {
			sqltType = "text"
		}

		sqlStmt = fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tabl, k, sqltType)
		executeNonQuery(db, sqlStmt)
	}
}

func executeNonQuery(db *sql.DB, sqlStmt string) {
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		fmt.Println("Error preparing SQL statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error executing SQL statement:", err)
		return
	}
}

// 将 map 转换为 SQL 插入语句的参数
func ArrToSqlPrms4insert(sortedList map[string]interface{}) string {
	var columns []string
	var values []string

	// 遍历 sortedList 的键和值
	for key, value := range sortedList {
		columns = append(columns, key)

		if value == nil {
			values = append(values, "null")
		} else {
			values = append(values, fmt.Sprintf("'%v'", value))
		}
	}

	// 将列和值转换为字符串
	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")

	return fmt.Sprintf("(%s) values (%s)", columnsStr, valuesStr)
}

//var dbgpad int

func setDbgFunEnter(method string, funcGetArgs interface{}) {
	dbgpad += 4

	funcGetArgsJSON, err := json.Marshal(funcGetArgs)
	if err != nil {
		fmt.Println("Error serializing arguments:", err)
		return
	}

	msglog := fmt.Sprintf("\n\n\n%s FUN %s((%s))", strings.Repeat(" ", dbgpad), method, string(funcGetArgsJSON))
	fmt.Println(msglog)
}

func main() {
	// Example usage of the save function
	//sortedList1 := NewSortedList()
	//sortedList1.Add("column1", "value1")
	//sortedList1.Add("column2", "value2")
	//dbFileName := "example.db"
	//save(sortedList1, dbFileName)
}
