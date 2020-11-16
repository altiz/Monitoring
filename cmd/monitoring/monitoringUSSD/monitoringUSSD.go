package monitoringUSSD

import (
	mon "Monitoring/cmd/monitoring/common"
	"Monitoring/cmd/monitoring/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func monDB() (models.TData_resp, error) {

	const url string = "http://192.168.143.207/v1/get_balance"
	const msisdn string = "9872305570"
	const host string = "192.168.143.207"
	type TData_req struct {
		Msisdn string `gorm:"column:msisdn" json:"msisdn"`
	}

	type TData_resp struct {
		CurrentBalance string `gorm:"column:vCurrentBalance" json:"vCurrentBalance"`
		Client         string `gorm:"column:nClient" json:"nClient"`
	}

	var req_ TData_req
	var resp_ models.TData_resp
	var bodyObj TData_resp

	beginTime_ := time.Now().UnixNano()

	resp_, err := mon.Ping(host)
	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime_
		return resp_, errors.New(fmt.Sprint("Destination Host Unreachable - " + host))
	}

	req_.Msisdn = msisdn
	beginTime := time.Now().UnixNano()
	dat, err := json.Marshal(req_)

	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dat))
	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		resp_.Status = "504"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &bodyObj); err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, errors.New(fmt.Sprint("invalid body"))
	}
	if bodyObj.Client != "1282042" {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, errors.New(fmt.Sprint("client is not identified"))
	}
	resp_.Status = resp.Status
	resp_.RunTime = time.Now().UnixNano() - beginTime
	return resp_, nil
}

func monUSSD() (models.TData_resp, error) {

	const url string = "http://192.168.143.190/v1/smsc"
	const host string = "192.168.143.190"
	const msisdn string = "9872305570"
	const serviceNumber string = "*100#"
	const sessionId string = "test1234"

	type TData_req struct {
		Msisdn        string `gorm:"column:msisdn" json:"msisdn"`
		ServiceNumber string `gorm:"column:serviceNumber" json:"serviceNumber"`
		Request       string `gorm:"column:request" json:"request"`
		SessionId     string `gorm:"column:sessionId" json:"sessionId"`
	}

	type TData_resp struct {
		Text       string `gorm:"column:text" json:"text"`
		SessionId  string `gorm:"column:sessionId" json:"sessionId"`
		EndSession int    `gorm:"column:endSession" json:"endSession"`
	}

	var req_ TData_req
	var resp_ models.TData_resp
	var bodyObj TData_resp
	beginTime_ := time.Now().UnixNano()

	resp_, err := mon.Ping(host)
	if err != nil {
		resp_.Status = "0"
		resp_.RunTime = time.Now().UnixNano() - beginTime_
		return resp_, errors.New(fmt.Sprint("Destination Host Unreachable - " + host))
	}

	req_.Msisdn = msisdn
	req_.ServiceNumber = serviceNumber
	req_.SessionId = sessionId
	beginTime := time.Now().UnixNano()
	dat, err := json.Marshal(req_)

	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dat))
	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 8 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		resp_.Status = "504"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &bodyObj); err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, errors.New(fmt.Sprint("invalid body - USSD"))
	}

	if bodyObj.Text[:12] != "Баланс" {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, errors.New(fmt.Sprint("client is not identified - USSD"))
	}

	resp_.Status = resp.Status
	resp_.RunTime = time.Now().UnixNano() - beginTime
	return resp_, nil
}

//start
func Run(debug bool) (int, error) {
	//var pespDB models.TData_resp
	//var pespUSSD models.TData_resp
	_, err1 := monDB()
	if err1 != nil {
		return 0, err1
	}

	_, err2 := monUSSD()
	if err2 != nil {
		return 0, err2
	}

	return 1, nil
}
