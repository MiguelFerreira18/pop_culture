package ctx

import "context"

const KeyRequestId = "requestID"

func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(KeyRequestId).(string)

	return requestID
}
