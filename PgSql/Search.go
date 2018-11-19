package main

import (
    "database/sql"
    "fmt"
    _ "github.com/bmizerany/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=postgres password=12589 dbname=test sslmode=disable")
    checkErr(err)
    //查询数据
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    fmt.Println("success!")
    db.Close()
    fmt.Println("seccess!")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}