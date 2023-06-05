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

type Schedule struct {
	EmployeeDto     interface{} `json:"employeeDto"`
	StudentGroupDto struct {
		Name        string `json:"name"`
		FacultyId   int    `json:"facultyId"`
		FacultyName string `json:"facultyName"`
		Course      int    `json:"course"`
		Id          int    `json:"id"`
		CalendarId  string `json:"calendarId"`
	} `json:"studentGroupDto"`
	Schedules struct {
		Monday    []interface{} `json:"Понедельник"`
		Tuesday   []interface{} `json:"Вторник"`
		Wednesday []interface{} `json:"Среда"`
		Thursday  []interface{} `json:"Четверг"`
		Friday    []interface{} `json:"Пятница"`
		Saturday  []interface{} `json:"Суббота"`
	} `json:"schedules"`
	Exams          []interface{} `json:"exams"`
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	StartExamsDate interface{}   `json:"startExamsDate"`
	EndExamsDate   interface{}   `json:"endExamsDate"`
}

func GetBody(url string, client *http.Client) ([]byte, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func FacultiesParse(client *http.Client) {
	body, err := GetBody("https://iis.bsuir.by/api/v1/faculties", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
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
	body, err := GetBody("https://iis.bsuir.by/api/v1/student-groups", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
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
	body, err := GetBody("https://iis.bsuir.by/api/v1/employees/all", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return
	}
	var data []Employee
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
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

func ShowDaysLessons(item interface{}) {
	scheduleItem := item.(map[string]interface{})
	if subject, ok := scheduleItem["subject"].(string); ok {
		fmt.Print("Subject : ", subject)
	}
	if startTime, ok := scheduleItem["startTime"].(string); ok {
		fmt.Print("Subject : ", startTime)
	}
	if endTime, ok := scheduleItem["endTime"].(string); ok {
		fmt.Print("Subject : ", endTime)
	}
	if location, ok := scheduleItem["location"].(string); ok {
		fmt.Print("Subject : ", location)
	}
	fmt.Println()
}

func ScheduleParse(client *http.Client, groupNumber int) {
	body, err := GetBody("https://iis.bsuir.by/api/v1/schedule?studentGroup="+strconv.Itoa(groupNumber), client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return
	}
	var data Schedule
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return
	}

	fmt.Println("--------------------------------Schedule--------------------------------")
	fmt.Println("Monday : ")
	for _, item := range data.Schedules.Monday {
		ShowDaysLessons(item)
	}
	fmt.Println("Tuesday : ")
	for _, item := range data.Schedules.Tuesday {
		ShowDaysLessons(item)
	}
	fmt.Println("Wednesday : ")
	for _, item := range data.Schedules.Wednesday {
		ShowDaysLessons(item)
	}
	fmt.Println("Thursday : ")
	for _, item := range data.Schedules.Thursday {
		ShowDaysLessons(item)
	}
	fmt.Println("Friday : ")
	for _, item := range data.Schedules.Friday {
		ShowDaysLessons(item)
	}
	fmt.Println("Saturday : ")
	for _, item := range data.Schedules.Saturday {
		ShowDaysLessons(item)
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
				fmt.Println("Enter group number : ")
				fmt.Scan(&groupNumber)
				ScheduleParse(&client, groupNumber)
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

		fmt.Println("\nDo you want to continue?")
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

func newgcd(a ...int) int {
	var s []int
	s = a
	if len(s) == 2 {
		return gcd(s[0], s[1])
	} else {
		tmp := s[len(s)-1]
		s = s[:len(s)-1]
		return gcd(newgcd(s...), tmp)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func mcd(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	client := http.Client{}
	fmt.Println("Current weak number : ", GetWeakNumber(&client))
	ShowMenu(client)

	fmt.Println(gcd(2022, 831))
}
