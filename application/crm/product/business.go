package product

import (
	"context"

	domain "github.com/isaqueveras/servers-microservices-backend/domain/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/infrastructure/persistence/crm/product"
	"github.com/isaqueveras/servers-microservices-backend/services/grpc"
)

// GetProducts contains the logic to fetch a list of products
func GetProducts(ctx context.Context) (res *ListProducts, err error) {
	res = new(ListProducts)

	var (
		data *domain.ListProducts
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	if data, err = repo.GetProducts(); err != nil {
		return res, err
	}

	defer conn.Close()

	res.Data = make([]Product, len(data.Products))
	for i := range data.Products {
		res.Data[i] = Product{
			ID:          data.Products[i].ID,
			Name:        data.Products[i].Name,
			Description: data.Products[i].Description,
			Price:       data.Products[i].Price,
		}
	}

	return
}

// GetDetailsProduct contains the logic to fetch the details of product
func GetDetailsProduct(ctx context.Context, id *int64) (*Product, error) {
	var (
		data *domain.Product
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	data, err := repo.GetDetailsProduct(id)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return (*Product)(data), nil
}

// CreateProduct contains the logic to create a product
func CreateProduct(ctx context.Context, in *Product) error {
	var (
		conn = grpc.GetProductConnection()
		repo = product.New(ctx, conn)
		prod = &domain.Product{
			ID:          in.ID,
			Name:        in.Name,
			Description: in.Description,
			Price:       in.Price,
		}
	)

	if err := repo.CreateProduct(prod); err != nil {
		return err
	}

	defer conn.Close()

	return nil
}

// ListAllProductsWithMinimumQuantity contains the logic to fetch a list of products with minimum quantity
func ListAllProductsWithMinimumQuantity(ctx context.Context) (res *ListProducts, err error) {
	res = new(ListProducts)

	var (
		data *domain.ListProducts
		conn = grpc.GetProductConnection()
	)

	repo := product.New(ctx, conn)
	if data, err = repo.ListAllProductsWithMinimumQuantity(); err != nil {
		return res, err
	}

	defer conn.Close()

	res.Data = make([]Product, len(data.Products))
	for i := range data.Products {
		res.Data[i] = Product{
			ID:     data.Products[i].ID,
			Name:   data.Products[i].Name,
			Amount: data.Products[i].Amount,
		}
	}

	return
}
