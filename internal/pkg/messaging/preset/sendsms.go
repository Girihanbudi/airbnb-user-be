package preset

import (
	"airbnb-user-be/internal/pkg/json"
)

type SendSms struct {
	Type    string   `json:"type"`
	Context string   `json:"context"`
	Payload json.Raw `json:"payload"`
}

type SendSmsPayload struct {
	Recipients []string `json:"recipients"`
	Body       string   `json:"body"`
}
