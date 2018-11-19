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
    res, err := stmt.Exec("postgres", "研发部门", "2012-12-09")
    checkErr(err)
    //pg不支持这个函数，因为他没有类似MySQL的自增ID
    // id, err := res.LastInsertId()
    // checkErr(err)

    // fmt.Println(id)

    //更新数据
    stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
    checkErr(err)
	fmt.Println("sul")
    res, err = stmt.Exec("astaxieupdate", 1)
    checkErr(err)
	fmt.Println("sul")
    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

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

    //删除数据
    stmt, err = db.Prepare("delete from userinfo where uid=$1")
    checkErr(err)

    res, err = stmt.Exec(1)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}