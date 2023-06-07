package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

type Exams struct {
	WeekNumber    []int `json:"weekNumber"`
	StudentGroups []struct {
		SpecialityName   string `json:"specialityName"`
		SpecialityCode   string `json:"specialityCode"`
		NumberOfStudents int    `json:"numberOfStudents"`
		Name             string `json:"name"`
	} `json:"studentGroups"`
	NumSubgroup      int         `json:"numSubgroup"`
	Auditories       []string    `json:"auditories"`
	StartLessonTime  string      `json:"startLessonTime"`
	EndLessonTime    string      `json:"endLessonTime"`
	Subject          string      `json:"subject"`
	SubjectFullName  string      `json:"subjectFullName"`
	Note             interface{} `json:"note"`
	LessonTypeAbbrev string      `json:"lessonTypeAbbrev"`
	DateLesson       string      `json:"dateLesson"`
	Employees        []Employee  `json:"employees"`
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
	Exams          []Exams     `json:"exams"`
	StartDate      string      `json:"startDate"`
	EndDate        string      `json:"endDate"`
	StartExamsDate interface{} `json:"startExamsDate"`
	EndExamsDate   interface{} `json:"endExamsDate"`
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

func FacultiesParse(client *http.Client) []Faculties {
	body, err := GetBody("https://iis.bsuir.by/api/v1/faculties", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return nil
	}
	var facultiesData []Faculties
	err = json.Unmarshal(body, &facultiesData)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return nil
	}

	fmt.Println("----------------------FacultiesInfo--------------------------------------")
	for _, elem := range facultiesData {
		fmt.Println(elem.Abbrev, " ", elem.Id, " ", elem.Name)
	}
	return facultiesData
}

func StudentGroupsParse(client *http.Client) []Specialities {
	body, err := GetBody("https://iis.bsuir.by/api/v1/student-groups", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return nil
	}
	var data []Specialities
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON: %s", err)
		return nil
	}

	fmt.Println("----------------------Students Groups---------------------------")
	for _, elem := range data {
		fmt.Println(elem.Id, " ", elem.CalendarId, " ", elem.Course, " ", elem.Name, " ", elem.FacultyName, " ", elem.FacultyId, " ", elem.SpecialityName, " ", elem.SpecialityDepartmentEducationFormId)
	}
	return data
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

func ScheduleParse(client *http.Client, groupNumber int) ([]Exams, error) {
	body, err := GetBody("https://iis.bsuir.by/api/v1/schedule?studentGroup="+strconv.Itoa(groupNumber), client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return nil, err
	}
	var data Schedule
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return nil, err
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

	fmt.Println("---------------------------Students group---------------------------")
	el := data.StudentGroupDto
	fmt.Println(el.Id, " ", el.CalendarId, " ", el.Name, " ",
		el.FacultyId, " ", el.FacultyName, " ", el.Name, " ", el.Course)

	fmt.Println("Start date : ", data.StartDate)
	fmt.Println("End date : ", data.EndDate)
	fmt.Println("Exams start date : ", data.StartExamsDate.(string))
	fmt.Println("End exams date : ", data.EndExamsDate.(string))

	return data.Exams, nil
}

func SeeExams(exams []Exams) {
	fmt.Println("---------------------------Exams---------------------------")
	for _, exams := range exams {
		for _, el := range exams.WeekNumber {
			fmt.Println("Week number :")
			fmt.Print(el)
			fmt.Println()
		}
		fmt.Println("Student groups : ")
		for _, el := range exams.StudentGroups {
			fmt.Println(el.NumberOfStudents, " ", el.Name, " ",
				el.NumberOfStudents, " ", el.SpecialityName, " ", el.SpecialityCode)
		}
		fmt.Println("---------------------------------------------------------")
		fmt.Println("Num of subgroup : ", exams.NumSubgroup)
		fmt.Println("Subject : ", exams.Subject)
		fmt.Println("Date lesson : ", exams.DateLesson)
		fmt.Println("Start Lesson Time : ", exams.StartLessonTime)
		fmt.Println("End Lesson Time : ", exams.EndLessonTime)
		fmt.Println("Lesson Type : ", exams.LessonTypeAbbrev)
		fmt.Print("Auditories : ")
		for _, el := range exams.Auditories {
			fmt.Print(el)
		}
		fmt.Println()
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

func CreateReport(data interface{}, name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("There are some error with the file : %s", err)
		return nil, err
	}

	if obj, ok := data.([]Faculties); ok {
		for _, el := range obj {
			file.WriteString(el.Name + ", ")
			file.WriteString(el.Abbrev + ", ")
			file.WriteString(strconv.Itoa(el.Id) + "\n")
		}
	}

	if obj, ok := data.([]Specialities); ok {
		for _, el := range obj {
			file.WriteString(el.Name + ", ")
			file.WriteString(el.CalendarId + ", ")
			file.WriteString(strconv.Itoa(el.Id) + "\n")
			file.WriteString(el.FacultyName + "\n")
			file.WriteString(el.SpecialityName + "\n")
			file.WriteString(strconv.Itoa(el.Course) + "\n")
			file.WriteString(strconv.Itoa(el.FacultyId) + "\n")
			file.WriteString(strconv.Itoa(el.SpecialityDepartmentEducationFormId) + "\n")
		}
	}

	return file, nil
}

func ShowMenu(client http.Client) {
	for {
		fmt.Println("-----------------------Menu-----------------------\n\n", "Choose option")
		fmt.Println("0) Exit")
		fmt.Println("1) Schedule")
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
				tmp, _ := ScheduleParse(&client, groupNumber)

				fmt.Println("Do you wanna see exams?")
				var ans string
				fmt.Scan(&ans)
				if ans == "yes" {
					SeeExams(tmp)
				}
			}

		case 2:
			data := StudentGroupsParse(&client)
			CreateReport(data, "output.txt")
		case 3:
			data := FacultiesParse(&client)
			CreateReport(data, "output.txt")
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

func main() {
	client := http.Client{}
	fmt.Println("Current weak number : ", GetWeakNumber(&client))
	ShowMenu(client)
}
