package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "port=4432 user=postgres password=Oushu6@China dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Student struct {
	Id     int
	Name   string
	Gender string
	Age    int
	City   string
}

func main() {
	// 插入数据
	//stmt, err := Db.Prepare("INSERT INTO students (name, gender, age, city) VALUES ($1, $2, $3, $4);")
	//checkError(err)
	//res, err := stmt.Exec("zhangsan", "M", 18, "Beijing")
	//checkError(err)
	//lastInsertId, err := res.LastInsertId() // 最后插入数据的 id（PG 没实现）
	//rowsAffected, err := res.RowsAffected() // 影响行数
	//fmt.Printf("LastInsertId: %d\n", lastInsertId)
	//fmt.Printf("RowsAffected: %d\n", rowsAffected)
	//// 多次使用 stmt
	//res, err = stmt.Exec("lisi", "F", 20, "Shanghai")
	//checkError(err)
	//lastInsertId, err = res.LastInsertId() // 最后插入数据的 id（PG 没实现）
	//rowsAffected, err = res.RowsAffected() // 影响行数
	//fmt.Printf("LastInsertId: %d\n", lastInsertId)
	//fmt.Printf("RowsAffected: %d\n", rowsAffected)

	// 要想获取最后插入的数据的 id，可以在 SQL 语句中使用 returning 指定返回值
	//var lastInsertId int
	//row := Db.QueryRow("INSERT INTO students (name, gender, age, city) VALUES ($1, $2, $3, $4) RETURNING id;", "wangwu", "M", 35, "Jiangshu")
	//err := row.Scan(&lastInsertId)
	//defer rows.Close()
	//checkError(err)
	//fmt.Println("LastInsertId: ", lastInsertId)

	// 查询单个数据
	//stu := Student{}
	//row := Db.QueryRow("select id, name, gender, age, city from students where name = $1;", "wangwu")
	//err := row.Scan(&stu.Id, &stu.Name, &stu.Gender, &stu.Age, &stu.City)
	//defer rows.Close()
	//checkError(err)
	//fmt.Println(stu)

	// 查询多条数据
	//var stus []Student
	//rows, err := Db.Query("select id, name, gender, age, city from students;")
	//defer rows.Close()
	//checkError(err)
	//for rows.Next() {
	//	stu := Student{}
	//	_err := rows.Scan(&stu.Id, &stu.Name, &stu.Gender, &stu.Age, &stu.City)
	//	checkError(_err)
	//	stus = append(stus, stu)
	//}
	//fmt.Println(stus)

	// 更新数据
	//res, err := Db.Exec("update students set name = $1 where name = $2", "zhangsirui", "zhangsan")
	//checkError(err)
	//fmt.Println(res.LastInsertId()) // 0 LastInsertId is not supported by this driver
	//fmt.Println(res.RowsAffected()) // 0 <nil>

	// 删除数据
	res, err := Db.Exec("delete from students where name = $1", "lisi")
	checkError(err)
	fmt.Println(res.RowsAffected()) // 2 <nil>
}
