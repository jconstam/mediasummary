package mediasummary

import (
	"html/template"
	"log"
	"os"
)

func generate(sonarrData sonarrSeriesData, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Printf("Couldn't write to file %v", file)
	}

	defer f.Close()

	tHeader, _ := template.ParseFiles("header.html")
	tHeader.Execute(f, nil)

	for _, show := range sonarrData {
		tBody, _ := template.ParseFiles("body.html")
		tBody.Execute(f, show)
	}

	tFooter, _ := template.ParseFiles("footer.html")
	tFooter.Execute(f, nil)
}
