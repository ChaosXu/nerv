package golang

import (
	"fmt"
	"regexp"
	"testing"
)

func TestReg(t *testing.T) {

	reg := regexp.MustCompile(`\$\{(.+)\}`)
	//fmt.Println(reg.FindStringSubmatch("aa${a_b+c}bb")[1])
	fmt.Println(reg.ReplaceAllStringFunc("aa${a_b+c}bb", func(name string) string {
		fmt.Println(name[2 : len(name)-1])
		return "-"
	}))

	regex := regexp.MustCompile(`.*\t([0-9]+).*`)
	id := regex.FindStringSubmatch(`
	[
		123,
		null
	]
	`)[1]
	fmt.Println(id)
}
