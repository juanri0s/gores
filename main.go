package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func info() {
	app.Name = "GoRes"
	app.Version = "1.0.0"
	app.Author = "Juan Rios"
	app.Email = "juansebrios@gmail.com"
	app.Usage = "gores resume-section"
}

type Resume struct {
	Contact      string `json:"contact"`
	Skills       string `json:"skills"`
	Work         string `json:"work"`
	Projects     string `json:"projects"`
	Education    string `json:"education"`
	Publications string `json:"publications"`
	Awards       string `json:"awards"`
}

type Contact struct {
	Section         string `json:"section"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Site            string `json:"site"`
	Github          string `json:"github"`
	Linkedin        string `json:"linkedin"`
	IsLookingForJob bool   `json:"isLookingForJob"`
}

type About struct {
	Section string `json:"section"`
	Short   string `json:"short"`
	Long    string `json:"long"`
}

type Education struct {
	Section    string  `json:"section"`
	University string  `json:"university"`
	Major      string  `json:"major"`
	GPA        float32 `json:"gpa"`
	Year       int     `json:"year"`
}

type Skills struct {
	Section string  `json:"section"`
	Skills  []Skill `json:"skills"`
}

type Skill struct {
	Skill    string `json:"skillName"`
	Years    string `json:"years"`
	Projects string `json:"projects"`
}

type Work struct {
	Section      string `json:"section"`
	Company      string `json:"company"`
	Start        string `json:"start"`
	Finish       string `json:"finish"`
	Descriptions string `json:"descriptions"`
}

type Projects struct {
	Section     string `json:"section"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Link        string `json:"link"`
}

type Publications struct {
	Section string `json:"section"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Year    string `json:"year"`
	Venue   string `json:"venue"`
}

type Awards struct {
	Section     string `json:"section"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        string `json:"year"`
	Venue       string `json:"venue"`
}

var app = cli.NewApp()

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func resume() {
	r := Resume{
		contact(),
		skills(),
		work(),
		projects(),
		education(),
		publications(),
		awards(),
	}

	fmt.Println("\n Summer 2019 Resume", r)
}

func contact() string {
	c := Contact{
		"Contact",
		"Juan Rios",
		"juansebrios@gmail.com",
		"https://juanri0s.github.io/",
		"https://github.com/juanri0s",
		"https://www.linkedin.com/in/jsrios/",
		true,
	}

	st := PrettyPrintStruct(c)

	return st
}

func about() string {
	a := About{
		"About",
		"",
		"",
	}

	st := PrettyPrintStruct(a)

	return st
}

func education() string {
	e := Education{
		"Contact",
		"New Jersey Institute of Technology",
		"Human-computer Interaction",
		3.84,
		2018,
	}

	st := PrettyPrintStruct(e)

	return st
}

func skills() string {
	s := Skills{
		Section: "Skills",
		Skills: []Skill{
			{
				"Angular",
				"2",
				"website",
			},
			{
				"Angular",
				"2",
				"website",
			},
			{
				"Angular",
				"2",
				"website",
			},
			{
				"Angular",
				"2",
				"website",
			},
		},
	}

	st := PrettyPrintStruct(s)

	return st
}

func work() string {
	w := Work{
		"Work Experience",
		"",
		"",
		"",
		"",
	}

	st := PrettyPrintStruct(w)

	return st
}

func projects() string {
	p := Projects{
		"Projects",
		"",
		"",
		"",
		"",
	}

	st := PrettyPrintStruct(p)

	return st
}

func publications() string {
	p := Publications{
		"Publications",
		"",
		"",
		"",
		"",
	}

	st := PrettyPrintStruct(p)

	return st
}

func awards() string {
	a := Awards{
		"Awards",
		"",
		"",
		"",
		"",
	}

	st := PrettyPrintStruct(a)

	return st
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "resume",
			Aliases: []string{"r"},
			Usage:   "Full Resume",
			Action: func(c *cli.Context) {
				resume()
			},
		},
		{
			Name:    "contact",
			Aliases: []string{"c"},
			Usage:   "Contact me",
			Action: func(c *cli.Context) {
				ct := contact()
				fmt.Println(ct)
			},
		},
		{
			Name:    "about",
			Aliases: []string{"ab"},
			Usage:   "About me",
			Action: func(c *cli.Context) {
				a := about()
				fmt.Println(a)
			},
		},
		{
			Name:    "education",
			Aliases: []string{"e"},
			Usage:   "View my education",
			Action: func(c *cli.Context) {
				e := education()
				fmt.Println(e)
			},
		},
		{
			Name:    "skills",
			Aliases: []string{"s"},
			Usage:   "View my skills",
			Action: func(c *cli.Context) {
				s := skills()
				fmt.Println(s)
			},
		},
		{
			Name:    "work",
			Aliases: []string{"w"},
			Usage:   "View my work experience",
			Action: func(c *cli.Context) {
				w := work()
				fmt.Println(w)
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "View my projects",
			Action: func(c *cli.Context) {
				p := projects()
				fmt.Println(p)
			},
		},
		{
			Name:    "publications",
			Aliases: []string{"pubs"},
			Usage:   "View my publications",
			Action: func(c *cli.Context) {
				p := publications()
				fmt.Println(p)
			},
		},
		{
			Name:    "awards",
			Aliases: []string{"a"},
			Usage:   "View my awards",
			Action: func(c *cli.Context) {
				a := awards()
				fmt.Println(a)
			},
		},
	}
}

func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "\n", "    ")
	d := fmt.Sprintf("%s", s)
	return d
}
