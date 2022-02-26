package main

import (
	"errors"
	"net/http"
	"time"
)

func IntegrationTestsForCheckingRollout() error {
	fail := 0

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)

		resp, err := http.Get("http://sample-app-service:8080/stefan")
		if err != nil {
			fail = fail + 1
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fail = fail + 1
		}
	}

	if fail != 0 {
		return errors.New("failure with intg tests")
	}
	return nil
}
