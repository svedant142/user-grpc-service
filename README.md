# user-grpc-service

HOW TO RUN :

Three ways to run it -

1. run the following command on terminal in the project directory - go run main.go
   it will start the server

2. Since the project has a docker container. run the following commands in project directory in case you have docker installed.
   STEP 1 - docker build -t user-service .
   STEP 2 - docker run -p 8080:8080 user-service

3. If using vscode or any IDE -
   run the main.go file in debug mode.

HOW TO HIT THE API -

1. Use postman to hit the grpc request by using the .proto file.
   STEP 1 - Add new gRPC request
   Step 2 - Add new API
   STEP 3 - use the .proto file at "business/user/grpc/proto/user.proto"
   STEP 5 - Set URL as localhost:8080 along with the API
   STEP 6 - Hit the INVOKE

OR  
For sample, a client is made to hit these endpoints in the following Github repo -
https://github.com/svedant142/client-user-service

NOTE :

1. The project is implemented in Clean Architecture DESIGN PATTERN.
2. It consists of unit and integration tests.
3. Mocked the database by maintaining a list of user details in a variable.
4. Dockerized the whole application.

ABOUT PROJECT -
A grpc service with two endpoints for fetching individual and list of details respectively.
