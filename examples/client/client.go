package main

import (
	"log"
	"math/rand"

	context "golang.org/x/net/context"

	pb "github.com/didrocks/grpc-piglow/proto"
	"github.com/oleksandr/bonjour"

	"time"

	"strconv"

	"google.golang.org/grpc"
)

//go:generate protoc -I ../../proto/ ../../proto/piglow.proto --go_out=plugins=grpc:proto
func main() {

	// get the service ip and port
	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		log.Fatalf("failed to initialize mdns resolver: %v\n", err.Error())
	}
	results := make(chan *bonjour.ServiceEntry)
	go func() {
		err = resolver.Lookup("PiGlowGRPC", "_piglow._tcp", "", results)
		if err != nil {
			log.Fatalf("failed to find grpc piglow: %v\n", err)
		}
	}()
	// we only get the first result and connect to it
	var m *bonjour.ServiceEntry
	select {
	case m = <-results:
	case <-time.After(20 * time.Second):
		log.Fatalf("no PiGlow service found on the network")
	}
	resolver.Exit <- true
	conn, err := grpc.Dial(m.AddrIPv4.String()+":"+strconv.Itoa(m.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPiGlowClient(conn)

	// seed random generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var num, prevnum int32 = -1, -1
	for {
		var brightness uint32
		if prevnum == -1 {
			num = (int32)(r1.Intn(18))
			brightness = 100
			prevnum = num
		} else {
			brightness = 0
			prevnum = -1
		}
		log.Printf("Send light command to %d\n", num)
		_, err = c.SetLED(context.Background(), &pb.LedRequest{Num: num, Brightness: brightness})
		log.Printf("Done")
		if err != nil {
			log.Fatalf("Error in setting led: %v", err)
		}
		time.Sleep(time.Second)
	}

}
