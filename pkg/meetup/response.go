package meetup

type responseBody struct {
	Responses []*responses `json:"responses"`
}

type webActions struct {
	GroupProfileLink string `json:"group_profile_link"`
}

type eventContext struct {
	Host bool `json:"host"`
}

type member struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	WebActions   *webActions   `json:"web_actions"`
	Role         string        `json:"role"`
	EventContext *eventContext `json:"event_context"`
}

type value struct {
	Updated  int     `json:"updated"`
	Member   *member `json:"member"`
	Response string  `json:"response"`
}

type responses struct {
	Value []*value `json:"value"`
}
