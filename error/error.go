package error

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
