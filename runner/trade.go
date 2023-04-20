package runner

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// Trade can be extended by different type of DEX trades
type Trade interface {
	ExecuteBackrun(ctx context.Context, tx common.Hash)
}
