package tui

import (
	"fmt"

	"github.com/a3ylf/plight/db"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	selected map[int]struct{}
	sessions db.Sessions
	choices []string
}

func initialModel(sessions db.Sessions) model {
    var choices []string
    for a := range sessions {
        choices = append(choices, a)

    }
	return model{
		selected: make(map[int]struct{}),
        sessions: sessions,
        choices: choices,
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < len(m.sessions)-1 {
                m.cursor++
            }

        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    return m, nil
}
func (m model) View() string {
    // The header
    s := "Available Sessions\n\n"

    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        s += fmt.Sprintf("%s %s:\n", cursor, choice)
        if _, ok := m.selected[i]; ok {
            //todo
            x := m.sessions[m.choices[i]]
            for day,d := range x {
                s += fmt.Sprintf("%s:\nTotal time: %s\nPeriods:\n",day,d.Day_Total)
                for _,period := range d.Periods{
                    s += fmt.Sprintf("%s -> %s\n",period.From,period.To)

                }
                
            }
            
        }

        // Render the row
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s

}
func StartTui(sessions db.Sessions) {
    p := tea.NewProgram(initialModel(sessions))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        return
    }
}
