package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a util integer [min, max + 1)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomNetworkType generate a sync type which is one of these values: {0, 1, 2, 3}
//
//	"0" -> NetworkType.NOT_REQUIRED
//	"2" -> NetworkType.UNMETERED
//	"3" -> NetworkType.NOT_ROAMING
//	"4" -> NetworkType.METERED
//	"5" -> if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) { NetworkType.TEMPORARILY_UNMETERED}
//	 else { NetworkType.CONNECTED }
func RandomNetworkType() string {
	return strconv.Itoa(int(RandomInt(0, 7)))
}

// RandomTimeUnit generate a TimeUnit type which is one of these values: {0, 1, 2, ..., 6}
//
//	"0" -> TimeUnit.NANOSECONDS
//	"1" -> TimeUnit.MICROSECONDS
//	"2" -> TimeUnit.MICROSECONDS
//	"3" -> TimeUnit.SECONDS
//	"4" -> TimeUnit.MINUTES
//	"5" -> TimeUnit.HOURS
//	else -> TimeUnit.DAYS
func RandomTimeUnit() string {
	return strconv.Itoa(int(RandomInt(0, 7)))
}

// RandomBoolean generate a boolean
func RandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func RandomUrlHashGenerator(n int) []string {
	var result []string
	result = make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = strconv.Itoa(int(GenerateCR32(RandomString(int(RandomInt(0, 200))))))
	}
	return result
}

// RandomString generates a util string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomStringList generates a list of strings of length n
func RandomStringList(n int, length int) []string {
	var sb strings.Builder
	k := len(alphabet)
	var result []string
	result = make([]string, n)
	for p := 0; p < length; p++ {
		for i := 0; i < n; i++ {
			c := alphabet[rand.Intn(k)]
			sb.WriteByte(c)
		}
		result[p] = sb.String()
	}
	return result
}
