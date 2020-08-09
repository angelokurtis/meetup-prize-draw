package main

import (
	"flag"
	"fmt"
	"github.com/angelokurtis/meetup-prize-draw/pkg/meetup"
	"log"
	"math/rand"
	"time"
)

var (
	link string
)

func init() {
	flag.StringVar(&link, "event", "", "the meetup event link")
	flag.Parse()
}

func main() {
	event, err := meetup.NewEvent(link)
	if err != nil {
		log.Fatal(err)
	}
	attendees, err := event.Attendees()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(attendees), func(i, j int) { attendees[i], attendees[j] = attendees[j], attendees[i] })

	winning := attendees[0]
	fmt.Printf("%s (%s)", winning.Name, winning.ProfileLink)
}
