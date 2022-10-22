package repositories

import (
	"microservice/src/apps"
	"microservice/src/dto/request"
	"microservice/src/dto/response"
	"microservice/src/models/entities"

	"gorm.io/gorm"
)

// Example : https://gorm.io/docs/query.html

// Create
func InsertProducts(data *request.CreateProduct) *gorm.DB {
	db := apps.DatabaseConnect()
	result := db.Table("products").Create(&entities.Products{Code: data.Code, Price: data.Price})
	if result.Error != nil {
		panic(result.Error)
	}
	return result
}

// Read
func ReadProducts() []response.ReadProducts {
	db := apps.DatabaseConnect()
	var products []response.ReadProducts
	db.Table("products").Order("id desc").Find(&products)
	return products
}
func ReadProductsById(id int) entities.Products {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Table("products").First(&products, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}
func ReadProductsByCode(code string) entities.Products {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Table("products").Order("id desc").First(&products, "code = ?", code)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}

// Update
func UpdateProducts(data *request.UpdateProduct) {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Table("products").Model(&products).Where("id = ?", data.ID).Updates(entities.Products{Price: data.Price, Code: data.Code})
	if result.Error != nil {
		panic(result.Error)
	}
}

// Delete
func DeleteProductsById(id int) {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Table("products").Delete(&products, id)
	if result.Error != nil {
		panic(result.Error)
	}
}
