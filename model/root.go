package model

import (
	"boost/internal/generator"
	"boost/model/component"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Root struct {
	questions []component.Question
	index     int
	done      bool
}

func NewRoot() tea.Model {
	// 定义问题列表，混合了文本输入和选择列表
	qs := []component.Question{
		component.NewShortText("What is the project name?", "e.g. my-awesome-project"),
		component.NewShortText("What is the module name?", "e.g. github.com/user/project"),
		component.NewListSelect("Choose project type:", []string{"Web Application", "CLI Tool", "Library", "gRPC Service"}),
		component.NewList("Choose Dependencies:", func() []string {
			file, err := os.ReadFile("/home/l-pig/Coding/Go/boost/model/dependence.txt")
			if err != nil {
				panic(err)
			}
			split := strings.Split(string(file), "\n")

			print(split)
			return split
		}()),
	}

	// 聚焦第一个问题
	qs[0].Focus()

	return Root{
		questions: qs,
		index:     0,
	}
}

func (r Root) Init() tea.Cmd {
	return nil
}

func (r Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 全局退出键
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyCtrlC {
			return r, tea.Quit
		}
	}

	if r.done {
		return r, nil
	}

	current := r.questions[r.index]

	// 检查是否按下回车，且当前组件有值（简单的校验逻辑）
	if keyMsg, ok := msg.(tea.KeyMsg); ok && keyMsg.Type == tea.KeyEnter {
		// 这里可以加校验逻辑，比如 Value() 不能为空
		if current.Validate() {
			current.Blur()

			if r.index < len(r.questions)-1 {
				r.index++
				return r, r.questions[r.index].Focus()
			}

			r.done = true
			return r, tea.Quit
		}
	}

	// 将消息传递给当前组件
	var cmd tea.Cmd
	r.questions[r.index], cmd = current.Update(msg)

	return r, cmd
}

func (r Root) View() string {
	var b strings.Builder

	// 标题
	b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Render("Boost Project Initializer"))
	b.WriteString("\n\n")

	for i, q := range r.questions {
		if i < r.index {
			// 历史问题：显示结果
			b.WriteString(q.Result())
			b.WriteString("\n")
		} else if i == r.index {
			// 当前问题：显示交互界面
			b.WriteString(q.View())
			b.WriteString("\n")
		} else {
			// 未来问题：不显示
			// 如果你想显示即将到来的步骤（灰色），可以在这里写逻辑
		}
	}

	if r.done {
		for _, item := range r.questions {
			b.WriteString(item.Result() + "\n")
		}
		b.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("42")).Render("\nAll set! Generating project...\n"))
		// 生成代码

		config := generator.ProjectConfig{
			ProjectName: r.questions[0].Value(),
			Module:      r.questions[1].Value(),
			Dependencies: func() []string {
				return strings.Split(r.questions[3].Value(), ",")
			}(),
		}

		fmt.Printf("Creating project %s...\n", config.ProjectName)
		if err := generator.Generate(config); err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			os.Exit(1)
		}
	}

	return b.String()
}
