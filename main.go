package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	userloginv1 "github.com/rongfengliang/go-connect-app/pkg/userlogin/v1"
	userloginv1connect "github.com/rongfengliang/go-connect-app/pkg/userlogin/v1/userloginv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type UserLogin struct{}

func (s *UserLogin) Login(
	ctx context.Context,
	req *connect.Request[userloginv1.LoginRequest],
) (*connect.Response[userloginv1.LoginResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request username: ", req.Msg.Username)
	res := connect.NewResponse(&userloginv1.LoginResponse{
		Token:   "demoapp",
		Exprise: 100,
	})
	res.Header().Set("userlogin-Version", "v1")
	return res, nil
}

func main() {
	userlogin := &UserLogin{}
	api := http.NewServeMux()
	api.Handle(userloginv1connect.NewUserLoginServiceHandler(userlogin))
	mux := http.NewServeMux()
	mux.Handle("/grpc/", http.StripPrefix("/grpc", api))
	http.ListenAndServe(
		"0.0.0.0:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
