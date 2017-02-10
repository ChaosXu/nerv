package golang

import (
	"testing"
	"regexp"
	"fmt"
)

func TestReg(t *testing.T) {

	reg := regexp.MustCompile(`^\$\{(.+)\}$`)
	fmt.Println(reg.FindStringSubmatch("${a_b+c}")[1])

	regex := regexp.MustCompile(`.*\t([0-9]+).*`)
	id := regex.FindStringSubmatch(`
	[
		123,
		null
	]
	`)[1]
	fmt.Println(id)
}
