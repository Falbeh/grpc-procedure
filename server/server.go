package main

import (
	"context"
	"flag"
	"fmt"
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

func (c *Server) AskForCourse(ctx context.Context, in *course.CourseInfoRequest) (*course.CourseInfoResponse, error) {
	// Retrieve the requested course from the mock data
	courseData, ok := mockCourses[in.CourseId]

	log.Printf("Client with ID %d asked for info about course with ID %d", in.ClientId, in.CourseId)

	if !ok {
		return nil, fmt.Errorf("Course with ID %d not found", in.CourseId)
	}

	return courseData, nil
}

// Defining some mock courses to choose from
var mockCourses = map[int64]*course.CourseInfoResponse{
	1: {
		Id:          1,
		Name:        "Course 1",
		Description: "Description for Course 1",
		Teachers:    []string{"Teacher 1", "Teacher 2"},
		Workload:    5,
	},
	2: {
		Id:          2,
		Name:        "Course 2",
		Description: "Description for Course 2",
		Teachers:    []string{"Teacher 3", "Teacher 4"},
		Workload:    7,
	},
	3: {
		Id:          3,
		Name:        "Course 3",
		Description: "Description for Course 3",
		Teachers:    []string{"Teacher 5", "Teacher 6"},
		Workload:    6,
	},
}
