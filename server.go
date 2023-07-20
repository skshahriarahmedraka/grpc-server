package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/skshahriarahmedraka/grpc-server/logs"
	"github.com/skshahriarahmedraka/grpc-server/proto"

	"google.golang.org/grpc"
)



type server struct{

}

func (s *server) Speaking(ctx context.Context, in *messagepb.SpeakRequest)(*messagepb.SpeakResponse,error){
	log.Printf("Received Client Request : %v \n",in.GetClient_Request())
	log.Printf("server LIstenning 50051 ...\n")

	return &messagepb.SpeakResponse{Server_Response : "Hello 👋 !!! What are you doing "+in.GetClient_Request()+" ?"},nil
}


func Server(){

	lis ,err := net.Listen("tcp",":50051")
	logs.Error("Error in listening",err)

	MyServer := grpc.NewServer()

	messagepb.RegisterConversationServer(MyServer,&server{})
	
	fmt.Println("listening on 50051 ... ")
	err = MyServer.Serve(lis) 

	logs.Error("Error in serving",err)


}