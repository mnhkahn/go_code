package service

import (
	"bitbucket.org/tebeka/base62"
	"fmt"
)

var (
	DEFAULT_BLOCK_SIZE = uint(24)
	DEFAULT_ENCODER    *UrlEncoder
	MIN_LENGTH         = 5
)

type UrlEncoder struct {
	BlockSize uint
	Mask      uint64
	Mapping   []uint
}

func NewUrlEncoder(blocksize uint) (u *UrlEncoder) {
	u = new(UrlEncoder)
	u.BlockSize = blocksize
	u.Mask = (1 << blocksize) - 1
	for i := uint(0); i < blocksize; i++ {
		u.Mapping = append(u.Mapping, blocksize-i-1)
	}
	return u
}

func (this *UrlEncoder) Encode(n uint64) uint64 {
	return (n & ^this.Mask) | this.encode(n&this.Mask)
}

func (this *UrlEncoder) encode(n uint64) (result uint64) {
	for i, b := range this.Mapping {
		if n&(1<<uint(i)) > 0 {
			result |= (1 << b)
		}
	}
	return result

}

func (this *UrlEncoder) Decode(n uint64) uint64 {
	return (n & ^this.Mask) | this.decode(n&this.Mask)
}

func (this *UrlEncoder) decode(n uint64) (result uint64) {
	for i, b := range this.Mapping {
		if n&(1<<uint(b)) > 0 {
			result |= (1 << uint(i))
		}
	}
	return result
}

func (this *UrlEncoder) Enbase(n uint64) string {
	return base62.Encode(n)
}

func (this *UrlEncoder) Debase(n string) uint64 {
	return base62.Decode(n)
}

func (this *UrlEncoder) EncodeUrl(n uint64) (short_url string) {
	return DEFAULT_ENCODER.Enbase(DEFAULT_ENCODER.Encode(n))
}

func DecodeUrl(short_url string) (n uint64) {
	return DEFAULT_ENCODER.Decode(DEFAULT_ENCODER.Debase(short_url))
}

func init() {
	DEFAULT_ENCODER = NewUrlEncoder(DEFAULT_BLOCK_SIZE)
}

func Disperse(i uint64) (disperse_i uint64) {
	fmt.Println(DEFAULT_ENCODER)
	a := DEFAULT_ENCODER.Encode(i)
	b := DEFAULT_ENCODER.Enbase(a)
	c := DEFAULT_ENCODER.Debase(b)
	disperse_i = DEFAULT_ENCODER.Decode(c)
	fmt.Println(i, a, b, c, disperse_i)
	return disperse_i
}
