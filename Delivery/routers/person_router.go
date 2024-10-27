package routers

import( 
	"github.com/gin-gonic/gin"
	"github.com/Afomiat/GoCrudChallange/Delivery/controllers"
	"github.com/Afomiat/GoCrudChallange/Usecase"
	"github.com/Afomiat/GoCrudChallange/Repository"
	"github.com/Afomiat/GoCrudChallange/Database"
)
func SetupPersonRouter(router *gin.Engine) {
	personRepo := Repository.NewPersonRepoImplement(Database.PersonColletion)
    personUsecase := Usecase.NewPersonUsecase(personRepo)
    personController := controllers.NewPersonController(personUsecase)

	personRoutes := router.Group("/person")
	personRoutes.GET("", personController.GetPersons)
	personRoutes.GET("/:id", personController.GetPersonById)
	personRoutes.POST("", personController.AddPerson)
	personRoutes.PUT("/:id", personController.UpdatePerson)
	personRoutes.DELETE(":id", personController.DeletePerson)
}