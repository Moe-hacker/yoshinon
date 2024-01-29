// SPDX-License-Identifier: MIT
/*
 *
 * This file is part of yoshinon, with ABSOLUTELY NO WARRANTY.
 *
 * MIT License
 *
 * Copyright (c) 2023 Moe-hacker
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 *
 *
 */

package checklist

import (
	"fmt"
	"os"
	"syscall"
	. "yoshinon/error"
	. "yoshinon/ui"

	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/sys/unix"
)

var (
	tags        []string
	items       []string
	status      []int
	width       int
	height      int
	listheight  int
	listwidth   int
	message     string
	title       string
	border      string
	bgcolor     = "\033[1;48;2;100;149;237m"
	cursorcolor = "\033[1;48;2;152;245;225m"
	boxcolor    = "\033[1;48;2;254;228;208m\033[1;38;2;0;0;0m"
)

type Checklist_config struct {
	Tags        []string
	Items       []string
	Status      []int
	Width       int
	Height      int
	Listheight  int
	Message     string
	Title       string
	Bgcolor     string
	Cursorcolor string
	Boxcolor    string
	Border      string
}
type model struct {
	// Current chosen choice.
	cursor int
	// Curson position on the screen.
	position int
	// For yesno buttons.
	yesno int
}

func (m model) Init() tea.Cmd {
	// See ui.go
	Draw_borders(bgcolor, boxcolor, border, height, width)
	// Correct the value of listheight.
	if listheight > len(items) {
		listheight = len(items)
	}
	// Set value of listwidth.
	for i := 0; i < len(items); i++ {
		if len(items[i])+len(tags[i])+2 > listwidth {
			listwidth = len(items[i]) + len(tags[i]) + 2
		}
	}
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		case "tab":
			if m.yesno < 3 {
				m.yesno++
			} else {
				m.yesno = 1
			}
		case "down", "j":
			if m.yesno == 1 {
				if m.cursor < len(items)-1 {
					m.cursor++
				}
				if m.position < listheight-1 {
					m.position++
				}
			} else {
				if m.yesno < 3 {
					m.yesno++
				} else {
					m.yesno = 1
				}
			}
		case "up", "k":
			if m.yesno == 1 {
				if m.cursor > 0 {
					m.cursor--
				}
				if m.position > 0 {
					m.position--
				}
			} else {
				if m.yesno > 1 {
					m.yesno--
				} else {
					m.yesno = 1
				}
			}
		case " ":
			if status[m.cursor] == 1 {
				status[m.cursor] = 0
			} else {
				status[m.cursor] = 1
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	fmt.Fprint(os.Stderr, "\033[1H"+bgcolor)
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	control := ""
	row := (wsrow-height)/2 + 3
	col := ((wscol / 2) - len(message)/2) + 2
	control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	// See ui.go
	Show_message(message, title, boxcolor, width, height)
	row = ((wsrow / 2) - (listheight / 2)) + 2
	col = ((wscol / 2) - (listwidth / 2)) + 2
	for i := m.cursor - m.position; i < m.cursor+(listheight-m.position); i++ {
		control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
		if m.cursor == i && m.yesno == 1 {
			if status[i] == 1 {
				fmt.Fprintf(os.Stderr, control+cursorcolor+"[x]"+tags[i]+" "+items[i]+"\033[0m"+boxcolor)
			} else {
				fmt.Fprintf(os.Stderr, control+cursorcolor+"[ ]"+tags[i]+" "+items[i]+"\033[0m"+boxcolor)
			}
		} else {
			if status[i] == 1 {
				fmt.Fprintf(os.Stderr, control+boxcolor+"[x]"+tags[i]+" "+items[i])
			} else {
				fmt.Fprintf(os.Stderr, control+boxcolor+"[ ]"+tags[i]+" "+items[i])
			}
		}
		for j := 0; j <= listwidth-len(tags[i])-len(items[i])-2; j++ {
			fmt.Fprintf(os.Stderr, " ")
		}
		fmt.Fprintf(os.Stderr, "\033[0m"+bgcolor)
		fmt.Fprintf(os.Stderr, "\n")
		row++
	}
	row = (wsrow-height)/2 + height
	col = (wscol-width)/2 + 3
	control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	space := ""
	for i := 0; i < ((width/2)-7)*2; i++ {
		space += " "
	}
	if m.yesno == 2 {
		fmt.Fprintf(os.Stderr, cursorcolor+control+"<OK>\033[0m"+boxcolor+space+"<CANCEL>"+"\n")
	} else if m.yesno == 3 {
		fmt.Fprintf(os.Stderr, "\033[0m"+control+boxcolor+"<OK>"+space+cursorcolor+"<CANCEL>"+"\n")
	} else {
		fmt.Fprintf(os.Stderr, control+boxcolor+"<OK>"+space+"<CANCEL>\n")
	}
	return boxcolor
}

func Checklist(m Checklist_config) string {
	height = m.Height
	width = m.Width
	listheight = m.Listheight
	tags = make([]string, len(m.Tags))
	for i := 0; i < len(m.Tags); i++ {
		tags[i] = m.Tags[i]
	}
	items = make([]string, len(m.Items))
	for i := 0; i < len(m.Items); i++ {
		items[i] = m.Items[i]
	}
	status = make([]int, len(m.Status))
	for i := 0; i < len(m.Items); i++ {
		status[i] = m.Status[i]
	}
	message = m.Message
	title = m.Title
	if m.Bgcolor != "" {
		bgcolor = m.Bgcolor
	}
	if m.Boxcolor != "" {
		boxcolor = m.Boxcolor
	}
	if m.Cursorcolor != "" {
		cursorcolor = m.Cursorcolor
	}
	if m.Border == "" {
		m.Border = "rounded"
	}
	border = m.Border
	// Check window size.
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	if wsrow < height+4 || wscol < width+4 {
		Error("window size is too small")
		os.Exit(1)
	}
	p := tea.NewProgram(model{yesno: 1}, tea.WithOutput(os.Stderr), tea.WithAltScreen())
	r, _ := p.Run()
	fmt.Fprint(os.Stderr, "\033[?25h\033c")
	if r.(model).yesno == 3 {
		os.Exit(1)
	}
	ret := ""
	for i := 0; i < len(tags); i++ {
		if status[i] == 1 {
			ret += "\"" + tags[i] + "\"" + " "
		}
	}
	return ret
}
