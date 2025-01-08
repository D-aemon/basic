# Basic Service

This project is a basic service example that demonstrates API communication and service integration using Go. The service is built with Go, using Protocol Buffers (Protobuf) for API definitions, and provides a simple example of service interaction.

## Table of Contents

- [Project Overview](#project-overview)
- [Directory Structure](#directory-structure)
- [Build and Setup](#build-and-setup)
    - [Prerequisites](#prerequisites)
    - [Building the Project](#building-the-project)
    - [Running the Project](#running-the-project)
- [API Documentation](#api-documentation)
- [License](#license)

## Project Overview

This service provides a simple example of an API that can be extended for real-world use cases. In this case, the example demonstrates managing **camera devices** (as an example of a resource) through CRUD operations. While "camera" is used for demonstration purposes, the architecture and API can be easily adapted for other entities.

### Features
- Basic CRUD API for managing camera devices (as an example).
- Protobuf-based service definitions for structured API interactions.
- Supports Docker-based deployment for easy setup.

## Directory Structure

```plaintext
./
├── api
│   ├── docs                # Swagger JSON documentation files for the API
│   ├── proto               # Protobuf definitions and generated files
├── build
│   ├── Dockerfile          # Dockerfile for building and running the service in a container
│   ├── Makefile            # Makefile to automate build and deploy processes
│   └── config.yml          # Configuration file for the service
├── cmd
│   └── basic               # Main application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksum file
└── internal                # Internal packages containing core logic
    ├── config              # Configuration parsing and handling
    ├── db                  # Database interaction and models
    ├── handler             # API request handlers
    ├── logger              # Logging utilities
    └── utils               # General utility functions
```

## Build and Setup
Prerequisites
Before you can build and run the project, ensure you have the following installed:

- Go (version 1.21 or later)
- Docker
- Protobuf Compiler
- Make
### Building the Project

```shell
cd build && make build
```

## Running the Project
Once the project is built, you can run the application:

```bash
./build/basic
```
If you want to run the service in a Docker container, you can build and run it using the Dockerfile:

```bash
docker run -p 8080:8080 registry.cn-hangzhou.aliyuncs.com/daemon_public/basic:v1.0.0
```
This will start the service on port 8080 inside the container.

## API Documentation
The API for the service is defined using Protocol Buffers (Protobuf) and can be found in the api/proto/ directory. The API is also documented using Swagger, and the generated .swagger.json files can be found in the api/docs directory.

The example API includes CRUD operations for a "camera" entity, but can be adapted for different resources as needed.

For more information on how to use the API, refer to the generated Swagger documentation or the .proto files for detailed endpoint definitions.