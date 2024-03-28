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
package main

import (
	"fmt"
	"os"
	"strconv"
	. "yoshinon/checklist"
	. "yoshinon/error"
	. "yoshinon/gauge"
	. "yoshinon/help"
	. "yoshinon/inputbox"
	. "yoshinon/menu"
	. "yoshinon/msgbox"
	. "yoshinon/passwordbox"
	. "yoshinon/radiolist"
	. "yoshinon/version"
	. "yoshinon/yesno"
)

var Version_code = "1.1"
var Commit_id = ""

func run_gauge(args []string) {
	m := Gauge_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			break
		}
	}
	Gauge(m)
}
func run_passwordbox(args []string) string {
	m := Passwordbox_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			break
		}
	}
	return Passwordbox(m)
}
func run_inputbox(args []string) string {
	m := Inputbox_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			if len(args) > i {
				m.Init = args[i]
			} else {
				m.Init = ""
			}
			break
		}
	}
	return Inputbox(m)
}
func run_menu(args []string) string {
	m := Menu_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			m.Listheight, _ = strconv.Atoi(args[i])
			i++
			m.Items = make([]string, (len(args)-i)/2)
			m.Tags = make([]string, (len(args)-i)/2)
			k := 0
			for j := i; j < len(args)-1; j++ {
				m.Tags[k] = args[j]
				m.Items[k] = args[j+1]
				j++
				k++
			}
			break
		}
	}
	return Menu(m)
}
func run_radiolist(args []string) string {
	m := Radiolist_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			m.Listheight, _ = strconv.Atoi(args[i])
			i++
			m.Items = make([]string, (len(args)-i)/3)
			m.Tags = make([]string, (len(args)-i)/3)
			m.Status = make([]int, (len(args)-i)/3)
			k := 0
			for j := i; j < len(args); j++ {
				m.Tags[k] = args[j]
				m.Items[k] = args[j+1]
				if args[j+2] == "ON" {
					m.Status[k] = 1
				} else {
					m.Status[k] = 0
				}
				j += 2
				k++
			}
			break
		}
	}
	return Radiolist(m)
}
func run_checklist(args []string) string {
	m := Checklist_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			m.Listheight, _ = strconv.Atoi(args[i])
			i++
			m.Items = make([]string, (len(args)-i)/3)
			m.Tags = make([]string, (len(args)-i)/3)
			m.Status = make([]int, (len(args)-i)/3)
			k := 0
			for j := i; j < len(args); j++ {
				m.Tags[k] = args[j]
				m.Items[k] = args[j+1]
				if args[j+2] == "ON" {
					m.Status[k] = 1
				} else {
					m.Status[k] = 0
				}
				j += 2
				k++
			}
			break
		}
	}
	return Checklist(m)
}
func run_msgbox(args []string) {
	m := Msgbox_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			break
		}
	}
	Msgbox(m)
}
func run_yesno(args []string) int {
	m := Yesno_config{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--bgcolor" {
			i++
			m.Bgcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--boxcolor" {
			i++
			m.Boxcolor = "\033[1;48;2;" + args[i] + "m" + "\033[1;38;2;0;0;0m"
		} else if args[i] == "--cursorcolor" {
			i++
			m.Cursorcolor = "\033[1;48;2;" + args[i] + "m"
		} else if args[i] == "--border" {
			i++
			m.Border = args[i]
		} else if args[i] == "--title" {
			i++
			m.Title = args[i]
		} else {
			m.Message = args[i]
			i++
			m.Height, _ = strconv.Atoi(args[i])
			i++
			m.Width, _ = strconv.Atoi(args[i])
			i++
			break
		}
	}
	return Yesno(m)
}
func main() {
	if len(os.Args) < 2 {
		Help()
		Error("Missing arguments")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "--menu":
		ret := run_menu(os.Args[1:])
		fmt.Fprint(os.Stdout, ret)
	case "--inputbox":
		ret := run_inputbox(os.Args[1:])
		fmt.Fprint(os.Stdout, ret)
	case "--passwordbox":
		ret := run_passwordbox(os.Args[1:])
		fmt.Fprint(os.Stdout, ret)
	case "--checklist":
		ret := run_checklist(os.Args[1:])
		fmt.Fprint(os.Stdout, ret)
	case "--radiolist":
		ret := run_radiolist(os.Args[1:])
		fmt.Fprint(os.Stdout, ret)
	case "--msgbox":
		run_msgbox(os.Args[1:])
	case "--gauge":
		run_gauge(os.Args[1:])
	case "--yesno":
		os.Exit(run_yesno(os.Args[1:]))
	case "--version", "-v":
		Version(Version_code, Commit_id)
	case "--help", "-h":
		Help()
	case "-V":
		fmt.Print(Version_code)
	}
}
