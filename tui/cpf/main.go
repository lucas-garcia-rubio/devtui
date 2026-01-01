package cpf

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
	"github.com/skatkov/devtui/internal/generator/cpf"
	"github.com/skatkov/devtui/internal/ui"
)

const (
	Title        = "CPF Generator"
	InstrMessage = "Press 'g' to generate CPF"
)

type CPFGeneratorModel struct {
	ui.BasePagerModel
}

func NewCPFGeneratorModel(common *ui.CommonModel) CPFGeneratorModel {
	model := CPFGeneratorModel{
		BasePagerModel: ui.NewBasePagerModelWithStatus(common, Title, InstrMessage),
	}

	return model
}

func (m CPFGeneratorModel) Init() tea.Cmd {
	return m.BasePagerModel.Init()
}

func (m CPFGeneratorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if cmd, handled := m.HandleCommonKeys(msg); handled {
			return m, cmd
		}
		switch msg.String() {
		case "g":
			cpf := cpf.Generate()
			m.Viewport.SetContent(cpf)
			m.Content = cpf
			m.FormattedContent = cpf
			m.State = ui.PagerStateBrowse
			m.StatusMessage = InstrMessage
		}
	}

	return m, nil
}

func (m CPFGeneratorModel) View() string {
	var b strings.Builder

	fmt.Fprint(&b, m.Viewport.View()+"\n")
	fmt.Fprint(&b, m.StatusBarView())

	if m.ShowHelp {
		fmt.Fprint(&b, "\n"+m.helpView())
	}

	return b.String()
}

func (m CPFGeneratorModel) helpView() (s string) {
	col1 := []string{
		"c              copy CPF",
		"q/ctrl+c       quit",
	}

	s += "\n"
	s += "k/↑      up                  " + col1[0] + "\n"
	s += "j/↓      down                " + col1[1] + "\n"
	s += "b/pgup   page up             " + "\n"
	s += "f/pgdn   page down           " + "\n"
	s += "u        ½ page up           " + "\n"
	s += "d        ½ page down         "

	if len(col1) > 5 {
		s += col1[5]
	}

	s = ui.Indent(s, 2)

	if m.Common.Width > 0 {
		lines := strings.Split(s, "\n")
		for i := range lines {
			l := runewidth.StringWidth(lines[i])
			n := max(m.Common.Width-l, 0)
			lines[i] += strings.Repeat(" ", n)
		}

		s = strings.Join(lines, "\n")
	}

	return ui.HelpViewStyle(s)
}
