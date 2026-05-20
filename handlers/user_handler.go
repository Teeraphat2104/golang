package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"hello/models"
)

// Mock data
var mockUsers = []models.User{
	{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	},
	{
		ID:    2,
		Name:  "Jane Smith",
		Email: "jane@example.com",
	},
	{
		ID:    3,
		Name:  "Bob Johnson",
		Email: "bob@example.com",
	},
}

var nextID = 4

// GetAllUsers retrieves all users
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": mockUsers,
		"total": len(mockUsers),
	})
}

// GetUserByID retrieves a single user by ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	for _, user := range mockUsers {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser := models.User{
		ID:    nextID,
		Name:  req.Name,
		Email: req.Email,
	}

	mockUsers = append(mockUsers, newUser)
	nextID++

	c.JSON(http.StatusCreated, gin.H{
		"data": newUser,
		"message": "User created successfully",
	})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, user := range mockUsers {
		if user.ID == id {
			if req.Name != "" {
				mockUsers[i].Name = req.Name
			}
			if req.Email != "" {
				mockUsers[i].Email = req.Email
			}
			c.JSON(http.StatusOK, gin.H{
				"data": mockUsers[i],
				"message": "User updated successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	for i, user := range mockUsers {
		if user.ID == id {
			mockUsers = append(mockUsers[:i], mockUsers[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}
