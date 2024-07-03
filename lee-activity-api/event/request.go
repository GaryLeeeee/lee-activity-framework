package event

// 送礼物
type SendGiftReq struct {
	Uuid       string `form:"uuid"`
	FromUid    int    `form:"from_uid"`    // 送礼方uid
	ToUid      int    `form:"to_uid"`      // 收礼方uid
	GiftId     int    `form:"gift_id"`     // 礼物id
	GiftAmount int    `form:"gift_amount"` // 礼物数量
}
type SendGiftResp struct {
}

// 发消息
type SendMessageReq struct {
	Uuid           string `form:"uuid"`
	FromUid        int    `form:"from_uid"`        // 发送方uid
	ToUid          int    `form:"to_uid"`          // 接收方uid
	MessageId      int    `form:"message_id"`      // 消息id
	MessageContent int    `form:"message_content"` // 消息文本
	MessageType    int    `form:"message_type"`    // 消息类型
}
type SendMessageResp struct {
}
