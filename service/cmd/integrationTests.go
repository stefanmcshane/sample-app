package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func IntegrationTestsForCheckingRollout(address string) error {
	fail := 0

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		url := fmt.Sprintf("http://%s/%d", address, i)
		fmt.Println(url)
		resp, err := http.Get(url)
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
