package main

import (
	"log"
	"math/rand"

	context "golang.org/x/net/context"

	pb "github.com/didrocks/grpc-piglow/remote"

	"time"

	"google.golang.org/grpc"
)

//go:generate protoc -I ../remote/ ../remote/piglow.proto --go_out=plugins=grpc:remote
func main() {
	conn, err := grpc.Dial("192.168.0.150:9875", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
