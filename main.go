package main

import (
	"flag"
	"log"
	"os/exec"

	"github.com/mmcdole/gofeed"
	"github.com/rivo/tview"
)

var url = flag.String("u", "", "Get RSS url.")

func main() {
	flag.Parse()

	if *url == "" {
		log.Fatalln("No url.")
	}

	app := tview.NewApplication()
	list := tview.NewList()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(*url)
	if err != nil {
		log.Fatalln(err)
	}

	items := feed.Items

	list.AddItem("Quit", "Press to exit", 'q', func() { app.Stop() })

	for _, item := range items {
		link := item.Link
		list.AddItem(item.Title, item.Published, 0, func() {
			if err := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", link).Start(); err != nil {
				log.Fatalln(err)
			}
		})
	}

	if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
		log.Fatalln(err)
	}
}
