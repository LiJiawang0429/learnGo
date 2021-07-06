package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type Person interface {
	SayHi() interface{} // 如果此方法返回值换成[]interface{} 会报错
}

type Man struct {
	Name string
}

func (m Man) SayHi() interface{} {
	return []byte("Hi, my name is: " + m.Name)
}

type Women struct {
	Name string
}

func (w Women) SayHi() interface{} {
	return []byte("Hi, my name is: " + w.Name)
}

func sayHi(p Person) {
	fmt.Println(string(p.SayHi().([]byte))) // 转换为[]byte类型
}

func main() {
	m := Man{Name: "Bruce"}
	w := Women{
		Name: "Anna",
	}

	sayHi(m)
	sayHi(w)
}

func executeSQL(queryStr string) []byte {
	connString := createConnectString()
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error while opening database connection:", err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(queryStr)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	count := len(columns)

	var v struct {
		Data []interface{} // `json:"data"`
	}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		//Created a map to handle the issue
		var m map[string]interface{}
		m = make(map[string]interface{})
		for i := range columns {
			m[columns[i]] = values[i]
		}
		v.Data = append(v.Data, m)
	}
	jsonMsg, err := json.Marshal(v)
	return jsonMsg
}
