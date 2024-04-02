package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

type model struct {
	textInput textinput.Model
	help      help.Model
	keymap    keymap
}

type Procedures struct {
	Procedures []string `json:"procedures"`
}

type keymap struct{}

type gotReposSuccessMsg []repo
type gotReposErrMsg error

type repo struct {
	Names []string `json:"procedures"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	// client, err := supabase.InitSupabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, count, err := client.From("profiles").Select("*", "exact", false).Execute()
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Printf("Count: %d\n", count)
	// }
	// fmt.Printf("Data: %v\n", data)

	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}

func getRepos() tea.Msg {
	// Open the JSON file
	file, err := os.Open("lists/urological_procedures.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return gotReposErrMsg(err)
	}
	defer file.Close()

	// Decode the JSON data into a Procedures struct
	var data Procedures
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return gotReposErrMsg(err)
	}

	var repos []repo // will only have 1 repo
	repos = append(repos, repo{Names: data.Procedures})
	fmt.Println(data)
	return gotReposSuccessMsg(repos)
}

// Keybindings
func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete")),
		key.NewBinding(key.WithKeys("ctrl+n"), key.WithHelp("ctrl+n", "next")),
		key.NewBinding(key.WithKeys("ctrl+p"), key.WithHelp("ctrl+p", "prev")),
		key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "quit")),
	}
}

// Keybindings
func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func initialModel() model {
	textInput := textinput.New()
	textInput.Placeholder = "repository"
	textInput.Prompt = "charmbracelet/"
	textInput.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	textInput.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	textInput.Focus()
	// ti.CharLimit = 50
	// ti.Width = 20
	textInput.ShowSuggestions = true

	help := help.New()

	keymap := keymap{}

	return model{textInput: textInput, help: help, keymap: keymap}
}

func (self model) Init() tea.Cmd {
	return tea.Batch(getRepos, textinput.Blink)
}

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return self, tea.Quit
		}
	case gotReposSuccessMsg:
		var suggestions []string

		for _, r := range msg {
			for _, procedure := range r.Names {
				suggestions = append(suggestions, procedure)
			}
		}
		fmt.Println(suggestions)

		self.textInput.SetSuggestions(suggestions)
	}

	var cmd tea.Cmd
	self.textInput, cmd = self.textInput.Update(msg)
	return self, cmd
}

func (self model) View() string {
	return fmt.Sprintf(
		"Pick a Charm™ repo:\n\n  %s\n\n%s\n\n",
		self.textInput.View(),
		self.help.View(self.keymap),
	)
}
