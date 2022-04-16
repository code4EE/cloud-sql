package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/code4EE/cloud-sql/server"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func InitMysql(addr string) (err error) {
	db, err = sql.Open("mysql", addr)
	if err != nil {
		return
	}
	registerHTTPHandler()
	return
}

func registerHTTPHandler() {
	// example e.g http://localhost:8080/runsql?query=insert into Cities(name, population) Values('beijing', 200000)
	server.RegisterHandler("/runsql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 解析query语句
		q := r.URL.Query()
		var (
			sql                   string
			err                   error
			rowsByte, showDBsByte []byte
			handleError           = func(err error) {
				if err != nil {
					http.Error(w, fmt.Sprintf("sql running unsuccessfully: %+s", err.Error()), http.StatusInternalServerError)
					return
				}
			}
		)
		sql = q.Get("query")
		log.Printf("接收到sql参数: %s", sql)
		// 处理具体的sql语句
		newSQL := strings.ToLower(sql)
		if strings.HasPrefix(newSQL, "create") {
			// 如果是创建DB的语句
			err = doCreateDB(sql)
			handleError(err)
		} else if strings.HasPrefix(newSQL, "insert") {
			// 如果是插入语句
			err = doInsert(sql)
			handleError(err)
		} else if strings.HasPrefix(newSQL, "select") {
			// 如果是查询语句
			rowsByte, err = doSelect(sql)
			handleError(err)
			if rowsByte != nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write(rowsByte)
				return
			}
		} else if strings.Contains(newSQL, "show") && strings.Contains(newSQL, "databases") {
			showDBsByte, err = doShowDBs(sql)
			handleError(err)
			if showDBsByte != nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write(showDBsByte)
				return
			}
		}
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("sql running unsuccessfully: %+s", err.Error()), http.StatusInternalServerError)
		// 	return
		// }
		// if len(rowsByte) != 0 {

		// }
	}))
}
