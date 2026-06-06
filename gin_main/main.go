package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.taskmanager.gin_chi_fiber/gin_main/handler"
	"go.taskmanager.gin_chi_fiber/repository"
)

func main() {
	// 1. Конфиг и БД
	cfg, err := repository.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}
	if err := repository.InitDB(cfg.DatabaseURL()); err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
	defer repository.Pool.Close()

	// 2. Создаём роутер Gin
	r := gin.Default() // уже включает Logger() и Recovery()

	// 3. Настраиваем CORS глобально (из библиотеки gin-contrib/cors)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: false,
	}))

	// Группа маршрутов taskmanager
	tm := r.Group("/taskmanager")
	{
		// Эпики
		tm.GET("/task", handler.GetEpicHandler)              // GET с query id
		tm.GET("/tasks", handler.GetAllTasksHandler)         // GET все эпики
		tm.POST("/tasks", handler.CreateEpicHandler)         // POST создание эпика
		tm.PUT("/epic/update", handler.UpdateEpicHandler)    // PUT обновление эпика
		tm.DELETE("/epic/delete", handler.DeleteEpicHandler) // DELETE удаление эпика

		// Подзадачи
		tm.POST("/subtask/create", handler.CreateSubtaskHandler)   // POST создание подзадачи
		tm.PUT("/subtask/update", handler.UpdateSubtaskHandler)    // PUT обновление подзадачи
		tm.DELETE("/subtask/delete", handler.DeleteSubtaskHandler) // DELETE удаление подзадачи
	}

	// Если нужен кастомный логер с именем обработчика – оборачиваем:
	// tm.GET("/task", handler.HandlerLogger("GetEpicHandler"), handler.GetEpicHandler)
	// Но стандартный gin.Logger() уже даёт много информации.

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
