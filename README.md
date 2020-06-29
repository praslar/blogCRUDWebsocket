# blogCRUDWebsocket
## An an HTTP server that implements an in-memory CRUD database with an endpoint to subscribe to all changes.

## Requirments:
- Create an HTTP server that implements an in-memory CRUD database with an endpoint to subscribe to all changes.
- So you'd have endpoints to create, read, update and delete items and then another endpoint to register a subscriber to receive all changes in real time. 
- A change is just a JSON message indicating what item was changed and how it was changed. As in whether it was created, updated or deleted.

## Goals
1. Include at least one automated test.
2. Adhere to Go best practices.
3. No race conditions.

## How to run project?
```golang
go run main.go
```

## Integration test with postman: 
- GET: http://localhost:8080/api/v1/blog
- POST: http://localhost:8080/api/v1/blog
     
     body: { "title" : "testingCreate", "contect" : "testingCreate" }
- PUT: http://localhost:8080/api/v1/blog/{{blog_id}}
     
     body: { "title" : "testingUpdate", "contect" : "testingUpdate" }
- DELETE: http://localhost:8080/api/v1/blog/{{blog_id}}

## Monitoring database change with browsers (you can have multiple monitors):
- http://localhost:8080/api/v1/monitor

## Documents:
Websocket: go-programming-blueprints-2nd (Mat Ryer Book)

HTTP server: Go project layout standard
