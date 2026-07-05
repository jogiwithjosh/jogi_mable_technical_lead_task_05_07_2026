package dto

type TrackEventsRequest struct {
	Events    []TrackEvent `json:"events" binding:"required,min=1"`
	CreatedAt string       `json:"createdAt"`
}

type TrackEvent struct {
	ID          string         `json:"id" binding:"required"`
	Type        string         `json:"type" binding:"required"`
	Timestamp   string         `json:"timestamp" binding:"required"`
	Name        string         `json:"name"`
	Page        string         `json:"page"`
	URL         string         `json:"url"`
	Referrer    string         `json:"referrer"`
	UserID      string         `json:"userId"`
	AnonymousID string         `json:"anonymousId"`
	SessionID   string         `json:"sessionId"`
	Properties  map[string]any `json:"properties"`
}
