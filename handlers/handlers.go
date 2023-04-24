package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yujen77300/curd-practice/models"
	"github.com/yujen77300/curd-practice/services"
	"github.com/yujen77300/curd-practice/utils"
)

// GetAllItems returns all items from the storage
func GetAllItems(c *fiber.Ctx) error {
	// get all items
	var items []models.Item = services.GetAllItems()

	// return the response
	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items data",
		Data:    items,
	})
}

// GetItemByID returns item's data by ID
func GetItemByID(c *fiber.Ctx) error {
	// get the id from the request parameter
	var itemID string = c.Params("id")

	// get the item by ID
	item, err := services.GetItemByID(itemID)

	// if error is exists, return the error response
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// return the item
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item found",
		Data:    item,
	})
}

// CreateItem returns created item in a storage
func CreateItem(c *fiber.Ctx) error {
	// check the token
	isValid, err := utils.CheckToken(c)

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// create a variable to store the request
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	// parse the request into "itemInput" variable
	if err := c.BodyParser(itemInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := itemInput.ValidateStruct()

	// if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	// create a new item
	var createdItem models.Item = services.CreateItem(*itemInput)

	// return the created item in a storage
	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})
}

// UpdateItem returns updated item
func UpdateItem(c *fiber.Ctx) error {
	// create a variable to store the request
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	// parse the request into "itemInput" variable
	if err := c.BodyParser(itemInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := itemInput.ValidateStruct()

	// if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	// get the item's ID from the request parameter
	var itemID string = c.Params("id")

	// update the item's data
	updatedItem, err := services.UpdateItem(*itemInput, itemID)
	if err != nil {
		// if update is failed, return an error
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// return the updated item
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item updated",
		Data:    updatedItem,
	})
}

// DeleteItem returns deletion result
func DeleteItem(c *fiber.Ctx) error {
	// get the item's ID from the request parameter
	var itemID string = c.Params("id")

	// delete the item data
	var result bool = services.DeleteItem(itemID)

	if result {
		// if successful, return the result
		return c.JSON(models.Response[any]{
			Success: true,
			Message: "item deleted",
		})
	}

	// return the failed result
	return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "item failed to delete",
	})
}
