package meetup

import (
	"errors"
	"strings"
)

type event struct {
	ID    string
	Group string
}

func NewEvent(link string) (*event, error) {
	s := strings.Split(link, "/events/")
	if len(s) != 2 {
		return nil, errors.New("wrong Meetup Event URL format")
	}

	g := strings.Split(s[0], "/")
	group := g[len(g)-1]

	i := strings.Split(s[1], "/")
	id := i[0]

	return &event{ID: id, Group: group}, nil
}
