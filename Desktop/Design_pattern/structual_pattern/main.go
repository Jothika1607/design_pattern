package main

import (
	"fmt"
)

// Product is a leaf object that represents a single product
type Product struct {
	ID    int
	Name  string
	Price float64
}

// Category is a composite object that can hold Products or Subcategories
type Category struct {
	ID       int
	Name     string
	Products []Product
	Subcategories []Category
}

// Component interface defines common operations for both Products and Categories
type Component interface {
	GetDetails() string
	AddProduct(product Product)
	AddSubcategory(subcategory Category)
	GetProducts() []Product
	GetSubcategories() []Category
}

// GetDetails prints the product details
func (p Product) GetDetails() string {
	return fmt.Sprintf("Product [ID: %d, Name: %s, Price: %.2f]", p.ID, p.Name, p.Price)
}

// GetDetails prints the category details and lists the products and subcategories
func (c Category) GetDetails() string {
	categoryDetails := fmt.Sprintf("Category [ID: %d, Name: %s]", c.ID, c.Name)
	if len(c.Products) > 0 {
		categoryDetails += "\n  Products:"
		for _, product := range c.Products {
			categoryDetails += "\n    " + product.GetDetails()
		}
	}
	if len(c.Subcategories) > 0 {
		categoryDetails += "\n  Subcategories:"
		for _, subcategory := range c.Subcategories {
			categoryDetails += "\n    " + subcategory.GetDetails()
		}
	}
	return categoryDetails
}

// AddProduct adds a product to the category
func (c *Category) AddProduct(product Product) {
	c.Products = append(c.Products, product)
}

// AddSubcategory adds a subcategory to the category
func (c *Category) AddSubcategory(subcategory Category) {
	c.Subcategories = append(c.Subcategories, subcategory)
}

// CRUD Operations for Categories and Products

// CreateProduct simulates creating a new product
func CreateProduct(id int, name string, price float64) Product {
	return Product{ID: id, Name: name, Price: price}
}

// CreateCategory simulates creating a new category
func CreateCategory(id int, name string) Category {
	return Category{ID: id, Name: name}
}

// UpdateProduct simulates updating a product's price
func UpdateProduct(p Product, newPrice float64) Product {
	p.Price = newPrice
	return p
}

// DeleteProduct simulates deleting a product
func DeleteProduct(p Product) {
	fmt.Println("Deleting product:", p.GetDetails())
}

// ReadCategory simulates reading the details of a category
func ReadCategory(c Category) {
	fmt.Println(c.GetDetails())
}

func main() {
	fmt.Println("~~~~~~~~~  E-Commerce ~~~~~~~~~~~")
	fmt.Println("Create the Products")
	fmt.Println("*********************************************************")

	// Create Products
	product1 := CreateProduct(1, "Laptop", 60000)
	product2 := CreateProduct(2, "Smartphone", 20000)
	product3 := CreateProduct(3, "Tablet", 70000)

	fmt.Println("Product1 : ",product1 )
	fmt.Println("Product2 : ",product2 )
	fmt.Println("Product3 : ",product3 )


	fmt.Println("Create the Products Category")
	fmt.Println("*********************************************************")	
	// Create Categories
	category1 := CreateCategory(1, "Electronics")
	category2 := CreateCategory(2, "Mobile Devices")
	fmt.Println("Category1 : ",category1 )
	fmt.Println("Category1 : ",category1 )

	// Add products to categories
	category1.AddProduct(product1)
	category1.AddProduct(product2)
	category2.AddProduct(product3)

	// Add subcategory to Electronics category
	category1.AddSubcategory(category2)

	// Read the details of categories and products
	fmt.Println("Category Details:")
	ReadCategory(category1)

	fmt.Println("Update the Products")
	fmt.Println("*********************************************************")	

	// Update product price
	product1 = UpdateProduct(product1, 899.99)

	// Read updated product and category details
	fmt.Println("\nUpdated Product and Category Details:")
	
	ReadCategory(category1)

	fmt.Println("Delete the Products")
	fmt.Println("*********************************************************")	
	// Delete product
	DeleteProduct(product2)	
}
