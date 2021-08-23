package main

import (
	"context"
	"fmt"
	"grpcb/greet"
	"grpcb/greetsvc"
	"io"
	"log"
	"math/rand"
)

var (
	greets = []string{
		"Hello!",
		"你好!",
		"hola!",
		"konichiwa!",
	}
)

type greetingService struct {
	greetsvc.UnimplementedGreetingServiceServer
	/* --- other custom fields --- */
}

func chooseGreet() string {
	s := greets[int(rand.Uint32())%len(greets)]
	return s
}

func (s *greetingService) GreetOnce(ctx context.Context, req *greetsvc.GreetRequest) (*greet.Greet, error) {
	return &greet.Greet{Name: req.GetName(), Message: chooseGreet()}, nil
}

func (s *greetingService) GreetAtOnce(stream greetsvc.GreetingService_GreetAtOnceServer) error {
	var (
		gat    = &greetsvc.GreetGathered{}
		gSlice []*greet.Greet
		gr     *greetsvc.GreetRequest
		err    error
	)
recv_loop:
	for {
		gr, err = stream.Recv()
		switch err {
		case nil: // pass
		case io.EOF:
			break recv_loop
		default:
			log.Fatal(err)
		}
		gs := chooseGreet()
		gSlice = append(gSlice,
			&greet.Greet{Name: gr.GetName(), Message: gs},
		)
	}
	gat.Greets = gSlice
	stream.SendAndClose(gat)
	return nil
}

func (s *greetingService) GreetMany(req *greetsvc.GreetRequest, stream greetsvc.GreetingService_GreetManyServer) error {
	var (
		err error
	)
	if req == nil {
		return fmt.Errorf("nil request")
	}

	for _, gs := range greets {
		err = stream.Send(
			&greet.Greet{Name: req.GetName(), Message: gs},
		)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	return nil
}

func (s *greetingService) GreetStream(stream greetsvc.GreetingService_GreetStreamServer) error {
	var (
		err error
		req *greetsvc.GreetRequest
		//gp *greet.Greet
	)
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		err = stream.Send(
			&greet.Greet{Name: req.GetName(), Message: chooseGreet()},
		)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}
