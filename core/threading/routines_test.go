package threading

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/351423113/go-zero/core/lang"
	"github.com/stretchr/testify/assert"
)

func TestRoutineId(t *testing.T) {
	assert.True(t, RoutineId() > 0)
}

func TestRunSafe(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	i := 0

	defer func() {
		assert.Equal(t, 1, i)
	}()

	ch := make(chan lang.PlaceholderType)
	go RunSafe(func() {
		defer func() {
			ch <- lang.Placeholder
		}()

		panic("panic")
	})

	<-ch
	i++
}
