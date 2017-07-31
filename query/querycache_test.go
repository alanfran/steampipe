package query

import (
	"testing"
	"time"

	servers "github.com/alanfran/SteamCondenserGo"
)

func TestPut(t *testing.T) {
	c := NewCache(time.Duration(time.Second * 10))
	c.Put("test", servers.Response{})

	if len(c.cache) != 1 {
		t.Error("Failed to put Response in cache.")
	}
}

func TestQuery(t *testing.T) {
	address := "zs.nekonet.xyz:27015"

	c := NewCache(time.Duration(time.Second * 10))
	_, err := c.Get(address)
	if len(c.cache) != 1 {
		t.Error("Failed to query server.", err)
	}
}

func TestDelete(t *testing.T) {
	c := NewCache(time.Duration(time.Second * 10))
	c.Put("test", servers.Response{})
	c.Del("test")
	if len(c.cache) != 0 {
		t.Error("Failed to delete Response from cache.")
	}
}

func TestPurge(t *testing.T) {
	c := NewCache(time.Duration(0))
	c.Put("test", servers.Response{})
	// wait 1 nanosecond for cache to expire
	time.Sleep(1)
	c.Purge()
	if len(c.cache) != 0 {
		t.Error("Failed to purge cache.")
	}
}
