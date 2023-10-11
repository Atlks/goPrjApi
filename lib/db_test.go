package lib

import (
	"testing"
)

func TestHandler100(t *testing.T) {
	//	rcd:={"id":1};

	rcd := map[string]any{
		"id":   1,
		"name": 2,
	}
	Pdo_Insert(rcd, "coll33", "c:/dbx/")
}
