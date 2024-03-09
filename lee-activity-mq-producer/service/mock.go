package service

import (
	"encoding/json"
	"lee-activity-framework/lee-activity-api/event"
	"lee-activity-framework/lee-activity-api/mq/kafka"
	kafkaUtil "lee-activity-framework/lee-activity-mq-producer/kafka"
)

func SendGift(req event.SendGiftReq) (*event.SendGiftResp, error) {
	kafkaMsg := kafka.SendGiftMsg{
		FromUid:    req.FromUid,
		ToUid:      req.ToUid,
		GiftId:     req.GiftId,
		GiftAmount: req.GiftAmount,
	}
	b, _ := json.Marshal(kafkaMsg)

	kafkaUtil.ProduceMessage(kafkaMsg.Topic(), string(b))
	return &event.SendGiftResp{}, nil
}

func SendMessage(req event.SendMessageReq) (*event.SendMessageResp, error) {
	kafkaMsg := kafka.SendMessageMsg{
		FromUid:        req.FromUid,
		ToUid:          req.ToUid,
		MessageId:      req.MessageId,
		MessageContent: req.MessageContent,
		MessageType:    req.MessageType,
	}
	b, _ := json.Marshal(kafkaMsg)

	kafkaUtil.ProduceMessage(kafkaMsg.Topic(), string(b))
	return &event.SendMessageResp{}, nil
}
