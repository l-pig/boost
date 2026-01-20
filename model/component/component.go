package component

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Question interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Question, tea.Cmd)
	View() string   // 交互时的显示（例如带光标的输入框）
	Result() string // 完成后的显示（例如仅显示用户输入的值）
	Focus() tea.Cmd
	Blur() tea.Cmd
	Value() string // 获取原始值
	Validate() bool
}

// styles
var (
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	NoStyle      = lipgloss.NewStyle()
)
