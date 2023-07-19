package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/Home", home)
	e.GET("/contact", contact)
	e.GET("/Project", Project)
	e.GET("/Testimonials", Testimonials)
	e.GET("/ProjectDetail/:id", ProjectDetail)
	e.POST("/AddProject", AddProject)

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

func Project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Project.html")

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

	ProjectDetail := map[string]interface{}{
		"Id":        id,
		"Title":     "Aku Suka Main Bola Hehee",
		"Content":   "yaaa begituu lahhh gaiiiiisssssss",
		"Duration":  "2 Bulan",
		"StartDate": "15 Januari 2021",
		"EndDate":   "15 Maret 2021",
	}

	return tmpl.Execute(c.Response(), ProjectDetail)
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

	fmt.Println("title: ", title)
	fmt.Println("content: ", content)
	fmt.Println("startdate", startdate)
	fmt.Println("enddate", enddate)
	fmt.Println("ReactJs", ReactJs)
	fmt.Println("Javascript", Javascript)
	fmt.Println("Android", Android)
	fmt.Println("NodeJs", NodeJs)
	fmt.Println("Duration", durationString)

	// fmt.Println("Duration", durationString)

	return c.Redirect(http.StatusMovedPermanently, "/Project")
}

func formatDuration(duration time.Duration) string {
	months := duration / (time.Hour * 24 * 30)
	duration %= time.Hour * 24 * 30
	days := duration / (time.Hour * 24)
	duration %= time.Hour * 24

	return fmt.Sprintf("%d month, %d day", months, days)

}
