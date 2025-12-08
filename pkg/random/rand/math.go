package rand

import (
	mrand "math/rand/v2" // used for non-crypto unique ID and random port selection
	"sync"
	"time"
)

// MathRandomGenerator is a random generator for non-crypto usage.
type MathRandomGenerator interface {
	// Intn returns random integer within [0:n).
	Intn(n int) int

	// Uint32 returns random 32-bit unsigned integer.
	Uint32() uint32

	// Uint64 returns random 64-bit unsigned integer.
	Uint64() uint64

	// Float64 returns random float64 in range [0.0, 1.0).
	Float64() float64

	// IntBetween returns random integer within [min, max].
	IntBetween(min, max int) int

	// Shuffle shuffles elements in slice using Fisher-Yates algorithm.
	Shuffle(n int, swap func(i, j int))

	// GenerateString returns random string using given set of runes.
	// It can be used for generating unique ID to avoid name collision.
	// If runes is empty, an empty string is returned.
	//
	// Caution: DO NOT use this for cryptographic usage.
	GenerateString(n int, runes string) string
}

type mathRandomGenerator struct {
	r  *mrand.Rand
	mu sync.Mutex
}

// NewMathRandomGenerator creates new mathmatical random generator.
// Random generator is seeded by crypto random.
func NewMathRandomGenerator() MathRandomGenerator {
	seed, err := CryptoUint64()
	if err != nil {
		// crypto/rand is unavailable. Fallback to seed by time.
		seed = uint64(time.Now().UnixNano())
	}

	seed1, err := CryptoUint64()
	if err != nil {
		seed1 = uint64(time.Now().UnixNano())
	}
	return &mathRandomGenerator{r: mrand.New(mrand.NewPCG(seed, seed1))}
}

func (g *mathRandomGenerator) Intn(n int) int {
	g.mu.Lock()
	v := g.r.IntN(n)
	g.mu.Unlock()
	return v
}

func (g *mathRandomGenerator) Uint32() uint32 {
	g.mu.Lock()
	v := g.r.Uint32()
	g.mu.Unlock()
	return v
}

func (g *mathRandomGenerator) Uint64() uint64 {
	g.mu.Lock()
	v := g.r.Uint64()
	g.mu.Unlock()
	return v
}

func (g *mathRandomGenerator) GenerateString(n int, runes string) string {
	letters := []rune(runes)
	// 处理空字符集或长度为0的情况
	if len(letters) == 0 || n == 0 {
		return ""
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[g.Intn(len(letters))]
	}
	return string(b)
}

// Float64 returns random float64 in range [0.0, 1.0).
func (g *mathRandomGenerator) Float64() float64 {
	g.mu.Lock()
	v := g.r.Float64()
	g.mu.Unlock()
	return v
}

// IntBetween returns random integer within [min, max].
func (g *mathRandomGenerator) IntBetween(min, max int) int {
	// 处理边界情况
	if min >= max {
		return min
	}

	g.mu.Lock()
	// rand.Intn生成[0,n)范围的数，所以需要+1来包含max
	v := min + g.r.IntN(max-min+1)
	g.mu.Unlock()
	return v
}

// Shuffle shuffles elements in slice using Fisher-Yates algorithm.
func (g *mathRandomGenerator) Shuffle(n int, swap func(i, j int)) {
	g.mu.Lock()
	g.r.Shuffle(n, swap)
	g.mu.Unlock()
}
