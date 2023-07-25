package main

import (
	"context"
	"day10/connection"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Title        string
	Content      string
	Duration     string
	StartDate    time.Time
	EndDate      time.Time
	Technologies []string
	ReactJs      bool
	Javascript   bool
	Android      bool
	NodeJs       bool
	Image        string
}

// var dataProjects = []Project{
// {
// 	Id:         0,
// 	Title:      "aku suka main bola",
// 	Content:    "aku suka main bola",
// 	Duration:   "2 Bulan",
// 	StartDate:  "2000/09/08",
// 	EndDate:    "2000/10/08",
// 	ReactJs:    true,
// 	Javascript: false,
// 	// 	Android:    true,
// 	// 	NodeJs:     true,
// 	// },
// }

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/Project", project)
	e.GET("/FormProject", FormProject)
	e.GET("/Testimonials", Testimonials)
	e.GET("/ProjectDetail/:id", ProjectDetail)
	e.GET("/FormUpdateProject/:id", FormUpdateProject)

	e.POST("/AddProject", AddProject)
	e.POST("/DeleteProject/:id", DeleteProject)
	e.POST("/UpdateProject", UpdateProject)

	e.Logger.Fatal(e.Start("localhost: 5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func FormProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/FormProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func Testimonials(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, image, start_date, end_date, content, technologies FROM tb_project")

	dataProjects := []Project{}
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Title, &each.Image, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Duration = countDuration(each.StartDate, each.EndDate)

		if checkValue(each.Technologies, "ReactJs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "Javascript") {
			each.Javascript = true
		}
		if checkValue(each.Technologies, "Android") {
			each.Android = true
		}
		if checkValue(each.Technologies, "NodeJs") {
			each.NodeJs = true
		}

		dataProjects = append(dataProjects, each)
	}

	projects := map[string]interface{}{
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(), projects)
}

func ProjectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	ProjectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Image, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Content, &ProjectDetail.Technologies)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = countDuration(ProjectDetail.StartDate, ProjectDetail.EndDate)

	if checkValue(ProjectDetail.Technologies, "ReactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "Javascript") {
		ProjectDetail.Javascript = true
	}
	if checkValue(ProjectDetail.Technologies, "Android") {
		ProjectDetail.Android = true
	}
	if checkValue(ProjectDetail.Technologies, "NodeJs") {
		ProjectDetail.NodeJs = true
	}

	data := map[string]interface{}{
		"Id":              id,
		"Project":         ProjectDetail,
		"startDateString": ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString":   ProjectDetail.EndDate.Format("2006-01-02"),
	}

	return tmpl.Execute(c.Response(), data)
}

func AddProject(c echo.Context) error {
	title := c.FormValue("title")
	image := c.FormValue("image")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	content := c.FormValue("content")
	technoReactJs := c.FormValue("ReactJs")
	technoJavascript := c.FormValue("Javascript")
	technoAndroid := c.FormValue("Android")
	technoNodeJs := c.FormValue("NodeJs")

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (title, image, start_date, end_date, content, technologies[1], technologies[2], technologies[3], technologies[4]) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", title, image, startdate, enddate, content, technoReactJs, technoJavascript, technoAndroid, technoNodeJs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/Project")
}

func DeleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)

	return c.Redirect(http.StatusMovedPermanently, "/Project")
}

func FormUpdateProject(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/FormUpdateProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Project Not Found"})
	}

	idToInt, _ := strconv.Atoi(id)

	ProjectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Image, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Content, &ProjectDetail.Technologies)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = countDuration(ProjectDetail.StartDate, ProjectDetail.EndDate)

	if checkValue(ProjectDetail.Technologies, "ReactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "Javascript") {
		ProjectDetail.Javascript = true
	}
	if checkValue(ProjectDetail.Technologies, "Android") {
		ProjectDetail.Android = true
	}
	if checkValue(ProjectDetail.Technologies, "NodeJs") {
		ProjectDetail.NodeJs = true
	}

	data := map[string]interface{}{
		"Id":              id,
		"Project":         ProjectDetail,
		"startDateString": ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString":   ProjectDetail.EndDate.Format("2006-01-02"),
	}

	return tmpl.Execute(c.Response(), data)
}

func UpdateProject(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	image := c.FormValue("image")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	content := c.FormValue("content")
	technoReactJs := c.FormValue("ReactJs")
	technoJavascript := c.FormValue("Javascript")
	technoAndroid := c.FormValue("Android")
	technoNodeJs := c.FormValue("NodeJs")

	_, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	connection.Conn.Exec(context.Background(), "UPDATE tb_project SET title=$1, image=$2, start_date=$3, end_date=$4, content=$5, technologies[1]=$6, technologies[2]=$7, technologies[3]=$8, technologies[4]=$9 WHERE id=$10", title, image, startdate, enddate, content, technoReactJs, technoJavascript, technoAndroid, technoNodeJs, id)

	return c.Redirect(http.StatusMovedPermanently, "/Project")

}

func countDuration(d1 time.Time, d2 time.Time) string {

	diff := d2.Sub(d1)
	days := int(diff.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months >= 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}

func checkValue(slice []string, object string) bool {
	for _, data := range slice {
		if data == object {
			return true
		}
	}
	return false
}
