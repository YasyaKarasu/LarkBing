package messageCardDispatcher

import "encoding/json"

type FeishuMessageCardRequestRaw struct {
	OpenId        string         `json:"open_id"`
	UserId        string         `json:"user_id"`
	OpenMessageId string         `json:"open_message_id"`
	TenantKey     string         `json:"tenant_key"`
	Token         string         `json:"token"`
	Action        map[string]any `json:"action"`
	Challenge     string         `json:"challenge"`
}

type FeishuMessageCardRequest struct {
	OpenMessageId string
	Action        map[string]any
	Challenge     string
}

func deserializeRequest(dataStr string, request *FeishuMessageCardRequest) {
	var data FeishuMessageCardRequestRaw
	json.Unmarshal([]byte(dataStr), &data)

	request.OpenMessageId = data.OpenMessageId
	request.Action = data.Action
	request.Challenge = data.Challenge
}
