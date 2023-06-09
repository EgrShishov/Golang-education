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

type studentGroups struct {
	SpecialityName   string `json:"specialityName"`
	SpecialityCode   string `json:"specialityCode"`
	NumberOfStudents int    `json:"numberOfStudents"`
	Name             int    `json:"name"`
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

func EmployeeParse(client *http.Client) []Employee {
	body, err := GetBody("https://iis.bsuir.by/api/v1/employees/all", client)
	if err != nil {
		fmt.Printf("Problem with response body %s", err)
		return nil
	}
	var data []Employee
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return nil
	}
	return data
}

func PrintEmployees(data []Employee) {
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
	studentGroups, ok := scheduleItem["studentGroups"].([]interface{})
	if ok {
		for _, el := range studentGroups {
			group, ok := el.(map[string]string)
			if ok {
				fmt.Println("Speciality Name : ", group["SpecialityName"])
				fmt.Println("Name of group : ", group["Name"])
				fmt.Println("Number of students : ", group["NumberOfStudents"])
			}
		}
	}

	if numSubgroup, ok := scheduleItem["numSubgroup"].(int); ok {
		fmt.Println("Num of subgroup : ", numSubgroup)
	}
	if subject, ok := scheduleItem["subject"].(string); ok {
		fmt.Println("Subject : ", subject)
	}
	if auditories, ok := scheduleItem["auditories"].([]string); ok {
		fmt.Print("Auditories : ")
		for _, el := range auditories {
			fmt.Print(el, " ")
		}
		fmt.Println()
	}
	if lessonTypeAbbrev, ok := scheduleItem["lessonTypeAbbrev"].(string); ok {
		fmt.Println("lesson Type : ", lessonTypeAbbrev)
	}
	if startLessonTime, ok := scheduleItem["startLessonTime"].(string); ok {
		fmt.Println("Start Lesson Time : ", startLessonTime)
	}
	if endLessonTime, ok := scheduleItem["endLessonTime"].(string); ok {
		fmt.Println("End Lesson Time : ", endLessonTime)
	}
	if startTime, ok := scheduleItem["startTime"].(string); ok {
		fmt.Println("Start time : ", startTime)
	}
	if endTime, ok := scheduleItem["endTime"].(string); ok {
		fmt.Print("End time : ", endTime)
	}
	if location, ok := scheduleItem["location"].(string); ok {
		fmt.Print("Location : ", location)
	}
	if startExamsDate, ok := scheduleItem["startExamsDate"].(string); ok {
		fmt.Print("Exams Start : ", startExamsDate)
	}
	if endExamsDate, ok := scheduleItem["endExamsDate"].(string); ok {
		fmt.Print("Exams End : ", endExamsDate)
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

	ShowSchedule(&data)
	return data.Exams, nil
}

func ExamsParse(client *http.Client, groupNumber int) {
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
	SeeExams(data.Exams)
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

func EmployeesScheduleParse(EmployeeId string, client *http.Client) {
	body, err := GetBody("https://iis.bsuir.by/api/v1/employees/schedule/"+EmployeeId, client)
	if err != nil {
		fmt.Printf("Problem with response body : %s", err)
		return
	}

	var data Schedule
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Cant parse JSON : %s", err)
		return
	}
	ShowSchedule(&data)
}

func ShowSchedule(data *Schedule) {
	fmt.Println("--------------------------------Schedule--------------------------------")
	fmt.Println("Monday : ")
	for _, item := range data.Schedules.Monday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}
	fmt.Println("Tuesday : ")
	for _, item := range data.Schedules.Tuesday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}
	fmt.Println("Wednesday : ")
	for _, item := range data.Schedules.Wednesday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}
	fmt.Println("Thursday : ")
	for _, item := range data.Schedules.Thursday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}
	fmt.Println("Friday : ")
	for _, item := range data.Schedules.Friday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}
	fmt.Println("Saturday : ")
	for _, item := range data.Schedules.Saturday {
		if item != nil {
			ShowDaysLessons(item)
		}
	}

	fmt.Println("---------------------------Students group---------------------------")
	el := data.StudentGroupDto
	fmt.Println(el.Id, " ", el.CalendarId, " ", el.Name, " ",
		el.FacultyId, " ", el.FacultyName, " ", el.Name, " ", el.Course)

	fmt.Println("Start date : ", data.StartDate)
	fmt.Println("End date : ", data.EndDate)
	if _, ok := data.StartExamsDate.(string); ok {
		fmt.Println("Exams start date : ", data.StartExamsDate.(string))
	}
	if _, ok := data.EndExamsDate.(string); ok {
		fmt.Println("End exams date : ", data.EndExamsDate.(string))
	}
}

