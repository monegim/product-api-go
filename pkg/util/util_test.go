package util

import (
	"log"
	"testing"
)

func TestEncodeMD5(t *testing.T) {
	value := "mostafa"
	log.Println(EncodeMD5(value))
}
