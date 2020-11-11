package commands

import (
	db "Monitoring/cmd/monitoring/database"
	"Monitoring/cmd/monitoring/models"
	"strconv"
)

func Run(debug bool) (int, error) {
	//var pespDB models.TData_resp
	//var pespUSSD models.TData_resp
	var resp_ models.TData_resp
	sqlTxt := "select count(*) S from billing.commands_queue where when_to_run < sysdate"
	resp_, err1 := db.QuerySQL(sqlTxt)
	if err1 != nil {
		return 0, err1
	}
	if n, err := strconv.Atoi(resp_.Status); err != nil {
		return 0, err
	} else {
		return n, nil
	}
}
