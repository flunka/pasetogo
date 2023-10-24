package message

import (
	"testing"
)

func TestPAE(t *testing.T) {

	paeTests := []struct {
		pieces []string
		want   string
	}{
		{[]string{}, string([]byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'})},
		{[]string{""}, string([]byte{'\x01', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'})},
		{[]string{"test"}, string([]byte{'\x01', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x04', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', 't', 'e', 's', 't'})},
	}
	for _, tt := range paeTests {
		got := PAE(tt.pieces)
		if got != tt.want {
			t.Errorf("got %x want %x", got, tt.want)
		}
	}

}
