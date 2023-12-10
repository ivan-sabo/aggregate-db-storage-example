package user

import (
	"aggregate-db-storage-example/helper"
	"testing"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db := helper.GetTestDBConnection()
	db.AutoMigrate(&User{}, &Address{})
	db.Debug()

	testUser := getTestUser()
	r := db.Create(&testUser)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	var user User
	r = db.Model(&User{}).Preload("AddressCollection").Find(&user)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	testAddressCollection := getTestAddressCollection()
	for i := range user.AddressCollection {
		if user.AddressCollection[i].City != testAddressCollection[i].City {
			t.Errorf("wanted: %v, got: %v", testAddressCollection[i].City, user.AddressCollection[i].City)
		}
	}

	helper.CleanTestDB()
}

func TestUpdate(t *testing.T) {
	db := helper.GetTestDBConnection()
	db.AutoMigrate(&User{}, &Address{})
	db.Debug()

	testUser := getTestUser()
	r := db.Create(&testUser)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	var user User
	r = db.Model(&User{}).Preload("AddressCollection").Find(&user)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	user.AddressCollection[0].City = "Frankfurt"
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	r = db.Model(&User{}).Preload("AddressCollection").Find(&user)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if user.AddressCollection[0].City != "Frankfurt" {
		t.Errorf("expected: Frankfurt, got: %v", user.AddressCollection[0].City)
	}

	helper.CleanTestDB()
}

func getTestUser() User {
	return User{Name: "John", AddressCollection: getTestAddressCollection()}
}

func getTestAddressCollection() []Address {
	return []Address{
		{
			City: "Berlin",
		},
		{
			City: "Bremen",
		},
	}
}
