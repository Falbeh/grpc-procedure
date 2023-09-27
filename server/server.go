package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"

	course "github.com/Falbeh/grpc-procedure.git/grpc"
	"google.golang.org/grpc"
)

type Server struct {
	course.UnimplementedCourseAskServer // Necessary
	name                                string
	port                                int
}

var port = flag.Int("port", 0, "server port number")

func main() {
	flag.Parse()

	server := &Server{
		name: "ITU-server",
		port: *port,
	}

	// Start server
	go startServer(server)

	// Keep server running until manual quit
	for {

	}
}

func startServer(server *Server) {

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	course.RegisterCourseAskServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

func (c *Server) AskForCourse(ctx context.Context, in *course.CourseInfoRequest) (*course.CourseAskClient, error) {
	return nil, nil
}
