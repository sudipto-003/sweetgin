package repository

import (
	"context"
	"time"
)

var otp_expire = time.Duration(5) * time.Minute

func (client *Repos) SetOTP(ctx context.Context, pid, otp string) error {
	err := client.RedisClient.Set(ctx, pid, otp, otp_expire).Err()

	return err
}

func (client *Repos) GetOTP(ctx context.Context, pid string) (string, error) {
	value, err := client.RedisClient.Get(ctx, pid).Result()

	return value, err
}
