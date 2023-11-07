package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/robfig/cron"
)

func main() {
	cronJob()
	initFiber()
}

func cronJob() {
	c := cron.New()

	// @every 00h00m10s
	c.AddFunc("@every 00h00m10s", cronStart)
	c.Start()

}
func cronStart() {
	res, err := http.Get("http://localhost:8080/cron")
	if err != nil {
		log.Fatalf("error for geting response: %s",err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error geting body: %s",err)
	}

	log.Println(string(body))
}

func initFiber() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/cron", getMessage)
	app.Listen(":8080")
}

func getMessage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Hellooo"})
}
