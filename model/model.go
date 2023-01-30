package model

type Student struct {
	StudentID int      `json:"stuid"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Dept      string   `json:"dept"`
	DOB       string   `json:"dob"`
	PhoneNo   string   `json:"phoneno"`
	Course    []Course `json:"course"`
}

type Course struct {
	CourseID   int    `json:"courid"`
	CourseName string `json:"couname"`
	CourseFee  string `json:"courfee"`
}

type Enrollment struct {
	Id        int `json:"id"`
	StudentID int `json:"stuid"`
	CourseID  int `json:"courid"`
}

type StudentResponse struct {
	StudentID int    `json:"stuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Dept      string `json:"dept"`
	DOB       string `json:"dob"`
	PhoneNo   string `json:"phoneno"`
	Course    Course `json:"course"`
}
