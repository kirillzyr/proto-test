package uppercaser

import (
	"fmt"
	"strings"
)

func uuu(s string) string {
	fmt.Printf("Got {%v}", s)
	return strings.ToUpper(s)
}
