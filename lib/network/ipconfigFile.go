package network

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"../cp866"
	"../data"
)

func Ipconfig() (err error) {
	inter := map[int]data.InterfaceNet{}

	out, err := exec.Command("ipconfig", "/all").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer := strings.Split(string(cp866.Decode(out)), "\n")
	empty := data.InterfaceNet{}
	for _, line := range answer {
		switch {
		case strings.Index(line, "Aдаптер") != -1 || strings.Index(line, "Ethernet adapter") != -1:
			empty.Name = string([]rune(line)[8 : len([]rune(line))-2])
		case strings.Index(line, "Описание") != -1:
			description := strings.Split(line, ":")[1]
			empty.Description = string([]rune(description)[1 : len([]rune(description))-1])
		case strings.Index(line, "DHCP включен") != -1:
			dhcp := strings.Split(line, ":")[1]
			if strings.Index(dhcp, "Да") != -1 {
				empty.DHCP = true
			} else {
				empty.DHCP = false
			}
		case strings.Index(line, "IPv4-адрес") != -1:
			ip := strings.Split(line, ":")[1]
			separator := strings.Split(ip, "(")[0]
			empty.IPv4 = string([]rune(separator)[1:])
		case strings.Index(line, "Маска подсети") != -1:
			mask := strings.Split(line, ":")[1]
			empty.Mask = string([]rune(mask)[1 : len([]rune(mask))-1])
		case strings.Index(line, "Основной шлюз") != -1:
			gateway := strings.Split(line, ":")[1]
			empty.Gateway = string([]rune(gateway)[1 : len([]rune(gateway))-1])
		case strings.Index(line, "DHCP-сервер") != -1:
			serverDHCP := strings.Split(line, ":")[1]
			empty.ServerDHCP = string([]rune(serverDHCP)[1 : len([]rune(serverDHCP))-1])
		case strings.Index(line, "DNS-серверы") != -1:
			serverDNS := strings.Split(line, ":")[1]
			empty.ServerDNS = string([]rune(serverDNS)[1 : len([]rune(serverDNS))-1])

		case strings.Index(line, "NetBios через TCP/IP") != -1:
			inter[len(inter)] = empty
			empty = data.InterfaceNet{}
			empty.ID = len(inter)
		}
	}
	data.DBInfo.Interfaces = inter

	return
}
