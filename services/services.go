package services

import (
	"errors"
	"time"
	"github.com/google/uuid"
	"github.com/yujen77300/curd-practice/models"
)

//  storage is created to store the item’s data.
var storage []models.Item = []models.Item{}


func GetAllItems() []models.Item {
		return storage
}

// GetItemByID returns get item's data by ID
func GetItemByID(id string) (models.Item, error){
	 for _,item := range storage{
		 if item.ID == id{
			 return item,nil
		 }
	 }
	 return models.Item{}, errors.New("item not found")
}

// CreateItem returns created item in the storage
func CreateItem(itemRequest models.ItemRequest) models.Item {
    // create a new item
    var newItem models.Item = models.Item{
        ID:        uuid.New().String(),
        Name:      itemRequest.Name,
        Price:     itemRequest.Price,
        Quantity:  itemRequest.Quantity,
        CreatedAt: time.Now(),
    }

    // store the created item into storage
    storage = append(storage, newItem)

    // return the item that already created
    return newItem
}

// UpdateItem returns updated item
func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
    // iterate through all items
    for index, item := range storage {
        // if item is found
        if item.ID == id {
            // update the item's data
            item.Name = itemRequest.Name
            item.Price = itemRequest.Price
            item.Quantity = itemRequest.Quantity
            item.UpdatedAt = time.Now()

            storage[index] = item

            // return the updated item
            return item, nil
        }
    }

    // return error if update is failed
    return models.Item{}, errors.New("item update failed, item not found")
}

// DeleteItem returns deletion result
func DeleteItem(id string) bool {
    // create a new slice for storing items data
    // after deletion
    var newItems []models.Item = []models.Item{}

    // iterate through all items
    for _, item := range storage {
        // if current item's ID is not equal ID in parameter
        // insert the item into "newItems" slice
        if item.ID != id {
            newItems = append(newItems, item)
        }
    }

    // assign the "newItems" into storage
    storage = newItems

    return true
}