package main

import (
	"fmt"
)

// Product is the "Subject" that maintains a list of observers
type Product struct {
	ID       int
	Name     string
	Price    float64
	Observers []Observer
}

// Observer is the interface for objects that need to be notified when the Product's price changes
type Observer interface {
	Update(price float64)
}

// Customer represents an observer that wants to be notified about price changes
type Customer struct {
	ID       int
	Email    string
	InterestedIn string // The product that the customer is interested in
}

// Update is called when the product's price changes
func (c *Customer) Update(price float64) {
	fmt.Printf("Customer %s has been notified. New price of %s: %.2f\n", c.Email, c.InterestedIn, price)
}

// AddObserver adds an observer to the Product's list of observers
func (p *Product) AddObserver(observer Observer) {
	p.Observers = append(p.Observers, observer)
}

// RemoveObserver removes an observer from the Product's list
func (p *Product) RemoveObserver(observer Observer) {
	var index int
	var found bool
	for i, o := range p.Observers {
		if o == observer {
			index = i
			found = true
			break
		}
	}
	if found {
		p.Observers = append(p.Observers[:index], p.Observers[index+1:]...)
	}
}

// NotifyObservers notifies all registered observers about a price change
func (p *Product) NotifyObservers() {
	for _, observer := range p.Observers {
		observer.Update(p.Price)
	}
}

// CRUD Operations for Products

// CreateProduct creates a new product
func CreateProduct(id int, name string, price float64) *Product {
	return &Product{ID: id, Name: name, Price: price}
}

// UpdateProductPrice updates the price of a product and notifies observers
func UpdateProductPrice(p *Product, newPrice float64) {
	p.Price = newPrice
	fmt.Printf("Product %s price updated to %.2f\n", p.Name, newPrice)
	p.NotifyObservers()
}

// ReadProductDetails reads the details of a product
func ReadProductDetails(p *Product) {
	fmt.Printf("Product Details: [ID: %d, Name: %s, Price: %.2f]\n", p.ID, p.Name, p.Price)
}

// DeleteProduct simulates deleting a product
func DeleteProduct(p *Product) {
	fmt.Printf("Product %s has been deleted\n", p.Name)
}

func main() {

	fmt.Println("~~~~~~~~~~~~~~~~ E-Commerce ~~~~~~~~~~~~~~~~~")
	// Create a product
	product := CreateProduct(1, "Smartphone", 499.99)

	// Create customers who are interested in the product
	customer1 := &Customer{ID: 1, Email: "customer1@example.com", InterestedIn: "Smartphone"}
	customer2 := &Customer{ID: 2, Email: "customer2@example.com", InterestedIn: "Smartphone"}

	// Add customers as observers of the product
	product.AddObserver(customer1)
	product.AddObserver(customer2)

	// Read the product details before the price change
	ReadProductDetails(product)

	// Update the product's price and notify observers
	UpdateProductPrice(product, 449.99)

	// Read the product details after the price change
	ReadProductDetails(product)

	// Remove an observer (Customer1)
	product.RemoveObserver(customer1)

	// Update the product's price again and notify remaining observers
	UpdateProductPrice(product, 399.99)

	// Delete the product
	DeleteProduct(product)
}
