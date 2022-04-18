package stringwidth_test

import (
	"fmt"
	"github.com/shibukawa/stringwidth"
)

func ExampleCalc() {
	// Detect emoji sequence
	// []rune{ 0x1F3F4 0xE0067 0xE0062 0xE0065 0xE006E 0xE0067 0xE007F }
	fmt.Println(stringwidth.Calc("🏴󠁧󠁢󠁥󠁮󠁧󠁿"))

	// It can modify east asian ambiguous width with option
	// false is default.
	fmt.Println(stringwidth.Calc("¼", stringwidth.Opt{IsAmbiguousWide: false}))
	fmt.Println(stringwidth.Calc("¼", stringwidth.Opt{IsAmbiguousWide: true}))

	// It can ignore ANSI escape codes
	fmt.Println(stringwidth.Calc("\x1b[38;5;140m foo\x1b[0m bar"))
	// Output: 2
	// 1
	// 2
	// 8
}
