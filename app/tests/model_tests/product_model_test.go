package model_tests

import (
	"Rest/app/models"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
)

func TestGetroducts(t *testing.T) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = seedUsersAndProducts()
	if err != nil {
		log.Fatalf("Error seeding user and product  table %v\n", err)
	}
	var productTest models.Product
	posts, err := productTest.GetProducts(appTest.DB)
	if err != nil {
		t.Errorf("this is the error getting the products: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 2)
}

func TestCreateProduct(t *testing.T) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newProduct := models.Product{
		ID:       1,
		Name:    "Milk",
		Price:  12,
		UserID: int(user.ID),
	}
	savedProd, err := newProduct.CreateProduct(appTest.DB)
	if err != nil {
		t.Errorf("this is the error getting the product: %v\n", err)
		return
	}
	assert.Equal(t, newProduct.ID, savedProd.ID)
	assert.Equal(t, newProduct.Name, savedProd.Name)
	assert.Equal(t, newProduct.Price, savedProd.Price)
	assert.Equal(t, newProduct.UserID, savedProd.UserID)

}

func TestGetProductByID(t *testing.T) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}
	p, err := seedOneUserAndOneProduct()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundProduct, err := p.GetProduct(appTest.DB)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundProduct.ID, p.ID)
	assert.Equal(t, foundProduct.Name, p.Name)
	assert.Equal(t, foundProduct.Price, p.Price)
}

func TestUpdateAProduct(t *testing.T) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}
	p, err := seedOneUserAndOneProduct()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	prodUpdate := models.Product{
		ID:       p.ID,
		Name:    "apple",
		Price:  30,
		UserID: p.UserID,
	}

	updatedProd, err := prodUpdate.UpdateProduct(appTest.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedProd.ID, prodUpdate.ID)
	assert.Equal(t, updatedProd.Name, prodUpdate.Name)
	assert.Equal(t, updatedProd.Price, prodUpdate.Price)
	assert.Equal(t, updatedProd.UserID, prodUpdate.UserID)
}

func TestDeleteAProduct(t *testing.T) {

	_,err := appTest.DB.Exec("truncate table users")
	if err != nil {
		log.Fatal(err)
	}
	_,err = appTest.DB.Exec("truncate table products")
	if err != nil {
		log.Fatal(err)
	}



	p, err := seedOneUserAndOneProduct()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := p.DeleteProduct(appTest.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	//one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	//Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}
