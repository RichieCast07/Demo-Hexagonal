package controllers

import (
	"net/http"
	"strconv"
	"time"

	application "demo/src/users/application_users"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	useCase *application.GetAllUsersUseCase
}

func NewGetAllUsersController(uc *application.GetAllUsersUseCase) *GetAllUsersController {
	return &GetAllUsersController{useCase: uc}
}

func (controller *GetAllUsersController) Execute(c *gin.Context) {
	results, err := controller.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo obtener la lista de usuarios",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": results,
	})
}

func (controller *GetAllUsersController) LongPollingExecute(c *gin.Context) {
	lastUpdateStr := c.Query("lastUpdate")
	var clientLastUpdate int64 = 0
	if lastUpdateStr != "" {
		var err error
		clientLastUpdate, err = strconv.ParseInt(lastUpdateStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parametro 'lastUpdate' invalido"})
			return
		}
	}

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			results, err := controller.useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
				return
			}
			var maxUpdated int64 = clientLastUpdate
			for _, user := range results {
				if user.UpdatedAt > maxUpdated {
					maxUpdated = user.UpdatedAt
				}
			}
			if maxUpdated > clientLastUpdate {
				c.JSON(http.StatusOK, gin.H{
					"users":      results,
					"lastUpdate": maxUpdated,
				})
				return
			}
		case <-timeout:
			c.JSON(http.StatusOK, gin.H{
				"message":    "No hay cambios",
				"lastUpdate": clientLastUpdate,
			})
			return
		}
	}
}
