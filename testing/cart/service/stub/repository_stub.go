package stub

import "context"


type RepositoryStub struct {}

func (r RepositoryStub) AddToCart(ctx context.Context, userID string, productID string) error {
	// another logic reside here
	return nil
}

func NewRepositoryStub() RepositoryStub{
	return RepositoryStub{}
}