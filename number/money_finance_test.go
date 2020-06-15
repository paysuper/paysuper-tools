package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoneyRound_Ok(t *testing.T) {
	money := New()
	numbers := []struct {
		In  float64
		Out float64
		Mod int64
	}{
		{In: 0.123456789, Out: 0.12, Mod: 3456789},
		{In: 1.234567890, Out: 1.23, Mod: 8024679},
		{In: 2.345678901, Out: 2.35, Mod: 3703580},
		{In: 3.456789012, Out: 3.46, Mod: 492592},
		{In: 4.567890123, Out: 4.56, Mod: 8382715},
		{In: 5.678901234, Out: 5.68, Mod: 7283949},
		{In: 6.789012345, Out: 6.79, Mod: 6296294},
		{In: 17.890123456, Out: 17.89, Mod: 6419750},
		{In: 128.901234567, Out: 128.90, Mod: 7654317},
		{In: 1239.012345678, Out: 1239.01, Mod: 9999995},
		{In: 12390, Out: 12390, Mod: 9999995},
	}

	for _, val := range numbers {
		result, err := money.Round(val.In, int64(2))
		assert.NoError(t, err)
		assert.Equal(t, val.Out, result)
		assert.Equal(t, val.Mod, money.accumulatedMod)
	}
}

func TestMoneyRound_PrecisionError(t *testing.T) {
	money := New()
	result, err := money.Round(1239.012345678, int64(-1))
	assert.Error(t, err)
	assert.Equal(t, ErrorPrecisionToSmall, err)
	assert.EqualValues(t, 0, result)
}

func TestMoneyRound_CeilError(t *testing.T) {
	money := New()
	money.floatStringMask = "aaaaa%f"
	result, err := money.Round(1239.012345678, int64(1))
	assert.Error(t, err)
	assert.Equal(t, ErrorFloatConversion, err)
	assert.EqualValues(t, 0, result)
}

func TestMoneyCeil_Ok(t *testing.T) {
	money := New()
	masks := []string{"%.0f", "%.0f."}

	for _, val := range masks {
		money.floatStringMask = val
		val, err := money.ceil(123, 2)
		assert.NoError(t, err)
		assert.EqualValues(t, 123, val)
	}
}
