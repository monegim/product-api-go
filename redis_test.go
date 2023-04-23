package main

import (
	"testing"

	"github.com/monegim/product-api-go/pkg/gredis"
	"github.com/monegim/product-api-go/pkg/setting"
)

func TestSetup(t *testing.T) {
	setting.Setup()
	err := gredis.Setup()
	if err != nil {
		t.Logf("error:%v\n", err)
	}
}

func TestExists(t *testing.T) {
	key := "name1"
	exists := gredis.Exists(key)
	if !exists {
		t.Error("key does not exists")
	}
}
