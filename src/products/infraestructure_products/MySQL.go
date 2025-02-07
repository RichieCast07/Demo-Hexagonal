package infraestructureproducts

import (
	"demo/src/core"
	domain "demo/src/products/domain_products"
	"log"
	"fmt"
)

type ProductRepoMySQL struct {
	connection core.ConectionMySQL
}

func NewCreateProductRepoMySQL() *ProductRepoMySQL {
	conn := core.MySQLConection()
	if conn.Err != "" {
		log.Fatal("Error al configurar la pool de conexiones", conn.Err)
	}
	return &ProductRepoMySQL{connection: *conn}
}

func (r *ProductRepoMySQL) Save(product *domain.Product) error {
	query := "INSERT INTO products (name, price, amount) VALUES (?, ?, ?)"
	_, err := r.connection.ExecPreparedQuerys(query, product.Name, product.Price, product.Amount)
	if err != nil {
		log.Printf("Error al registrar producto: %v", err)
		return err
	}
	return nil
}

func (r *ProductRepoMySQL) GetAll() ([]domain.Product, error) {
	query := "SELECT product_id, name, price, amount FROM products"
	rows, err := r.connection.FetchRows(query)
	var products []domain.Product
	if err != nil {
		log.Printf("Error al obtener productos: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Product_id, &product.Name, &product.Price, &product.Amount); err != nil {
			log.Printf("Error al escanear la fila: %v", err)
			continue
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar sobre las filas: %v", err)
		return nil, err
	}
	return products, nil
}


func (r *ProductRepoMySQL) GetByID(id int32) (domain.Product, error) {
	query := "SELECT product_id, name, price, amount FROM products WHERE product_id = ?"
	rows, err := r.connection.FetchRows(query, id)
	if err != nil {
		log.Printf("Error al obtener producto: %v", err)
		return domain.Product{}, err
	}
	defer rows.Close()

	var product domain.Product
	if rows.Next() {
		if err := rows.Scan(&product.Product_id, &product.Name, &product.Price, &product.Amount); err != nil {
			log.Printf("Error al escanear el producto: %v", err)
			return product, err
		}
		return product, nil
	}

	return product, fmt.Errorf("producto con ID %d no encontrado", id)
}

func (r *ProductRepoMySQL) Edit(name string, price float32, amount float32, id int32) error {
	query := "UPDATE products SET name = ?, price = ?, amount = ? WHERE product_id = ?"
	_, err := r.connection.ExecPreparedQuerys(query, name, price, amount, id)
	if err != nil {
		log.Printf("Error al editar info. de producto: %v", err)
		return err
	}
	return nil
}

func (r *ProductRepoMySQL) Delete(id int32) error {
	query := "DELETE FROM products WHERE product_id = ?"
	_, err := r.connection.ExecPreparedQuerys(query, id)
	if err != nil {
		log.Printf("Error al eliminar producto: %v", err)
		return err
	}
	return nil
}
