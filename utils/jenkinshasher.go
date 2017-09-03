package utils

type JenkinsHasher struct {
}

func NewJenkinsHasher() *JenkinsHasher {
	return &JenkinsHasher{}
}

func (hasher *JenkinsHasher) Hash64(data string) uint64 {
	return computeHash(&data)
}

func computeHash(data *string) uint64 {
	bytes := []byte(*data)

	var a, b, c uint64
	a, b = 0x9e3779b9, 0x9e3779b9
	c = 0
	i := 0

	for i = 0; i < len(bytes)-12; {
		a += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)
		i += 4
		b += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)
		i += 4
		c += uint64(bytes[i]) | uint64(bytes[i+1]<<8) | uint64(bytes[i+2]<<16) | uint64(bytes[i+3]<<24)

		a, b, c = mix(a, b, c)
	}

	c += uint64(len(bytes))

	if i < len(bytes) {
		a += uint64(bytes[i])
		i++
	}
	if i < len(bytes) {
		a += uint64(bytes[i]) << 8
		i++
	}
	if i < len(bytes) {
		a += uint64(bytes[i]) << 16
		i++
	}
	if i < len(bytes) {
		a += uint64(bytes[i]) << 24
		i++
	}

	if i < len(bytes) {
		b += uint64(bytes[i])
		i++
	}
	if i < len(bytes) {
		b += uint64(bytes[i]) << 8
		i++
	}
	if i < len(bytes) {
		b += uint64(bytes[i]) << 16
		i++
	}
	if i < len(bytes) {
		b += uint64(bytes[i]) << 24
		i++
	}

	if i < len(bytes) {
		c += uint64(bytes[i]) << 8
		i++
	}
	if i < len(bytes) {
		c += uint64(bytes[i]) << 16
		i++
	}
	if i < len(bytes) {
		c += uint64(bytes[i]) << 24
		i++
	}

	a, b, c = mix(a, b, c)
	return c
}

func mix(a, b, c uint64) (uint64, uint64, uint64) {
	a -= b
	a -= c
	a ^= c >> 13
	b -= c
	b -= a
	b ^= a << 8
	c -= a
	c -= b
	c ^= b >> 13
	a -= b
	a -= c
	a ^= c >> 12
	b -= c
	b -= a
	b ^= a << 16
	c -= a
	c -= b
	c ^= b >> 5
	a -= b
	a -= c
	a ^= c >> 3
	b -= c
	b -= a
	b ^= a << 10
	c -= a
	c -= b
	c ^= b >> 15
	return a, b, c
}
