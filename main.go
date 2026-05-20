package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://admin:210447@localhost:5432/mydb?sslmode=disable",
	)

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL Connected")

	// Create Gin Router
	r := gin.Default()

	// API Route
	r.GET("/users", func(c *gin.Context) {

		rows, err := conn.Query(
			context.Background(),
			"SELECT id, name, email FROM users",
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer rows.Close()

		var users []gin.H

		for rows.Next() {

			var id int
			var name string
			var email string

			err := rows.Scan(&id, &name, &email)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			users = append(users, gin.H{
				"id":    id,
				"name":  name,
				"email": email,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start Server
	r.Run(":3000")
}
