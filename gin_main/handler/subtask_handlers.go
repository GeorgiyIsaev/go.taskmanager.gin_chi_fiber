package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go.taskmanager.gin_chi_fiber/model"
	"go.taskmanager.gin_chi_fiber/repository"
)

// CreateSubtaskHandler создаёт подзадачу (POST /taskmanager/subtask/create).
func CreateSubtaskHandler(c *gin.Context) {
	var req CreateSubtaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if req.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Название подзадачи не может быть пустым"})
		return
	}
	if req.EpicID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Указан некорректный ID эпика"})
		return
	}

	sub := &model.Subtask{
		Title:       req.Title,
		Description: req.Description,
		Status:      parseStatus(req.Status),
		EpicID:      req.EpicID,
	}
	if err := repository.AddSubtask(sub); err != nil {
		log.Printf("Ошибка создания подзадачи: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Внутренняя ошибка сервера"})
		return
	}

	c.JSON(http.StatusCreated, SubtaskResponse{
		ID:          sub.ID,
		Title:       sub.Title,
		Description: sub.Description,
		Status:      sub.Status.String(),
		EpicID:      sub.EpicID,
	})
}

// UpdateSubtaskHandler обновляет подзадачу (PUT /taskmanager/subtask/update).
func UpdateSubtaskHandler(c *gin.Context) {
	var req UpdateSubtaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	if req.ID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID подзадачи"})
		return
	}
	if req.Title == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Название не может быть пустым"})
		return
	}

	sub := &model.Subtask{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      parseStatus(req.Status),
	}
	if err := repository.UpdateSubtask(sub); err != nil {
		log.Printf("Ошибка обновления подзадачи: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Подзадача обновлена"})
}

// DeleteSubtaskHandler удаляет подзадачу (DELETE /taskmanager/subtask/delete?id=...).
func DeleteSubtaskHandler(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID подзадачи"})
		return
	}
	if err := repository.DeleteSubtask(id); err != nil {
		log.Printf("Ошибка удаления подзадачи %d: %v", id, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Подзадача удалена"})
}
