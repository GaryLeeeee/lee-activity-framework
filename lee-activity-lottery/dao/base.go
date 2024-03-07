package dao

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const UidLockKey = "lee-activity:lottery:uid_lock:%d"

func UidLock(uid int) (string, error) {
	lockKey := fmt.Sprintf(UidLockKey, uid)
	uuid := uuid.New().String()
	lock, err := _rd.SetNX(lockKey, uuid, 10*time.Second).Result()
	if err != nil {
		return "", err
	}
	if !lock {
		return "", errors.New("uid already lock")
	}

	return uuid, nil
}

func UidUnlock(uid int, uuid string) error {
	lockKey := fmt.Sprintf(UidLockKey, uid)
	value, err := _rd.Get(lockKey).Result()
	if err != nil {
		return err
	}
	if value != uuid {
		return errors.New("uuid check fail")
	}
	_, err = _rd.Del(lockKey).Result()
	return err
}
