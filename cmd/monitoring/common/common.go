package common

import (
	"Monitoring/cmd/monitoring/models"
	"fmt"
	"os/exec"
	"time"
)

func Ping(IP string) (models.TData_resp, error) {
	var resp models.TData_resp
	Command := fmt.Sprintf("ping -c 1 " + IP + " > /dev/null && echo 1 || echo 0")
	beginTime := time.Now().UnixNano()
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		resp.Status = "0"
		resp.RunTime = time.Now().UnixNano() - beginTime
		return resp, err
	}
	resp.Status = string(output[0])
	resp.RunTime = time.Now().UnixNano() - beginTime
	return resp, nil
}
