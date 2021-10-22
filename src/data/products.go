package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (products *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.Id = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, error := findProduct(id)

	if error != nil {
		return error
	}

	p.Id = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.Id == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	productTemp := productList[len(productList)-1]
	return productTemp.Id + 1
}

func (p *Product) FromJson(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

var productList = []*Product{
	&Product{
		Id:          1,
		Name:        "Latte",
		Description: "Coffee with milk",
		Price:       2.50,
		Sku:         "ltt01",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		Id:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee",
		Price:       3.10,
		Sku:         "xprs01",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
