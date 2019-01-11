package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const (
	seed = "ThisIsSecretSeedForTest"
)

var (
	md5Seed = fmt.Sprintf("%x", md5.Sum([]byte(seed)))
)

func Encrypt(timestamp int64) string {
	start := timestamp % 24
	return md5Seed[start:start + 8] + strconv.FormatInt(timestamp, 36)
}