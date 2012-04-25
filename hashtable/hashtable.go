package main

const (
	s            = 2654435769 // A*2^32 A=(sqrt5-1)/2 = 0.6180339887 
	P            = 5          // Initial default Hashtable size 2^5=32
	MAXTHRESHOLD = 0.7        // if above 0.7 utilization resize
	MINTHRESHOLD = 0.2        // if below 0.2 utilization resize
	MAX_PROBING  = 24
)

type bucket struct {
	k string
	v interface{}
}

type Hashtable struct {
	buckets []*bucket
	old     []*bucket
	items   int
	min     int
	max     int
	p       uint32
}

func hash(k, p uint32) uint32 {
	x := k * s
	return x >> (32 - p)
}

func rehash(i, m uint32) uint32 {
	return (i + 1) % m
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

func NewHashTable() *Hashtable {
	return &Hashtable{make([]*bucket, 32), nil, 0, 0, int(32 / MAXTHRESHOLD), P}
}

// Get the value associated with name k
func (h *Hashtable) Get(k string) interface{} {
	index := hash(hashcode(k), h.p)
	for i := 0; (h.buckets[index] == nil || !h.buckets[index].k == k) && i < MAX_PROBING; i++ {
		index = rehash(index, len(h.buckets))
	}
	if h.buckets[index] == nil {
		return nil
	}
	return h.v
}

// Put associates name k with value v
func (h *Hashtable) Put(k string, v interface{}) {
	index := hash(hashcode(k), h.p)
	for i := 0; h.buckets[index] != nil && i < MAX_PROBING; i++ {
		index = rehash(index, len(h.buckets))
	}
	if h.buckets[index] != nil {
		resize(h, len(h.buckets))
		Put(k, v)
		return
	}
	h.buckets[index] = &bucket{k, v}
	h.items++
	if h.items > h.max {
		resize(h, true)
	}
}

// Remove the association for name k 
func (h *Hashtable) Remove(k string) {
	index := hash(hashcode(k), h.p)
	for i := 0; (h.buckets[index] == nil || !h.buckets[index].k == k) && i < MAX_PROBING; i++ {
		index = rehash(index, len(h.buckets))
	}
	if h.buckets[index] == nil {
		return
	}
	h.buckets[index] = nil
	h.items--
	if h.items < h.min {
		resize(h, false)
	}
}

// Len returns the number of associations stored
func (h *Hashtable) Len() {
	return h.items
}

func resize(h *Hashtable, grow bool) {
	newsize := len(h.buckets)
	newp := h.p
	if grow {
		newsize += newsize
		newp++
	} else {
		newsize = newsize >> 1
		newp--
	}
	h.old=h.buckets
	h.buckets=make([]*bucket, newsize)
	h.p=newp
}
