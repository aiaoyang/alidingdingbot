package main

// AliBotHeader 阿里云机器人请求头
type AliBotHeader struct {
	ContentType string `json:"Content-Type" form:"Content-Type"`
	Timestamp   string `json:"timestamp" form:"timestamp"`
	Sign        string `json:"sign" form:"sign"`
}

type AliBotRequsetMSGBody struct {
	MsgType           string   `json:"msgtype" form:"msgtype"`
	Text              Content  `json:"text" form:"text"`
	MsgID             string   `json:"msgId" form:"msgId"`
	CreateAt          int64    `json:"createAt" form:"createAt"`
	ConversationType  string   `json:"conversationType" form:"conversationType"`
	ConversationID    string   `json:"conversationId" form:"conversationId"`
	ConversationTitle string   `json:"conversationTitle" form:"conversationTitle"`
	SenderID          string   `json:"senderId" form:"senderId"`
	SenderNick        string   `json:"senderNick" form:"senderNick"`
	SenderCorpID      string   `json:"senderCorpId" form:"senderCorpId"`
	SenderStaffID     string   `json:"senderStaffId" form:"senderStaffId"`
	ChatbotUserID     string   `json:"chatbotUserId" form:"chatbotUserId"`
	AtUsers           []AtUser `json:"atUsers" form:"atUsers"`
}
type AliBotResponse struct {
	MsgType string  `json:"msgtype" form:"msgtype"`
	Text    Content `json:"text" form:"text"`
	At      `json:"at" form:"at"`
}
type At struct {
	AtMobiles []string `json:"atMobiles" form:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll" form:"isAtAll"`
}
type AtUser struct {
	DingtalkID string `json:"dingtalkId" form:"dingtalkId"`
	StaffID    string `json:"staffId" form:"staffId"`
}
type Content struct {
	Msg string `json:"content"`
}
