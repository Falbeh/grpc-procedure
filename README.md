1. Define service endpoints student, course, teacher using gRPC
- These are defined in course.proto, student.proto and teacher.proto files.

2. Discuss whether the operations should use one way, or bi-directional streaming
- I believe bi-directional streaming would be optimal for this setup, as the client needs to request information about courses, students, and teachers, which the server then responds to. 

3. Implement the RPC / gRPC service in Golang, that exposes your course endpoint
- 

4. Consume the RPC / gRPC course service endpoint by creating a client in Golang
