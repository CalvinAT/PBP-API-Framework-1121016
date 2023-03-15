package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM product"
	id := c.Request.URL.Query()["id"]
	if id != nil {
		query += " WHERE id=" + id[0]
	}

	rows, err := db.Query(query)

	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	var product Product
	var products []Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			c.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		} else {
			products = append(products, product)
		}
	}

	var response ResponseData
	response.Status = 200
	response.Message = "Select Success"
	response.Data = products
	c.IndentedJSON(200, response)
}

func InsertProducts(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	var product Product
	name := c.Request.Form.Get("name")
	price, _ := strconv.Atoi(c.Request.Form.Get("price"))

	_, errQuery := db.Exec("INSERT INTO product(name, price) values (?,?)", name, price)

	rows, err := db.Query("SELECT id FROM product")
	for rows.Next() {
		if err := rows.Scan(&product.ID); err != nil {
			c.IndentedJSON(400, gin.H{"message": err.Error()})
			return
		}
	}

	var response ResponseData
	if errQuery == nil {
		product.Name = name
		product.Price = price
		response.Status = 200
		response.Message = "Insert Success"
		response.Data = product
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}
	c.IndentedJSON(response.Status, response)
}

func UpdateProducts(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	var product Product
	name := c.Request.Form.Get("name")
	price, _ := strconv.Atoi(c.Request.Form.Get("price"))

	productId := c.Param("prodId")

	_, errQuery := db.Exec("UPDATE product SET name = ?,price = ? WHERE id = ?", name, price, productId)
	var response ResponseData
	if errQuery == nil {
		productIdInt, _ := strconv.Atoi(productId)
		product.ID = productIdInt
		product.Name = name
		product.Price = price
		response.Status = 200
		response.Message = "Update Success"
		response.Data = product
	} else {
		response.Status = 400
		response.Message = "Update Failed"
	}
	c.IndentedJSON(response.Status, response)
}

func DeleteProducts(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	productId := c.Param("prodId")

	_, errQuery := db.Exec("DELETE FROM product WHERE id=?", productId)

	var response ResponseData
	if errQuery == nil {
		response.Status = 200
		response.Message = "Delete Success"
	} else {
		response.Status = 400
		response.Message = "Delete Failed"
	}
	c.IndentedJSON(response.Status, response)
}
