package mysql

import (
	"encoding/json"
	"fmt"
	"log"
)

type ClassTable struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// create databases
func doCreateDB(sql string) error {
	return fmt.Errorf("not implemented")
}

// show databases
func doShowDBs(sql string) ([]byte, error) {
	results, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer func() {
		results.Close()
	}()
	var temp = struct {
		Databases []interface{} `json:"Databases"`
	}{}
	cols, _ := results.Columns()
	colNum := len(cols)
	log.Printf("一共有%d列", colNum)
	slice := make([]interface{}, colNum)
	for results.Next() {
		results.Scan(slice...)
		temp.Databases = append(temp.Databases, slice...)
	}
	return json.Marshal(&temp)
}

// insert
func doInsert(sql string) error {
	insertRows, err := db.Query(sql)
	defer func() {
		insertRows.Close()
	}()
	return err
}

// select
func doSelect(sql string) ([]byte, error) {
	selectRows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer func() {
		selectRows.Close()
	}()

	var temp = []*ClassTable{}
	for selectRows.Next() {
		c := ClassTable{}
		err = selectRows.Scan(&c.Id, &c.Name, &c.Age)
		if err != nil {
			return nil, err
		}
		temp = append(temp, &c)
	}
	// cols, _ := selectRows.Columns()
	// colNum := len(cols)
	// log.Printf("一共有%d列", colNum)
	// slice := make([]interface{}, colNum)
	// for selectRows.Next() {
	// 	var i interface{}
	// 	err = selectRows.Scan(&slice[0], &slice[1], &slice[2])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	temp.Results = append(temp.Results, i)
	// }
	return json.Marshal(&temp)
}

// update

// drop/delete
