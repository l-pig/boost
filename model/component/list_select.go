package component

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ListSelect struct {
	label   string
	options []string
	cursor  int
	choice  string
}

func NewListSelect(label string, options []string) *ListSelect {
	return &ListSelect{
		label:   label,
		options: options,
		cursor:  0,
	}
}

func (l *ListSelect) Init() tea.Cmd {
	return nil
}

func (l *ListSelect) Update(msg tea.Msg) (Question, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.cursor > 0 {
				l.cursor--
			}
		case "down", "j":
			if l.cursor < len(l.options)-1 {
				l.cursor++
			}
		case "enter", " ":
			l.choice = l.options[l.cursor]
		}
	}
	return l, nil
}

func (l *ListSelect) View() string {
	var s strings.Builder
	s.WriteString(l.label + "\n")
	for i, opt := range l.options {
		cursor := " "
		style := NoStyle
		if i == l.cursor {
			cursor = "➜"
			style = FocusedStyle
		}

		checked := " "
		if opt == l.choice {
			checked = "√"
		}

		s.WriteString(fmt.Sprintf("%s %s %s\n", cursor, style.Render(fmt.Sprintf("[%s]", checked)), style.Render(opt)))
	}
	return s.String()
}

func (l *ListSelect) Result() string {
	label := strings.TrimSuffix(strings.TrimSpace(l.label), ":")
	return fmt.Sprintf("%s: %s", label, BlurredStyle.Render(l.choice))
}

func (l *ListSelect) Focus() tea.Cmd {
	return nil
}

func (l *ListSelect) Blur() tea.Cmd {
	return nil
}

func (l *ListSelect) Value() string {
	return l.choice
}

func (l *ListSelect) Validate() bool {
	return strings.TrimSpace(l.choice) != ""
}
