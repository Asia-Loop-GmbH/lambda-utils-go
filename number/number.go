package number

import (
	"fmt"
	"math/big"
)

func NewBigNumber(number string) (*big.Float, error) {
	value := &big.Float{}
	value, ok := value.SetString(number)
	if !ok {
		return nil, fmt.Errorf("could not parse '%s' to float", number)
	}
	return value, nil
}
