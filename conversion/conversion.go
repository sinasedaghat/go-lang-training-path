package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, 0, len(strings))

	for _, str := range strings {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("can't convert one of string content to float64")
		}

		floats = append(floats, val)
	}
	return floats, nil
}
