package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure of a product
// swagger:model
type Product struct {
	// Product ID
	// example: 1
	Id          int     `json:"id"`
	// Name of the product
	// required: true
	Name        string  `json:"name" validate:"required"`
	// Description of the product
	// required: true
	Description string  `json:"description"`
	// Price of the product
	// required: true
	// example: 9.99
	Price       float32 `json:"price" validate:"gt=0"`
	// SKU of the product
	// required: true
	// example: `^[a-z]+-[0-9]+$`
	Sku         string  `json:"sku" validate:"required,skuFormat"`
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

func (p *Product) Validate() error {
	// Initialize Validate
	validate := validator.New()
	// Registering Custom Validation Function
	validate.RegisterValidation("skuFormat", validateSkuFormat)

	return validate.Struct(p)
}

func validateSkuFormat(fl validator.FieldLevel) bool {

	rgx := regexp.MustCompile(`^[a-z]+-[0-9]+$`)
	matches := rgx.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
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
