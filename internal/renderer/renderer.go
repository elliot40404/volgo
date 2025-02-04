package renderer

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/elliot40404/volgo/internal/controller"
)

type (
	Charm struct {
		controls *controller.Controller
	}
	model struct {
		controls *controller.Controller
		percent  float64
		progress progress.Model
		muted    bool
	}
	Renderer interface {
		Render() error
	}
	tickMsg time.Time
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func NewRenderer(c *controller.Controller) Renderer {
	return Charm{
		controls: c,
	}
}

func (model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "enter":
			return m, tea.Quit
		case "up", "k", "l", "=":
			m.controls.IncreaseVolume(5)
		case "down", "j", "h", "-":
			m.controls.DecreaseVolume(5)
		case "m":
			if m.muted {
				m.controls.Unmute()
			} else {
				m.controls.Mute()
			}
			m.muted = !m.muted
		}
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil
	case tickMsg:
		pvol := m.controls.GetVolume()
		if pvol > -1 {
			m.percent = float64(pvol) / 100
			m.muted = m.controls.GetMuted()
		}
		return m, tickCmd()
	default:
		return m, nil
	}
	return m, nil
}

func getMuteIcon(mute bool) string {
	if mute {
		return "󰸈  "
	}
	return "󰕾  "
}

func getMuteText(mute bool) string {
	if mute {
		return "unmute"
	}
	return "mute"
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	str := strings.Builder{}
	str.WriteString("\n")
	str.WriteString(pad)
	str.WriteString(getMuteIcon(m.muted))
	str.WriteString(m.progress.ViewAs(m.percent))
	str.WriteString("\n\n")
	str.WriteString(pad)
	str.WriteString(helpStyle("q - quit, l/k/+/⇧ - increase, h/j/-/⇩ - decrease, m -", getMuteText(m.muted)))
	return str.String()
}

func (c Charm) Render() error {
	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	vol := c.controls.GetVolume()
	m := model{
		controls: c.controls,
		progress: prog,
		percent:  float64(vol) / 100,
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func tickCmd() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
