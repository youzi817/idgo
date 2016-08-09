package config

import (
	"fmt"
	"testing"
)

func TestLoadCfg(t *testing.T) {
	c := NewCfg("test.conf")
	if err := c.Load(); err != nil {
		//		t.Error(err)
	}
	fmt.Printf("host:%s port:%s mode:%s\n", c.m["host"], c.m["port"], c.m["mode"])
}

func TestReadString(t *testing.T) {
	c := NewCfg("test.conf")
	if err := c.Load(); err != nil {
		t.Error(err)
	}
	//	fmt.Printf("host:%s port:%s mode:%s\n", c.m["host"], c.m["port"], c.m["mode"])
	host, err := c.ReadString("host")
	if host == "" {
		t.Error(err)
	}
	fmt.Printf("host:%s\n", host)

	hostNum, err := c.ReadString("hostNum")
	if hostNum == "" {
		t.Error(err)
	}
	fmt.Printf("hostNum:%s\n", hostNum)
	fmt.Printf("\n")
}
