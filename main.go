package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var app = cli.NewApp()

// TODO: add colors for logging

func main() {
	fmt.Println("Starting GoRes")

	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func info() {
	app.Name = "GoRes"
	app.Version = "1.0.0"
	app.Author = "Juan Rios"
	app.Email = "juansebrios@gmail.com"
	app.Usage = "gores resume-section"
}

type Contact struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Site     string `json:"site"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
}

type About struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Education struct {
	University string `json:"university"`
	Major      string `json:"major"`
	GPA        string `json:"gpa"`
	Year       string `json:"year"`
}

type Skills struct {
	Skill    string `json:"skillName"`
	Years    string `json:"years"`
	Projects string `json:"projects"`
}

type Work struct {
	Company      string `json:"company"`
	Start        string `json:"start"`
	Finish       string `json:"finish"`
	Descriptions string `json:"descriptions"`
}

type Projects struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Link        string `json:"link"`
}

type Publications struct {
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Year    string `json:"year"`
	Venue   string `json:"venue"`
}

type Awards struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Venue       string `json:"venue"`
}

// TODO: fix types and create models? for them
func contact() {
	co := Contact{
		"Juan Rios",
		"sdasd",
		"sdfdf",
		"asdfdfsd",
		"sdfsdf",
	}

	PrettyPrintStruct(&co)
}

func about() string {
	// return quick description, maybe add flags for a short, medium, and long version
	return "return"
}

func education() string {
	// return kv pair of education including major, year graduated, GPA
	return "return"
}

func skills() string {
	// return kv pair of skills, years experience, and maybe a project that shows it
	return "return"
}

func work() string {
	// return kv pair of company, years, and bullet descriptions
	return "return"
}

func projects() string {
	// return kv pair of project, link, skills, year
	return "return"
}

func publications() string {
	// create kv of publication title, year, venue, authors
	return "return"
}

func awards() string {
	// create kv pair of award, year, company, description
	return "return"
}

func download() string {
	// download resume as pdf, might be hard to store on here so maybe we'll find a way to retrieve the latest
	return "return"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "contact",
			Aliases: []string{"c"},
			Usage:   "Contact me",
			Action: func(c *cli.Context) {
				contact()
			},
		},
		{
			Name:    "about",
			Aliases: []string{"ab"},
			Usage:   "About me",
			Action: func(c *cli.Context) {
				about()
			},
		},
		{
			Name:    "education",
			Aliases: []string{"e"},
			Usage:   "View my education",
			Action: func(c *cli.Context) {
				education()
			},
		},
		{
			Name:    "skills",
			Aliases: []string{"s"},
			Usage:   "View my skills",
			Action: func(c *cli.Context) {
				skills()
			},
		},
		{
			Name:    "work",
			Aliases: []string{"w"},
			Usage:   "View my work experience",
			Action: func(c *cli.Context) {
				work()
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "View my projects",
			Action: func(c *cli.Context) {
				projects()
			},
		},
		{
			Name:    "publications",
			Aliases: []string{"pubs"},
			Usage:   "View my publications",
			Action: func(c *cli.Context) {
				publications()
			},
		},
		{
			Name:    "awards",
			Aliases: []string{"a"},
			Usage:   "View my awards",
			Action: func(c *cli.Context) {
				awards()
			},
		},
		{
			Name:    "pdf",
			Aliases: []string{"pdf"},
			Usage:   "Download resume as PDF",
			Action: func(c *cli.Context) {
				download()
			},
		},
	}
}

func PrettyPrintStruct(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("%s \n", s)
}
