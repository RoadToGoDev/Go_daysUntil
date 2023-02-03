package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {

	s := echo.New()

	s.Use(Wrapper)

	s.GET("status", Handler)

	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
func Handler(ctx echo.Context) error {
	d := time.Date(2023, time.July, 28, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)

	s := fmt.Sprintf("Количество дней до моего дня рождения: %d", int64(dur.Hours())/24)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func Wrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-role")

		if val == "Hr" {
			log.Println("Get me a JOB")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
