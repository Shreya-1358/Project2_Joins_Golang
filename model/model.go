package model

type Student struct {
	StudentID int      `json:"stuid"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Dept      string   `json:"dept"`
	DOB       string   `json:"dob"`
	PhoneNo   string   `json:"phoneno"`
	Course    []Course `json:"course,omitempty"`
}

type Course struct {
	CourseID   int    `json:"courid,omitempty"`
	CourseName string `json:"couname,omitempty"`
	CourseFee  string `json:"courfee,omitempty"`
}

type Enrollment struct {
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
}
