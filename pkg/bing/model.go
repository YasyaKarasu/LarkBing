package bing

type SourceAttribution struct {
	ProviderDisplayName string `json:"providerDisplayName"`
	SeeMoreURL          string `json:"seeMoreUrl"`
}

type SuggestedResponse struct {
	Text string `json:"text"`
}

type ChatMessage struct {
	Text               string              `json:"text"`
	Author             string              `json:"author"`
	MessageID          string              `json:"messageId"`
	MessageType        string              `json:"messageType"`
	SourceAttributions []SourceAttribution `json:"sourceAttributions"`
	SuggestedResponses []SuggestedResponse `json:"suggestedResponses"`
	RequestID          string              `json:"requestId"`
}

type Item struct {
	Messages  []ChatMessage `json:"messages"`
	RequestID string        `json:"requestId"`
}

type ChatResponse struct {
	Item Item `json:"item"`
}
