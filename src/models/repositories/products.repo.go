package repositories

import (
	"main-service/src/apps"
	"main-service/src/dto/request"
	"main-service/src/dto/response"
	"main-service/src/models/entities"

	"gorm.io/gorm"
)

// Example : https://gorm.io/docs/query.html

// Create
func InsertProduct(data *request.CreateProduct) *gorm.DB {
	db := apps.DatabaseConnect()
	result := db.Debug().Table("products").Create(&entities.Products{Code: data.Code, Price: data.Price})
	if result.Error != nil {
		panic(result.Error)
	}
	return result
}

// Read
func ReadProducts() []response.ReadProducts {
	db := apps.DatabaseConnect()
	var products []response.ReadProducts
	db.Debug().Table("products").Order("id desc").Find(&products)
	return products
}
func ReadProductById(id int) entities.Products {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Debug().Table("products").First(&products, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}
func ReadProductByCode(code string) entities.Products {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Debug().Table("products").Order("id desc").First(&products, "code = ?", code)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}

// Update
func UpdateProduct(data *request.UpdateProduct) {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Debug().Table("products").Model(&products).Where("id = ?", data.ID).Updates(entities.Products{Price: data.Price, Code: data.Code})
	if result.Error != nil {
		panic(result.Error)
	}
}

// Delete
func DeleteProductById(id int) {
	db := apps.DatabaseConnect()
	var products entities.Products
	result := db.Debug().Table("products").Delete(&products, id)
	if result.Error != nil {
		panic(result.Error)
	}
}
