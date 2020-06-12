package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/PuloV/ics-golang"
)

func main() {

	// calendarFile = flag.String("f", os.Env, "the calendar to convert")
	help := flag.Bool("help", false, "this help.")
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	parser := ics.New()
	parserChan := parser.GetInputChan()
	outputChan := parser.GetOutputChan()
	go func() {
		nowYear := time.Now().Year()
		for event := range outputChan {
			if event.GetStart().Year() == nowYear {
				printEvent(event)
			}
		}
	}()

	parserChan <- "https://calendar.google.com/calendar/ical/colin%40jchenry.me/private-ff5ffa18eb856032d166c7f410fe33c0/basic.ics"

	parser.Wait()
}

func printEvent(evt *ics.Event) {
	fmt.Printf("%s - %s : %s (%s)\n", fmtTime(evt.GetStart()), fmtTime(evt.GetEnd()), evt.GetSummary(), evt.GetLocation())
}

func fmtTime(t time.Time) string {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic("bad timezone")
	}

	return t.In(loc).Format("Jan 02\t2006 15:04 MST")
}
