package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func Pdo_Insert(rcd map[string]any, collName string, storedir string) any {

	// log_enterFun(arguments)

	dbdir := storedir + collName
	//递归创建文件夹
	os.MkdirAll(dbdir, os.ModePerm)
	rowid := CURRENT_TIME_forID()
	rcd["_fil"] = rowid
	str := Json_encode(rcd)

	//rowid := rand.Intn(999999999)
	//strconv.Itoa(rowid)

	dbf := dbdir + "/" + rowid + ".json"

	Write(dbf, str)

	return rcd
}

func Pdo_qry(storedir string, collName string, f func(any) bool) any {

	var arr []any

	dbdir := storedir + collName
	err := filepath.Walk(dbdir, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		conStr := ReadToStr(path)
		cont_map := Json_decode(conStr)
		//	if
		//追加一个元素
		arr = append(arr, cont_map)
		return nil
	})
	fmt.Println(err)
	rows := Filter(arr, f)

	return rows
}

func pdo_queryV5(qryDsl any, collName string, storedir string) {
	//log_enterFun(arguments)
	//// var  _ = await import("lodash")  not err,but cant use _.flt said cant find this fun
	//// import  lds = require('lodash') ;
	//console.log("::23")
	//let data;
	//var {readFileSync, writeFileSync, appendFileSync} = require("fs")
	//let rzt_all = [];
	//let ptns = ptn_getPartnsV5(collName,storedir)
	//require("./arr")
	//for (let pt of ptns) {
	//let dbfile = pt.dbf
	//var txt = readFileSync(dbfile).toString();
	////  console.log(" dbtxt len100 =>" + txt.substring(0, 100))
	//data = JSON.parse(txt)
	//require("esm-hook");
	////  const _ = require('lodash').default
	//const _ = require('lodash')
	//let rzt = _.filter(data, qryDsl)
	//rzt_all = array_merge(rzt_all, rzt)
	//}
	//console.log("[pdo_query]rzt is=>" + JSON.stringify(rzt_all).substring(0, 300))
	//return rzt_all;
}
