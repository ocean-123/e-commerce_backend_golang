package seed

import (
	"ecommerce_new/models"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func Seed(db *gorm.DB) {
	seedCategories(db)
	seedUsers(db)
	seedProducts(db)
	seedOrders(db)
}

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Electronics"},
		{Name: "Books"},
		{Name: "Clothing"},
		{Name: "Home & Kitchen"},
		{Name: "Sports"},
		{Name: "Toys"},
		{Name: "Health"},
		{Name: "Beauty"},
		{Name: "Grocery"},
		{Name: "Automotive"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			log.Fatalf("could not seed categories: %v", err)
		}
	}
}

func seedUsers(db *gorm.DB) {
	users := []models.User{
		{Name: "User 1", Email: "user1@example.com", Password: "password"},
		{Name: "User 2", Email: "user2@example.com", Password: "password"},
		{Name: "User 3", Email: "user3@example.com", Password: "password"},
		{Name: "User 4", Email: "user4@example.com", Password: "password"},
		{Name: "User 5", Email: "user5@example.com", Password: "password"},
		{Name: "User 6", Email: "user6@example.com", Password: "password"},
		{Name: "User 7", Email: "user7@example.com", Password: "password"},
		{Name: "User 8", Email: "user8@example.com", Password: "password"},
		{Name: "User 9", Email: "user9@example.com", Password: "password"},
		{Name: "User 10", Email: "user10@example.com", Password: "password"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("could not seed users: %v", err)
		}
	}
}

func seedProducts(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())
	var categories []models.Category
	db.Find(&categories)

	products := []models.Product{}
	for i := 1; i <= 50; i++ {
		products = append(products, models.Product{
			Name:        "Product " + string(i),
			Description: "Description for product " + string(i),
			Price:       float64(rand.Intn(1000)),
			CategoryID:  categories[rand.Intn(len(categories))].ID,
		})
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("could not seed products: %v", err)
		}
	}
}

func seedOrders(db *gorm.DB) {
	var users []models.User
	var products []models.Product

	db.Find(&users)
	db.Find(&products)

	for i := 1; i <= 10; i++ {
		orderItems := []models.OrderItem{}
		total := float64(0)

		for j := 0; j < rand.Intn(5)+1; j++ {
			product := products[rand.Intn(len(products))]
			quantity := rand.Intn(5) + 1
			price := product.Price * float64(quantity)
			total += price

			orderItems = append(orderItems, models.OrderItem{
				ProductID: product.ID,
				Quantity:  quantity,
				Price:     price,
			})
		}

		order := models.Order{
			UserID:     users[rand.Intn(len(users))].ID,
			OrderItems: orderItems,
			Total:      total,
		}

		if err := db.Create(&order).Error; err != nil {
			log.Fatalf("could not seed orders: %v", err)
		}
	}
}
