package utils

import (
	"context"

	"github.com/newrelic/go-agent"
)

type contextKey string

// NewRelicTransactionKey is used to retrieve NewRelic transaction from context
const NewRelicTransactionKey = contextKey("NewRelicTransactionKey")

// GetNewRelicTransaction encapsulates type-checking routines used for retrieving transaction from context
func GetNewRelicTransaction(ctx context.Context) newrelic.Transaction {
	v := ctx.Value(NewRelicTransactionKey)
	if txn, ok := v.(newrelic.Transaction); ok {
		return txn
	}
	return nil
}
