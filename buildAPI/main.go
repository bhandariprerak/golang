package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake  DB
var courses []Course

//middleware, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //we'll manually create courseId instead of relying on customer
}

func main() {
	fmt.Println("welcome to building api in go")

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice: 999, Author: &Author{Fullname: "Prerak Bhandari", Website: "pb.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Python", CoursePrice: 499, Author: &Author{Fullname: "Dev Bhandari", Website: "db.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "JS", CoursePrice: 299, Author: &Author{Fullname: "Manish Bhandari", Website: "metro.in"}})

	r := mux.NewRouter()
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	fmt.Println("port is up and running at : 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello and welcome prerak</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Printf("params content is : %v and type is %T\n", params, params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No matching ID found.")
	// return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create  one course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty?
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some body ")
	}

	// what if data is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in JSON ")
		return
	}
	//check if course already exists only by title name
	for _, mycourses := range courses {
		if course.CourseName == mycourses.CourseName {
			json.NewEncoder(w).Encode("Course Already Exists!")
		}
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update  one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some body ")
	}

	//first grab id from request
	params := mux.Vars(r)

	//loop, find id, remove that course, add with myid
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send a reponse when ID not found
	json.NewEncoder(w).Encode("No matching ID found.")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted successfully")
			return //to avoid no matching id found
		}
	}
	json.NewEncoder(w).Encode("No matching ID found.")
}
