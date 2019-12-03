package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/aizu-garbage/usecase"
)

type controller struct {
	scheduleCase *usecase.ScheduleUseCase
	placeCase    *usecase.PlaceUseCase
	deleteKey    string
}

func NewController(sc *usecase.ScheduleUseCase, pc *usecase.PlaceUseCase, deleteKey string) *controller {
	return &controller{
		scheduleCase: sc,
		placeCase:    pc,
		deleteKey:    deleteKey,
	}
}

func (c *controller) Listen(port string) {
	router := gin.Default()

	router.GET("/places", c.getPlaces)
	router.GET("/calendar/:id", c.getCallender)
	router.DELETE("/schedule", c.deleteSchedule)

	router.Run(port)
}

func (c *controller) getPlaces(ctx *gin.Context) {
	places, err := c.placeCase.GetAll(ctx.Request.Context())
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, places)
}

func (c *controller) getCallender(ctx *gin.Context) {
	id := ctx.Param("id")
	place, err := c.placeCase.Get(ctx.Request.Context(), id)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	if place == nil && err == nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	schedules, err := c.scheduleCase.Get(ctx.Request.Context(), place)
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	resSchedules := make([]*responseSchedule, len(schedules))
	for i, schedule := range schedules {
		resSchedules[i] = convertSchedule(schedule)
	}

	calendar := &calendar{
		Place:     place,
		Schedules: resSchedules,
	}

	ctx.JSON(http.StatusOK, calendar)
}

func (c *controller) deleteSchedule(ctx *gin.Context) {
	key := ctx.Query("key")

	log.Printf("collectKey = %s", c.deleteKey)
	log.Printf("sendKey = %s", key)

	if key != c.deleteKey {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Key did not match"})
		return
	}

	err := c.scheduleCase.DeleteAll(ctx.Request.Context())
	if err != nil {
		log.Fatal(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
