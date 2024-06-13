package state

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorkflow(t *testing.T) {
	wf := NewWorkflow()
	assert.Equal(t, toConfirm, wf.Current())
	wf.Print()

	t.Run("Next", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			wf.Next()
			wf.Print()
		}
	})

	t.Run("RollBack", func(t *testing.T) {
		var stat int

		wf.Next()
		wf.Print()

		for {
			wf.Next()
			wf.Print()
			stat = wf.Current()

			if stat == rejected {
				wf.Rollback()
				wf.Print()
				stat = wf.Current()
				continue
			} else if stat == approved {
				break
			}
		}

		for {
			wf.Next()
			wf.Print()
			stat = wf.Current()

			if stat == failed {
				wf.Rollback()
				wf.Print()
				stat = wf.Current()
				continue
			} else if stat == finished {
				break
			}
		}

	})

	t.Run("RollBack", func(t *testing.T) {
		var stat int
		var counter int

		wf.Next()
		wf.Print()

		for counter = 0; wf.Current() != canceled; counter++ {
			if counter > 1 {
				fmt.Println("尝试次数超过限制 2")
				wf.Cancel()
				wf.Print()
				break
			}

			fmt.Println("尝试次数: ", counter+1)
			wf.Next()
			wf.Print()
			stat = wf.Current()

			if stat == rejected {
				wf.Rollback()
				wf.Print()
				stat = wf.Current()
				continue
			} else if stat == approved {
				break
			}
		}

		for counter = 0; wf.Current() != canceled; counter++ {
			if counter > 2 {
				fmt.Println("尝试次数超过限制 3")
				wf.Cancel()
				wf.Print()
				break
			}

			fmt.Println("尝试次数: ", counter+1)
			wf.Next()
			wf.Print()
			stat = wf.Current()

			if stat == failed {
				wf.Rollback()
				wf.Print()
				stat = wf.Current()
				continue
			} else if stat == finished {
				break
			}
		}

	})

}
