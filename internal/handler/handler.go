package handler

import (
	"cripta_course_work/internal/model"
	"cripta_course_work/internal/service"
)

type Handler struct {
	services *service.Service

	//TODO КОСТЫЛИ ЕБУЧИЕ, ЗА ЭТО ТЫ ПОПАДЁШЬ В АД!
	questionOfCurrentUser *model.ShitUser
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

/*
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Handle("/create")

	//На будущее для разработки нового api
	//api := router.Group("/api")
	//{
	//
	//}
	return router
}

*/

//TODO: переделать на gorilla/mux
