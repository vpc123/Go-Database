package main

import (
    "database/sql"
    "fmt"
    _ "github.com/bmizerany/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres password=12589 dbname=test sslmode=disable")
    checkErr(err)
    //删除数据
    stmt, err = db.Prepare("delete from userinfo where uid=$1")
    res, err = stmt.Exec(3)
    affect, err = res.RowsAffected()
    fmt.Println(affect)
    db.Close()
    fmt.Println("seccess!")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}