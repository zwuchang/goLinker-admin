package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	r "math/rand"
	"time"
)

// GenerateUniqueInviteCode 生成唯一邀请码，带重试机制
func GenerateUniqueInviteCode(checkUnique func(string) bool) (string, error) {
	const maxRetries = 10
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	for i := 0; i < maxRetries; i++ {
		// 生成随机邀请码
		result := make([]byte, length)
		for j := range result {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return "", fmt.Errorf("failed to generate random number: %v", err)
			}
			result[j] = charset[num.Int64()]
		}

		code := string(result)

		// 检查唯一性
		if checkUnique(code) {
			return code, nil
		}

		// 如果接近最大重试次数，添加时间戳确保唯一性
		if i >= maxRetries-2 {
			timestamp := time.Now().UnixNano() % 10000 // 取时间戳的后4位
			// 确保时间戳不会超过字符集长度
			if timestamp >= int64(len(charset)) {
				timestamp = timestamp % int64(len(charset))
			}
			code = fmt.Sprintf("%s%c", code[:length-1], charset[timestamp])
			if checkUnique(code) {
				return code, nil
			}
		}
	}

	return "", fmt.Errorf("failed to generate unique invite code after %d attempts", maxRetries)
}

// ValidateInviteCode 验证邀请码格式
func ValidateInviteCode(code string) bool {
	if len(code) < 6 || len(code) > 20 {
		return false
	}

	// 检查是否包含非法字符
	for _, char := range code {
		if !((char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}

	return true
}

// 生成指定位长度的随机数字
func GenerateRandomNumber(length int) string {
	charset := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}
