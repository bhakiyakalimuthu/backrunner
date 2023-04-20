package types

import (
	"math/big"
)

type UniswapTrade struct {
	Deadline          *big.Int
	TokenIn           string
	TokenOut          string
	Fee               *big.Int
	Recipient         string
	AmountIn          *big.Int
	AmountOutMin      *big.Int
	AmountOut         *big.Int
	AmountInMax       *big.Int
	SqrtPriceLimitX96 *big.Int
}
