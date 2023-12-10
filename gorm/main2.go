package main

//
//import (
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"log"
//)
//
//type Product struct {
//	// 组合 gorm.Model 的字段，如：ID、CreatedAt、UpdatedAt、DeletedAt
//	gorm.Model
//	Code  string
//	Price uint
//}
//
//func main() {
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	// 迁移 schema，其实是用来创建表的语句
//	err = db.AutoMigrate(&Product{})
//	if err != nil {
//		log.Fatalln("can not create table")
//		return
//	}
//
//	// Create
//	db.Create(&Product{Code: "D42", Price: 100})
//
//	// Read
//	var product Product
//	// 根据整型主键查找，并将结果写入 product
//	db.First(&product, 1)
//	// 查找 code 字段值为 "D42" 的记录，并将结果写入到 product
//	db.First(&product, "code = ?", "D42")
//
//	// Update - 将 product 的 price 字段更新为 200
//	db.Model(&product).Update("price", 200)
//	// Update - 更新多个字段
//	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})                    // 仅更新 struct 中的非零值字段
//	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "DF2"}) // 不指定类型，写哪个 key 更新哪个 field
//
//	// Delete
//	db.Delete(&product, 1)
//}
