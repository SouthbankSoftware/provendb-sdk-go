package anchor

import (
	"testing"
)

func TestDecodeCHP_PATH(t *testing.T) {
	data := "eJykkbGS0zAQQH+G1rG0K9lWqszwC1Q0nt3VGmsIlsfWBa4EGtp8w4XhYCiBkv/I3zC5wHU00D7pvZ3Z/XC/kzwVfVN+jqXM67auX2OKm7y8qGWkNM05TaU+4Knczvrl6SM6jbSO5x3IAC4y+YENoVFotEMGEs9BIahl21j1YgdqHXht0TjEBgJ2HDpwXy+ZPsV+ylHPTyIZQPFSSRd8Za2GihCGCghIIbKIuh8PynrDr1IpejV7Kt/AgK0MVuCfAW6d3Vp8/piXvPxj/mL+LX/PC00y6np893FPrPvvc+Rey9hfcF766/tdntfj209Xth7fP+zyTmj/+fe3FM+7VnCAltk2rlWN1nhpPEb0A6m6BkNk10Egb7yG0EqnTL7V0HaN+GCa082S1uP55Z8zXtObqIfNvOQ8EO91k3J90CUNt7WWsf7fkb8CAAD//5Arw98="
	p, err := DecodeProof(data)
	if err != nil {
		t.FailNow()
	}
	_ = p.(map[string]interface{})
}

func TestDecodeETH_TRIE(t *testing.T) {
	data := "eJxqW5WYl5yRXxRSWZC62DXEY2lJRZ5nyk0HA7OUNMskcwOTZEvjtLREQ6MkA4u0ZFODFGNzM1Mzs8QkE8NkgyRzI0PjZDMLc3ODVAvjRKMky5TkRJPURHOTtGUlFXmhRZk3kzJKSgqKrfT1izLzslOTKvVSSzJSi4qTE/P0MvP1Syr0DSootWplUk5+cnZIZm7qeQYGBoaEWEXJ1WAhv9LcpNQisCBD3cLolSVFmal++SmpxROPaL5YqKDzyXdX/IbFfGovrLWO7nw0T1pMLkZ+uVNcMYtx0uTuGSqtvt7OTt6AAAAA///4m2Mk"
	p, err := DecodeProof(data)
	if err != nil {
		t.FailNow()
	}

	_ = p.(map[string]interface{})
}

func TestDecodeCHP_PATH_SIGNED(t *testing.T) {
	data := "eJykk81u42QUhsW9sItS299f7EojNQlJbadN0tSZZiKk6Ps5jp36b/w5cexdYcO24gpgimZALIEl91GuBnUCs2PDbN+j5zmL95zvPlzIPKvgWP0ZVVWhzw2jxrE6y8utISMeZ0UeZ5VxwE9VU8Avw0/RU8R19HyBZIiIEpyGwuTYBMTAxgJxSYUDyAFLWMwCKq2Q9wii0MMmwZghB9vCsRH59UWzidUmyxU8f6m4ibCksitth3YtC5wuxyjsIo44ICWkBPLHR0TvRRpXFZzIDa9+QyayuibuIhogfE6tc+ysP+llXv5P/Qv5X/oPouSZjEA/fvNTwgUkvxdKbKCKNi9xXm5O83d5ob9/+KH864uvdbw9X/kjHGz7xjoOgjv3zTbo93eLIYui3d2g09BVGvaYTGa7vY/T7cHvpO5xWI5XSZiWC1/O7th2r+KJumuKW/92dDRZHXlXU1Td3tdXFVt49k2THjqGo/3w8v4rgia05lmg+6OlnBNbFfF+/bbvCjdctrvFZVJgx7ePK6VeW7Ry0d5r0CCrr83aa66u7NYdIrM93t7PsY0c10+HMhgVLVgMrrPgcsCaTm6k48FgeOP0SxqtlokdGrCnenFIp3FPYqjzdBr0k2vYLjrRdHIcJ7YPi3w+iiZtvhedYzvy9PoNuG7ribH3VuPVzJwJFs7DpdDptJnK4evlbDdOm3o6h/lCUS+ptD+HzNuhm0vTayfr/qtXDz/mxXsd8S6i7OH9qQH9+O3Hy30nefLzP6XE6vkiBIlB2hAiEFgqIQEJAjaTXHLGCev1ECGOgzChWDk9RnrEIcABGMHCtKynfRnrx+f7f5/mpD5TcDgryjwPuUjgLM6NA5Rx2BhQRcbnrvw7AAD//ywtQbc="
	p, err := DecodeProof(data)
	if err != nil {
		t.FailNow()
	}
	_ = p.(map[string]interface{})
}

func TestDecodeETH_TRIE_SIGNED(t *testing.T) {
	data := "eJxqW5WYl5yRXxRSWZC62DXEY2lJRZ5nyk0HA1Mz09Qko5QUg2SLZPNUC1MLM5NUS2PzJEtLS/MksxQDC4M0M8s0Q4PkNHPDxBSTJGNDQ+M0y0RzA/NksxTzxGUlFXmhRZk3kzJKSgqKrfT1izLzslOTKvVSSzJSi4qTE/P0MvP1Syr0DSootWplUk5+cnZIZm7qeQYGBoaEWOXg1WAhv9LcpNQisCBD3cLGlSVFmal++SmpxVOOBP8IbICABTV/Y1rfR3uelcie8eibm5lEteLq2U0de5eXMTbtFjQ6vwasiHnz3IdfmRg/r19Zn3/roWr8fbWVIptPXfvjY8iwb+bfB/kgRUc0ni8w6f/IHd3jd1Z373t1qf6jLBYFjze8ZVrq//Pca9XT7okfWn29nZ28jzJy/2TkaDHOTLfayciwdMsWk9JjPHbbb5va260/0Wx6OutlWR/PdLfvMuX3V1akW63K3fnC/onVbM9bJasUU42kfr5nFN3gX800/daVkN97b1qe0tsiWLGd165u8Xln9nPfT8ovSxFqf5wrIK07uejmk5WXDxf39a53EKr5+/1Un2V/hFB1nvl8RbPAtTt8n8guK7HZv6R3969nE/RX2XgmTboa9fRqhs5G66anCnwp81YXCC6K2X3Zbm71m20+/5ao8lmb/pbVOjRHTGlKwMzKu9HT7wpnmzCzsBh/33KoYm1LqUmr/Y7036yZm9Qj5ewmXks5+uqQaX3DtkSGRzcyLkzPfPNNUJJf3dw9dfP9rIy+RTu2BGqf3250RPPFQgWdT7674jcs5lN7Ya11dOejedJicjHyy53iilmMkyZ3z1CBBBwgAAD//0CEJ0A="
	p, err := DecodeProof(data)
	if err != nil {
		t.FailNow()
	}
	_ = p.(map[string]interface{})
}
