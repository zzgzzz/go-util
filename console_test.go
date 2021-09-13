package util

import (
	"flag"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {

	args := []string{
		"-conf=args测试",
	}
	flag.CommandLine.Parse(args)

	c:=GetConfig()
	fmt.Println(c)
}
