package portfolio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Portfolio struct holding all details
type Portfolio struct {
	Name     string
	About    string
	Skills   []string
	Projects map[string]string // project name -> description
	Contact  map[string]string // method -> info
}

// Show menu and handle user input
func (p *Portfolio) Show() {
	reader := bufio.NewReader(os.Stdin)

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	for {
		fmt.Println()
		color.Magenta("====== My Portfolio CLI ======")
		fmt.Printf("%s: %s\n\n", green("Name"), p.Name)
		fmt.Println(cyan("Select a section to view:"))
		fmt.Println(" 1) About Me")
		fmt.Println(" 2) Skills")
		fmt.Println(" 3) Projects")
		fmt.Println(" 4) Contact")
		fmt.Println(" 5) Exit")
		fmt.Print(yellow("Enter choice [1-5]: "))

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			p.printAbout()
		case "2":
			p.printSkills()
		case "3":
			p.printProjects()
		case "4":
			p.printContact()
		case "5":
			fmt.Println(red("Goodbye! 👋"))
			time.Sleep(time.Second)
			return
		default:
			fmt.Println(red("Invalid choice, please try again."))
		}
	}
}

func (p *Portfolio) printAbout() {
	color.Cyan("\n--- About Me ---")
	fmt.Println(p.About)
}

func (p *Portfolio) printSkills() {
	color.Cyan("\n--- Skills ---")
	for i, skill := range p.Skills {
		fmt.Printf(" %d. %s\n", i+1, skill)
	}
}

func (p *Portfolio) printProjects() {
	color.Cyan("\n--- Projects ---")
	for name, desc := range p.Projects {
		color.Green(" • %s:", name)
		fmt.Printf("   %s\n", desc)
	}
}

func (p *Portfolio) printContact() {
	color.Cyan("\n--- Contact ---")
	for method, info := range p.Contact {
		fmt.Printf(" %s: %s\n", method, info)
	}
}

// Helper to create & return a populated Portfolio instance
func NewPortfolio() *Portfolio {
	return &Portfolio{
		Name:  "Bikash", // Replace with your name or keep it dynamic later
		About: "I am a passionate Go developer building awesome CLI tools and APIs. I love concurrency and clean code.",
		Skills: []string{
			"Go",
			"REST APIs",
			"Concurrency",
			"Microservices",
			"SQL & NoSQL",
			"Docker & Kubernetes",
		},
		Projects: map[string]string{
			"API Benchmark Tool": "A CLI app to test API performance with concurrency.",
			"Portfolio CLI":      "This interactive CLI portfolio you're using right now.",
			"Chatbot":           "A GPT-powered assistant built in Go.",
		},
		Contact: map[string]string{
			"Email":    "youremail@example.com",
			"GitHub":   "github.com/yourusername",
			"LinkedIn": "linkedin.com/in/yourprofile",
		},
	}
}
