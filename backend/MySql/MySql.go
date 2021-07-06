package MySql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

type info struct {
	Uid      int
	Uname    string
	Ucontent string
}

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

func SqlSelect() (list []info) {
	//查询数据
	rows, err := db.Query("SELECT uid,uname,ucontent FROM commentinfo")
	checkErr(err)

	var uid int
	var uname string
	var ucontent string

	println("-----------")
	for rows.Next() {
		err = rows.Scan(&uid, &uname, &ucontent)
		checkErr(err)
		var row info
		row.Uid = uid
		row.Uname = uname
		row.Ucontent = ucontent
		list = append(list, row)
	}
	return
}

func SqlInsert(uname string, ucontent string) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO commentinfo(uname,ucontent) VALUES($1,$2) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec(uname, ucontent)
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

	// SqlInsert()
	// SqlSelect()
	// println(sep, "*sqlInsert")

	SqlClose()
	println(sep, "*sqlClose")
}
