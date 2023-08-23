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
package version

import (
	"fmt"
)

func Version(Version_code string, Commit_id string) {
	fmt.Print("\n\033[1;38;2;254;228;208m")
	fmt.Print("●   ●  ●●●   ●●●  ●   ●  ●●●  ●   ●  ●●●  ●   ●\n")
	fmt.Print(" ● ●  ●   ● ●     ●   ●   ●   ●●  ● ●   ● ●●  ●\n")
	fmt.Print("  ●   ●   ●  ●●●  ●●●●●   ●   ● ● ● ●   ● ● ● ●\n")
	fmt.Print("  ●   ●   ●     ● ●   ●   ●   ●  ●● ●   ● ●  ●●\n")
	fmt.Print("  ●    ●●●   ●●●  ●   ●  ●●●  ●   ●  ●●●  ●   ●\n")
	fmt.Print("             #Meet the modern TUI\n")
	fmt.Print("        Licensed under the MIT License\n")
	fmt.Print("          <https://mit-license.org>\n")
	fmt.Print("        Copyright (C) 2023 Moe-hacker\n")
	fmt.Print("             yoshinon : " + Version_code + "\n")
	fmt.Print("            Commit ID : " + Commit_id + "\n")
	fmt.Print("           Powered by Bubble Tea\n\033[0m")
}
