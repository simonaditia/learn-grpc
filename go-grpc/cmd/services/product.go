package services

import (
	"context"
	"go_grpc/cmd/helpers"
	pagingPb "go_grpc/go_grpc/pb/pagination"
	productPb "go_grpc/go_grpc/pb/product"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductService) GetProductsHardcoded(context.Context, *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{
		Pagination: &pagingPb.Pagination{
			Total:       10,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    2,
		},
		Data: []*productPb.Product{
			{
				Id:    1,
				Name:  "Naruto T-Shirt",
				Price: 10000.00,
				Stock: 10,
				Category: &productPb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Doraemon T-Shirt",
				Price: 15000.00,
				Stock: 20,
				Category: &productPb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
		},
	}

	return products, nil
}

func (p *ProductService) GetProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {
	var page int64 = 1

	var pagination pagingPb.Pagination
	var products []*productPb.Product

	// rows, err := p.DB.Table("products AS p").
	// 	Joins("LEFT JOIN categories AS c on c.id = p.category_id").
	// 	Select("p.id", "p.name", "p.price", "p.stock", "c.id AS category_id", "c.name AS category_name").
	// 	Rows()

	sql := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c on c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id AS category_id", "c.name AS category_name")

	offset, limit := helpers.Pagination(sql, page, &pagination)

	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Gagal mengambil row data %v", err.Error())
		}
		product.Category = &category
		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return response, nil
}
