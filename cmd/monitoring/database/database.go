package database

import (
	"Monitoring/cmd/monitoring/models"
	"database/sql"
	"time"

	_ "github.com/godror/godror"
)

func QuerySQL(sqlTxt string) (models.TData_resp, error) {
	var resp_ models.TData_resp
	beginTime := time.Now().UnixNano()
	db, err := sql.Open("godror", `user="ttk_billing" password="wdbip" connectString="e-scan:1521/irbis"`)
	if err != nil {
		resp_.Status = "20000"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}

	rows, err := db.Query(sqlTxt)
	if err != nil {
		resp_.Status = "20000"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	defer db.Close()
	resp_.Status = thedate
	resp_.RunTime = time.Now().UnixNano() - beginTime
	return resp_, nil

}
