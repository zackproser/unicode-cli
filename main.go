package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type model struct {
	current rune
	last    rune
}

func initialModel() model {
	return model{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			m.current = []rune(msg.String())[0]
		}
	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("# Key to unicode point converter\n")
	b.WriteString(fmt.Sprintf("## Pressed: %s\n", string(m.current)))
	b.WriteString(fmt.Sprintf("## Unicode point: %U\n", m.current))
	rendered, err := glamour.Render(b.String(), "dark")
	if err != nil {
		b.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
	}
	return rendered
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
