package net

import (
	"net"
)

func CheckIp(str string) bool {
	result := net.ParseIP(str)
	if result == nil {
		return false
	}
	
	return true
}

