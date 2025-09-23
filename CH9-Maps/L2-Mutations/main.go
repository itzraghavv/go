package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, exists := users[name]
	deleted = true
	if !exists {
		return deleted, errors.New("not found")
	}
	if !user.scheduledForDeletion {
		deleted = false
		return deleted, nil
	}
	delete(users, user.name)
	return deleted, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
