package domain

type Product struct {
    Product_id    int32   `gorm:"primaryKey;autoIncrement" json:"product_id"`
    Name  string  `gorm:"size:255" json:"name"`
    Price float32 `gorm:"type:decimal(10,2)" json:"price"`
}

func NewProduct(name string, price float32) *Product {
    return &Product{Name: name, Price: price}
}


func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}