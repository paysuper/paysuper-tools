package number

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToPrecise(val float64) float64 {
	p := math.Pow(10, 6)
	return math.Round(val*p) / p
}

func GetPercentPartFromAmount(amount, rate float64) float64 {
	return amount / (1 + rate) * rate
}

func FormatAmount(amount float64) float64 {
	return ToFixed(amount, 2)
}

func ToFixed(num float64, precision int) float64 {
	parts := strings.Split(fmt.Sprintf("%f", num), ".")
	v := parts[0]
	if precision > 0 && parts[1] != "" {
		v += "." + parts[1][0:precision]
	}
	val, _ := strconv.ParseFloat(v, 64)
	return val
}
