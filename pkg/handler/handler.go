package handler

import (
	"github.com/bigbob004/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Когда-нибудь здесь будет красивый сайт, а пока почитайте стихи Пушкина\n\n"+
			"Я помню чудное мгновенье:\nПередо мной явилась ты,\nКак мимолетное виденье,\nКак гений чистой красоты.\n\nВ томленьях грусти безнадежной,\nВ тревогах шумной суеты,\nЗвучал мне долго голос нежный\nИ снились милые черты.\n\nШли годы. Бурь порыв мятежный\nРассеял прежние мечты,\nИ я забыл твой голос нежный,\nТвои небесные черты.\n\nВ глуши, во мраке заточенья\nТянулись тихо дни мои\nБез божества, без вдохновенья,\nБез слез, без жизни, без любви.\n\nДуше настало пробужденье:\nИ вот опять явилась ты,\nКак мимолетное виденье,\nКак гений чистой красоты.\n\nИ сердце бьется в упоенье,\nИ для него воскресли вновь\nИ божество, и вдохновенье,\nИ жизнь, и слезы, и любовь.")
	})

	//На будущее для разработки нового api
	//api := router.Group("/api")
	//{
	//
	//}
	return router
}
