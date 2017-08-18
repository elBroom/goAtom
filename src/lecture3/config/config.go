package config

import (
	"../workers"
	"time"
)

const RequestWaitInQueueTimeout = time.Millisecond * 100
var Wp workers.IPool = workers.NewPool(20)
