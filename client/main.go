package main

import (
	"context"
	"grpcb/greet"
	"grpcb/greetsvc"
	"io"
	"log"

	"google.golang.org/grpc"
)

var (
	names = []string{
		"James", "Jack", "Jason", "Jacob", "Jobs", "Julia",
	}
)

func callOnce(client greetsvc.GreetingServiceClient) error {
	var (
		g   *greet.Greet
		err error
	)
	g, err = client.GreetOnce(context.Background(), &greetsvc.GreetRequest{Name: "Jason"})
	if err != nil {
		return err
	} else {
		log.Println(g)
	}
	return err
}

func callAtOnce(client greetsvc.GreetingServiceClient) error {
	var (
		gat *greetsvc.GreetGathered
	)
	stream, err := client.GreetAtOnce(context.Background())
	if err != nil {
		return err
	}
	for _, name := range names {
		err = stream.Send(
			&greetsvc.GreetRequest{Name: name},
		)
		if err != nil {
			return err
		}
	}
	gat, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Println(gat)
	return nil
}

func callMany(client greetsvc.GreetingServiceClient) error {
	var (
		gp  *greet.Greet
		err error
	)
	stream, err := client.GreetMany(
		context.Background(), &greetsvc.GreetRequest{Name: "Jessie"},
	)
	if err != nil {
		return err
	}
recv_loop:
	for {
		gp, err = stream.Recv()
		switch err {
		case nil:
		case io.EOF:
			break recv_loop
		default:
			return err
		}
		log.Println(gp)
	}
	return nil
}

func callStream(client greetsvc.GreetingServiceClient) error {
	var (
		err  error
		wait = make(chan struct{})
	)
	stream, err := client.GreetStream(context.Background())
	if err != nil {
		return err
	}
	/* Receiver */
	go func(stream greetsvc.GreetingService_GreetStreamClient, waitc chan<- struct{}) {
		var (
			gp  *greet.Greet
			err error
		)
		/* signal waiting goroutines on close, multiple close will cause panic. */
		defer close(waitc)
	recv_loop:
		for {
			gp, err = stream.Recv()
			switch err {
			case nil:
			case io.EOF:
				break recv_loop
			default:
				log.Fatal(err)
			}
			log.Println(gp)
		}
	}(stream, wait)

	for _, name := range names {
		err = stream.Send(
			&greetsvc.GreetRequest{Name: name},
		)
		if err != nil {
			return err
		}
	}
	// Must manually close the channel.
	stream.CloseSend()
	/* sync with receiver. */
	/* This is a useful pattern simulating a broadcast condition lock with chan. */
	<-wait
	return nil
}

func main() {
	var (
		clientOpt []grpc.DialOption
		err       error
	)

	clientOpt = append(clientOpt, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:55100", clientOpt...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := greetsvc.NewGreetingServiceClient(conn)

	calls := []func(greetsvc.GreetingServiceClient) error{
		callOnce, callAtOnce, callMany, callStream,
	}
	for _, call := range calls {
		call(client)
		if err != nil {
			log.Fatal(err)
		}
	}

}
