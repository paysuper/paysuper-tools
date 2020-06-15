package number

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	MultiplierDefault = 1000000000
)

var (
	ErrorFloatConversion  = errors.New("unable to convert float with rounding")
	ErrorPrecisionToSmall = errors.New("precision must be a greater then 0")
)

type Money struct {
	mul             int64
	accumulatedMod  int64
	mx              sync.Mutex
	floatStringMask string
}

func New() *Money {
	money := &Money{
		mul:             MultiplierDefault,
		floatStringMask: "%f",
	}

	return money
}

func (m *Money) Round(val float64, precision int64) (float64, error) {
	if precision < 0 {
		return 0, ErrorPrecisionToSmall
	}

	valCeil, err := m.ceil(val, precision)

	if err != nil {
		return 0, err
	}

	valInt := m.toInt(valCeil)
	precMul := m.toInt(1 / math.Pow(10, float64(precision)))

	m.mx.Lock()
	val1Int := m.toInt(val)
	dif := val1Int - valInt
	m.accumulatedMod += dif

	for m.accumulatedMod >= precMul {
		valInt += precMul
		m.accumulatedMod -= precMul
	}
	m.mx.Unlock()

	valCeil, _ = big.NewRat(valInt, m.mul).Float64()
	return valCeil, nil
}

func (m *Money) ceil(val float64, precision int64) (float64, error) {
	valStr := fmt.Sprintf(m.floatStringMask, val)
	index := strings.Index(valStr, ".")

	if index < 0 {
		return val, nil
	}

	index += int(precision + 1)
	valLen := len(valStr)

	if index > valLen {
		index = valLen
	}

	valBig, ok := big.NewFloat(0).SetString(valStr[:index])

	if !ok {
		return 0, ErrorFloatConversion
	}

	result, _ := valBig.Float64()
	return result, nil
}

func (m *Money) toInt(val float64) int64 {
	result, accuracy := big.NewFloat(0).Mul(big.NewFloat(val), big.NewFloat(0).SetInt64(m.mul)).Int64()

	if accuracy != 0 {
		result -= int64(accuracy)
	}

	return result
}
