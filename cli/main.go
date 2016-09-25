package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	template := flag.String("t", "", "The path of service template")

	flag.Parse();

	if *template == "" {
		fmt.Println("Please set the path of serive template by -t. eg. -t=~/st.json")
		os.Exit(-1);
	}
}