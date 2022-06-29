package client

import (
	"context"
	"fmt"
	"log"

	pb "github.com/York-Shawn/micro-practice/proto"
	"github.com/gin-gonic/gin"
)

type Products struct {
	cc pb.ProductServiceClient
}

func NewProducts(cc pb.ProductServiceClient) *Products {
	return &Products{
		cc: cc,
	}
}

func (p *Products) GetProductList(c *gin.Context) {
	list, err := p.cc.GetProductList(context.Background(), &pb.GetProductListRequest{
		Page:     1,
		PageSize: 3,
	})
	if err != nil {
		log.Fatalf("Get product list failed: %v", err)
	}

	for _, product := range list.List {
		fmt.Println(product)
	}
}
