package handler

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerLogger возвращает middleware, который логирует имя обработчика, метод и время.
func HandlerLogger(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // выполняем следующий обработчик
		log.Printf("Хендлер: %s, Метод: %s, Время: %v", name, c.Request.Method, time.Since(start))
	}
}
