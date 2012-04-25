package main

const (
	s = 2654435769 // A*2^32 A=(sqrt5-1)/2 = 0.6180339887 
	P = 5          // Initial default Hashtable size 2^5=32
)

type HashBucket struct {
	k string
	v interface{}
}

type Hashtable struct {
	buckets []HashBucket
}

func hash(k, p uint32) uint32 {
	x := k * s
	return x >> (32 - p)
}

func hashcode(s string) uint32 {
	k := []byte(s)
	h := uint32(0)
	for i := 0; i < len(k); i += 4 {
		highorder := h & 0xf8000000 // extract high-order 5 bits from h
		// 0xf8000000 is the hexadecimal representation
		//   for the 32-bit number with the first five 
		//   bits = 1 and the other bits = 0   	
		h = h << 5                // shift h left by 5 bits
		h = h ^ (highorder >> 27) // move the highorder 5 bits to the low-order end and XOR into h
		ki := uint32(k[i])
		for j := 1; j < 4; j++ {
			shift := uint32(j << 3) // 8,16,24
			if i+j < len(k) {
				ki = ki | uint32(k[i+j])<<shift
			}
		}

		h = h ^ ki
	}
	return h
}

