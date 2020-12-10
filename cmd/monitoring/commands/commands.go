package commands

import (
	db "Monitoring/cmd/monitoring/database"
	"Monitoring/cmd/monitoring/models"
	"errors"
	"fmt"
	"time"
)

func Run(debug bool) (int, error) {
	start := time.Now()
	start.Format("2006.01.02-15.04.05")
	var resp_ models.TData_resp
	sqlTxt := "select count(*) S from billing.commands_queue where when_to_run < sysdate"
	resp_, err1 := db.QuerySQL(sqlTxt)
	if err1 != nil {
		elapsed := time.Since(start)
		return 0, errors.New(fmt.Sprint(start.Format("2006.01.02-15.04.05") + "; RunTime : " + fmt.Sprint(elapsed) + "; Error : " + fmt.Sprint(err1.Error())))
	}
	return resp_.Value, nil
}
