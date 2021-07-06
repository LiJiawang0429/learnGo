package Sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SqlOpen() {
	var err error
	//打开数据库
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456789++ dbname=postgres sslmode=disable")

	checkErr(err)
}

func SqlSelect()(uid int,uname string,ucontent string) {

	//查询数据
	rows, err := db.Query("SELECT * FROM commentinfo where uid = 1")
	checkErr(err)

	println("-----------")
	fmt.Println(rows)
	fmt.Printf("rowstype%T\n", rows)
	for rows.Next() {
		err = rows.Scan(&uid, &uname, &ucontent)
		checkErr(err)
		fmt.Println("uid = ", uid, "\nuname = ", uname, "\nucon = ", ucontent, "\n-----------")
	}
	return uid,uname,ucontent
}

func SqlInsert() {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO commentinfo(uname,ucontent) VALUES($1,$2) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("小张", "加油加油")
	//参数对应上面的$1,$2

	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func SqlClose() {
	db.Close()
}

func SqlTest() {

	sep := "----------\n"
	SqlOpen()
	println(sep, "*sqlOpen")

	SqlInsert()
	SqlSelect()
	println(sep, "*sqlInsert")

	SqlClose()
	println(sep, "*sqlClose")
}
