package lab

import "testing"

func TestCreateConfigs(t *testing.T) {
	store := NewStore()
	if err := store.CreateTable("configs"); err != nil {
		t.Log(err)
	}
}
