package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		myMsg := messages[:]
		return myMsg, nil
	}
	if plan == planFree {
		fMsg := messages[0:2]
		return fMsg, nil
	}
	return nil, errors.New("unsupported plan")
}
