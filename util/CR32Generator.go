package util

import (
	"hash/crc32"
)

func GenerateCR32(url string) uint32 {
	crc32q := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum([]byte(url), crc32q)
}
