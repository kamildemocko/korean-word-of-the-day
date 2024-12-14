# Korean Word of the Day

Korean Word of the Day is a Go application that scrapes a Korean dictionary website to fetch the word of the day and a conversation sentence of the day. It then sends these details to a specified Discord webhook.

## Installation Guide

### Prerequisites
- Docker
- Docker Compose
- Go (if running locally)
- ROD Manager (if running remotely)

### Steps
1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd korean-word-of-the-day
    ```

2. Create a `.env` file in the root directory with the following keys:
    ```env
    ROD_MANAGER_ADDR=http://rod-manager:port
    DISCORD_WEBHOOK=https://discord.com/api/webhooks/your-webhook-id/your-webhook-token
    ```

3. Build and run the application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

### Running Locally
If you prefer to run the application locally without Docker:
1. Ensure you have Go installed.
2. Build the application:
    ```sh
    go build -o bin/korean-word-of-a-day.exe ./cmd/app
    ```
3. Run the application:
    ```sh
    ./bin/korean-word-of-a-day.exe -local
    ```

## .env Keys
- `ROD_MANAGER_ADDR`: The address of the Rod manager for remote browser control. In this case it's a Docker image running on the same network.
- `DISCORD_WEBHOOK`: The Discord webhook URL to send the messages.

## Makefile Commands
- `build`: Builds the Go application.
- `start`: Builds and starts the application.
- `stop`: Stops the running application.

## License
This project is licensed under the MIT License.