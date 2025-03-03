# Packaging Application

This repository contains a React.js frontend and a Go backend for a packaging application. The application calculates the optimal number of packages required for a given number of items.

# Prerequisites

Before running the application, ensure you have the following installed:
Docker: Install Docker
Docker Compose: Install Docker Compose

# Running the Application

Follow these steps to set up and run the application:

1. Update the API Server Address
Before starting the application, you need to update the REACT_APP_API_URL environment variable in the docker-compose.yaml file.

Open the docker-compose.yaml file.

Locate the REACT_APP_API_URL variable under the client service.

Change the value to match your environment:

For local development, use:
REACT_APP_API_URL=http://localhost:8080

2. Build and Run the Application

Run the following command to build and start the application:
sudo docker-compose up --build
This command will:

Build the Docker images for the React frontend and Go backend.

Start the containers for both services.

3. Access the Application
Once the build is complete, you can access the application in your browser:

React Frontend: Open http://localhost:3000.

Go Backend: The backend API will be available at http://localhost:8080.
