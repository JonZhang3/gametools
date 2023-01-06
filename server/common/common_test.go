package common

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	data := &struct{ Date DateTime }{Date: DateTime(time.Now())}
	bytes, _ := json.Marshal(data)
    t.Log(string(bytes))
}
