package timestring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/timestring"
)

func TestTimeStringSum(t *testing.T) {
	assert.Equal(t, "0:0:0", timestring.SumOfTimes(""))
	//assert.Equal(t, "0:0:1", timestring.SumOfTimes("0:1"))
	//assert.Equal(t, "0:1:0", timestring.SumOfTimes("1:0"))
	//assert.Equal(t, "1:0:0", timestring.SumOfTimes("60:0"))
	//assert.Equal(t, "1:1:1", timestring.SumOfTimes("0:1 0:60 0:3600"))
	//assert.Equal(t, "2:32:41", timestring.SumOfTimes("12:32 34:01 15:23 9:27 55:22 25:56"))
}
