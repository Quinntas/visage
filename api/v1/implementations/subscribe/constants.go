package v1Subscribe

import "time"

const (
	ACTIVE_CHANNELS_CTX_KEY = "active_channels"
	HEALTHCHECK_STRING      = "i_am_alive"
	HEALTHCHECK_TIMEOUT     = time.Second * 10
)
