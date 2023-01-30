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
	GetAllStudents() ([]*model.Student, error)
	GetStudentsWithCourses() ([]*model.Student, error)
	CreateStudent(*model.Student) error
	GetStudentById(string) ([]*model.Student, error)
	GetStudentByCourseId(string) ([]*model.Student, error)
}

type StudentDaoImpl struct {
}

func NewStudentDaoImpl() *StudentDaoImpl {
	return &StudentDaoImpl{}
}

func (studdao *StudentDaoImpl) GetAllStudents() ([]*model.Student, error) {
	studmap := make(map[int]model.StudentResponse)
	courmap := make(map[int][]model.Course)

	var studentlist []*model.Student
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,COALESCE(Student.course.courseID,-1),COALESCE(courseName,''),COALESCE(courseFee,-1) FROM Student.student left join Student.enrollment on Student.student.studentID = Student.enrollment.studentID  left join Student.course on Student.enrollment.courseID = Student.course.courseID")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse
		var cour model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &cour.CourseID, &cour.CourseName, &cour.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if cour.CourseID != -1 {
				courmap[stud.StudentID] = append(courmap[stud.StudentID], cour)
			}

			studmap[stud.StudentID] = stud

		}
	}
	for i, j := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: j.StudentID, Name: j.Name, Email: j.Email, Dept: j.Dept, DOB: j.DOB, PhoneNo: j.PhoneNo, Course: courmap[i],
		})
	}
	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentsWithCourses() ([]*model.Student, error) {
	studmap := make(map[int]model.StudentResponse)
	courmap := make(map[int][]model.Course)

	var studentlist []*model.Student
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,Student.course.courseID,courseName,courseFee FROM Student.student  join Student.enrollment on Student.student.studentID = Student.enrollment.studentID   join Student.course on Student.enrollment.courseID = Student.course.courseID")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse
		var cour model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &cour.CourseID, &cour.CourseName, &cour.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if cour.CourseID != -1 {
				courmap[stud.StudentID] = append(courmap[stud.StudentID], cour)
			}
			studmap[stud.StudentID] = stud
		}
	}
	for i, j := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: j.StudentID, Name: j.Name, Email: j.Email, Dept: j.Dept, DOB: j.DOB, PhoneNo: j.PhoneNo, Course: courmap[i],
		})
	}
	return studentlist, nil

}
func (studdao *StudentDaoImpl) GetStudentById(id string) ([]*model.Student, error) {
	studmap := make(map[int]model.StudentResponse)
	courmap := make(map[int][]model.Course)

	var studentlist []*model.Student = []*model.Student{}
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,COALESCE(Student.course.courseID,-1),COALESCE(courseName,''),COALESCE(courseFee,-1) FROM Student.student  left join Student.enrollment on Student.student.StudentID = Student.enrollment.studentID left  join Student.course on Student.enrollment.courseID = Student.course.courseID where Student.student.studentID =?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse
		var cour model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &cour.CourseID, &cour.CourseName, &cour.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if cour.CourseID != -1 {
				courmap[stud.StudentID] = append(courmap[stud.StudentID], cour)
			}
			studmap[stud.StudentID] = stud
		}
	}
	for i, j := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: j.StudentID, Name: j.Name, Email: j.Email, Dept: j.Dept, DOB: j.DOB, PhoneNo: j.PhoneNo, Course: courmap[i],
		})
	}
	return studentlist, nil
}
func (studdao *StudentDaoImpl) GetStudentByCourseId(id string) ([]*model.Student, error) {
	studmap := make(map[int]model.StudentResponse)
	courmap := make(map[int][]model.Course)

	var studentlist []*model.Student = []*model.Student{}
	rows, err := config.DB.Query("SELECT Student.student.studentID,Name,Email,Dept,DOB,PhoneNo,COALESCE(Student.course.courseID,-1),COALESCE(courseName,''),COALESCE(courseFee,-1) FROM Student.student  left join Student.enrollment on Student.student.StudentID = Student.enrollment.studentID left  join Student.course on Student.enrollment.courseID = Student.course.courseID where Student.course.courseID =?", id)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var stud model.StudentResponse
		var cour model.Course

		err = rows.Scan(&stud.StudentID, &stud.Name, &stud.Email, &stud.Dept, &stud.DOB, &stud.PhoneNo, &cour.CourseID, &cour.CourseName, &cour.CourseFee)
		if err != nil {
			return nil, err
		} else {
			if cour.CourseID != -1 {
				courmap[stud.StudentID] = append(courmap[stud.StudentID], cour)
			}
			studmap[stud.StudentID] = stud
		}
	}
	for i, j := range studmap {
		studentlist = append(studentlist, &model.Student{
			StudentID: j.StudentID, Name: j.Name, Email: j.Email, Dept: j.Dept, DOB: j.DOB, PhoneNo: j.PhoneNo, Course: courmap[i],
		})
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
