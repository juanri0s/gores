package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/urfave/cli"
	"os"
	"sort"
	"time"
)

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

type Resume struct {
	Contact        Contact        `json:"contact"`
	WorkExperience WorkExperience `json:"workExperience"`
	Projects       Projects       `json:"projects"`
	Education      Education      `json:"education"`
	Publications   Publications   `json:"publications"`
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
	Section        string  `json:"section"`
	UniversityName string  `json:"university_name"`
	UniversityAbbr string  `json:"university_abbr"`
	Major          string  `json:"major"`
	GPA            float32 `json:"gpa"`
	Year           int     `json:"year"`
}

type WorkExperience struct {
	Section    string       `json:"section"`
	Experience []Experience `json:"experience"`
}

type Experience struct {
	Company    string `json:"company"`
	Role       string `json:"role"`
	StartDate  string `json:"startDate"`
	FinishDate string `json:"finishDate"`
}

type Projects struct {
	Section  string    `json:"section"`
	Projects []Project `json:"projects"`
}

type Project struct {
	Name        string `json:"name"`
	Year        string `json:"year"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	Link        string `json:"link"`
}

type Publications struct {
	Section      string        `json:"section"`
	Publications []Publication `json:"publications"`
}

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
		[]Experience{
		},
	}

	for _, exp := range Config.Data.WorkExperience.Experience {
		w.Experience = append(w.Experience, exp)
	}

	sort.Slice(w.Experience, func(i, j int) bool {
		return ParseMonthYearToDate(w.Experience[i].StartDate).After(ParseMonthYearToDate(w.Experience[j].StartDate))
	})

	return w
}

func projects() Projects {
	p := Projects{
		Config.Data.Projects.Section,
		[]Project{
		},
	}

	for _, exp := range Config.Data.Projects.Projects {
		p.Projects = append(p.Projects, exp)
	}

	sort.Slice(p.Projects, func(i, j int) bool {
		return ParseYearToDate(p.Projects[i].Year).After(ParseYearToDate(p.Projects[j].Year))
	})

	return p
}

func publications() Publications {
	p := Publications{
		Config.Data.Publications.Section,
		[]Publication{
		},
	}

	for _, exp := range Config.Data.Publications.Publications {
		p.Publications = append(p.Publications, exp)
	}

	sort.Slice(p.Publications, func(i, j int) bool {
		return ParseYearToDate(p.Publications[i].Year).After(ParseYearToDate(p.Publications[j].Year))
	})

	return p
}

func ParseMonthYearToDate(d string) time.Time {

	time, err := time.Parse("January 2006", d)
	if err != nil {
		panic(err)
	}

	return time
}

func ParseYearToDate(d string) time.Time {

	time, err := time.Parse("2006", d)
	if err != nil {
		panic(err)
	}

	return time
}

// turns our structs into clean json
func PrettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
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
			Name:    "experience",
			Aliases: []string{"ex"},
			Usage:   "View my work experience",
			Action: func(c *cli.Context) {
				e := experience()
				PrettyPrint(e)
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
	}
}
