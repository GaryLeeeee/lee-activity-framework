package kafka

type SendMessageMsg struct {
	Uuid           string `json:"uuid"`
	FromUid        int    `json:"from_uid"`        // 发送方uid
	ToUid          int    `json:"to_uid"`          // 接收方uid
	MessageId      int    `json:"message_id"`      // 消息id
	MessageContent int    `json:"message_content"` // 消息文本
	MessageType    int    `json:"message_type"`    // 消息类型
}

func (msg SendMessageMsg) Topic() string {
	return "send_message_event"
}
