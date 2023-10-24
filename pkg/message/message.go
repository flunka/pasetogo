package message

import "strings"

type Message struct {
	version Version
	purpose Purpose
	payload string
	footer  string
}

type Version string

const (
	V3 Version = "v3"
	V4 Version = "v4"
)

type Purpose string

const (
	LOCAL  Purpose = "local"
	PUBLIC Purpose = "public"
)

func PAE(pieces []string) string {
	var sb strings.Builder
	count := uint64(len(pieces))
	sb.WriteString(LE64(count))
	for i := 0; i < int(count); i++ {
		sb.WriteString(LE64(uint64(len(pieces[i]))))
		sb.WriteString(pieces[i])
	}
	return sb.String()
}

func LE64(n uint64) string {
	var sb strings.Builder
	for i := 0; i < 8; i++ {
		if i == 7 {
			n &= 127
		}
		sb.WriteRune(rune(n & 255))
		n = n >> 8
	}
	return sb.String()
}
