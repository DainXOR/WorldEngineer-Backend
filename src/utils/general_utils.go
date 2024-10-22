package utils

import (
	"dainxor/we/logger"
	"time"
)

func Retry[T any](callback func() (T, error), retryAttemps int, failMsg string, finalMsg string) (T, error) {
	result, err := callback()
	failCount := 0

	for err != nil && failCount < retryAttemps {
		logger.Error(failMsg, err)
		logger.Warning("Fail count: ", failCount+1)
		failCount++
		seconds := 5 * failCount

		for seconds > 0 {
			logger.Debug("Trying again in ", seconds, "...")
			time.Sleep(1 * time.Second)
			seconds--
		}

		result, err = callback()
	}

	if err != nil {
		logger.Error(finalMsg, err)
	}

	return result, err
}

func RetryOrPanic[T any](callback func() (T, error), times int, failMsg string, finalMsg string) T {
	result, err := Retry(callback, times, failMsg, finalMsg)

	if err != nil {
		panic(err)
	}

	return result
}
