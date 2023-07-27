# Mailer Service Project

This is a mailer service implemented in Golang which integrates RabbitMQ for message queue processing. The service provides functionality to process emails with customizable HTML templates.

## Project Structure

The basic structure of the project is presented as follows:

```
.
├── Dockerfile
├── go.mod
├── go.sum
├── mailer
│   └── mailer.go
├── main.go
├── rabbitmq
│   └── rabbitmq.go
├── templates
│   ├── login_template.html
│   ├── payment_template.html
└── token
    └── token.go
```

## Usage

### Build & Run

Assuming that you have Golang installed on your system, you can build with: 

```bash
$ go build
```

Run the application:

```bash
$ ./main
```

### Docker Build

The Dockerfile allows for building a Docker image:

```bash
$ docker build -t your-docker-username/mailer-service .
```

Then run the image:

```bash
$ docker run -p 8080:8080 -d your-docker-username/mailer-service
```
Remember to replace "your-docker-username" with your Docker username.

## Templates

This project uses templates for generating email body. Current templates include `login_template.html` and `payment_template.html`. You can customize these templates or add more according to your requirements.

## RabbitMQ

The `main.go` file starts a server that connects to a RabbitMQ service. The `rabbitmq/rabbitmq.go` file contains logic for connecting to the RabbitMQ service and listening to the message queue.

## Mailer

Once messages are received from the RabbitMQ queue, they are processed and email is sent using the functionality defined in `mailer/mailer.go`. Customize this according to your SMTP server and credentials.

## Token

The `token/token.go` file is used to validate the token associated with each request. Write your own logic according to your use case.

Remember to set up your SMTP server credentials and RabbitMQ server address in their respective files or through environment variables.

## Tests

No automated test has been written for this project.

## Versioning

Check the `go.mod` and `go.sum` files for dependency versioning.

Keep updated with your requirements!

Note: Please set up SMTP server credentials and RabbitMQ properly for successful operation.