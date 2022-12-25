package models

import "test-auth/service/grpc/product"

type Auth struct {
	ClientID string
	Token    string
}

type AuthStatus struct {
	Status int64
}
type Product struct {
	ID    string
	Price float64
	Qty   int64
}

func FromAuthGRPCToAuth(auth *product.AuthStruct) Auth {
	return Auth{
		ClientID: auth.ClientID,
		Token:    auth.AuthToken,
	}
}

func FromAuthStatusToStatus(auth *product.CheckAuthResponse) AuthStatus {
	return AuthStatus{
		Status: auth.Status,
	}
}

func FromGRPCProductToProduct(prResp *product.ProductsResponse) []*Product {
	products := make([]*Product, len(prResp.Products))
	for i, pr := range prResp.Products {
		prod := &Product{
			ID:    pr.Id,
			Price: pr.Price,
			Qty:   pr.Qty,
		}
		products[i] = prod
	}
	return products
}
