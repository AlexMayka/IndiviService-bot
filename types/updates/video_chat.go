package updates

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type VideoChatStarted struct{}
