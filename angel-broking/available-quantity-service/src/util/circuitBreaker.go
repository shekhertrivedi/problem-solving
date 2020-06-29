package util

import (
	"hystrix-go/hystrix"
)

// ConfigCB configure circuit breaker
func ConfigCB() {

	hystrix.ConfigureCommand("addproductservice", hystrix.CommandConfig{
		// timeout for function execution, in milliseconds
		Timeout: 100,

		// max number of concurrent function execution
		MaxConcurrentRequests: 10000,

		// request volume threshold
		RequestVolumeThreshold: 50,

		// time when circuit is in open state before trying to execute function, in milliseconds
		SleepWindow: 60,

		// percent of failed function execution to open circuit
		ErrorPercentThreshold: 50,
	})

	hystrix.ConfigureCommand("allocateproductservice", hystrix.CommandConfig{
		// timeout for function execution, in milliseconds
		Timeout: 100,

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

/*

sample usage
*************

circuitBreaker := func() error {
		bo, err = c.Client.GetByPublicID(boID, publicID)
		if err != nil {
			metricError = err
			return filterCherwellConnectivityErrors(err)
		}

		if bo != nil {
			metricError = bo.ErrorData
		}
		return nil
	}

	hErr := hystrix.Do(hystrixCherwellServiceName, circuitBreaker, nil)
	if hErr != nil {
		err = hErr
	}

*/
