package main

import (
	"coba-grpc/pb"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	// fmt.Println("Hello World!")
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Baju Warna Hitam",
				Price: 50000,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Baju",
				},
			},
			{
				Id:    2,
				Name:  "Sepatu Merah",
				Price: 30000,
				Stock: 20,
				Category: &pb.Category{
					Id:   2,
					Name: "Sepatu",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	fmt.Println(data) // -> compact binary wire format

	testProducts := &pb.Products{}
	err = proto.Unmarshal(data, testProducts)
	if err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println(testProducts)

	for _, v := range testProducts.Data {
		fmt.Println(v)
	}

	for _, product := range testProducts.GetData() {
		// fmt.Println(product.Name)
		fmt.Println(product.GetName(), product.GetCategory())
	}
}
