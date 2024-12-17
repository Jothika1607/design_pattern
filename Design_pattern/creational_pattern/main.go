package main

import (
	"fmt"
	"time"
)

// Product is an interface for all product types
type Product interface {
	GetDetails() string
}

// Electronics is a concrete implementation of Product
type Electronics struct {
	ID    int
	Name  string
	Brand string
	Price float64
}

func (e Electronics) GetDetails() string {
	return fmt.Sprintf("Electronics [ID: %d, Name: %s, Brand: %s, Price: %.2f]", e.ID, e.Name, e.Brand, e.Price)
}

// Clothing is a concrete implementation of Product
type Clothing struct {
	ID    int
	Name  string
	Size  string​
	Color string
	Price float64
}

func (c Clothing) GetDetails() string {
	return fmt.Sprintf("Clothing [ID: %d, Name: %s, Size: %s, Color: %s, Price: %.2f]", c.ID, c.Name, c.Size, c.Color, c.Price)
}

// Books is a concrete implementation of Product
type Books struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func (b Books) GetDetails() string {
	return fmt.Sprintf("Books [ID: %d, Title: %s, Author: %s, Price: %.2f]", b.ID, b.Title, b.Author, b.Price)
}

// ProductFactory is the Factory Method that creates products
func ProductFactory(productType string, id int, name string, price float64) Product {
	switch productType {
	case "Electronics":
		return Electronics{ID: id, Name: name, Brand: "BrandX", Price: price}
	case "Clothing":
		return Clothing{ID: id, Name: name, Size: "M", Color: "Red", Price: price}
	case "Books":
		return Books{ID: id, Title: name, Author: "AuthorY", Price: price}
	default:
		return nil
	}
}

// CreateProduct simulates the creation of a product
func CreateProduct(productType, name string, price float64) Product {
	// Generate a random ID (just for simulation)
	id := time.Now().UnixNano()
	return ProductFactory(productType, int(id), name, price)
}

// ReadProduct retrieves details of a product
func ReadProduct(p Product) {
	if p != nil {
		fmt.Println(p.GetDetails())
	} else {
		fmt.Println("Product not found.")
	}
}

// UpdateProduct simulates updating a product's price
func UpdateProduct(p Product, newPrice float64) Product {
	switch v := p.(type) {
	case Electronics:
		v.Price = newPrice
		return v
	case Clothing:
		v.Price = newPrice
		return v
	case Books:
		v.Price = newPrice
		return v
	default:
		return nil
	}
}

// DeleteProduct simulates deleting a product (just sets to nil here)
func DeleteProduct(p Product) {
	fmt.Println("Deleting product:", p.GetDetails())
}

func main() {

	fmt.Println("Create the Products")
	fmt.Println("*********************************************************")
	// Create new products
	product1 := CreateProduct("Electronics", "Iphone", 100000)
	product2 := CreateProduct("Clothing", "Jean", 1000)
	product3 := CreateProduct("Books", "Go Programming", 500)
	fmt.Println("Product1 : ",product1 )
	fmt.Println("Product2 : ",product2 )
	fmt.Println("Product3 : ",product3 )

	fmt.Println("Read the Products")
	fmt.Println("*********************************************************")
	// Read product details
	ReadProduct(product1)
	ReadProduct(product2)
	ReadProduct(product3)

	fmt.Println("Update the Products")
	fmt.Println("*********************************************************")
	// Update product price
	product1 = UpdateProduct(product1, 120000)
	product2 = UpdateProduct(product2, 2500)
	fmt.Println("Product1 : ",product1 )
	fmt.Println("Product2 : ",product2 )

	fmt.Println("Read the updated Products")
	fmt.Println("*********************************************************")
	// Read updated product details
	ReadProduct(product1)
	ReadProduct(product2)

	fmt.Println("Delete the Products")
	fmt.Println("*********************************************************")
	// Delete a product
	DeleteProduct(product3)


}
