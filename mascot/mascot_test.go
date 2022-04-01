package mascot_test

import (
	"testing"

	"github.com/rawka/bootcamp/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Tux" {
		t.Fatal("It's not a mascot")
	}
}
