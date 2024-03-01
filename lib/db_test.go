package lib

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestMysqlInst(t *testing.T) {
	//	rcd:={"id":1};

	db, err := sql.Open("mysql", "root:123456@tcp(192.168.128.173:3306)/jianbanbw")
	CheckErr(err)

	res, err := db.Exec("create database1 db3")

	//	CheckErr(err)
	//查询删除多少条信息
	num, err := res.RowsAffected()
	CheckErr(err)
	fmt.Println(num) //1  if creted database ok ..

	//关闭数据库连接
	db.Close()

}

func TestHandler100(t *testing.T) {
	//	rcd:={"id":1};

	rcd := map[string]any{
		"id":   1,
		"name": 2,
	}

	Pdo_Insert(rcd, "coll33", "c:/dbx/")
}

func TestHandler_qry22(t *testing.T) {

	rows := Pdo_qry("c:/dbx/", "coll33", func(v any) bool {

		map1 := v.(map[string]any)
		map1 = AnyToMap(v)
		if map1["id"] == nil {
			return false
		}
		if AnyToNumFlt64(map1["id"]) == 2 {
			return true
		} else {
			return false
		}

	})

	//print(rows)
	fmt.Printf("%v\n", rows)
	fmt.Printf("%+v\n", rows)

}

func TestHandler_qry(t *testing.T) {
	//	rcd:={"id":1};

	rcd := map[string]any{
		"id":   1,
		"name": "name111",
	}

	rcd2 := map[string]any{
		"id":   2,
		"name": "name222",
	}

	var arr []any

	//追加一个元素
	arr = append(arr, rcd, rcd2)

	rows := Filter(arr,
		func(v any) bool {

			map1 := v.(map[string]any)
			map1 = AnyToMap(v)
			if map1["id"] == 1 {
				return true
			} else {
				return false
			}

		},
	)

	print(rows)
	//Pdo_Insert(rcd, "coll33", "c:/dbx/")
}
