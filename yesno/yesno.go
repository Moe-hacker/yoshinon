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
package yesno

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
	width       int
	height      int
	message     string
	title       string
	border      string
	bgcolor     = "\033[1;48;2;100;149;237m"
	cursorcolor = "\033[1;48;2;152;245;225m"
	boxcolor    = "\033[1;48;2;254;228;208m\033[1;38;2;0;0;0m"
)

type Yesno_config struct {
	Width       int
	Height      int
	Message     string
	Title       string
	Bgcolor     string
	Cursorcolor string
	Boxcolor    string
	Border      string
}

type model struct {
	status int
}

func (m model) Init() tea.Cmd {
	// See ui.go
	Draw_borders(bgcolor, boxcolor, border, height, width)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		case "left":
			if m.status > 0 {
				m.status = 0
			}
		case "right":
			if m.status < 1 {
				m.status = 1
			}
		case "tab":
			if m.status > 0 {
				m.status = 0
			} else {
				m.status = 1
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
	row = (wsrow-height)/2 + height
	col = (wscol-width)/2 + 3
	control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	space := ""
	for i := 0; i < ((width/2)-7)*2; i++ {
		space += " "
	}
	if m.status == 0 {
		fmt.Fprintf(os.Stderr, cursorcolor+control+"<OK>\033[0m"+boxcolor+space+"<CANCEL>"+"\n")
	} else {
		fmt.Fprintf(os.Stderr, control+boxcolor+"<OK>"+space+cursorcolor+"<CANCEL>\n")
	}
	return boxcolor
}

func Yesno(m Yesno_config) int {
	height = m.Height
	width = m.Width
	message = m.Message
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
	title = m.Title
	// Check window size.
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	if wsrow < height+4 || wscol < width+4 {
		Error("window is too small!")
		os.Exit(1)
	}
	p := tea.NewProgram(model{status: 0}, tea.WithOutput(os.Stderr), tea.WithAltScreen())
	r, _ := p.Run()
	fmt.Fprint(os.Stderr, "\033[?25h\033c")
	return r.(model).status
}
