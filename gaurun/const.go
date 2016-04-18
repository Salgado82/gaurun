package gaurun

const (
	Version = "0.4.3"
)

const (
	EpApnsProd    = "gateway.push.apple.com:2195"
	EpApnsSandbox = "gateway.sandbox.push.apple.com:2195"
)

const (
	PlatFormIos = iota + 1
	PlatFormAndroid
)

const (
	ErrorStatusUnknown = iota
	ErrorStatusNotRegistered
	ErrorStatusMismatchSenderId
	ErrorStatusCanonicalId
)

const (
	StatusAcceptedPush  = "accepted-push"
	StatusSucceededPush = "succeeded-push"
	StatusFailedPush    = "failed-push"
)
