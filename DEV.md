# TeachMeChess Backend

This is the backend server for TeachMeChess, a web application that allows users to learn and practice chess with a custom AI teacher.

## Architecture

The backend is built using a microservice architecture, with separate services for authentication, user management, game logic, and AI. Each microservice communicates with the others via REST APIs and WebSockets.

### Services

The following microservices are included in the backend:

- **Auth Service**: Provides authentication and authorization services, including login and token generation.
- **User Service**: Manages user accounts and profiles, including registration and profile updates.
- **Game Service**: Handles game logic and player matchmaking.
- **AI Service**: Provides a custom AI teacher for each user, using natural language processing to answer questions and provide guidance during gameplay.

## Technologies Used

The backend is built using the following technologies:

- **Go**: A high-performance language that is well-suited for building concurrent and networked applications.
- **Gorilla WebSocket**: A Go library that provides WebSocket functionality for the server.
- **Redis**: A fast and scalable in-memory data store used for caching and session management.
- **MongoDB**: A NoSQL document database used for storing user and game data.
- **JWT**: A standard for securely transmitting information between parties, used for token-based authentication.
- **Docker**: A platform for containerizing and deploying applications.

## Getting Started

To get started with the backend, follow these steps:

1. Install Go and Docker on your machine.
2. Clone the repository to your local machine.
3. Run `docker-compose up` to start the Redis and MongoDB containers.
4. Run `go run auth-service/main.go` to start the Auth Service.
5. Run `go run user-service/main.go` to start the User Service.
6. Run `go run game-service/main.go` to start the Game Service.
7. Run `go run ai-service/main.go` to start the AI Service.

## Contributing

Contributions to the project are welcome and appreciated! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them to the new branch.
4. Push the branch to your forked repository.
5. Open a pull request to the main repository.

Before submitting a pull request, please ensure that your changes are well-documented and tested.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
