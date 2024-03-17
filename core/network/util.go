package network

import (
	"app/core"
	"errors"
	"net"
	"strings"
)

type UtilNetwork struct {
}

func GetAllMacs() (map[string]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ret := map[string]string{}

	for _, ifa := range ifas {
		address := ifa.HardwareAddr.String()
		if address != "" {
			ret[ifa.Name] = address
		}
	}
	return ret, nil
}

func GetMac(iface string) (string, error) {
	ifas, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	for _, ifa := range ifas {
		if ifa.Name == iface {
			return ifa.HardwareAddr.String(), nil
		}
	}
	return "", errors.New("iface not found")
}

func (obj UtilNetwork) IsInSubnet(ip string, subnets []string) bool {
	for _, subnet := range subnets {
		_, ipnetA, _ := net.ParseCIDR(subnet)
		ipB, _, _ := net.ParseCIDR(ip)

		if ipnetA.Contains(ipB) {
			return true
		}
	}

	return false
}

func (obj UtilNetwork) RemoteIPAddress(c core.AppContext) string {
	xFF := c.Request.Header.Get("x-forwarded-for")
	remote := c.RemoteIP()

	if xFF == "" {
		return remote
	} else {
		xFFs := strings.Split(xFF, ",")

		if len(xFFs) == 0 || xFFs[0] == "" {
			return remote
		}

		return xFFs[0]
	}
}
