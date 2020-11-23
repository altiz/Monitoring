package main

import (
	"Monitoring/cmd/monitoring/commands"
	"Monitoring/cmd/monitoring/monitoringPAY"
	"Monitoring/cmd/monitoring/monitoringSMS"
	"Monitoring/cmd/monitoring/monitoringUSSD"
	"Monitoring/cmd/monitoring/telegramm"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Configuration struct {
	UrlDB string `gorm:"column:UrlDB" json:"UrlDB"`
	Host  string `gorm:"column:Host" json:"host"`
}

func setConf() Configuration {
	file, err_ := os.Open("configs/conf.json")
	if err_ != nil {
		fmt.Println("error:", err_)
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func main() {
	_, err := monitoringUSSD.Run(true)
	if err != nil {
		telegramm.Send("USSD service DOWN!!! " + err.Error())
	}

	var ncount int
	ncount, err1 := commands.Run(true)
	if err1 != nil {
		telegramm.Send("Commands_queue service DOWN!!! " + err1.Error())
	}
	if ncount > 4000 {
		telegramm.Send("Commands_queue service DOWN!!! Commands = " + strconv.Itoa(ncount))
	}

	_, err2 := monitoringSMS.Run(true)
	if err2 != nil {
		telegramm.Send("USSD service DOWN!!! " + err2.Error())
	}

	_, err3 := monitoringPAY.Run(true)
	if err3 != nil {
		telegramm.Send("PAY service DOWN!!! " + err3.Error())
	}

}
