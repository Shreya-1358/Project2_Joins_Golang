# Creating an API to join a student and course table using Golang, Gin and MySql.

## Steps to follow:-

### Prerequisites:-

#### Install go1.19:
To avoid any surprises in the project install go 1.19, please use the commands below to install go 1.19 in your machine:

```javascript
$ go install golang.org/dl/go1.19@latest

$ go1.19 download
```

#### Clone the project and install the dependencies:
Clone the project from Github using the following command:

```javascript
$ git clone git@github.com:Shreya-1358/Project2_Joins_Golang.git
```

Next run the following command to install dependencies.

```javascript
$ go get
```

## API Reference

#### Get all Students with their courses

```http
  GET /students
```

#### Get a Stduent with courses

```http
  GET /student/:id
```

#### Get Stduents enrolled with the course Id

```http
  GET /course/:id
```

#### POST Create Student, Course and Enrollment

```http
  POST /student
```

## Project Outline :-
Firstly storing the data in MySql database. I have a database configuration, util, model, routes, controller, service, dao and main file.

### main:- 
It is the entry point of the program. The package has the following functions:
* init(): This function initiallizes the database connection and sets it up to use the MySQL driver. The function also checks if the connection was successful. If it fails, it logs an error.
* main(): This is the main function that runs when the program is executed. It sets up the routing of the application using the "SetupRouter" function of the "routes" package and starts the application with the "Run" method.

### routes:- 
* This package is responsible for setting up the routing of a web application using the Gin framework.
* The code imports necessary packages for the application's controller, data access object (DAO), and service layer.
* The function "SetupRouter" creates instances of the Student DAO, Service, and Controller, and sets up the RESTful API endpoints for the Student resource.
* Finally, the function returns a Gin Engine instance with the routes set up, which can be used to start the application.

### controller:- 
The code is defining a struct called StudentController that contains a property studentservice of type service.StudentService. It has four methods:
* GetAllStudents
* GetStudentById
* GetStudentByCourseId
* CreateStudent

### service:- 
* This is the implementation of the service layer of the application the student service. 
* It defines a StudentService interface with methods to perform CRUD operations on student records. The interface is then implemented by the struct StudentServiceImpl which has a field of type dao.StudentDao. 
* The implementation methods call corresponding methods on the StudentDao object to perform the operations on the data. 
* A function NewStudentServiceImpl is also provided to create an instance of the StudentServiceImpl struct with a provided dao.StudentDao instance.

### dao:- 
* This package implements the Student DAO pattern. The StudentDao interface defines the four methods for accessing student data, GetAllStudents(), CreateStudent(), GetStudentById(), and GetStudentByCourseId(). 
* The StudentDaoImpl struct implements these methods and provides the actual logic to interact with a database. The methods use the config.DB object to execute SQL queries to retrieve and manipulate the student data.
* The NewStudentDaoImpl function returns a new instance of the StudentDaoImpl struct.

### model:- 
This package contains Go structs that represent different entities in the application.
* "Student" struct represents a student entity with some specified fields
* "Course" struct represents a course entity with some specified fields
* "Enrollment" struct represents an enrollment of a student in a course with some specified fields
* "StudentResponse" struct represents a simplified version of the "Student" struct, with only a subset of the fields, used for API responses.

### config:- 
* This package contains a variable DB of type pointer to sql.DB, which is a database connection handle. 
* It also defines a function DbURL() which returns a string that represents a database URL built from environment variables.

### util:- 
* The code defines a package util that provides a function GetEnvVariable to retrieve the value of an environment variable.
* The function first loads the .env file using the godotenv library to load environment variables from a file into the process's environment.
* After loading the environment variables, the function returns the value of the environment variable specified by the argument key by calling the os.Getenv function.
