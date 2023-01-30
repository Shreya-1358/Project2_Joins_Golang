package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"project_2/config"
	model "project_2/model"
	"time"
)

type StudentDao interface {
	GetAllStudents() ([]*model.StudentResponse, error)
	GetStudentsWithCourses() ([]*model.StudentResponse, error)
	CreateStudent(*model.Student) error
	GetStudentById(string) ([]*model.StudentResponse, error)
}

type StudentDaoImpl struct {
}

func NewStudentDaoImpl() *StudentDaoImpl {
	return &StudentDaoImpl{}
}

func (studdao *StudentDaoImpl) GetAllStudents() ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,COALESCE(Student.course.courseID,-1),COALESCE(courseName,''),COALESCE(courseFee,-1) FROM Student.student left join Student.enrollment on Student.student.studentID = Student.enrollment.studentID  left join Student.course on Student.enrollment.courseID = Student.course.courseID")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Course.CourseID, &stud.Course.CourseName, &stud.Course.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentsWithCourses() ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,Student.course.courseID,courseName,courseFee FROM Student.student  join Student.enrollment on Student.student.studentID = Student.enrollment.studentID   join Student.course on Student.enrollment.courseID = Student.course.courseID")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Course.CourseID, &stud.Course.CourseName, &stud.Course.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentById(id string) ([]*model.StudentResponse, error) {

	var studentlist []*model.StudentResponse = []*model.StudentResponse{}
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,COALESCE(Student.course.courseID,-1),COALESCE(courseName,''),COALESCE(courseFee,-1) FROM Student.student  left join Student.enrollment on Student.student.StudentID = Student.enrollment.studentID left  join Student.course on Student.enrollment.courseID = Student.course.courseID where Student.student.studentID =?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &stud.Course.CourseID, &stud.Course.CourseName, &stud.Course.CourseFee)
		if err != nil {
			return nil, err
		} else {
			studentlist = append(studentlist, &stud)

		}
	}

	return studentlist, nil
}
func (studdao *StudentDaoImpl) CreateStudent(student *model.Student) (err error) {
	dateString := student.DOB
	date, errs := time.Parse("2006-01-02", dateString)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	courses := student.Course

	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.Exec("INSERT INTO Student.student(studentID,Name,Email,Dept,DOB,PhoneNo)VALUES(?,?,?,?,?,?)", student.StudentID, student.Name, student.Email, student.Dept, date, student.PhoneNo)
	if err != nil {

		return err
	}

	for _, course := range courses {
		Id := uuid.New()
		fmt.Print(Id)
		if err != nil {

			return err
		}

		_, err = tx.Exec("INSERT INTO Student.course(courseID,courseName,courseFee)VALUES(?,?,?)", course.CourseID, course.CourseName, course.CourseFee)
		if err != nil {

			return err
		}
		_, err = tx.Exec("INSERT INTO Student.enrollment(Id,studentID,courseID)VALUES(?,?,?)", Id, student.StudentID, course.CourseID)
		if err != nil {

			return err
		}
	}

	return nil

}
