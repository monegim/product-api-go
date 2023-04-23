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

func TestSet(t *testing.T) {
	key := "name"
	value := "mostafaaaa"
	err := gredis.Set(key, value, 0)
	if err != nil {
		t.Logf("TestSet err:%v\n", err)
	}
}
func TestExists(t *testing.T) {
	key := "name1"
	exists := gredis.Exists(key)
	if !exists {
		t.Error("key does not exists")
	}
}

func TestGet(t *testing.T) {
	key := "name"
	val, err := gredis.Get(key)
	if err != nil {
		t.Errorf("error for TestGet: %v\n", err)
	}
	if string(val) != "mostafa" {
		t.Error("val does not match")
	}
}
