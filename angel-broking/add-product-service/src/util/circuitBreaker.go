package util

import (
	"hystrix-go/hystrix"
)

// ConfigCB configure circuit breaker
func ConfigCB() {

	hystrix.ConfigureCommand("availablequantityservice", hystrix.CommandConfig{
		// timeout for function execution, in milliseconds
		Timeout: 5000,

		// max number of concurrent function execution
		MaxConcurrentRequests: 10000,

		// request volume threshold
		RequestVolumeThreshold: 50,

		// time when circuit is in open state before trying to execute function, in milliseconds
		SleepWindow: 60,

		// percent of failed function execution to open circuit
		ErrorPercentThreshold: 50,
	})
}
