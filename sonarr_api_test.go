package mediasummary

import (
	"testing"
)

func Test_GetData(t *testing.T) {
	data := getData("http://192.168.1.23:8989", "dba566f126fb46e4ac785174d2ef3cd6")

	if len(data) == 0 {
		t.Error("Got no data")
	}
}
