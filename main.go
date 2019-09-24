package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/jinzhu/configor"
	"github.com/urfave/cli"
)

// Config represents the config file that values are read from.
var Config = struct {
	App     string `default:"app"`
	Version string `default:"version"`
	Author  string `default:"author"`
	Email   string `default:"email"`
	Usage   string `default:"usage"`

	Data struct {
		Contact        Contact        `default:"contact"`
		About          About          `default:"about"`
		Education      Education      `default:"education"`
		WorkExperience WorkExperience `default:"workExperience"`
		Projects       Projects       `default:"projects"`
		Publications   Publications   `default:"publications"`
	}
}{}

// Resume represents a full resume with every section returned
type Resume struct {
	Contact        Contact        `json:"contact"`
	WorkExperience WorkExperience `json:"workExperience"`
	Projects       Projects       `json:"projects"`
	Education      Education      `json:"education"`
	Publications   Publications   `json:"publications"`
}

// Contact represents the contact section of the resume.
type Contact struct {
	Section         string `json:"section"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Site            string `json:"site"`
	Github          string `json:"github"`
	Linkedin        string `json:"linkedin"`
	IsLookingForJob bool   `json:"isLookingForJob"`
}

// About represents the about section of the resume.
type About struct {
	Section string `json:"section"`
	Short   string `json:"short"`
	Long    string `json:"long"`
}

// Education represents the about section of the resume.
type Education struct {
	Section        string  `json:"section"`
	UniversityName string  `json:"university_name"`
	UniversityAbbr string  `json:"university_abbr"`
	Major          string  `json:"major"`
	GPA            float32 `json:"gpa"`
	Year           int     `json:"year"`
}

// WorkExperience represents the experience section of the resume
type WorkExperience struct {
	Section    string       `json:"section"`
	Experience []Experience `json:"experience"`
}

// Experience represents a single experience
type Experience struct {
	Company    string `json:"company"`
	Role       string `json:"role"`
	StartDate  string `json:"startDate"`
	FinishDate string `json:"finishDate"`
}

// Projects represents the projects section of the resume
type Projects struct {
	Section  string    `json:"section"`
	Projects []Project `json:"projects"`
}

// Project represents a single project
type Project struct {
	Name        string `json:"name"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	Link        string `json:"link"`
}

// Publications represents the publications section of the resume.
type Publications struct {
	Section      string        `json:"section"`
	Publications []Publication `json:"publications"`
}

// Publication represents a single publication.
type Publication struct {
	Title string `json:"title"`
	Year  string `json:"year"`
	Venue string `json:"venue"`
}

func getConf() {
	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&Config, "conf.json")
	if err != nil {
		fmt.Println("error parsing config:", err)
	}

	_ = configor.Load(&Config, "conf.json")
}

var app = cli.NewApp()

func main() {
	getConf()
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func info() {
	app.Name = Config.App
	app.Version = Config.Version
	app.Author = Config.Author
	app.Email = Config.Email
	app.Usage = Config.Usage
}

func resume() Resume {
	r := Resume{
		contact(),
		experience(),
		projects(),
		education(),
		publications(),
	}

	return r
}

func contact() Contact {
	c := Contact{
		Config.Data.Contact.Section,
		Config.Data.Contact.Name,
		Config.Data.Contact.Email,
		Config.Data.Contact.Site,
		Config.Data.Contact.Github,
		Config.Data.Contact.Linkedin,
		Config.Data.Contact.IsLookingForJob,
	}

	return c
}

func about() About {
	a := About{
		Config.Data.About.Section,
		Config.Data.About.Short,
		Config.Data.About.Long,
	}

	return a
}

func education() Education {
	e := Education{
		Config.Data.Education.Section,
		Config.Data.Education.UniversityName,
		Config.Data.Education.UniversityAbbr,
		Config.Data.Education.Major,
		Config.Data.Education.GPA,
		Config.Data.Education.Year,
	}

	return e
}

func experience() WorkExperience {
	w := WorkExperience{
		Config.Data.WorkExperience.Section,
		[]Experience{},
	}

	for _, exp := range Config.Data.WorkExperience.Experience {
		w.Experience = append(w.Experience, exp)
	}

	sort.Slice(w.Experience, func(i, j int) bool {
		return parseMonthYearToDate(w.Experience[i].StartDate).After(parseMonthYearToDate(w.Experience[j].StartDate))
	})

	return w
}

func projects() Projects {
	p := Projects{
		Config.Data.Projects.Section,
		[]Project{},
	}

	for _, exp := range Config.Data.Projects.Projects {
		p.Projects = append(p.Projects, exp)
	}

	sort.Slice(p.Projects, func(i, j int) bool {
		return parseYearToDate(p.Projects[i].Year).After(parseYearToDate(p.Projects[j].Year))
	})

	return p
}

func publications() Publications {
	p := Publications{
		Config.Data.Publications.Section,
		[]Publication{},
	}

	for _, exp := range Config.Data.Publications.Publications {
		p.Publications = append(p.Publications, exp)
	}

	sort.Slice(p.Publications, func(i, j int) bool {
		return parseYearToDate(p.Publications[i].Year).After(parseYearToDate(p.Publications[j].Year))
	})

	return p
}

// parseMonthYearToDate takes a month year string and converts to time.
func parseMonthYearToDate(d string) time.Time {
	t, err := time.Parse("January 2006", d)
	if err != nil {
		panic(err)
	}

	return t
}

// parseYearToDate takes a year string and converts to time.
func parseYearToDate(d string) time.Time {
	t, err := time.Parse("2006", d)
	if err != nil {
		panic(err)
	}

	return t
}

// prettyPrint turns our struct into clean json.
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "resume",
			Aliases: []string{"r"},
			Usage:   "Full Resume",
			Action: func(c *cli.Context) {
				r := resume()
				re := prettyPrint(r)
				fmt.Println(re)
			},
		},
		{
			Name:    "contact",
			Aliases: []string{"c"},
			Usage:   "Contact me",
			Action: func(c *cli.Context) {
				co := contact()
				con := prettyPrint(co)
				fmt.Println(con)
			},
		},
		{
			Name:    "about",
			Aliases: []string{"ab"},
			Usage:   "About me",
			Action: func(c *cli.Context) {
				a := about()
				ab := prettyPrint(a)
				fmt.Println(ab)

			},
		},
		{
			Name:    "education",
			Aliases: []string{"e"},
			Usage:   "View my education",
			Action: func(c *cli.Context) {
				e := education()
				ed := prettyPrint(e)
				fmt.Println(ed)
			},
		},
		{
			Name:    "experience",
			Aliases: []string{"ex"},
			Usage:   "View my work experience",
			Action: func(c *cli.Context) {
				e := experience()
				ex := prettyPrint(e)
				fmt.Println(ex)
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"p"},
			Usage:   "View my projects",
			Action: func(c *cli.Context) {
				p := projects()
				pr := prettyPrint(p)
				fmt.Println(pr)
			},
		},
		{
			Name:    "publications",
			Aliases: []string{"pubs"},
			Usage:   "View my publications",
			Action: func(c *cli.Context) {
				p := publications()
				pu := prettyPrint(p)
				fmt.Println(pu)
			},
		},
	}
}
