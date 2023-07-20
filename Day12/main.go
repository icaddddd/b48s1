package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id         int
	Title      string
	Content    string
	Duration   string
	StartDate  string
	EndDate    string
	ReactJs    string
	Javascript string
	Android    string
	NodeJs     string
}

var dataProjects = []Project{
	{
		Id:         0,
		Title:      "aku suka main bola",
		Content:    "aku suka main bola",
		Duration:   "2 Bulan",
		StartDate:  "2000/09/08",
		EndDate:    "2000/10/08",
		ReactJs:    "ReactJs",
		Javascript: "Javascript",
		Android:    "Android",
		NodeJs:     "NodeJs",
	},
}

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/Project", project)
	e.GET("/FormProject", FormProject)
	e.GET("/Testimonials", Testimonials)
	e.GET("/ProjectDetail/:id", ProjectDetail)

	e.POST("/FormUpdateProject/:id", FormUpdateProject)
	e.POST("/AddProject", AddProject)
	e.POST("/DeleteProject/:id", DeleteProject)
	e.POST("/UpdateProject", UpdateProject)

	e.Logger.Fatal(e.Start("localhost: 5000"))
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/home.html")

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

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(), data)
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

func ProjectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	ProjectDetail := Project{}

	for index, data := range dataProjects {
		if index == idToInt {
			ProjectDetail = Project{
				Id:         index,
				Title:      data.Title,
				Content:    data.Content,
				Duration:   data.Duration,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				ReactJs:    data.ReactJs,
				Javascript: data.Javascript,
				Android:    data.Android,
				NodeJs:     data.NodeJs,
			}
		}
	}

	data := map[string]interface{}{
		"Id":      id,
		"Project": ProjectDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func AddProject(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	ReactJs := c.FormValue("ReactJs")
	Javascript := c.FormValue("Javascript")
	Android := c.FormValue("Android")
	NodeJs := c.FormValue("NodeJs")

	start, _ := time.Parse("2006-01-02", startdate)
	end, _ := time.Parse("2006-01-02", enddate)
	duration := end.Sub(start)

	durationString := formatDuration(duration)

	newProject := Project{
		Id:         0,
		Title:      title,
		Content:    content,
		Duration:   durationString,
		StartDate:  startdate,
		EndDate:    enddate,
		ReactJs:    ReactJs,
		Javascript: Javascript,
		Android:    Android,
		NodeJs:     NodeJs,
	}

	dataProjects = append(dataProjects, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/Project")
}

func formatDuration(duration time.Duration) string {
	months := duration / (time.Hour * 24 * 30)
	duration %= time.Hour * 24 * 30
	days := duration / (time.Hour * 24)
	duration %= time.Hour * 24

	return fmt.Sprintf("%d month, %d day", months, days)
}

func DeleteProject(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	dataProjects = append(dataProjects[:idToInt], dataProjects[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/Project")
}

func FormUpdateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ProjectToUpdate := Project{}

	for index, data := range dataProjects {
		if id == index {
			ProjectToUpdate = Project{
				Id:         index,
				Title:      data.Title,
				Content:    data.Content,
				Duration:   data.Duration,
				StartDate:  data.StartDate,
				EndDate:    data.EndDate,
				ReactJs:    data.ReactJs,
				Javascript: data.Javascript,
				Android:    data.Android,
				NodeJs:     data.NodeJs,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectToUpdate,
	}

	var tmpl, err = template.ParseFiles("views/FormUpdateProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Project Not Found"})
	}

	return tmpl.Execute(c.Response(), data)
}

func UpdateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	title := c.FormValue("title")
	content := c.FormValue("content")
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	ReactJs := c.FormValue("ReactJs")
	Javascript := c.FormValue("Javascript")
	Android := c.FormValue("Android")
	NodeJs := c.FormValue("NodeJs")

	start, _ := time.Parse("2006-01-02", startdate)
	end, _ := time.Parse("2006-01-02", enddate)
	duration := end.Sub(start)

	durationString := formatDuration(duration)

	dataProjects[id].Title = title
	dataProjects[id].Content = content
	dataProjects[id].StartDate = startdate
	dataProjects[id].EndDate = enddate
	dataProjects[id].ReactJs = ReactJs
	dataProjects[id].Javascript = Javascript
	dataProjects[id].Android = Android
	dataProjects[id].NodeJs = NodeJs
	dataProjects[id].Duration = durationString

	return c.Redirect(http.StatusMovedPermanently, "/Project")

}
