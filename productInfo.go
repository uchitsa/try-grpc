package main

import "context"

type Server struct {
}

type Product struct {
}

func (s *Server) AddProduct(ctx context.Context, in *Product) (int, error) {
	//business logic
	return 0, nil
}

func (s *Server) GetProduct(ctx context.Context, id int) (*Product, error) {
	//business logic
	return nil, nil
}
