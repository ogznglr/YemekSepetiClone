package main

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	readDuration, err := strconv.ParseInt(os.Getenv("GOBIT_SERVER_REAT_TIMEOUT"), 10, 64)

	if err != nil {
		panic(err)
	}

	writeDuration, err := strconv.ParseInt(os.Getenv("GOBIT_SERVER_WRITE_TIMEOUT"), 10, 64)

	if err != nil {
		panic(err)
	}

	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readDuration),
		WriteTimeout: time.Second * time.Duration(writeDuration),
		BodyLimit:    50 * 1024 * 1024,
	}
}

func StartServerWithGracefulShotdown() {

}
