package main

import (
	"flag"
	"fmt"

	"github.com/jcyamacho/db-migrate-name/pkg/generator"
)

func main() {
	gen := generator.New()

	flag.StringVar(&gen.Separator, "sep", gen.Separator, "separator")
	flag.BoolVar(&gen.LowerCase, "lc", gen.LowerCase, "lower case")
	flag.BoolVar(&gen.TimePrefix, "tp", gen.TimePrefix, "time prefix")
	flag.StringVar(&gen.TimePrefixFormat, "tpf", gen.TimePrefixFormat, "time prefix format")
	flag.Parse()

	args := flag.Args()

	name := gen.Name(args)

	fmt.Println(name)
}
