// Package localhost
package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	log.SetFlags(0)

	local, err := GetLocalHostIp()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(local)
}

func GetLocalHostIp() (string, error) {
	conn, err := net.Dial("udp", "cyeam.com:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	LOCALHOST := strings.Split(conn.LocalAddr().String(), ":")[0]
	return LOCALHOST, nil
}
