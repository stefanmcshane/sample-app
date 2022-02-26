package main

import (
	"fmt"
	"net/http"
	"time"
)

func IntegrationTestsForCheckingRollout(address string) error {
	fail := 0
	u := fmt.Sprintf("http://%s", address)
	fmt.Println(u)
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		url := fmt.Sprintf("http://%s/%d", address, i)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err, url)
			fail = fail + 1
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode, url)
			fail = fail + 1
		}
	}

	if fail == 5 {
		return fmt.Errorf("failure with intg tests %d", fail)
	}
	fmt.Println("success with intg", fail)
	return nil
}
