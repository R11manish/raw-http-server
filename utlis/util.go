package utlis

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseIPV4(ipStr string) ([4]byte, error) {
	var ip [4]byte

	if ipStr == "localhost" {
		return [4]byte{127, 0, 0, 1}, nil
	}

	parts := strings.Split(ipStr, ".")

	if len(parts) != 4 {
		return ip, fmt.Errorf("invalid IPv4 address format : %s", ipStr)
	}

	for i, part := range parts {
		octet, err := strconv.Atoi(part)

		if err != nil {
			return ip, fmt.Errorf("invalid octet '%s' in IP address %s: %v", part, ipStr, err)
		}

		if octet < 0 || octet > 255 {
			return ip, fmt.Errorf("octet %d out of range (0-255) in IP address %s", octet, ipStr)
		}

		ip[i] = byte(octet)
	}

	return ip, nil
}
