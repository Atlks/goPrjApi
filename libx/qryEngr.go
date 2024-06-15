package libx

import (
	"path/filepath"
	"strings"
	"time"
)

// save24614 saves the sorted list and returns an integer based on the provided function.
func Qe_save(sortedListNew map[string]interface{}, dataDir string, callFunStrEngr func(map[string]interface{}) int, dbg map[string]interface{}) int {
	timestamp := time.Now().Format("20060102_150405_000")
	if _, exists := sortedListNew["id"]; !exists {
		sortedListNew["id"] = timestamp
	}
	if sortedListNew["id"].(string) == "" {
		sortedListNew["id"] = timestamp
	}

	mereed := sortedListNew
	str := callFunStrEngr(mereed)
	return str
}

// qry queries and returns a list of sorted lists based on the provided functions.
func Qe_qry(fromDdataDir, partnsExprs string, whereFun func(map[string]interface{}) bool, cfgStrEngr func(string) []map[string]interface{}) []map[string]interface{} {
	if cfgStrEngr == nil {
		panic("cfgStrEngr cannot be null")
	}

	rztLi := []map[string]interface{}{}
	patnsDbfs := _calcPatnsV3(fromDdataDir, partnsExprs)
	arr := strings.Split(patnsDbfs, ",")
	for _, dbf := range arr {
		li := _qryBySnglePart(dbf, whereFun, cfgStrEngr)
		rztLi = arrayMerge(rztLi, li)
	}

	return rztLi
}

// _qryBySnglePart is a placeholder for the function that queries a single part.
func _qryBySnglePart(dbf string, whereFun func(map[string]interface{}) bool, cfgStrEngr func(string) []map[string]interface{}) []map[string]interface{} {
	// Implement the function according to your needs
	return []map[string]interface{}{}
}

// arrayMerge merges two slices of map[string]interface{}
//
//	func arrayMerge(arr1, arr2 []map[string]interface{}) []map[string]interface{} {
//		return append(arr1, arr2...)
//	}
func _calcPatnsV3(dir string, partfile string) string {
	method := "calcPatnsV3"

	setDbgFunEnter(method, Func_get_args())

	if partfile == "" {
		rzt := getFilePathsCommaSeparated(dir)
		setDbgValRtval(method, rzt)
		return rzt
	}

	var arrayList []string
	dbArr := strings.Split(partfile, ",")
	for _, dbf := range dbArr {
		path := filepath.Join(dir, dbf)
		arrayList = append(arrayList, path)
	}

	result := strings.Join(arrayList, ",")

	setDbgValRtval(method, result)

	return result
}

func getFilePathsCommaSeparated(dir string) string {
	// Implement this function according to your requirements
	return ""
}

//func _qryBySnglePart(dbf string, whereFun func(map[string]interface{}) bool, cfgStrEngr func(string) []map[string]interface{}) []map[string]interface{} {
//	li := cfgStrEngr(dbf)
//	// Assuming qryV7 is a function that queries the data based on whereFun
//	li = qryV7(li, whereFun)
//	return li
//}
