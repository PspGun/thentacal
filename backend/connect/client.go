package connect

import (
	"fmt"
	"github.com/fexcel/fexcel-backend/pb"
	"google.golang.org/grpc"
	"log"
	"os"
)

var Client *pb.PromptServiceClient

func Client_setup() {

	port := os.Getenv("PORT")
	// Connect to GRPC server using port given in .env file
	conn, err := grpc.Dial(fmt.Sprintf(":%s", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing grpc server %v", err)
	}
	client := pb.NewPromptServiceClient(conn)
	fmt.Println("connect lew jaa")
	Client = &client
}
