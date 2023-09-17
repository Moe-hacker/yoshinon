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
package ui

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strings"
	"syscall"
)

type borderStruct struct {
	Top         string
	Bottom      string
	Left        string
	Right       string
	TopLeft     string
	TopRight    string
	BottomRight string
	BottomLeft  string
}

var roundedBorder = borderStruct{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "╰",
	BottomRight: "╯",
}
var normalBorder = borderStruct{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "└",
	BottomRight: "┘",
}
var thickBorder = borderStruct{
	Top:         "━",
	Bottom:      "━",
	Left:        "┃",
	Right:       "┃",
	TopLeft:     "┏",
	TopRight:    "┓",
	BottomLeft:  "┗",
	BottomRight: "┛",
}

var doubleBorder = borderStruct{
	Top:         "═",
	Bottom:      "═",
	Left:        "║",
	Right:       "║",
	TopLeft:     "╔",
	TopRight:    "╗",
	BottomLeft:  "╚",
	BottomRight: "╝",
}

var hiddenBorder = borderStruct{
	Top:         " ",
	Bottom:      " ",
	Left:        " ",
	Right:       " ",
	TopLeft:     " ",
	TopRight:    " ",
	BottomLeft:  " ",
	BottomRight: " ",
}
var border = roundedBorder

func Draw_borders(bgcolor string, boxcolor string, bordertype string, height int, width int) {
	switch bordertype {
	case "rounded":
		border = roundedBorder
	case "normal":
		border = normalBorder
	case "thick":
		border = thickBorder
	case "double":
		border = doubleBorder
	case "hidden":
		border = hiddenBorder
	}
	// Clear the screen.
	fmt.Fprint(os.Stderr, "\033c")
	// Set background color.
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	fmt.Fprint(os.Stderr, bgcolor+"\033[?25l")
	for i := 0; i < wsrow; i++ {
		for j := 0; j < wscol; j++ {
			fmt.Fprint(os.Stderr, " ")
		}
	}
	// Draw the border.
	fmt.Fprint(os.Stderr, "\033[2H")
	if wsrow-height > 0 {
		for i := 0; i < (wsrow-height)/2-1; i++ {
			fmt.Fprint(os.Stderr, "\n")
		}
	}
	space_before := ""
	for i := 0; i < (wscol-width)/2; i++ {
		space_before += " "
	}
	fmt.Fprint(os.Stderr, space_before)
	fmt.Fprint(os.Stderr, "\033[0m"+boxcolor)
	fmt.Fprint(os.Stderr, border.TopLeft)
	for i := 0; i < width; i++ {
		fmt.Fprint(os.Stderr, border.Top)
	}
	fmt.Fprint(os.Stderr, border.TopRight+bgcolor+"\n")
	space_after := ""
	for i := 0; i < width; i++ {
		space_after += " "
	}
	for i := 0; i < height; i++ {
		fmt.Fprint(os.Stderr, space_before)
		fmt.Fprint(os.Stderr, "\033[0m"+boxcolor)
		fmt.Fprint(os.Stderr, border.Left)
		fmt.Fprint(os.Stderr, space_after)
		fmt.Fprint(os.Stderr, border.Right)
		fmt.Fprint(os.Stderr, bgcolor+"\n")
	}
	fmt.Fprint(os.Stderr, space_before)
	fmt.Fprint(os.Stderr, "\033[0m"+boxcolor)
	fmt.Fprint(os.Stderr, border.BottomLeft)
	for i := 0; i < width; i++ {
		fmt.Fprint(os.Stderr, border.Bottom)
	}
	fmt.Fprint(os.Stderr, border.BottomRight+bgcolor+"\n")
}
func Show_message(message string, title string, boxcolor string, width int, height int) {
	ws, err := unix.IoctlGetWinsize(syscall.Stderr, unix.TIOCGWINSZ)
	if err != nil {
		panic(err)
	}
	wsrow := int(ws.Row)
	wscol := int(ws.Col)
	control := ""
	row := (wsrow-height)/2 + 1
	col := (wscol-len(title))/2 + 1
	control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	if title != "" {
		fmt.Fprintf(os.Stderr, boxcolor+control)
		fmt.Fprintf(os.Stderr, " "+title+" ")
	}
	row = (wsrow-height)/2 + 3
	col = ((wscol / 2) - len(message)/2) + 2
	control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
	if len(message) > width-1 {
		// The message is larger than width.
		col = (wscol-width)/2 + 2
		control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
		// Split message by word.
		buf := strings.Fields(message)
		length := 0
		width -= 1
		index := 0
		for i := 0; i < len(buf); i++ {
			// The word is too long.
			if len(buf[i]) > width {
				fmt.Fprintf(os.Stderr, boxcolor+control)
				for j := index; j < i; j++ {
					fmt.Fprintf(os.Stderr, buf[j])
				}
				length = 0
				for {
					if length < len(buf[i]) {
						control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
						fmt.Fprintf(os.Stderr, boxcolor+control)
						if length+width <= len(buf[i]) {
							fmt.Fprintf(os.Stderr, buf[i][length:length+width])
						} else {
							fmt.Fprintf(os.Stderr, buf[i][length:len(buf[i])])
						}
						length += width
						row++
					} else {
						break
					}
				}
				index = i + 1
				length = 0
			} else {
				// Output the words.
				length += len(buf[i])
				if length > width {
					control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
					fmt.Fprintf(os.Stderr, boxcolor+control)
					for j := index; j < i; j++ {
						fmt.Fprintf(os.Stderr, buf[j]+" ")
					}
					index = i
					length = len(buf[i]) + 1
					row++
				} else {
					length += 1
				}
			}
		}
		// The last line's length is less than width.
		if index != len(buf) {
			control = "\033[" + fmt.Sprint(row) + "H" + "\033[" + fmt.Sprint(col) + "G"
			fmt.Fprintf(os.Stderr, boxcolor+control)
			for j := index; j < len(buf); j++ {
				fmt.Fprintf(os.Stderr, buf[j]+" ")
			}
		}
	} else {
		// The message is small.
		fmt.Fprintf(os.Stderr, boxcolor+control+message+"\n")
	}
}
