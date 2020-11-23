package monitoringPAY

import (
	"Monitoring/cmd/monitoring/models"
	"net/http"
	"time"
)

func monPAY() (models.TData_resp, error) {
	const url string = "https://pmnt.tattelecom.ru:4443/osmp.php?command=check&txn_id=1&account=8432633006&sum=10.05&bank=test&prv_id=2&txn_date=20161121000000"
	const host string = "192.168.143.207"

	var resp_ models.TData_resp

	beginTime := time.Now().UnixNano()

	/*resp_, err := mon.Ping(host)
	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime_
		return resp_, errors.New(fmt.Sprint("Destination Host Unreachable - " + host))
	}*/

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		resp_.Status = "500"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/text")

	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		resp_.Status = "504"
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)

	if resp.Status != "200 OK" {
		resp_.Status = resp.Status
		resp_.RunTime = time.Now().UnixNano() - beginTime
		return resp_, err
	}

	resp_.Status = resp.Status
	resp_.RunTime = time.Now().UnixNano() - beginTime
	return resp_, nil
}

//start
func Run(debug bool) (int, error) {
	//var pespDB models.TData_resp
	//var pespUSSD models.TData_resp
	_, err := monPAY()
	if err != nil {
		return 0, err
	}

	return 1, nil
}
