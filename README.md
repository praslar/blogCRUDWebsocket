# blogCRUDWebsocket
## An an HTTP server that implements an in-memory CRUD database with an endpoint to subscribe to all changes.

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

## Monitoring database change with browsers:
- http://localhost:8080/api/v1/monitor
