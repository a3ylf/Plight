package tui


import (
	"fmt"

	"github.com/a3ylf/plight/db"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	selected map[int]struct{}
	data     *db.Data
	sessions []string
	hits     []string
	end      bool
}

func initialModel(data *db.Data) model {
	var sessions []string
	for a := range data.Sessions {
		sessions = append(sessions, a)

	}
	var hits []string
	for b := range data.Hits {
		hits = append(hits, b)
	}
	return model{
		selected: make(map[int]struct{}),
		data:     data,
		sessions: sessions,
		hits:     hits,
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
			if m.cursor < len(m.data.Sessions)-1+len(m.data.Hits)-1 {
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
	for i, choice := range m.sessions {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s:\n", cursor, choice)
		if _, ok := m.selected[i]; ok {
			x := m.data.Sessions[m.sessions[i]]
			for day, d := range x.Days {
				s += fmt.Sprintf("%s:\nTotal time: %s\nPeriods:\n", day, d.Day_Total)
				for _, period := range d.Periods {
					s += fmt.Sprintf("%s -> %s\n", period.From, period.To)
				}
			}
		}

		// Render the row
	}
	// s += "\nHits\n"
	//    for i, choice := range m.hits {
	// 	s += fmt.Sprintf("%s %s:\n", cursor, choice)
	// 	if _, ok := m.selected[i]; ok {
	// 		x := m.data.Hits[m.hits[i]]
	// 		for day, d := range x {
	// 		    s += fmt.Sprintf("%v, %v\n",day,d)
	// 			}
	// 		}
	//    }

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s

}
func StartTui(data *db.Data) {
	p := tea.NewProgram(initialModel(data))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return
	}
}
