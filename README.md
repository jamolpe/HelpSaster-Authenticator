# Golang-Sesioner

A simple microservice to register and login users with session managment.

## Tech

-Docker
-GithubActions
-Postman

External libraries:

Logger library(also developed by me) -> https://github.com/jamolpe/gologger
GoMod

Inside the docs folder you can find a .json postman request template

## Structure

This proyect follows a DDD architecture aproach to allow easier functional testing.
For more information you can check the followed article I used to start a DDD aproach https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5

This microservice follows the most common golang microservice structure:

```bash
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
```

## Env-Variables

| variable           | Description                      | Example                   |
| ------------------ | -------------------------------- | ------------------------- |
| GO_ENV             | Actual environment               | DEV                       |
| CONNECTION_STRING  | mongo database connection string | mongodb://localhost:27017 |
| SESSION_COLLECTION | database collection for session  | Session                   |
| USERS_COLLECTION   | database collection for users    | Users                     |
| LOG_COLLECTION     | database collection for logs     | logs                      |
| LOG_LEVEL          | log level for the logger library | DEV                       |

## Testing

This library has test focused on functional testing (to be improved)
