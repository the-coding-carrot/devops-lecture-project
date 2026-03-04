package pkg

import "testing"

var products = []Product{
	{ID: 1, Name: "Minecraft", Price: 24.99},
	{ID: 2, Name: "Grand Theft Auto V", Price: 59.99},
	{ID: 3, Name: "Assassin's Creed II", Price: 39.99},
	{ID: 4, Name: "Star Wars: Jedi Survivor", Price: 89.99},
}

func TestFindProductByID(t *testing.T) {

	for _, p := range products {
		product := FindProductByID(products, p.ID)
		if product == nil {
			t.Errorf("Expected to find product with ID %d, but got nil", p.ID)
			continue
		}
		if product.ID != p.ID {
			t.Errorf("Expected product ID %d, but got %d", p.ID, product.ID)
		}
		if product.Name != p.Name {
			t.Errorf("Expected product name %s, but got %s", p.Name, product.Name)
		}
		if product.Price != p.Price {
			t.Errorf("Expected product price %.2f, but got %.2f", p.Price, product.Price)
		}

	}
	product := FindProductByID(products, -1)
	if product != nil {
		t.Errorf("Expected to get nil for non-existing product ID, but got %v", product)
	}

	emptyList := []Product{}
	product = FindProductByID(emptyList, 1)
	if product != nil {
		t.Errorf("Expected to get nil for empty product list, but got %v", product)
	}
}
