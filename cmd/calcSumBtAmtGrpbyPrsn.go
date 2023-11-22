package main

import (
	"fmt"
	"goapiPrj/lib"
	"log"
)

func main() {

	fmt.Println("hello22233")
	f := "C:\\modyfing\\apiprj\\jbbot\\zmng\\db\\111523\\userColl.json"
	jsonRows := lib.ReadToJsonArr(f)
	idx := lib.RdmMax(len(jsonRows))

	fmt.Println(idx)
	idx = 128
	row := jsonRows[idx]

	uid := row["userid"]

	uidStr := fmt.Sprintf("%.0f", uid)

	fmt.Println("uid==>" + uidStr)
	rzt := sumAllbet(uidStr)
	fmt.Println("sumbet==>" + fmt.Sprintf("%.0f", rzt))
	lib.WriteJsonArr("ucol.json", jsonRows)

}

func sumAllbet(userid string) float64 {
	file := "C:\\modyfing\\apiprj\\jbbot\\zmng\\db_zhudan\\zhudan_uid" + userid + ".json"
	data_rows := lib.ReadToJsonArr(file)

	modifiedArr := lib.Array_column(data_rows, "ValidBet")

	arraySum, err := lib.ArraySum(modifiedArr)
	if err != nil {
		log.Fatal("Error ==>", err)
	}
	return arraySum

}
