# Golang-Sesioner

A simple microservice to register and login users with session managment.

## Structure

This microservice follows the most common golang microservice structure:

.
+--.github
| +--...
+--CMD
| +--auth-service
| | +--main.go
+--docs
| +--AuthService.postman_collection.json
+--internal
| +--api
| | +--...
| +--authorization
| | +--...
| +--authorization-core
| | +--...
| +--middlewares
| | +--...
| +--repository
| | +--...
| +--session-core
| | +--...
+--pkg
| +--errors
| | +--...
| +--models
| | +--...
+--test
| +--authroization
| | +--...
| +--...
+--.env
+--Dockerfile
+--go.mod
+--go.sum
+--README.md

## Env-Variables

| variable           | Description                      | Example                   |
| ------------------ | -------------------------------- | ------------------------- |
| GO_ENV             | Actual environment               | DEV                       |
| CONNECTION_STRING  | mongo database connection string | mongodb://localhost:27017 |
| SESSION_COLLECTION | database collection for session  | Session                   |
| USERS_COLLECTION   | database collection for users    | Users                     |
| LOG_COLLECTION     | database collection for logs     | logs                      |
| LOG_LEVEL          | log level for the logger library | DEV                       |
