package utils

import (
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

func GetSha512(key []byte, salt []byte) (string, error) {
	crypt := sha512_crypt.New()
	return crypt.Generate([]byte(key), []byte(salt))
}
