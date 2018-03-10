package net

import "net"

func GetHostAddr() (ip net.IP, err error) {
	ifaces, err := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.To4(), nil
				}
			}
		}
	}
	return
}
