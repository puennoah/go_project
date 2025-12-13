package convert

import "testing"

func Test_centigradeToFahrenheit(t *testing.T) {
	result := centigradeToFahrenheit(1.0)
	if result != 33.8 {
		t.Errorf("expected 33.8, got %v", result)
	}
}