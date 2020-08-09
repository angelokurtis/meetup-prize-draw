package meetup

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var client = &http.Client{}

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

func (e *event) Attendees() ([]*Attendee, error) {
	url := fmt.Sprintf("https://www.meetup.com/mu_api/urlname/events/eventId/attendees?queries=(endpoint:%s/events/%s/rsvps,meta:(method:get),params:(desc:!t,fields:'answers,pay_status,self,web_actions,attendance_status',only:'answers,response,attendance_status,guests,member,pay_status,updated',order:time),ref:eventAttendees_%s_%s,type:attendees)", e.Group, e.ID, e.Group, e.ID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var target *responseBody
	err = json.NewDecoder(res.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	attendees := make([]*Attendee, 0, 0)
	for _, attendee := range target.Responses {
		for _, value := range attendee.Value {
			member := value.Member
			if !member.EventContext.Host {
				attendees = append(attendees, &Attendee{
					ID:          member.ID,
					Name:        member.Name,
					ProfileLink: member.WebActions.GroupProfileLink,
				})
			}
		}
	}
	return attendees, nil
}

type Attendee struct {
	ID          int
	Name        string
	ProfileLink string
}
