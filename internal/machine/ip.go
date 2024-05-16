package machine

import (
	"log"
	"net"
	"time"
)

func Reachable(ip string) error {
	conn, err := net.DialTimeout("tcp", ip, 3*time.Second)
	if err != nil {
		log.Fatalln("Failed to connect to the server:", err)
		return err
	} else {
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatalln("Failed to close the connection:", err)
			}
		}(conn)
		return nil
	}
}
