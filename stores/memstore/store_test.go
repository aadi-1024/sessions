package memstore_test

import (
	"github.com/aadi-1024/sessions/stores"
	"github.com/aadi-1024/sessions/stores/memstore"
	"testing"
	"time"
)

var storeInstance stores.Store

// replace storeInstance with whatever Store you're testing for
func init() {
	storeInstance = memstore.NewMemStore()
}

func TestMemStore_Save(t *testing.T) {
	table := []struct {
		name string
		sid  string
		err  bool //should it return error
	}{
		{"Should be nil", "s1", false},
		{"Should be nil", "s2", false},
		{"Should return err", "s1", true},
	}
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			err := storeInstance.Save(v.sid, time.Now())
			if (err == nil) == v.err {
				t.Errorf("error doesnt match %v expected %v", err, v.err)
			}
		})
	}
}

func TestMemStore_Load(t *testing.T) {
	table := []struct {
		name  string
		sid   string
		isNil bool
	}{
		{"isnt nil", "s1", false},
		{"is nil", "dont copy this anywhere", true},
	}
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			err := storeInstance.Load(v.sid)
			if (err == nil) != v.isNil {
				t.Errorf("returned unexpected %v", err)
			}
		})
	}
}

func TestMemStore_Delete(t *testing.T) {
	//s1 and s2 are already in the store
	err := storeInstance.Delete("s1")
	if err != nil {
		t.Errorf("s1 should have been present")
	}
	if storeInstance.Load("s1") != nil {
		t.Errorf("element deleted but still present")
	}
	err = storeInstance.Delete("s1")
	if err == nil {
		t.Errorf("s1 already deleted, should return err")
	}
}

func TestMemStore_Expire(t *testing.T) {
	if err := storeInstance.Save("s3", time.Now()); err != nil {
		t.Errorf("should'nt have happened, how did the Save test pass")
	}
	t.Logf("this test would take a few seconds")
	go func() {
		time.Sleep(time.Second * 3)
		storeInstance.Expire(time.Second * 3)
	}()
	time.Sleep(time.Second * 5)
	if ret := storeInstance.Load("s3"); ret != nil {
		t.Errorf("s3 not deleted")
	}
}
