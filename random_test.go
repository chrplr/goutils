package goutils
 
import "testing"
 
func TestRandomU(t *testing.T) {
	result := RandomU(10, 0)
	if len(result) != 10 {
		t.Error("len(result) should be 10, got", len(result))
	}
}

func TestRandomIntn(t *testing.T) {
	result := RandomIntn(10, 10, 0)
	if len(result) != 10 {
		t.Error("len(result) should be 10, got", len(result))
	}
}

