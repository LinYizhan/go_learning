package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestRemotePackage(t *testing.T) {
	m := cm.CreateConcurrentMap(100)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
