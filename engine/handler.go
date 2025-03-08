package engine

import (
	"net"
	"strings"
	"time"
)

func HandleRequest(conn *net.UDPConn, clientAddr *net.UDPAddr, request []byte) {
	command := strings.TrimSpace(string(request))
	commandParts := strings.Fields(command)
	checker := true
	var response string
	if len(commandParts) == 0 {
		conn.WriteToUDP([]byte("(error) ERR empty command\n"), clientAddr)
		return
	}
	switch strings.ToUpper(commandParts[0]) {
	case "PING":
		response = "PONG\n"
	case "SET":
		if len(commandParts) < 3 {
			response = "(error) ERR wrong number of arguments for 'SET' command\n"
		} else {
			key := commandParts[1]
			value := strings.Join(commandParts[2:], " ")
			expiry := int64(0)

			if pxVal, i, bl := ifPXexists(commandParts); bl {
				if len(commandParts) > i+2 {
					response = "(error) ERR wrong number of arguments for 'SET' command\n"
					checker = false
				}
				if pxVal[:1] == "-" {
					response = "(error) ERR invalid PX value\n"
					checker = false
				}
				expiryMillis, err := time.ParseDuration(pxVal + "ms")
				if err != nil {
					response = "(error) ERR invalid PX value\n"
					checker = false
				} else {
					expiry = time.Now().Add(expiryMillis).UnixMilli()
					value = strings.Join(commandParts[2:i], " ")
				}
			}
			if checker {
				setKeyValue(key, value, expiry)
				response = "OK\n"
			}
		}
	case "GET":
		if len(commandParts) != 2 {
			response = "(error) ERR wrong number of arguments for 'GET' command\n"
		} else {
			key := commandParts[1]
			response = getValue(key) + "\n"
		}
	default:
		response = "(error) ERR unknown command\n"
	}

	conn.WriteToUDP([]byte(response), clientAddr)
}
