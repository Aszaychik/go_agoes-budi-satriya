package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)
type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type Products []Product

const productsAPIURL = "https://fakestoreapi.com/products"

func (p *Products) GetAllProducts(productChannel chan<- Product) (Products, error) {
	// Mengirim HTTP GET request ke productsAPIURL
	response, err := http.Get(productsAPIURL)
	if err != nil {
		return nil, err
	}

	// Menutup response.Body setelah selesai dibaca
	defer response.Body.Close()

	// Cek status code response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d", response.StatusCode)
	}

	// Decode JSON ke slice Product
	var products Products
	if err := json.NewDecoder(response.Body).Decode(&products); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// Kirim data ke channel
	for _, product := range products {
		productChannel <- product
	}

	return products, nil
}

func main() {
	var wg sync.WaitGroup
	productChannel := make(chan Product)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := (&Products{}).GetAllProducts(productChannel)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	// Menunggu Go routine sebelumnya selesai, jika selesai maka close channel
	go func() {
		wg.Wait()
		close(productChannel)
	}()

	fmt.Println("Product Data")
	for product := range productChannel {
		fmt.Println("===")
		fmt.Printf("Title: %s\nPrice: %.2f\nCategory: %s\n", product.Title, product.Price, product.Category)
		fmt.Println("===")
	}
}
