package types

import (
	"math/big"
	"time"
)

type UniswapTrade struct {
	Deadline     time.Duration `json:"deadline"`
	*TradeParams `json:"params"`
}

type TradeParams struct {
	Params struct {
		TokenIn           string   `json:"tokenIn"`
		TokenOut          string   `json:"tokenOut"`
		Fee               *big.Int `json:"fee"`
		Recipient         string   `json:"recipient"`
		AmountIn          *big.Int `json:"amountIn"`
		AmountOutMinimum  *big.Int `json:"amountOutMinimum"`
		AmountOut         *big.Int `json:"amountOut"`
		AmountInMaximum   *big.Int `json:"amountInMaximum"`
		SqrtPriceLimitX96 *big.Int `json:"sqrtPriceLimitX96"`
	} `json:"params"`
}
