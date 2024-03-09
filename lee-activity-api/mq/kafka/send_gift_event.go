package kafka

type SendGiftMsg struct {
	FromUid    int `json:"from_uid"`    // 送礼方uid
	ToUid      int `json:"to_uid"`      // 收礼方uid
	GiftId     int `json:"gift_id"`     // 礼物id
	GiftAmount int `json:"gift_amount"` // 礼物数量
}

func (msg SendGiftMsg) Topic() string {
	return "send_gift_event"
}
