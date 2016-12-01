package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/didrocks/grpc-piglow/piglowservice"
	"github.com/oleksandr/bonjour"
	"google.golang.org/grpc"
)

//go:generate protoc -I proto/ proto/piglow.proto --go_out=plugins=grpc:proto
func main() {
	var l net.Listener
	var err error

	// listen on given port
	port := 3146
	for l == nil {
		l, err = net.Listen("tcp", ":"+strconv.Itoa(port))
		if err != nil {
			log.Printf("failed to listen: %v\n", err)
			port++
		}
	}
	log.Printf("listening on port %d\n", port)

	// get machine and ip parameters
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not determine host: %v", err)
	}
	ip, err := getOutboundIP()
	if err != nil {
		log.Fatalf("could not determine ip: %v", err)
	}

	// register our mdns service
	b, err := bonjour.RegisterProxy("PiGlowGRPC", "_piglow._tcp", "", port, hostname, ip, nil, nil)
	if err != nil {
		log.Fatalf("couldn't publicize our service: %v", err)
	}
	defer b.Shutdown()

	// register and start our grpc server
	s := grpc.NewServer()
	if err := piglowservice.RegisterPiGlowService(s); err != nil {
		log.Fatalf("couldn't create piglow proxy: %v", err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer s.Stop()
}

// Get preferred outbound ip of this machine
func getOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx], nil
}
