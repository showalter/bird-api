package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Bird struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	app := fiber.New()

	buf, err := ioutil.ReadFile("birds.json")
	if err != nil {
		fmt.Print(err)
	}

	s := string(buf)

	var birds []Bird
	json.Unmarshal([]byte(s), &birds)

	fmt.Printf("Birds : %+v", birds)

	randSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randSource)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(birds[random.Intn(len(birds))])
	})

	app.Listen(":8080")
}