func CreateReport(data interface{}, name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("There are some error with the file : %s", err)
		return nil, err
	}

	if obj, ok := data.([]Faculties); ok {
		for _, el := range obj {
			_, _ = file.WriteString(el.Name + ", ")
			_, _ = file.WriteString(el.Abbrev + ", ")
			_, _ = file.WriteString(strconv.Itoa(el.Id) + "\n")
		}
	}

	if obj, ok := data.([]Specialities); ok {
		for _, el := range obj {
			_, _ = file.WriteString(el.Name + ", ")
			_, _ = file.WriteString(el.CalendarId + ", ")
			_, _ = file.WriteString(strconv.Itoa(el.Id) + "\n")
			_, _ = file.WriteString(el.FacultyName + "\n")
			_, _ = file.WriteString(el.SpecialityName + "\n")
			_, _ = file.WriteString(strconv.Itoa(el.Course) + "\n")
			_, _ = file.WriteString(strconv.Itoa(el.FacultyId) + "\n")
			_, _ = file.WriteString(strconv.Itoa(el.SpecialityDepartmentEducationFormId) + "\n")
		}
	}

	return file, nil
}

func WriteJSONExams(name string, data []Exams) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}
	asJson, _ := json.MarshalIndent(data, "", "\t")
	_, _ = file.Write(asJson)
	return nil
}

func ReadFromFile(name string) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("There are error with reading from file : %s", err)
		return
	}

	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("There are some error : %s", err)
		return
	}
	println(string(content))
	println("Reading status: Success!")
}

func WriteExams(data []Exams, file *os.File) {
	for _, el := range data {
		file.WriteString(el.SubjectFullName + " " + el.StartLessonTime + " " + el.EndLessonTime + " " +
			el.DateLesson + " " + el.LessonTypeAbbrev + "\n")
	}
}

func WriteIntoFile(name string, data interface{}) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Printf("There are error with the file : %s", err)
		return
	}

	if tmp, ok := data.([]Exams); ok {
		WriteExams(tmp, file)
	}
}

func ShowMenu(client http.Client) {
	for {
		fmt.Println("-----------------------Menu-----------------------\n", "Current Week Number : ",
			GetWeakNumber(&client), "\nChoose Option")
		fmt.Println("0) Exit")
		fmt.Println("1) Schedule")
		fmt.Println("2) Exams")
		fmt.Println("3) Student Group")
		fmt.Println("4) Faculties")
		fmt.Println("5) Employees")
		fmt.Println("6) Show Employees Schedule")
		fmt.Println("7) Write Into file")
		fmt.Println("8) Read From File")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("Input error : %s", err)
			return
		}

		switch choice {

		case 1:
			{
				var groupNumber int
				fmt.Println("Enter group number : ")
				_, err := fmt.Scan(&groupNumber)
				if err != nil {
					fmt.Printf("Input error : %s", err)
					return
				}
				ScheduleParse(&client, groupNumber)
			}

		case 2:
			{
				var groupNumber int
				fmt.Println("Enter group number : ")
				_, err := fmt.Scan(&groupNumber)
				if err != nil {
					fmt.Printf("Input error : %s", err)
					return
				}
				ExamsParse(&client, groupNumber)
				//if err := WriteJSONExams("json_test.txt", tmp); err != nil {
				//	fmt.Printf("There are some error with writting in the file : %s", err)
				//	return
				//}
				//fmt.Println("Successfully wrote!")
			}

		case 3:
			data := StudentGroupsParse(&client)
			_, _ = CreateReport(data, "output.txt")
		case 4:
			data := FacultiesParse(&client)
			_, _ = CreateReport(data, "output.txt")
		case 5:
			data := EmployeeParse(&client)
			PrintEmployees(data)
		case 6:
			{
				fmt.Println("Enter employees FIO")
				var lstNme, frstNme, mdlNme string
				fmt.Scan(&lstNme, &frstNme, &mdlNme)

				data := EmployeeParse(&client)
				for _, el := range data {
					if el.FirstName == frstNme && el.MiddleName == mdlNme && el.LastName == lstNme {
						EmployeesScheduleParse(el.UrlId, &client)
					}
				}
			}

		case 7:
		case 8:
			var name string
			fmt.Println("Input file name with .txt")
			_, _ = fmt.Scan(&name)
			ReadFromFile(name)
		case 0:
			return
		}

		fmt.Println("\nDo you want to continue?")
		var answer string
		_, err = fmt.Scan(&answer)
		if err != nil {
			fmt.Printf("Input error : %s", err)
			return
		}

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
	ShowMenu(client)
}
