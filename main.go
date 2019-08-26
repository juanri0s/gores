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
	Skill string `json:"skillName"`
}

type Works struct {
	Section string `json:"section"`
	Works   []Work `json:"work"`
}

type Work struct {
	Company      string        `json:"company"`
	Start        string        `json:"startDate"`
	Finish       string        `json:"finishDate"`
	Descriptions []Description `json:"description"`
}

type Description struct {
	Description string `json:"description"`
}

type Projects struct {
	Section  string    `json:"section"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Name        string  `json:"name"`
	Year        string  `json:"year"`
	Description string  `json:"description"`
	Skills      []Skill `json:"skillsUsed"`
	Link        string  `json:"link"`
}

type Publications struct {
	Section      string        `json:"section"`
	Publications []Publication `json:"publications"`
}

type Publication struct {
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Year    string `json:"year"`
	Venue   string `json:"venue"`
}

type Awards struct {
	Section string  `json:"section"`
	Awards  []Award `json:"awards"`
}

type Award struct {
	Name        string `json:"name"`
	Year        string `json:"year"`
	Description string `json:"description"`
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

	fmt.Println("\n Summer 2019 Resume:", r)
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

	st := PrettyPrint(c)

	return st
}

func about() string {
	a := About{
		"About",
		"",
		"",
	}

	st := PrettyPrint(a)

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

	st := PrettyPrint(e)

	return st
}

func skills() string {
	s := Skills{
		"Skills",
		[]Skill{
			{
				"Angular",
			},
			{
				"Angular",
			},
		},
	}

	st := PrettyPrint(s)

	return st
}

func work() string {
	w := Works{
		"Work Experience",
		[]Work{
			{
				"UPS",
				"2019",
				"2019",
				[]Description{
					{
						"adfasffasfagfagdfgdsfg",
					},
				},
			},
			{
				"Carnegie Mellon",
				"2019",
				"2019",
				[]Description{
					{
						"sdgsdfgdsfgsdgsdgsdfgsdf",
					},
				},
			},
		},
	}

	st := PrettyPrint(w)

	return st
}

func projects() string {
	p := Projects{
		"Projects",
		[]Project{
			{
				"",
				"",
				"df",
				[]Skill{
					{
						"",
					},
				},
				"www.hello.com",
			},
		},
	}

	st := PrettyPrint(p)

	return st
}

func publications() string {
	p := Publications{
		"Publications",
		[]Publication{
			{
				"",
				"",
				"",
				"",
			},
		},
	}

	st := PrettyPrint(p)

	return st
}

func awards() string {
	a := Awards{
		"Awards",
		[]Award{
			{
				"",
				"",
				"",
			},
		},
	}

	st := PrettyPrint(a)

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

// We have to marshall the structs so that we can pretty print it as json
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "    ")
	d := fmt.Sprintf("\n%s", s)
	return d
}
