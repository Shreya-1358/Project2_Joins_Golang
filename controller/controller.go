package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project_2/model"
	"project_2/service"
)

type StudentController struct {
	studentservice service.StudentService
}

func NewStudentController(studsvc service.StudentService) *StudentController {
	return &StudentController{studentservice: studsvc}
}

func (studcont *StudentController) GetAllStudents(c *gin.Context) {

	student, err := studcont.studentservice.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) GetStudentById(c *gin.Context) {
	id := c.Param("id")

	student, err := studcont.studentservice.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) GetStudentByCourseId(c *gin.Context) {
	id := c.Param("id")

	student, err := studcont.studentservice.GetStudentByCourseId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
func (studcont *StudentController) CreateStudent(c *gin.Context) {
	var student model.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(student)
	err := studcont.studentservice.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, student)
	}
}
