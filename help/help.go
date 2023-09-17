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
 * of this software and associated documentation files (the "Software\n"), to deal
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
//TODO

package help

import (
	"fmt"
	"os"
)

func Error(msg string) {
	fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m\n", msg)
	// A very cute catgirl nya~~
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", "  .^.   .^.")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", "  /⋀\\_ﾉ_/⋀\\")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", " /ﾉｿﾉ\\ﾉｿ丶)|")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", " ﾙﾘﾘ >  x )ﾘ")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", "ﾉノ㇏  ^ ﾉ|ﾉ")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", "      ⠁⠁")
	fmt.Fprintf(os.Stderr, "\033[1;38;2;254;228;208m%s\033[0m\n", "If you think something is wrong, please report at:")
	fmt.Fprintf(os.Stderr, "\033[4;1;38;2;254;228;208m%s\033[0m\n", "https://github.com/Moe-hacker/yoshinon/issues")
	os.Exit(1)

}
func Help() {
	fmt.Print("\033[1;38;2;254;228;208mBox options:\n")
	fmt.Print("\n")
	fmt.Print("  --msgbox      [options] <text> <height> <width>\n")
	fmt.Print("  --yesno       [options] <text> <height> <width>\n")
	fmt.Print("  --infobox     [options] <text> <height> <width>\n")
	fmt.Print("  --inputbox    [options] <text> <height> <width> [init]\n")
	fmt.Print("  --passwordbox [options] <text> <height> <width>\n")
	fmt.Print("  --menu        [options] <text> <height> <width> <listheight> [tag item] ...\n")
	fmt.Print("  --checklist   [options] <text> <height> <width> <listheight> [tag item status]...\n")
	fmt.Print("  --radiolist   [options] <text> <height> <width> <listheight> [tag item status]...\n")
	fmt.Print("  --gauge       [options] <text> <height> <width>\n")
	fmt.Print("\n")
	fmt.Print("Options:\n")
	fmt.Print("\n")
	fmt.Print("  --bgcolor     <R;G;B>\n")
	fmt.Print("  --boxcolor    <R;G;B>\n")
	fmt.Print("  --cursorcolor <R;G;B>\n")
	fmt.Print("  --border      <rounded/normal/thick/double/hidden>\n")
	fmt.Print("  --title       <title>\n")
	fmt.Print("  --help        print this message\n")
	fmt.Print("  --version     print version information\033[0m\n")

}
