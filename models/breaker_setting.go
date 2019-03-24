package models

// BreakerSetting ... ref to https://github.com/sony/gobreaker,
// casue BreakerSetting is used to gobreaker.CircutBreaker
// Name is the name of the CircuitBreaker.
// MaxRequests is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open. If MaxRequests is 0, CircuitBreaker allows only 1 request.
// Interval is the cyclic period of the closed state for CircuitBreaker to clear the internal Counts, described later in this section. If Interval is 0, CircuitBreaker doesn't clear the internal Counts during the closed state.
// Timeout is the period of the open state, after which the state of CircuitBreaker becomes half-open. If Timeout is 0, the timeout value of CircuitBreaker is set to 60 seconds.
type BreakerSetting struct {
	Name             string  `json:"name"`
	MaxRequests      uint32  `json:"max_requests"`
	ClearInterval    uint32  `json:"interval"`           // time.Millisecond
	Timeout          uint32  `json:"timeout"`            // time.Millisecond
	TripRequestCnt   uint32  `json:"trip_request_cnt"`   // setting for ready
	TripFailureRatio float64 `json:"trip_failure_ratio"` //

	// ReadyToTrip   func(counts Counts) bool // should be used with default func
	// OnStateChange func(name string, from State, to State) // should be used with default func
}

// const (
// 	// CntRequests ...
// 	CntRequests = 10
// 	// FailureRatio ...
// 	FailureRatio = 0.6
// )

// func readyToTrip(counts gobreaker.Counts) bool {
// 	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
// 	return counts.Requests >= CntRequests && failureRatio >= FailureRatio
// }
