package product

import (
	"database/sql"

	"github.com/nulfrost/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("select * from products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.CreatedAt)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	return nil, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Exec("insert into products(name, price, description, image, quantity) values (?, ?, ?, ?, ?)", product.Name, product.Price, product.Description, product.Image, product.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateProduct(id int) error {
	return nil
}

func (s *Store) DeleteProduct(id int) error {
	return nil
}
