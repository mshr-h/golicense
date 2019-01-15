package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"
)

// License contains information for license representation
type License struct {
	Author string
	Year   int
}

func main() {
	var (
		authorName  = flag.String("author", "none", "Author name")
		licenseType = flag.String("license", "unknown", "License type")
	)
	flag.Parse()

	var thisYear = time.Now().Year()
	var license = License{Author: *authorName, Year: thisYear}

	var f, err = os.Create("LICENSE")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var tmpl *template.Template
	switch *licenseType {
	case "MIT":
		tmpl = template.Must(template.ParseFiles("template/MIT"))
	default:
		tmpl = nil
	}

	if tmpl == nil {
		fmt.Println("Invalid license type")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err = tmpl.Execute(f, &license)
	if err != nil {
		panic(err)
	}
}
