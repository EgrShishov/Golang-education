package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Specialities struct {
	Name                                string `json:"name"`
	FacultyId                           int    `json:"facultyId"`
	FacultyName                         string `json:"facultyName"`
	SpecialityDepartmentEducationFormId int    `json:"specialityDepartmentEducationFormId"`
	SpecialityName                      string `json:"specialityName"`
	Course                              int    `json:"course"`
	Id                                  int    `json:"id"`
	CalendarId                          string `json:"calendarId"`
}

type Faculties struct {
	Name   string `json:"name"`
	Abbrev string `json:"abbrev"`
	Id     int    `json:"id"`
}

type Employee struct {
	FirstName          string      `json:"firstName"`
	LastName           string      `json:"lastName"`
	MiddleName         string      `json:"middleName"`
	Degree             string      `json:"degree"`
	Rank               string      `json:"rank"`
	PhotoLink          string      `json:"photoLink"`
	CalendarId         string      `json:"calendarId"`
	AcademicDepartment interface{} `json:"academicDepartment"`
	Id                 int         `json:"id"`
	UrlId              string      `json:"urlId"`
	FIO                string      `json:"fio"`
}

type Shedule struct {
	EmployeeDto     interface{}      `json:"employeeDto"`
	StudentGroupDto interface{}      `json:"studentGroupDto"`
	Schedules       interface{}      `json:"schedules"`
	Exams           interface{}      `json:"exams"`
	StartDate       string           `json:"startDate"`
	EndDate         string           `json:"endDate"`
	StartExamsDate  interface{}      `json:"startExamsDate"`
	EndExamsDate    interface{}      `json:"endExamsDate"`
	SheduleField    []SheduleLessons `json:"shedule"`
}

type SheduleLessons struct {
	WeekNumber       interface{} `json:"weekNumber"`
	StudentGroups    interface{} `json:"studentGroups"`
	NumSubgroup      int         `json:"numSubgroup"`
	Auditories       interface{} `json:"auditories"`
	StartLessonTime  string      `json:"startLessonTime"`
	EndLessonTime    string      `json:"endLessonTime"`
	Subject          string      `json:"subject"`
	SubjectFullName  string      `json:"subjectFullName"`
	Note             interface{} `json:"note"`
	LessonTypeAbbrev string      `json:"lessonTypeAbbrev"`
	DateLesson       interface{} `json:"dateLesson"`
	Employees        interface{} `json:"employees"`
}

func FacultiesParse(client *http.Client) {
	response, err := client.Get("https://iis.bsuir.by/api/v1/faculties")
	if err != nil {
		fmt.Printf("There are some error with response body : %s", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Cant read response body %s", err)
		return
	}

	var facultiesData []Faculties
	err = json.Unmarshal(body, &facultiesData)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return
	}

	fmt.Println("----------------------FacultiesInfo--------------------------------------")
	for _, elem := range facultiesData {
		fmt.Println(elem.Abbrev, " ", elem.Id, " ", elem.Name)
	}
}

func StudentGroupsParse(client *http.Client) {
	response, err := client.Get("https://iis.bsuir.by/api/v1/student-groups")
	if err != nil {
		fmt.Printf("There are some errors with requst : %s", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Cant read response body : %s", err)
		return
	}

	var data []Specialities
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON: %s", err)
		return
	}

	fmt.Println("----------------------Students Groups---------------------------")
	for _, elem := range data {
		fmt.Println(elem.Id, " ", elem.CalendarId, " ", elem.Course, " ", elem.Name, " ", elem.FacultyName, " ", elem.FacultyId, " ", elem.SpecialityName, " ", elem.SpecialityDepartmentEducationFormId)
	}
}

func EmployeeParse(client *http.Client) {
	response, err := client.Get("https://iis.bsuir.by/api/v1/employees/all")
	if err != nil {
		fmt.Printf("There are some error with the response body : %s", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Cant read response body : %s", err)
		return
	}

	var data []Employee
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		fmt.Printf("Cant parse JSON : %s", err1)
		return
	}

	fmt.Println("---------------------------Employees---------------------------")
	i := 1
	for _, el := range data {
		fmt.Println(i, " ", el.FirstName, " ", el.MiddleName, " ", el.LastName, " ", el.Id, " ", el.AcademicDepartment, " ", el.Degree,
			" ", el.PhotoLink, " ", el.Rank, " ", el.UrlId, " ", el.CalendarId, " ", el.FIO)
		i++
	}
}

func SheduleParse(client *http.Client, groupNumber int) {
	response, err := client.Get("https://iis.bsuir.by/api/v1/schedule?studentGroup=" + strconv.Itoa(groupNumber))
	if err != nil {
		fmt.Printf("There are some error with response body : %s", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Cant read response body : %s", err)
		return
	}

	var data []Shedule
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return
	}

	fmt.Println("--------------------------------Shedule--------------------------------")
	for _, el := range data {
		fmt.Println(el.EmployeeDto)
	}
}

func GetWeakNumber(client *http.Client) int {
	response, _ := client.Get("https://iis.bsuir.by/api/v1/schedule/current-week")

	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	var weakNumber int
	err := json.Unmarshal(body, &weakNumber)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return -1
	}
	return weakNumber
}

func ShowMenu(client http.Client) {

	for {
		fmt.Println("-----------------------Menu-----------------------\n\n", "Choose option")
		fmt.Println("0) Exit")
		fmt.Println("1) Shedule")
		fmt.Println("2) Student group")
		fmt.Println("3) Faculties")
		fmt.Println("4) Employees")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			{
				var groupNumber int
				fmt.Scan(&groupNumber)
				SheduleParse(&client, groupNumber)
			}

		case 2:
			StudentGroupsParse(&client)

		case 3:
			FacultiesParse(&client)

		case 4:
			EmployeeParse(&client)

		case 0:
			return

		}

		fmt.Println("Do you want to continue?")
		var answer string
		fmt.Scan(&answer)

		if answer == "yes" {
			continue
		} else if answer == "no" {
			break
		} else {
			fmt.Println("There are no such command")
		}

	}
}

func main() {

	client := http.Client{}

	fmt.Println("Current weak number : ", GetWeakNumber(&client))
	ShowMenu(client)
}
