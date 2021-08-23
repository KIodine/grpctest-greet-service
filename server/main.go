package main

import (
	"grpcb/greetsvc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var (
		serverOpt []grpc.ServerOption
		err       error
	)
	listener, err := net.Listen("tcp", "localhost:55100")
	if err != nil {
		log.Fatal(err)
	}
	// no need to manually close, user handles all control to `server.Serve`, the
	// server will close it.
	//defer listener.Close()

	server := grpc.NewServer(serverOpt...)

	greetsvc.RegisterGreetingServiceServer(server, &greetingService{})

	log.Println("GreetingService starts...")
	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
