package server

import (
	"testing"
	"regexp"
	"fmt"
)

func TestReg(t *testing.T) {

	reg := regexp.MustCompile(`^\$\{(.+)\}$`)
	fmt.Println(reg.FindStringSubmatch("${a_b+c}")[1])
}
