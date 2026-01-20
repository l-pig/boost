package component

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func NewShortText(label, placeholder string) *ShortText {
	t := textinput.New()
	t.Placeholder = placeholder
	t.CharLimit = 156
	t.Width = 30
	t.Prompt = "➜ "
	t.TextStyle = FocusedStyle
	t.Cursor.Style = CursorStyle

	return &ShortText{
		label: label,
		input: t,
	}
}

type ShortText struct {
	label string
	input textinput.Model
}

func (s *ShortText) Init() tea.Cmd {
	return textinput.Blink
}

func (s *ShortText) Update(msg tea.Msg) (Question, tea.Cmd) {
	var cmd tea.Cmd
	s.input, cmd = s.input.Update(msg)
	return s, cmd
}

func (s *ShortText) View() string {
	return fmt.Sprintf("%s\n%s", s.label, s.input.View())
}

func (s *ShortText) Result() string {
	label := strings.TrimSuffix(strings.TrimSpace(s.label), ":")
	return fmt.Sprintf("%s: %s", label, BlurredStyle.Render(s.input.Value()))
}

func (s *ShortText) Focus() tea.Cmd {
	return s.input.Focus()
}

func (s *ShortText) Blur() tea.Cmd {
	s.input.Blur()
	return nil
}

func (s *ShortText) Value() string {
	return s.input.Value()
}

func (s *ShortText) Validate() bool {
	return strings.TrimSpace(s.input.Value()) != ""
}
