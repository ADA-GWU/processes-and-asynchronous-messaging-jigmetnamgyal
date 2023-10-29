[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/qg4qXfSB)
## Description
This Github repo contains code implementation of the asynchronous concurrent messaging.

- The code is written in **Golang**.
- This code implemented a sender and reader software that reads the list of Database server IPs and connect to all the DBs in different threads.
- In the sender software, it creates a threads equal to the DB servers, and it will request user's input. 
- Every time the user enters a text, one of the threads will insert a record into ASYNC_MESSAGE table with your SENDER_NAME
- In the reader software, it check available messages in each DB. 
- Avail message is the one that has (RECEIVED_TIME IS NULL and SENDER_NAME != yours).
- It also block the record while reading to prevent readers access and show the same message all the time.

## Installation

Need a golang install in your machine @latest version
find [Link](https://go.dev/doc/install)

Install all the dependencies with the following command

```shell
go mod download && go mod tidy
```

Run The following command to run the sender software

```shell
go run sender.go
```

This will prompt user to write a message. The user message will then be saved to database.

In order to exit the program. Press:
**control + c**

Run the reader software with the following command:

```shell
go run reader.go
```

This will query the message from the list of database and display on the terminal. With the Format

**Sender XXX sent XXX at time XXXX.**

All of this are handled in go routine.