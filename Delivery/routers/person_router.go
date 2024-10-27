package routers

import( 
	"github.com/gin-gonic/gin"
	"github.com/Afomiat/GoCrudChallange/Delivery/controllers"
	"github.com/Afomiat/GoCrudChallange/Usecase"
	"github.com/Afomiat/GoCrudChallange/Repository"
)
func SetupPersonRouter(router *gin.Engine) {
	personRepo := Repository.NewSongRepoImplement(Database.personCollection)
    personUsecase := Usecase.NewSongUsecase(personRepo)
    personController := controllers.NewSongController(personUsecase)

	personRouter := router.Group("/person")
	personRouter.GET("/person", PersonController.GetPerson)
	personRouter.GET("/person/:id", PersonController.GetPersonById)
	personRouter.POST("/person", PersonController.CreatePerson)
	personRouter.PUT("/person/:id", PersonController.UpdatePerson)
	personRouter.DELETE("/person/:id", PersonController.DeletePerson)
}