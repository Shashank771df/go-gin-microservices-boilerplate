# go-gin-microservices-boilerplate

This repository contains the source code for go-gin-microservices-boilerplate. It is structured into two main packages:

## core

The `core` package contains base implementations of utilities, SDKs, or integrations that any of the microservices in the project may need.

## microservices

The `microservices` package contains all microservices of the project. Each microservice is organized into the following directories:

- **collections**: Postman files for testing purposes.
- **config**: Global environments of the applications.
- **controller**: Controllers for each main path of the microservice.
- **db_migrations**: SQL scripts in timestamp order for updating databases.
- **lib**: Interface to core utilities like alerts (Telegram), database, SDKs, etc.
- **log**: Interface to logger utilities.
- **pipes**: Pipes/inputs of each endpoint in the routes.
- **repository**: Database communication for the controllers. Contains generic functions like `Item`, `Items`, `NewItem`, `UpdateItem`, `RemoveItem`, and custom functions to create your own SQL scripts.
- **routes**: Endpoints of the microservice.
- **start.go**: Main file of the microservice.
- **startup.go**: Asynchronous file for starting anything that takes a long time for the microservice, for example initializing REDIS variables from SQL database.

### Starting Microservices

To start any microservice, you can use one of the following methods:

1. Navigate to the microservice directory:
   ```bash
   cd microservices/{micro_name}
   ```

   Then run using nodemon:
   ```bash
   nodemon
   ```

2. Navigate to the microservice directory:
   ```bash
   cd microservices/{micro_name}
   ```

   Then run using Go:
   ```bash
   go run *.go
   ```

Feel free to explore each microservice's directory for more detailed instructions and specific implementation details.

---

This template provides an overview of the project structure and instructions for starting microservices. You can customize it further by adding details about dependencies, installation instructions, contribution guidelines, etc., based on your project's specific requirements.