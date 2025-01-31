package domain

type IProduct interface {
    Save(product *Product) error
    GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(productId int32) error
}