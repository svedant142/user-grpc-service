# user-grpc-service

HOW TO RUN :

Three ways to run it -

1. run the following command on terminal - go run main.go
   it will start the server

2. Since the project has a docker container. run the following command in project directory in case you have docker installed.
   docker run -p 8080:8080 user-service

3. If using vscode or any IDE -
   run the main.go file in debug mode.

NOTE :

1. The project is implemented in clean architecture design pattern.
2. It consists of unit and integration tests.
3. Mocked the database by maintaining a list of user details in a variable.
4. Dockerized the whole application.

ABOUT PROJECT -
A grpc service with the following capabilities-
● Mock the database by maintaining a list of user details in a variable.
● An endpoint to fetch user details based on user id.
● An endpoint to fetch a list of user details based on a list of ids

For sample, a client is made to hit these endpoints in the following Github repo -
https://github.com/svedant142/client-user-service
