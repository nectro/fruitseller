package services

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

func AddProduct(name string, price float32, imageid string, description string) uuid.UUID {
	newuuid := uuid.New()
	currentTime := time.Now().Format(time.RFC3339)

	addUserQuery := fmt.Sprintf("INSERT INTO PRODUCTS (ID,NAME,PRICE,IMAGEID,DESCRIPTION,CREATEDAT,UPDATEDAT) VALUES ('%s','%s','%f','%s','%s','%s','%s')", newuuid, name, price, imageid, description, currentTime, currentTime)
	_, er := db.Exec(addUserQuery)
	checkError(er)

	return newuuid
}

func DeleteProduct(id uuid.UUID) bool {
	addUserQuery := fmt.Sprintf("DELETE FROM PRODUCTS WHERE ID='%s'", id)
	_, er := db.Exec(addUserQuery)
	checkError(er)

	if er != nil {
		return false
	} else {
		return true
	}
}

func UpdateProduct(id uuid.UUID, name string, price float32, imageid string, description string) bool {
	var updateProductsQuery string = "UPDATE PRODUCTS SET"
	currentTime := time.Now().Format(time.RFC3339)

	if name != "" {
		updateProductsQuery = fmt.Sprintf("%s NAME='%s',", updateProductsQuery, name)
	}
	if price != 0 {
		updateProductsQuery = fmt.Sprintf("%s PRICE=%f,", updateProductsQuery, price)
	}
	if imageid != "" {
		updateProductsQuery = fmt.Sprintf("%s IMAGEID='%s',", updateProductsQuery, imageid)
	}
	if description != "" {
		updateProductsQuery = fmt.Sprintf("%s DESCRIPTION='%s',", updateProductsQuery, description)
	}

	updateProductsQuery = fmt.Sprintf("%s UPDATEDAT='%s' WHERE ID='%v'", updateProductsQuery, currentTime, id)

	// defer fmt.Printf(updateProductsQuery)

	// return true
	_, er := db.Exec(updateProductsQuery)
	checkError(er)

	if er != nil {
		return false
	} else {
		return true
	}
}

func ReadProducts() []Product {
	readQuery := fmt.Sprintf("SELECT ID,NAME,PRICE,IMAGEID,DESCRIPTION FROM PRODUCTS")

	R, er := db.Query(readQuery)
	checkError(er)
	defer R.Close()

	var productList []Product

	for R.Next() {
		var product Product
		R.Scan(&product.Id, &product.Name, &product.Price, &product.ImageId, &product.Description)
		// fmt.Print(R)
		productList = append(productList, product)
	}

	return productList
}

func ReadProduct(prodId string) Product {
	readQuery := fmt.Sprintf("SELECT ID,NAME,PRICE,IMAGEID,DESCRIPTION FROM PRODUCTS WHERE ID='%s'", prodId)

	R, er := db.Query(readQuery)
	checkError(er)
	defer R.Close()

	var product Product

	for R.Next() {
		R.Scan(&product.Id, &product.Name, &product.Price, &product.ImageId, &product.Description)
	}

	return product
}

func ReadFilteredProducts(prodIdList []string) []Product {
	readQuery := fmt.Sprintf("SELECT ID,NAME,PRICE,IMAGEID,DESCRIPTION FROM PRODUCTS WHERE ")

	for index, prodId := range prodIdList {
		if index == (len(prodIdList) - 1) {
			readQuery = fmt.Sprintf("%s ID='%v'", readQuery, prodId)
		} else {
			readQuery = fmt.Sprintf("%s ID='%v' OR", readQuery, prodId)
		}
	}

	// fmt.Print(prodIdList)

	R, er := db.Query(readQuery)
	checkError(er)
	defer R.Close()

	var productList []Product

	for R.Next() {
		var product Product
		R.Scan(&product.Id, &product.Name, &product.Price, &product.ImageId, &product.Description)
		// fmt.Print(R)
		productList = append(productList, product)
	}

	return productList
}
