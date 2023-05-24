package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	var groupNumber int
	fmt.Scan(&groupNumber)

	client := http.Client{}
	response, err := client.Get("https://iis.bsuir.by/api/v1/student-groups")
	if err != nil {
		fmt.Printf("There are some errors with requst : %s", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
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

	fmt.Println("Id CalendarId Course Name FacultyName FacultyId SpecialityName SpecialityDepartmentEducationFormId")
	for _, elem := range data {
		fmt.Println(elem.Id, " ", elem.CalendarId, " ", elem.Course, " ", elem.Name, " ", elem.FacultyName, " ", elem.FacultyId, " ", elem.SpecialityName, " ", elem.SpecialityDepartmentEducationFormId)
	}
}
