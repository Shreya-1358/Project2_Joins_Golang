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

	r.GET("Students", studContoller.GetStudents)
	r.GET("Student", studContoller.GetStudentsWithCourses)
	r.GET("Student/:id", studContoller.GetStudentById)
	r.GET("Students/:id", studContoller.GetStudentByCourseId)
	r.POST("Student", studContoller.CreateStudent)

	return r
}
