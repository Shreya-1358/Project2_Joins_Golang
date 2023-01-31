Aim :- To join a student and course table using MySql in golang

Task :- Firstly storing the data in MySql database. I have a database configuration, model, routes, controller, service, dao and main file.

Process :-

config:- This package contains a variable DB of type pointer to sql.DB, which is a database connection handle. It also defines a function DbURL() which returns a string that represents a database URL built from environment variables.
The environment variables include:
* DB_USER: username for the database
* DB_PASSWORD: password for the database
* DB_HOST: hostname or IP address of the database
* DB_PORT: port number of the database
* DB_NAME_STUDENT: name of the database
The util.GetEnvVariable() function retrieves the value of an environment variable, which is used to build the database URL string that is returned by the DbURL() function.

model:- This package contains Go structs that represent different entities in the application.
* "Student" struct represents a student entity with the following fields:
    * StudentID: integer identifier of the student
    * Name: name of the student
    * Email: email of the student
    * Dept: department of the student
    * DOB: date of birth of the student
    * PhoneNo: phone number of the student
    * Course: array of courses taken by the student
* "Course" struct represents a course entity with the following fields:
    * CourseID: integer identifier of the course
    * CourseName: name of the course
    * CourseFee: fee of the course
* "Enrollment" struct represents an enrollment of a student in a course with the following fields:
    * StudentID: identifier of the student
    * CourseID: identifier of the course
* "StudentResponse" struct represents a simplified version of the "Student" struct, with only a subset of the fields, used for API responses.

routes:- This package is responsible for setting up the routing of a web application using the Gin framework.
The code imports necessary packages for the application's controller, data access object (DAO), and service layer.
The function "SetupRouter" creates instances of the Student DAO, Service, and Controller, and sets up the following RESTful API endpoints for the Student resource:
* GET /students: returns a list of all students
* GET /student/:id: returns a student with a given id
* GET /course/:id: returns a list of students in a given course
* POST /student: creates a new student
Finally, the function returns a Gin Engine instance with the routes set up, which can be used to start the application.

dao:- This package implements the Student DAO pattern. The StudentDao interface defines the four methods for accessing student data, GetAllStudents(), CreateStudent(), GetStudentById(), and GetStudentByCourseId(). 
The StudentDaoImpl struct implements these methods and provides the actual logic to interact with a database. The methods use the config.DB object to execute SQL queries to retrieve and manipulate the student data.
The StudentDao interface specifies the methods to be implemented for accessing and manipulating student data. The implementation of the interface is provided in the struct StudentDaoImpl.
The NewStudentDaoImpl function returns a new instance of the StudentDaoImpl struct.

service:-

controller:-

main:- It is the entry point of the program. The package has the following functions:
* init(): This function initiallizes the database connection and sets it up to use the MySQL driver. The function also checks if the connection was successful. If it fails, it logs an error.
* main(): This is the main function that runs when the program is executed. It sets up the routing of the application using the "SetupRouter" function of the "routes" package and starts the application with the "Run" method.
