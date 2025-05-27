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

// created slice as db
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("")

	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh", Website: "git.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", CoursePrice: 499, Author: &Author{Fullname: "parth", Website: "gitmit.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createoneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOnceCourse ).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse ).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9090", r))
}

// serve routes
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello go world</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one couse")

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	// if nothing found
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func createoneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one couse")
	w.Header().Set("Content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is missing")
	}
	var course Course
	//once pass course to decode now we have data in course
	_ = json.NewDecoder(r.Body).Decode(&course) // passed &course as reference of course
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
	}

	// append course
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) //string conversion
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOnceCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one couse")
	w.Header().Set("Content-type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	//loop, id, remove, add, with my ID

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one couse")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
