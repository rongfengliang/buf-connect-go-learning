# buf connect-go leaning


## running

```code
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
go mod tidy
go run main.go
```

## testing

```code
curl \
    --header "Content-Type: application/json" \
    --data '{"username": "dalong","userpassword":"demoapp"}' \
    http://localhost:8080/userlogin.v1.UserLoginService/Login
```