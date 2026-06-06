package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go.taskmanager.gin_chi_fiber/model"
	"go.taskmanager.gin_chi_fiber/repository"
)

// GetEpicHandler возвращает один эпик с подзадачами (GET /taskmanager/task?id=...).
func GetEpicHandler(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Не указан id эпика"})
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id должен быть числом"})
		return
	}

	epic, err := repository.GetEpic(id)
	if err != nil {
		log.Printf("Ошибка получения Epic id=%d: %v", id, err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Эпик не найден"})
		return
	}

	subtasks, err := repository.GetSubtasksByEpicID(id)
	if err != nil {
		log.Printf("Ошибка получения подзадач для Epic id=%d: %v", id, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке подзадач"})
		return
	}

	subResponses := make([]SubtaskResponse, 0, len(subtasks))
	for _, s := range subtasks {
		subResponses = append(subResponses, SubtaskResponse{
			ID:          s.ID,
			Title:       s.Title,
			Description: s.Description,
			Status:      s.Status.String(),
			EpicID:      s.EpicID,
		})
	}

	c.JSON(http.StatusOK, GetEpicResponse{
		ID:          epic.ID,
		Title:       epic.Title,
		Description: epic.Description,
		Status:      epic.Status.String(),
		Subtasks:    subResponses,
	})
}

// GetAllTasksHandler возвращает все эпики с подзадачами (GET /taskmanager/tasks).
func GetAllTasksHandler(c *gin.Context) {
	epics, err := repository.GetAllEpics()
	if err != nil {
		log.Printf("Ошибка получения списка эпиков: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	var response []GetEpicResponse
	for _, epic := range epics {
		subtasks, err := repository.GetSubtasksByEpicID(epic.ID)
		if err != nil {
			log.Printf("Ошибка получения подзадач для эпика %d: %v", epic.ID, err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки подзадач"})
			return
		}
		subResponses := make([]SubtaskResponse, 0, len(subtasks))
		for _, s := range subtasks {
			subResponses = append(subResponses, SubtaskResponse{
				ID:          s.ID,
				Title:       s.Title,
				Description: s.Description,
				Status:      s.Status.String(),
				EpicID:      s.EpicID,
			})
		}
		response = append(response, GetEpicResponse{
			ID:          epic.ID,
			Title:       epic.Title,
			Description: epic.Description,
			Status:      epic.Status.String(),
			Subtasks:    subResponses,
		})
	}

	c.JSON(http.StatusOK, response)
}

// CreateEpicHandler создаёт новый эпик (POST /taskmanager/tasks).
func CreateEpicHandler(c *gin.Context) {
	var req CreateEpicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if req.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Название задачи не может быть пустым"})
		return
	}

	epic := &model.Epic{
		Title:       req.Title,
		Description: req.Description,
		Status:      model.StatusNew,
	}
	if err := repository.CreateEpic(epic); err != nil {
		log.Printf("Ошибка при сохранении Epic: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusCreated, CreateEpicResponse{
		ID:          epic.ID,
		Title:       epic.Title,
		Description: epic.Description,
		Status:      epic.Status.String(),
	})
}

// UpdateEpicHandler обновляет эпик (PUT /taskmanager/epic/update).
func UpdateEpicHandler(c *gin.Context) {
	var req UpdateEpicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if req.ID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID эпика"})
		return
	}
	if req.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Название не может быть пустым"})
		return
	}

	epic := &model.Epic{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      parseStatus(req.Status),
	}
	if err := repository.UpdateEpic(epic); err != nil {
		log.Printf("Ошибка обновления эпика: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Эпик обновлён"})
}

// DeleteEpicHandler удаляет эпик (DELETE /taskmanager/epic/delete?id=...).
func DeleteEpicHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID эпика"})
		return
	}
	if err := repository.DeleteEpic(id); err != nil {
		log.Printf("Ошибка удаления эпика %d: %v", id, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Эпик и все его подзадачи удалены"})
}
