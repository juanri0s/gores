package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

type Resume struct {
	Contact      Contact `json:"contact"`
	Skills       Skills `json:"skills"`
	Work         Works `json:"work"`
	Projects     Projects `json:"projects"`
	Education    Education `json:"education"`
	Publications Publications `json:"publications"`
	Awards       Awards `json:"awards"`
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

func info() {
	app.Name = "GoRes"
	app.Version = "1.0.0"
	app.Author = "Juan Rios"
	app.Email = "juansebrios@gmail.com"
	app.Usage = "gores resume-section"
}

func resume() Resume {

	r := Resume{
		contact(),
		skills(),
		work(),
		projects(),
		education(),
		publications(),
		awards(),
	}

	return r
}

func contact() Contact {
	c := Contact{
		"Contact",
		"Juan Rios",
		"juansebrios@gmail.com",
		"https://juanri0s.github.io",
		"https://github.com/juanri0s",
		"https://www.linkedin.com/in/jsrios",
		true,
	}

	return c
}

func about() About {
	a := About{
		"About",
		"",
		"",
	}

	return a
}

func education() Education {
	e := Education{
		"Education",
		"New Jersey Institute of Technology",
		"Human-computer Interaction",
		3.84,
		2018,
	}

	return e
}

func skills() Skills {
	s := Skills{
		"Skills",
		[]Skill{
			{
				"Angular",
			},
			{
				"React",
			},
		},
	}

	return s
}

func work() Works {
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

	return w
}

func projects() Projects {
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

	return p
}

func publications() Publications {
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

	return p
}

func awards() Awards {
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

	return a
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "resume",
			Aliases: []string{"r"},
			Usage:   "Full Resume",
			Action: func(c *cli.Context) {
				r := resume()
				PrettyPrint(r)
			},
		},
		{
			Name:    "contact",
			Aliases: []string{"c"},
			Usage:   "Contact me",
			Action: func(c *cli.Context) {
				ct := contact()
				PrettyPrint(ct)
			},
		},
		{
			Name:    "about",
			Aliases: []string{"ab"},
			Usage:   "About me",
			Action: func(c *cli.Context) {
				a := about()
				PrettyPrint(a)
			},
		},
		{
			Name:    "education",
			Aliases: []string{"e"},
			Usage:   "View my education",
			Action: func(c *cli.Context) {
				e := education()
				PrettyPrint(e)
			},
		},
		{
			Name:    "skills",
			Aliases: []string{"s"},
			Usage:   "View my skills",
			Action: func(c *cli.Context) {
				s := skills()
				PrettyPrint(s)
			},
		},
		{
			Name:    "work",
			Aliases: []string{"w"},
			Usage:   "View my work experience",
			Action: func(c *cli.Context) {
				w := work()
				PrettyPrint(w)
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "View my projects",
			Action: func(c *cli.Context) {
				p := projects()
				PrettyPrint(p)
			},
		},
		{
			Name:    "publications",
			Aliases: []string{"pubs"},
			Usage:   "View my publications",
			Action: func(c *cli.Context) {
				p := publications()
				PrettyPrint(p)
			},
		},
		{
			Name:    "awards",
			Aliases: []string{"a"},
			Usage:   "View my awards",
			Action: func(c *cli.Context) {
				a := awards()
				PrettyPrint(a)
			},
		},
	}
}

// turns out structs into clean json
func PrettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
