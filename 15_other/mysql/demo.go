package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func init() {
	dsn := "root:root@tcp(127.0.0.1)/db1"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panicf("dsn is wrong, err:%v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Panicf("ping failed, err:%v\n", err)
		return
	}
	// 设置最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大闲置连接数
	db.SetMaxIdleConns(0)
	fmt.Printf("connect success\n")
}

// 查询一行
func queryOneById(id int) {
	s := "select id, name, age from t_user where id=?"
	var u user
	err := db.QueryRow(s, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("query err: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", u)
}

// 查询多行
func queryMultiRows() {
	s := "select * from t_user"
	rows, err := db.Query(s)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("close conn failed, err: %v\n", err)
		}
	}(rows)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("get info failed, err: %v\n", err)
			return
		}
		fmt.Printf("%v\n", u)
	}
}

// 插入数据
func insert(name string, age int) {
	s := "insert into t_user (name,age) value(?,?)"
	result, err := db.Exec(s, name, age)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affedcted failed, err: %v\n", err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last inserted data id failed, err: %v\n", err)
		return
	}
	fmt.Printf("affected %d rows, last inserted %d", affected, lastInsertId)
}

// 删除数据
func del(id int) {
	s := "delete from t_user where id=?"
	result, err := db.Exec(s, id)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affedcted failed, err: %v\n", err)
		return
	}
	fmt.Printf("affected %d rows", affected)
}

// 修改数据
func update(id int, age int) {
	s := "update t_user set age=? where id=?"
	result, err := db.Exec(s, age, id)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affedcted failed, err: %v\n", err)
		return
	}
	fmt.Printf("affected %d rows", affected)
}

// 预处理，可重复使用，提高效率
// 防止 sql 注入
func prepareQuery(id int) {
	s := "select * from t_user where id>=?"
	stmt, err := db.Prepare(s)
	if err != nil {
		fmt.Printf("prepare statement failed, err: %v\n", err)
		return
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Printf("close conn failed, err: %v\n", err)
		}
	}(stmt)
	// 此处是查询，增删改操作使用 Exec
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("get user info failed, err: %v\n", err)
			return
		}
		fmt.Printf("%v\n", u)
	}
}

// 使用事务
func transaction(id1, age1, id2, age2 int) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin transaction failed, err: %v\n", err)
		return
	}

	s := "update t_user set age=? where id=?"

	_, err = tx.Exec(s, age1, id1)
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("set age of %d to %d failed, err: %v\n", id1, age1, err)
		return
	}

	_, err = tx.Exec(s, age2, id2)
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("set age of %d to %d failed, err: %v\n", id2, age2, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Printf("excute transaction succeed!\n")

}

func main() {
	//queryOneById(1)
	//queryMultiRows()
	//insert("金伟民", 21)
	//del(3)
	//update(2, 22)
	//prepareQuery(0)
	//transaction(1, 20, 2, 20)
}
