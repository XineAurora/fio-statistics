package api

import (
	"net/http"
	"os"
	"strconv"

	"github.com/XineAurora/fio-statistics/intrernal/database"
	"github.com/XineAurora/fio-statistics/intrernal/entities"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ApiServer(repo database.FIORepository) *http.Server {
	return &http.Server{
		Addr:    ":" + os.Getenv("API_PORT"),
		Handler: setupRouter(repo),
	}
}

func setupRouter(repo database.FIORepository) *gin.Engine {
	r := gin.Default()

	fioRoutes := r.Group("/api/fio")
	{
		fioRoutes.GET("/", getFIOs(repo))
		fioRoutes.POST("/", createFIO(repo))
		fioRoutes.DELETE("/", deleteFIO(repo))
		fioRoutes.PUT("/", updateFIO(repo))
	}

	return r
}

func getFIOs(repo database.FIORepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fiter database.FIOFilter
		if err := c.BindWith(&fiter, binding.JSON); err != nil {
			abortWithError(c, http.StatusBadRequest, "bad body")
		}
		page, _ := strconv.Atoi(c.Query("page"))
		onPage, _ := strconv.Atoi(c.Query("onpage"))
		fios, err := repo.GetFIOs(fiter, database.Pagination{Page: page, OnPage: onPage})
		if err != nil {
			abortWithError(c, http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, fios)
	}
}

func createFIO(repo database.FIORepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fio entities.FIO
		if err := c.BindWith(&fio, binding.JSON); err != nil {
			abortWithError(c, http.StatusBadRequest, "bad body")
		}
		if err := fio.Validate(); err != nil {
			abortWithError(c, http.StatusBadRequest, err.Error())
		}
		fio, err := repo.CreateFIO(fio)
		if err != nil {
			abortWithError(c, http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, fio)
	}
}

func deleteFIO(repo database.FIORepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id_str := c.Param("id")
		id, err := strconv.ParseUint(id_str, 10, 32)
		if err != nil || id <= 0 {
			abortWithError(c, http.StatusBadRequest, "id must be positive integer")
		}
		if err := repo.DeleteFIO(uint(id)); err != nil {
			abortWithError(c, http.StatusInternalServerError, err.Error())
		}
		c.Status(http.StatusOK)
	}
}

func updateFIO(repo database.FIORepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fio entities.FIO
		if err := c.BindWith(&fio, binding.JSON); err != nil {
			abortWithError(c, http.StatusBadRequest, "bad body")
		}
		if err := fio.Validate(); err != nil {
			abortWithError(c, http.StatusBadRequest, err.Error())
		}
		if fio.ID == 0 {
			abortWithError(c, http.StatusBadRequest, "id must be positive integer")
		}
		if err := repo.UpdateFIO(fio); err != nil {
			abortWithError(c, http.StatusInternalServerError, err.Error())
		}
		c.Status(http.StatusOK)
	}
}

func abortWithError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code":    status,
		"message": message,
	})
}
