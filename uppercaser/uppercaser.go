package uppercaser

import (
	"fmt"
	"strings"
)

func Uuu(s string) string {
	fmt.Printf("Got {%v}", s)
	return strings.ToUpper(s)
}
