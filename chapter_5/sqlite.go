package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./foo.s3db")
    checkErr(err)

    // 插入数据
    stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec("mangolau", "研发部门", "2021-06-14")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
    // 更新数据
    stmt, err = db.Prepare("UPDATE userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("mangolau34", id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // 查询数据
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created time.Time
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    // 删除数据
    stmt, err = db.Prepare("DELETE from userinfo WHERE uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
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