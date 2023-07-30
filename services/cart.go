package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

func AddCart() uuid.UUID {

	newuuid := uuid.New()
	currentTime := time.Now().Format(time.RFC3339)

	addCartQuery := fmt.Sprintf("INSERT INTO CART (ID,TOTALQUANTITY,CREATEDAT,UPDATEDAT) VALUES ('%s',0,'%s','%s')", newuuid, currentTime, currentTime)
	_, er := db.Exec(addCartQuery)
	checkError(er)

	if er != nil {
		return uuid.Nil
	}
	return newuuid
}

func UpdateCartAddProd(cartId string, prodId uuid.UUID) bool {
	currentTime := time.Now().Format(time.RFC3339)

	updateCartQuery := fmt.Sprintf("UPDATE CART SET PRODUCTID=PRODUCTID || ARRAY['%v'], UPDATEDAT='%s', TOTALQUANTITY=TOTALQUANTITY+1 WHERE ID='%s'", prodId, currentTime, cartId)
	_, er := db.Exec(updateCartQuery)
	checkError(er)

	if er != nil {
		return false
	} else {
		return true
	}

}

func UpdateCartRemoveProd(cartId string, prodId uuid.UUID) bool {
	currentTime := time.Now().Format(time.RFC3339)

	updateCartQuery := fmt.Sprintf("UPDATE CART SET PRODUCTID=ARRAY_REMOVE(PRODUCTID,'%v'), UPDATEDAT='%s', TOTALQUANTITY=TOTALQUANTITY-1 WHERE ID='%s'", prodId, currentTime, cartId)
	_, er := db.Exec(updateCartQuery)
	checkError(er)

	if er != nil {
		return false
	} else {
		return true
	}

}

func ReadProductsFromCart(cartId string) []Product {
	ProdIdListQuery := fmt.Sprintf("SELECT PRODUCTID FROM CART WHERE ID='%s'", cartId)
	R, er := db.Query(ProdIdListQuery)
	checkError(er)

	var ProdIdList []string

	for R.Next() {
		R.Scan(pq.Array(&ProdIdList))
		// fmt.Print(ProdIdList)
	}

	var prodList []Product

	prodList = ReadFilteredProducts(ProdIdList)

	return prodList
}

func ReadProductIdsFromCart(cartId string) []string {
	ProdIdListQuery := fmt.Sprintf("SELECT PRODUCTID FROM CART WHERE ID='%s'", cartId)
	R, er := db.Query(ProdIdListQuery)
	checkError(er)

	var ProdIdList []string

	for R.Next() {
		R.Scan(pq.Array(&ProdIdList))
		// fmt.Print(ProdIdList)
	}

	return ProdIdList
}
