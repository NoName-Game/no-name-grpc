package client

import (
	"context"
	"log"
	"time"

	pb "bitbucket.org/no-name-game/nn-ws/app/schema/examples/schema"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	response, err := c.SayHelloAgain(ctx, &pb.HelloAgainRequest{
		Name:    "Vito",
		Surname: "caste",
	})
	if err != nil {
		log.Panicln(err)
	}

	// Request streaming
	// stream, err := c.All(ctx, &pb.PlayerRequest{})
	// if err != nil {
	// 	log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
	// }
	//
	// for {
	// 	feature, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	//
	// 	if err != nil {
	// 		log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
	// 	}
	//
	// 	log.Println(feature)
	// }

	log.Printf("Responde: %v", response.GetMessage())
}
