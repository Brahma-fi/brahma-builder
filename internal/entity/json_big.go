package entity

import (
	"fmt"
	"math/big"
	"strings"
)

type JsonBigInt struct {
	big.Int
}

func (b JsonBigInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *JsonBigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(strings.ReplaceAll(string(p), `"`, ""), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}
