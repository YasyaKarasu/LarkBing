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
}

type Item struct {
	Messages  []ChatMessage `json:"messages"`
	RequestID string        `json:"requestId"`
}

type ChatResponse struct {
	Item Item `json:"item"`
}

type MessageEvent struct {
	Sender struct {
		Sender_id struct {
			Union_id string `json:"union_id"`
			Open_id  string `json:"open_id"`
			User_id  string `json:"user_id"`
		} `json:"sender_id"`
		Sender_type string `json:"sender_type"`
		Tenant_key  string `json:"tenant_key"`
	} `json:"sender"`
	Message struct {
		Message_id   string `json:"message_id"`
		Root_id      string `json:"root_id"`
		Parent_id    string `json:"parent_id"`
		Create_time  string `json:"create_time"`
		Chat_id      string `json:"chat_id"`
		Chat_type    string `json:"chat_type"`
		Message_type string `json:"message_type"`
		Content      string `json:"content"`
		Metions      struct {
			Key string `json:"key"`
			Id  struct {
				Union_id string `json:"union_id"`
				Open_id  string `json:"open_id"`
				User_id  string `json:"user_id"`
			} `json:"id"`
			Name       string `json:"name"`
			Tenant_key string `json:"tenant_key"`
		} `json:"metions"`
	} `json:"message"`
}
