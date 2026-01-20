package component

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
	label   string
	options []string
	cursor  int
	choice  map[string]struct{}
}

func NewList(label string, options []string) *List {
	return &List{
		label:   label,
		options: options,
		cursor:  0,
		choice:  make(map[string]struct{}),
	}
}

func (l *List) Init() tea.Cmd {
	return nil
}

func (l *List) Update(msg tea.Msg) (Question, tea.Cmd) {
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
			if _, ok := l.choice[l.options[l.cursor]]; ok {
				delete(l.choice, l.options[l.cursor])
			} else {
				l.choice[l.options[l.cursor]] = struct{}{}
			}
		}
	}
	return l, nil
}

func (l *List) View() string {
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
		if _, ok := l.choice[opt]; ok {
			checked = "√"
		}
		s.WriteString(fmt.Sprintf("%s %s %s\n", cursor, style.Render(fmt.Sprintf("[%s]", checked)), style.Render(opt)))
	}
	return s.String()
}

func (l *List) Result() string {
	label := strings.TrimSuffix(strings.TrimSpace(l.label), ":")
	items := make([]string, 0)
	for key := range l.choice {
		split := strings.Split(key, "，")
		items = append(items, split[0])
	}
	return fmt.Sprintf("%s: %s", label, BlurredStyle.Render(strings.Join(items, ",")))
}

func (l *List) Focus() tea.Cmd {
	return nil
}

func (l *List) Blur() tea.Cmd {
	return nil
}

func (l *List) Value() string {
	items := make([]string, 0)
	for key := range l.choice {
		split := strings.Split(key, "，")
		items = append(items, split[0])
	}
	return strings.Join(items, ",")
}

func (l *List) Validate() bool {
	return true
}
