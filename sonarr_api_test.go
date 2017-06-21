package mediasummary

import (
	"testing"
)

func Test_GetData(t *testing.T) {
	getData("http://192.168.1.23:8989", "dba566f126fb46e4ac785174d2ef3cd6")
}
