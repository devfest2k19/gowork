# echo server with standard libraries and handler implementation 


```go
$ go build -o main main.go
$ ./main
```

#### call the endpoints
```sh 
curl localhost:8001/helloFunc -d aruna
curl localhost:8001/helloFunc

curl localhost:8001/helloHandler -d aruna
curl localhost:8001/helloHandler 
```
