package runner

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// Trade is a abstraction layer which can be extended by different type of DEX trades
type Trade interface {
	ExecuteBackrun(ctx context.Context, tx common.Hash)
}
