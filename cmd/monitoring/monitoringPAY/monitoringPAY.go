package monitoringPAY

import (
	"Monitoring/cmd/monitoring/models"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func monPAY(n int32) (models.TData_resp, error) {
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

	client := &http.Client{Timeout: time.Duration(rand.Int31n(n)) * time.Second}

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
	start := time.Now()
	start.Format("2006.01.02-15.04.05")
	_, err := monPAY(2)
	if err != nil {
		_, err2 := monPAY(10)
		if err2 != nil {
			_, err3 := monPAY(30)
			if err3 != nil {
				elapsed := time.Since(start)
				fmt.Println(elapsed)
				return 0, errors.New(fmt.Sprint(start.Format("2006.01.02-15.04.05") + "; RunTime : " + fmt.Sprint(elapsed) + "; Error : " + fmt.Sprint(err.Error())))
			}
		}
	}

	return 1, nil
}
