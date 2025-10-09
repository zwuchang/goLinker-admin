package utils

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// MD5VShort 截取MD5哈希的前N位
func MD5VShort(str []byte, length int, b ...byte) string {
	hash := MD5V(str, b...)
	if len(hash) >= length {
		return hash[:length]
	}
	return hash
}

// MD5VWithLetters 将MD5哈希中的数字转换为字母
func MD5VShortString(str string) string {
	result := ""

	hash := []byte(str)

	for _, char := range hash {
		if char >= '0' && char <= '9' {
			// 0-9 转为 A-J
			result += string(char - '0' + 'A')
		} else {
			result += string(char)
		}
	}

	return result
}
