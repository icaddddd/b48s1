package main

import (
	"context"
	"day10/connection"
	"day10/middleware"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	Author       string
	UserId       int
	LoginName    bool
}

type User struct {
	Id             int
	Name           string
	Email          string
	HashedPassword string
}

type SessionData struct {
	IsLogin  bool
	Name     string
	NotLogin bool
}

var userData = SessionData{}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/Project", project)
	e.GET("/FormProject", FormProject)
	e.GET("/Testimonials", Testimonials)
	e.GET("/ProjectDetail/:id", ProjectDetail)
	e.GET("/FormUpdateProject/:id", FormUpdateProject)

	e.GET("/FormRegister", FormRegister)
	e.GET("/FormLogin", FormLogin)
	e.GET("/logout", logout)

	e.POST("/AddProject", middleware.UploadFile(AddProject))
	e.POST("/DeleteProject/:id", DeleteProject)
	e.POST("/UpdateProject", middleware.UploadFile(UpdateProject))

	e.POST("/register", registerUser)
	e.POST("/login", loginUser)

	e.Logger.Fatal(e.Start("localhost: 5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/Home.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)

}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func FormProject(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/FormProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func Testimonials(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	dataSession := map[string]interface{}{
		"dataSession": userData,
	}

	return tmpl.Execute(c.Response(), dataSession)
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/Project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_project.id, title, image, start_date, end_date, content, technologies, tb_user.name AS author, tb_user.id FROM tb_project LEFT JOIN tb_user ON tb_project.author = tb_user.id ORDER BY tb_project.id DESC;")

	if session.Values["isLogin"] != true {
		userData.NotLogin = true
	} else {
		userData.NotLogin = false
	}

	dataProjects := []Project{}
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Title, &each.Image, &each.StartDate, &each.EndDate, &each.Content, &each.Technologies, &each.Author, &each.UserId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if session.Values["name"] == each.Author {
			each.LoginName = true
		} else {
			each.LoginName = false
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
		"Projects":     dataProjects,
		"dataSession":  userData,
		"FlashStatus":  session.Values["status"],
		"FlashMessage": session.Values["message"],
		"FlashName":    session.Values["name"],
	}

	delete(session.Values, "message")
	session.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), projects)
}

func ProjectDetail(c echo.Context) error {

	tmpl, err := template.ParseFiles("views/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)

	ProjectDetail := Project{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT tb_project.id, title, image, start_date, end_date, content, technologies, tb_user.name AS author, tb_user.id FROM tb_project LEFT JOIN tb_user ON tb_project.author = tb_user.id WHERE tb_user.id = $1;", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Image, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Content, &ProjectDetail.Technologies, &ProjectDetail.Author)

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
		"dataSession":     userData,
	}

	return tmpl.Execute(c.Response(), data)
}

func AddProject(c echo.Context) error {
	session, _ := session.Get("session", c)

	title := c.FormValue("title")
	image := c.Get("dataFile").(string)
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	content := c.FormValue("content")
	technoReactJs := c.FormValue("ReactJs")
	technoJavascript := c.FormValue("Javascript")
	technoAndroid := c.FormValue("Android")
	technoNodeJs := c.FormValue("NodeJs")
	author := session.Values["id"]

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (title, image, start_date, end_date, content, technologies[1], technologies[2], technologies[3], technologies[4], author) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", title, image, startdate, enddate, content, technoReactJs, technoJavascript, technoAndroid, technoNodeJs, author)

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

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", idToInt).Scan(&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.Image, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Content, &ProjectDetail.Technologies, &ProjectDetail.Author)

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

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = session.Values["isLogin"].(bool)
		userData.Name = session.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Id":              id,
		"Project":         ProjectDetail,
		"startDateString": ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString":   ProjectDetail.EndDate.Format("2006-01-02"),
		"dataSession":     userData,
	}

	return tmpl.Execute(c.Response(), data)
}

func UpdateProject(c echo.Context) error {
	session, _ := session.Get("session", c)

	id := c.FormValue("id")
	title := c.FormValue("title")
	image := c.Get("dataFile").(string)
	startdate := c.FormValue("startdate")
	enddate := c.FormValue("enddate")
	content := c.FormValue("content")
	technoReactJs := c.FormValue("ReactJs")
	technoJavascript := c.FormValue("Javascript")
	technoAndroid := c.FormValue("Android")
	technoNodeJs := c.FormValue("NodeJs")
	author := session.Values["id"]

	_, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	connection.Conn.Exec(context.Background(), "UPDATE tb_project SET title=$1, image=$2, start_date=$3, end_date=$4, content=$5, technologies[1]=$6, technologies[2]=$7, technologies[3]=$8, technologies[4]=$9, author=$10 WHERE id=$11", title, image, startdate, enddate, content, technoReactJs, technoJavascript, technoAndroid, technoNodeJs, author, id)

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

// auth and session

func FormLogin(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	if session.Values["isLogin"] == true {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	messageFlash := map[string]interface{}{
		"FlashStatus":  session.Values["status"],
		"FlashMessage": session.Values["message"],
	}

	delete(session.Values, "status")
	delete(session.Values, "message")
	session.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), messageFlash)
}

func FormRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	session, _ := session.Get("session", c)

	messageFlash := map[string]interface{}{
		"FlashStatus":  session.Values["status"],
		"FlashMessage": session.Values["message"],
	}

	delete(session.Values, "status")
	delete(session.Values, "message")
	session.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), messageFlash)
}

func registerUser(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		redirectMessage(c, "RegistrationFailed, please try again!", false, "/FormRegister")
	}

	return redirectMessage(c, "Registration Success", true, "/FormLogin")
}

func loginUser(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	var user = User{}

	errEmail := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
	errPass := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if errEmail != nil {
		return redirectMessage(c, "Email or Password wrong!", true, "/FormLogin")
	}

	if errPass != nil {
		return redirectMessage(c, "Email or Password wrong!", true, "/FormLogin")
	}

	session, _ := session.Get("session", c)
	session.Options.MaxAge = 3600
	session.Values["message"] = "login Success"
	session.Values["status"] = true // show alert
	session.Values["name"] = user.Name
	session.Values["id"] = user.Id
	session.Values["isLogin"] = true // access login
	session.Save(c.Request(), c.Response())

	return redirectMessage(c, "Login Succes", true, "/Project")
}

func redirectMessage(c echo.Context, message string, status bool, path string) error {
	session, _ := session.Get("session", c)
	session.Values["message"] = message
	session.Values["status"] = status
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, path)
}

func logout(c echo.Context) error {
	session, _ := session.Get("session", c)
	session.Options.MaxAge = -1
	session.Values["isLogin"] = false
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
