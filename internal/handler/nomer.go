package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Futturi/testovoe/internal/models"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func (h *Handler) GetNomers(c *gin.Context) {
	querys := make(map[string]string)
	regnum := c.Query("regnum")
	if regnum != "" {
		querys["regNum"] = regnum
	}
	mark := c.Query("mark")
	if mark != "" {
		querys["mark"] = mark
	}
	model := c.Query("model")
	if model != "" {
		querys["model"] = model
	}
	page := c.Param("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		slog.Error("incorrect page", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	name := c.Query("name")
	if name != "" {
		querys["name"] = name
	}
	surname := c.Query("surname")
	if surname != "" {
		querys["surname"] = surname
	}
	cars, err := h.service.GetNomers(pageInt, querys)
	if err != nil {
		slog.Error("error with getting nomers", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, cars)
}

func (h *Handler) Delete(c *gin.Context) {
	var nomer models.Car
	err := c.BindJSON(&nomer)
	if err != nil {
		slog.Error("incorrect nomer", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	err = h.service.Delete(nomer)
	if err != nil {
		slog.Error("error", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
	slog.Info("success deleted nomer", slog.String("nomer", nomer.Nomer))
}

func (h *Handler) Update(c *gin.Context) {
	var nomer models.CarUpdate
	if err := c.BindJSON(&nomer); err != nil {
		slog.Error("error with getting nomers", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	err := h.service.Put(nomer)
	if err != nil {
		slog.Error("error", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) Create(c *gin.Context) {
	var nomers []models.Car
	var new []models.Car
	worker := http.Client{Timeout: 10 * time.Second}
	if err := c.BindJSON(&nomers); err != nil {
		slog.Error("error with getting nomers", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	jsonBody := `"regNums":[ `
	for ind, value := range nomers {
		if ind != len(nomers)-1 {
			jsonBody += fmt.Sprintf(`"%s"`, value.Nomer)
		} else {
			jsonBody += fmt.Sprintf(`"%s"]`, value.Nomer)
		}
	}
	bodyReader := bytes.NewReader([]byte(jsonBody))
	req, err := http.NewRequest("POST", fmt.Sprintf("localhost:%s/info", os.Getenv("ANOTHER_API")), bodyReader)
	if err != nil {
		slog.Error("error with getting info", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	resp, err := worker.Do(req)
	if err != nil {
		slog.Error("error with getting info", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error with getting info", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	err = json.Unmarshal(byt, &new)
	if err != nil {
		slog.Error("error with marshalling info", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	for i := 0; i < len(new); i++ {
		new[i] = models.Car{
			Nomer:   nomers[i].Nomer,
			Mark:    new[i].Mark,
			Model:   new[i].Model,
			Name:    new[i].Name,
			Surname: new[i].Surname,
		}
	}
	err = h.service.Insert(new)
	if err != nil {
		slog.Error("error with marshalling info", slog.Any("error", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
