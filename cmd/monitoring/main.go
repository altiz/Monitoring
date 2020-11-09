package main

import (
	"Monitoring/cmd/monitoring/monitoringUSSD"
	"encoding/json"
	"fmt"
	"os"
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
	s, _ := monitoringUSSD.Run(true)
	fmt.Println(s)

}
