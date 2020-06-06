package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "bitbucket.org/no-name-game/nn-ws/app/schema/examples/schema"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloAgainRequest) (*pb.HelloAgainReply, error) {
	log.Printf("Received: %v, %v", in.GetName(), in.GetSurname())

	response := new(pb.HelloAgainReply)
	response.Message = fmt.Sprintf("Hello %v, %v", in.GetName(), in.GetSurname())

	return response, nil
}

// func (p *PlayerController) All(in *pb.PlayerRequest, stream pb.NoName_AllServer) (err error) {
// 	log.Println("Message recived")
//
// 	var players models.Players
// 	err = players.All()
//
// 	log.Println(players)
//
// 	for _, player := range players {
// 		test := &pb.PlayerResponse{
// 			Id:       int32(player.ID),
// 			Username: player.Username,
// 		}
//
// 		if err := stream.Send(test); err != nil {
// 			return err
// 		}
// 	}
//
// 	return err
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
