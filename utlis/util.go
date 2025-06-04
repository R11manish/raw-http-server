package utlis

import (
	"fmt"
	"strconv"
	"strings"

	"r11manish.com/model"
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

func ParsedHttpRequest(data []byte) (*model.HTTPRequest, error) {
	dataStr := string(data)

	lines := strings.Split(dataStr, "\r\n")
	if len(lines) < 1 {
		return nil, fmt.Errorf("invalid Http Request")
	}

	requestLine := strings.Split(lines[0], " ")
	if len(requestLine) != 3 {
		return nil, fmt.Errorf("invalid request line")
	}

	req := &model.HTTPRequest{
		Method:  requestLine[0],
		Path:    requestLine[1],
		Version: requestLine[2],
		Headers: make(map[string]string),
	}

	bodyStart := 1

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			bodyStart = i + 1
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			req.Headers[strings.ToLower(key)] = value
		}
	}

	if bodyStart < len(lines) {
		req.Body = strings.Join(lines[bodyStart:], "\r\n")
	}
	return req, nil
}

func FormatHeaders(headers map[string]string) string {
	var parts []string
	for key, value := range headers {
		parts = append(parts, fmt.Sprintf(`    "%s": "%s"`, key, value))
	}
	return "{\n" + strings.Join(parts, ",\n") + "\n  }"
}
