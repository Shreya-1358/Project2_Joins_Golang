package routes

import (
	"github.com/gin-gonic/gin"
	"project_2/controller"
	"project_2/dao"
	"project_2/service"
)

func SetupRouter() *gin.Engine {

	studentDao := dao.NewStudentDaoImpl()
	studentService := service.NewStudentServiceImpl(studentDao)
	studContoller := controller.NewStudentController(studentService)

	r := gin.Default()

	r.GET("students", studContoller.GetAllStudents)
	r.GET("student/:id", studContoller.GetStudentById)
	r.GET("course/:id", studContoller.GetStudentByCourseId)
	r.POST("student", studContoller.CreateStudent)

	return r
}
