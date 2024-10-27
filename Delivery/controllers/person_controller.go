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

func (pc *PersonController) GetPersons(c *gin.Context) {
	persons, err := pc.PersonUsecase.GetPersons()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons) }

func (pc *PersonController) GetPersonById(c *gin.Context) {
	persons, err := pc.PersonUsecase.GetPersonById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons) }

func (pc *PersonController) UpdatePerson(c *gin.Context){
	id := c.Param("id")
	var person Domain.Person

	if err := c.ShouldBindJSON(&person); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.PersonUsecase.UpdatePerson(id, person); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person updated successfully"})
}

func (pc *PersonController) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	person, err := pc.PersonUsecase.GetPersonById(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if _,err := pc.PersonUsecase.DeletePerson(id); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "person deleted successfully", "task": person})

}