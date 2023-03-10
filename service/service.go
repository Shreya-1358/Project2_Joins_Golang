package service

import (
	"project_2/dao"
	"project_2/model"
)

type StudentService interface {
	GetAllStudents() ([]*model.Student, error)
	GetStudentById(string) ([]*model.Student, error)
	GetStudentByCourseId(string) ([]*model.Student, error)
	CreateStudent(*model.Student) error
}

type StudentServiceImpl struct {
	studdao dao.StudentDao
}

func NewStudentServiceImpl(dao dao.StudentDao) StudentService {
	return &StudentServiceImpl{studdao: dao}
}

func (studsvc *StudentServiceImpl) GetAllStudents() ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetAllStudents()
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) GetStudentById(id string) ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetStudentById(id)
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) GetStudentByCourseId(id string) ([]*model.Student, error) {

	studentlist, err := studsvc.studdao.GetStudentByCourseId(id)
	if err != nil {
		return nil, err
	}
	return studentlist, nil
}
func (studsvc *StudentServiceImpl) CreateStudent(student *model.Student) error {
	err := studsvc.studdao.CreateStudent(student)
	if err != nil {
		return err

	}
	return nil

}
