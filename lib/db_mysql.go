package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接至数据库
	// sql.Open()中的数据库连接串格式为："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"。
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mysql")
	CheckErr(err)
	insert(db)
	//关闭数据库连接
	db.Close()
}

// 检查错误信息
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 插入demo
func insert(db *sql.DB) {
	//准备插入操作
	//
	//  db.
	res, err := db.Exec("create database db3")

	CheckErr(err)
	//查询删除多少条信息
	num, err := res.RowsAffected()
	CheckErr(err)
	fmt.Println(num) //1  if creted database ok ..

}

// 查询操作
func Query(db *sql.DB) {

	rows, e := db.Query("select name from mysql.help_topic where name >'a' limit 3  ")
	CheckErr(e)

	//cols, _ := rows.Columns()

	for rows.Next() {

		var name string
		rows.Scan(&name)
		fmt.Println(name)
		fmt.Println(777)
		//fmt.Println(json.Marshal()

	}

}
