package lib

//
//
//
//
//global['pdo_queryV5'] = pdo_queryV5
//
///**
// *
// * @param qryDsl
// * @param collName
// * @param storedir
// * @returns {*[]}
// */
//function pdo_queryV5(qryDsl, collName,storedir) {
//    log_enterFun(arguments)
//    // var  _ = await import("lodash")  not err,but cant use _.flt said cant find this fun
//    // import  lds = require('lodash') ;
//    console.log("::23")
//    let data;
//    var {readFileSync, writeFileSync, appendFileSync} = require("fs")
//    let rzt_all = [];
//    let ptns = ptn_getPartnsV5(collName,storedir)
//    require("./arr")
//    for (let pt of ptns) {
//        let dbfile = pt.dbf
//        var txt = readFileSync(dbfile).toString();
//        //  console.log(" dbtxt len100 =>" + txt.substring(0, 100))
//        data = JSON.parse(txt)
//        require("esm-hook");
//        //  const _ = require('lodash').default
//        const _ = require('lodash')
//        let rzt = _.filter(data, qryDsl)
//        rzt_all = array_merge(rzt_all, rzt)
//    }
//    console.log("[pdo_query]rzt is=>" + JSON.stringify(rzt_all).substring(0, 300))
//    return rzt_all;
//}
//
//function  ptn_getPtnByName( name,storedir)
//{
//    let dbfile = storedir+name;
//
//    return {"ptn": name, "dbf": dbfile};
//
//}
//
//global['pdo_saveV5']=pdo_saveV5
//
///**
// *
// * @param data2
// * @param collName
// * @param storedir
// */
//function pdo_saveV5(data2, collName,storedir) {
//    let qryDsl = {"id": data2.id}
//    let rowsAll;
//    const {readFileSync, writeFileSync, appendFileSync} = require("fs");
//    let rzt_all = [];
//
//    let ptn=ptn_getPtnByName(data2._ptn,storedir)
//    let dbfile = ptn.dbf
//    let txt = readFileSync(dbfile).toString();
//    //  console.log(" dbtxt len100 =>" + txt.substring(0, 100))
//    rowsAll = JSON.parse(txt)
//
//
//    require("./sys")
//    //定位到所在rcd，更新数据，save
//    require("esm-hook");
//    //  const _ = require('lodash').default
//    const _ = require('lodash')
//    let rows = _.filter(rowsAll, qryDsl)
//    if (rows) {
//
//        let obj = rows[0]
//        copyProp(data2, obj)
//        require("./enc")
//        writeFileSyncx(dbfile, json_encode(rowsAll));
//        return
//    }
//
//
//}
//
//
//
///**
// *
// * @param param
// * @param collName
// * @param storedir
// * @returns {*|{data: *, dbf, write: write}}
// */
//function pdo_insertV5(rcd, collName,storedir) {
//
//    log_enterFun(arguments)
//
//    let lastPrtn = ptn_getLastPtn(collName,storedir)
//    let db_conn_dt = pdo_connV3(lastPrtn.dbf)
//
//    if (typeof rcd == "string")
//        rcd = JSON.parse(rcd)
//
//    rcd._ptn = lastPrtn.ptn
//    db_conn_dt.push(rcd)
//
//    pdo_save(db_conn_dt, lastPrtn.dbf)
//    return db_conn_dt;
//}
//
//
//global['pdo_insertV5'] = pdo_insertV5
//
//
///**
// *
// * @param uid123Coll99
// * @param storedir
// * @returns {*}
// */
//function ptn_getLastPtn(uid123Coll99, storedir) {
//
//    let partList = ptn_getPartnsV5(uid123Coll99,storedir)
//
//    let lastptn = partList[partList.length - 1];
//
// require("./file")
//    let dbfile = lastptn.dbf
//    if (filesize(dbfile) > 1024 * 1024) {
//        // if fi如果lesieze > 1M 物 ，new file
//
//        lastptn = ptn_newprtnV5(uid123Coll99,storedir)
//    }
//    return lastptn;
//
//}
//
//
//function ptn_newprtnV5(collName,dbdir) {
//
//    let ptnname = collName + "_" + idBasetime() + ".json";
//    let dbfile = dbdir + ptnname
//
//    return {"ptn": ptnname, "dbf": dbfile};
//}
//
//
//function ptn_getPartnsV5(collName, storedir) {
//    let dbdir =storedir
//    var {readFileSync, writeFileSync, appendFileSync} = require("fs");
//    const fs = require("fs");
//
//    fs.mkdirSync(dbdir, {recursive: true});
//    let partList = fs.readdirSync(dbdir);
//    partList = partList.filter(f => {
//        return f.startsWith(collName)
//    })
//    if (partList.length == 0) {
//        let dbfile = dbdir + collName + ".json"
//        writeFileSync(dbfile, "[]")
//        partList = [collName + ".json"]
//    }
//    let arr = []
//    for (let ptnname of partList) {
//        let dbfile = dbdir + ptnname
//        let ptn = {"ptn": ptnname, "dbf": dbfile};
//        arr.push(ptn)
//    }
//    return arr;
//}
