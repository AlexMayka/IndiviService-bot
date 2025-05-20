package updates

type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	QuestionEntities      []*MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption     `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"` // "regular" or "quiz"
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

type PollOption struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}
