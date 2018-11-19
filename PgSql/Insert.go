package main

import (
    "database/sql"
    "fmt"
    _ "github.com/bmizerany/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres password=12589 dbname=test sslmode=disable")
    checkErr(err)
    fmt.Println("sul")

    //插入数据
    stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
    checkErr(err)
    fmt.Println("sul")
    stmt.Exec("hhh1", "研发部门3", "2018-11-19")

    db.Close()
    fmt.Println("seccess!")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}