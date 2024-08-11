package humanize

import (
	"fmt"

	human "github.com/dustin/go-humanize"
)

func HelloHuman() {
	var number uint64 = 123456789
	fmt.Println("Size of file is", human.Bytes(number))
}
