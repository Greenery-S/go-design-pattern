package state

import "testing"

func TestDayContext_Next(t *testing.T) {
	weekDay := NewDayContext()

	for i := 0; i < 23; i++ {
		weekDay.Today()
		weekDay.Next()
	}
}
