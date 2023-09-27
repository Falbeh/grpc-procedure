package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	course "github.com/Falbeh/grpc-procedure.git/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id int
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number (should match the port used for the server)")
)

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id: *clientPort,
	}

	// Wait for the client (user) to ask for course info
	go waitForCourseRequest(client)

	for {

	}
}

func waitForCourseRequest(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		inputInt64, err := strconv.ParseInt(input, 10, 64)

		if err != nil {
			log.Printf("Input error: Input has to be the id of a course")
		} else {
			log.Printf("Client asked for course info with input: %s\n", input)

			// Ask the server for course info
			courseReturnMessage, err := serverConnection.AskForCourse(context.Background(), &course.CourseInfoRequest{
				ClientId: int64(client.id),
				CourseId: int64(inputInt64),
			})

			if err != nil {
				log.Printf(err.Error())
			} else {
				log.Printf("Server sent the following info about course with id: %d\n", courseReturnMessage.Id)
				log.Printf("Course Name: %s", courseReturnMessage.Name)
				log.Printf("Course Description: %s", courseReturnMessage.Description)
				log.Printf("Teachers:")
				for _, teacher := range courseReturnMessage.Teachers {
					log.Printf("  - %s", teacher)
				}
				log.Printf("Workload: %d", courseReturnMessage.Workload)
			}
		}
	}
}

func connectToServer() (course.CourseAskClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return course.NewCourseAskClient(conn), nil
}
