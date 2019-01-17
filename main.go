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
		err = tmpl.Execute(f, &license)
	case "GPL-3":
		tmpl = template.Must(template.ParseFiles("template/GPL-3"))
		err = tmpl.Execute(f, &license)
	case "Apache-2":
		tmpl = template.Must(template.ParseFiles("template/Apache-2"))
		err = tmpl.Execute(f, &license)
	default:
		fmt.Println("Invalid license type")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}
}
