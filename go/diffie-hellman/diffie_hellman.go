package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	private := big.NewInt(0)
	private.Rand(rnd, private.Add(p, big.NewInt(-2)))
	return private.Add(private, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	public := big.NewInt(g)
	return public.Exp(public, private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := big.NewInt(0)
	return secret.Exp(public2, private1, p)
}
