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

package gauge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
	. "yoshinon/error"
	. "yoshinon/ui"

	"golang.org/x/sys/unix"
)

type Gauge_config struct {
	Width       int
	Height      int
	Message     string
	Title       string
	Bgcolor     string
	Cursorcolor string
	Boxcolor    string
	Border      string
}

var (
	width       int
	height      int
	message     string
	title       string
	border      string
	bgcolor     = "\033[1;48;2;100;149;237m"
	cursorcolor = "\033[1;38;2;152;245;225m"
	boxcolor    = "\033[1;48;2;254;228;208m\033[1;38;2;0;0;0m"
)

func draw_gauge(per int) {
	gauge := ""
	i := per * (width - 3) / 100
	for j := 0; j < i; j++ {
		gauge += "█"
	}
	for j := i; j < (width - 4); j++ {
		gauge += "░"
	}
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	row := wsrow/2 + height/2
	col := ((wscol / 2) - (width / 2)) + 2
	control := "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	fmt.Fprint(os.Stderr, control+cursorcolor+bgcolor+gauge+boxcolor+strconv.Itoa(per)+"%")
}
func show_gauge() {
	// See ui.go
	Draw_borders(bgcolor, boxcolor, border, height, width)
	Show_message(message, title, boxcolor, width, height)
	bkstring := ""
	for {
		stat, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
			Error("not running with a pipe!")
			os.Exit(1)
		}
		reader := bufio.NewReader(os.Stdin)
		var b strings.Builder
		for {
			r, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			if r == '\n' {
				break
			}
			_, err = b.WriteRune(r)
			if err != nil {
				fmt.Println("Error getting input:", err)
				os.Exit(1)
			}
		}
		if b.String() != bkstring {
			i, err := strconv.Atoi(b.String())
			if err != nil {
				fmt.Fprint(os.Stderr, "\033[?25h\033c")
				Error("Not a number!")
				os.Exit(1)
			}
			draw_gauge(i)
			bkstring = b.String()
		}
		if b.String() == "100" {
			fmt.Fprint(os.Stderr, "\033[?25h\033c")
			break
		}
	}
}
func Gauge(m Gauge_config) {
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
	show_gauge()
}
