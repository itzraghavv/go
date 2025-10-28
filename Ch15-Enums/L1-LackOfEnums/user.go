package main

import "fmt"

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) (string, error) {
	if status != "email_bounced" {
		return u.status, fmt.Errorf("invalid status: %s", status)
	}
	u.status = status
	return u.status, nil
}

func (a *analytics) track(event string) (int, error) {
	if event != "email_bounced" {
		return 0, fmt.Errorf("invalid event: %s", event)
	}
	a.totalBounces++
	return a.totalBounces, nil
}
