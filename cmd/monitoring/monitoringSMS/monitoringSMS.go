package monitoringSMS

import (
	mon "Monitoring/cmd/monitoring/common"
	"errors"
	"fmt"
)

func Run(debug bool) (int, error) {
	const host string = "192.168.143.208"

	_, err := mon.Ping264(host)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Destination Host Unreachable - " + host))
	}
	return 1, nil
}
