package main

import (
	"net/http"

	"github.com/York-Shawn/micro-practice/client"
	"github.com/York-Shawn/micro-practice/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	cc := proto.NewProductServiceClient(conn)
	ph := client.NewProducts(cc)

	server := gin.Default()
	server.GET("/", ph.GetProductList)

	s := &http.Server{
		Addr:           ":8081",
		Handler:        server,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
