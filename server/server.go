package main

import (
	"context"
	"net"

	pb "github.com/York-Shawn/micro-practice/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	conn *gorm.DB
}

func (p *ProductServer) GetProductList(ctx context.Context, req *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	var products []pb.Product
	p.conn.Offset(int(req.Page) - 1).Limit(int(req.PageSize)).Find(&products)
	res := &pb.GetProductListResponse{}
	for _, product := range products {
		res.List = append(res.List, &pb.Product{
			Id:        product.Id,
			Name:      product.Name,
			Stock:     product.Stock,
			SKU:       product.SKU,
			IsDeleted: product.IsDeleted,
		})
	}
	return res, nil
}

func (p *ProductServer) setupDBEngine() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3307)/mysql?charset=utf8&parseTime=True&loc=Local",
	}))
	if err != nil {
		return err
	}

	p.conn = db
	return nil
}

type test struct {
}

func main() {
	server := grpc.NewServer()
	productServer := &ProductServer{}
	productServer.setupDBEngine()

	pb.RegisterProductServiceServer(server, productServer)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	server.Serve(lis)
}
