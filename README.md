Aim :- Creating an API to join a student and course table using MySql in golang.

Task :- Firstly storing the data in MySql database. I have a database configuration, util, model, routes, controller, service, dao and main file.

Process :-

config:- This package contains a variable DB of type pointer to sql.DB, which is a database connection handle. 
It also defines a function DbURL() which returns a string that represents a database URL built from environment variables.

util:- The code defines a package util that provides a function GetEnvVariable to retrieve the value of an environment variable.
The function first loads the .env file using the godotenv library to load environment variables from a file into the process's environment.
After loading the environment variables, the function returns the value of the environment variable specified by the argument key by calling the os.Getenv function.

model:- This package contains Go structs that represent different entities in the application.
* "Student" struct represents a student entity with the following fields
* "Course" struct represents a course entity with the following fields
* "Enrollment" struct represents an enrollment of a student in a course with the following fields
* "StudentResponse" struct represents a simplified version of the "Student" struct, with only a subset of the fields, used for API responses.

routes:- This package is responsible for setting up the routing of a web application using the Gin framework.
The code imports necessary packages for the application's controller, data access object (DAO), and service layer.
The function "SetupRouter" creates instances of the Student DAO, Service, and Controller, and sets up the RESTful API endpoints for the Student resource.
Finally, the function returns a Gin Engine instance with the routes set up, which can be used to start the application.

dao:- This package implements the Student DAO pattern. The StudentDao interface defines the four methods for accessing student data, GetAllStudents(), CreateStudent(), GetStudentById(), and GetStudentByCourseId(). 
The StudentDaoImpl struct implements these methods and provides the actual logic to interact with a database. The methods use the config.DB object to execute SQL queries to retrieve and manipulate the student data.
The StudentDao interface specifies the methods to be implemented for accessing and manipulating student data. The implementation of the interface is provided in the struct StudentDaoImpl.
The NewStudentDaoImpl function returns a new instance of the StudentDaoImpl struct.

service:- This is a Go programming language implementation of the service layer of an application, specifically the student service. 
It defines a StudentService interface with methods to perform CRUD operations on student records. The interface is then implemented by the struct StudentServiceImpl which has a field of type dao.StudentDao. 
The implementation methods call corresponding methods on the StudentDao object to perform the operations on the data. 
A function NewStudentServiceImpl is also provided to create an instance of the StudentServiceImpl struct with a provided dao.StudentDao instance.

controller:- The code is defining a struct called StudentController that contains a property studentservice of type service.StudentService. It has four methods:
NewStudentController: A constructor method that takes an object of service.StudentService as an argument and returns a pointer to a StudentController object with studentservice property set to the passed argument.
GetAllStudents
GetStudentById
GetStudentByCourseId
CreateStudent

main:- It is the entry point of the program. The package has the following functions:
* init(): This function initiallizes the database connection and sets it up to use the MySQL driver. The function also checks if the connection was successful. If it fails, it logs an error.
* main(): This is the main function that runs when the program is executed. It sets up the routing of the application using the "SetupRouter" function of the "routes" package and starts the application with the "Run" method.
