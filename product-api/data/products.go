package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return fmt.Errorf("unable to validate SKU")
	}
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU is of format abc-defg-hijkl
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Marshal as it does not
// have to buffer the output into an in-memory slice of bytes
//
// https://pkg.go.dev/encoding/base64@go1.18#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}

	productList[position] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextId() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
