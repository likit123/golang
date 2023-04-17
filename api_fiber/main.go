/*
http://localhost:3000/promotion
*/

package main

import (
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
)

var db1 *sqlx.DB

func main() {
	var err1 error
	db1, err1 = sqlx.Open("mysql", "_USER_:_PASSWORD_E@tcp(_IPADDRESS_:3306)/_DATABASE_")

	if err1 != nil {
		panic(err1)
	}

	app := fiber.New(fiber.Config{
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New())

	app.Get("/promotion", promotion)

	app.Listen(":3000")
}

func promotion(c *fiber.Ctx) error {
	results := []Line_result{}
	rows, _ := db1.Query("select id,promotion_name from promotion ")
	defer rows.Close()
	for rows.Next() {
		s := Line_result{}
		if err := rows.Scan(&s.Id, &s.Promotion_name); err != nil {
			return err
		}
		results = append(results, s)
	}
	return c.JSON(results)
}

type Line_result struct {
	Id             string `json:"id"`
	Promotion_name string `json:"promotion_name"`
}
