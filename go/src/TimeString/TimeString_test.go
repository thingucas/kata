package TimeStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeStringSum(t *testing.T) {
	s := TimeString{}

	assert.Equal(t, "0:0:0", s.PrintSum([]string{}))
	//assert.Equal(t, "1:1:1", s.PrintSum([]string{"0:1", "0:60", "0:3600"}))
}

func xTestAllTimeString(t *testing.T) {
	s := TimeString{}

	assert.Equal(t, "0:0:0", s.PrintSum([]string{}))
	assert.Equal(t, "2:32:41", s.PrintSum([]string{"12:32", "34:01", "15:23", "9:27", "55:22", "25:56"}))
}
