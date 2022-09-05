// Разработать программу, которая перемножает,
// делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package task

import "math/big"

// MulDivSumSubBigInt пример работы с big.Int
func MulDivSumSubBigInt() {
	var a, b, r *big.Int
	a.Exp(big.NewInt(2), big.NewInt(64), nil) // 2^64 (1 << 64)
	b.Exp(big.NewInt(2), big.NewInt(32), nil) // 2^32 (1 << 32)
	r.Mul(a, b)                               // умножаем a, b с записью в r
	r.Div(a, b)                               // делим a, b
	r.Add(a, b)                               // складываем a, b
	r.Sub(a, b)                               // вычитаем a, b
}
