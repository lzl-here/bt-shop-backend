package utils

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// 两种用法：

// 防止并发问题，执行完业务逻辑里吗调用cleaner:
// 同样的请求不能同时执行，但是可以先后执行

// 防止重复执行，不手动执行cleaner，让redis自动过期
// 同样的请求不能重复执行（有超时时间）

// 业务逻辑产生异常记得调用cleaner
func NoDuplicate(ctx context.Context, redisCli *redis.Client, prefix string, value any, ttl time.Duration) (ok bool, cleaner func() error, key string, err error) {
	// no panic guaranteed
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("NoDuplicates panic: %v", r)
		}
	}()

	// always return a non-nil cleaner
	cleaner = func() error { return nil }

	// security check
	if redisCli == nil {
		return false, nil, "", fmt.Errorf("redisCli is nil")
	}

	// serilize and hash
	serialized, err := json.Marshal(value)
	if err != nil {
		return false, nil, "", fmt.Errorf("failed to serialize value: %w", err)
	}
	checksum := fmt.Sprintf("%x", md5.Sum(serialized))

	// generate Redis key
	sb := new(strings.Builder)
	sb.WriteString(prefix)
	sb.WriteByte(':')
	sb.WriteString(checksum)
	key = sb.String()
	sb = nil

	// setnx
	ok, err = redisCli.SetNX(ctx, key, "🫡", ttl).Result() // value is \xf0\x9f\xab\xa1, or, U+1FAE1, or, "Saluting Face" emoji
	if err != nil {
		return false, nil, "", fmt.Errorf("failed to setnx key: %w", err)
	}

	// set cleaner to be a non-empty function, if SetNX succeeds
	if ok {
		cleaner = func() error { return delKey(ctx, redisCli, key, 5) }
	}

	return ok, cleaner, key, nil
}

// 删除key，由cleaner调用
func delKey(ctx context.Context, redisCli *redis.Client, key string, retry int) (err error) {
	// no panic guaranteed
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("NoDuplicates panic: %v", r)
		}
	}()

	// abort directly for safety
	if retry < 1 {
		return fmt.Errorf("retry must be at least 1")
	}

	// retry loop
	for i := 0; i < retry; i++ {
		_, err = redisCli.Del(ctx, key).Result()
		if err == nil {
			return nil
		}
	}

	// return the err from the last retry loop
	return fmt.Errorf("failed to delete key after %d retries: %w", retry, err)
}
