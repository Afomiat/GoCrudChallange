package controllers

import (
	"log"
	"net/http"

	"github.com/Afomiat/GoCrudChallange/Usecase"
	"github.com/Afomiat/GoCrudChallange/Domain"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	PersonUsecase *Usecase.PersonUsecase
}

func NewPersonController(personUsecase *Usecase.PersonUsecase) *PersonController {
	return &PersonController{
		PersonUsecase: personUsecase,
	}
}

func (pc *PersonController) AddPerson(c *gin.Context) {
	var person Domain.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		log.Println("Error binding JSON:", err) 

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received Song: %+v\n", person) 

	if err := pc.PersonUsecase.AddPerson(person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "person added successfully"})

}
