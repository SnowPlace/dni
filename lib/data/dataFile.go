package data

import "time"

var DBInfo Info = Info{}

type Info struct {
	Interfaces   map[int]InterfaceNet
	ResultsTests Results
	NetDevices   Devices
}
type InterfaceNet struct {
	ID          int
	Name        string
	Description string
	DHCP        bool
	IPv4        string
	Mask        string
	Gateway     string
	ServerDHCP  string
	ServerDNS   string
}

type Results struct {
	Test1 bool
	Test2 ResultPing
	Test3 bool
	Test4 ResultPing
	Test5 ResultPing
	Test6 ResultPing
	Test7 ResultPing
}
type ResultPing struct {
	Result  bool
	Message string
	MinTime time.Duration
	MaxTime time.Duration
	AvgTime time.Duration
	PCT     int
	State   bool
}
type Devices struct {
	Computers map[int]DeviceInfo
	POSs      map[int]DeviceInfo
	Scales    map[int]DeviceInfo
	Printers  map[int]DeviceInfo
	Video     map[int]DeviceInfo
	Outer     map[int]DeviceInfo
}
type DeviceInfo struct {
	Name      string
	IPAddress string
	State     bool
}

var DevicesMap map[int]string = map[int]string{
	2: "computer", 3: "computer", 4: "computer", 5: "computer", 6: "computer", 7: "computer", 8: "computer", 9: "computer",
	11: "pos", 12: "pos", 13: "pos", 14: "pos", 15: "pos", 16: "pos", 17: "pos", 18: "pos", 19: "pos",
	21: "scales", 22: "scales", 23: "scales", 24: "scales", 25: "scales", 26: "scales", 27: "scales", 28: "scales", 29: "scales",
	10: "printer", 20: "printer",
	50: "video", 51: "video", 52: "video", 53: "video", 54: "video",
}
