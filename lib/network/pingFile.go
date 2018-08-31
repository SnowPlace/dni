package network

import (
	"os/exec"
	"strconv"
	"strings"

	"../cp866"
	"../data"
)

func Ping(address string, pn, wtime int) (rp data.ResultPing, err error) {
	n := strconv.Itoa(pn)
	w := strconv.Itoa(wtime)
	out, err := exec.Command("ping", "-n", n, "-w", w, address).Output()
	if err != nil {
		rp.Result = false
		return
	}
	answer := strings.Split(string(cp866.Decode(out)), "\n")
	rp.Message = string(cp866.Decode(out))
	for _, line := range answer {

	}
	return
}
