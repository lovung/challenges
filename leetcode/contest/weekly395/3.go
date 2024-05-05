package weekly395

func minEnd(n int, x int) int64 {
	nbit := int64(n) - 1
	res := int64(x)
	xmask := int64(0x01)
	xmaskshifted := 0
	for nbit > 0 {
		for int64(x)&xmask > 0 {
			xmask <<= 1
			xmaskshifted++
		}
		res |= (nbit & 1) << xmaskshifted
		xmask <<= 1
		xmaskshifted++
		nbit >>= 1
	}
	return res
}
