# todo app
<img src="https://www.docker.com/wp-content/uploads/2023/08/logo-guide-logos-1.svg" alt="Docker" width="100" height="100">

This is a small project designed to introduce Docker and Docker Compose. It was created in approximately 3 hours.

## Prerequisites

- Docker
- Docker Compose

## How to Launch

1. Clone the repository: ` git clone https://github.com/CharlyRousseau/todo-app.git`

2. Navigate to the project directory: `cd todo-app/`

3. Edit the `docker-compose.yml` file and ensure that the port for the frontend service matches the `ALLOWED_PORT` in the backend service. For example, if the frontend service is running on port 3000, the `ALLOWED_PORT` should also be 3000.

4. Build and run the Docker containers:


Your application should now be running at `http://localhost:<FRONTEND_PORT>`, where `<FRONTEND_PORT>` is the port you set for the frontend service.

## Important Note

Please ensure that the port for the frontend service in the `docker-compose.yml` file matches the `ALLOWED_PORT` in the backend service. This is necessary for the CORS configuration to work correctly.
