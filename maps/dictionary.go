package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// notice the ok variable - a boolean that tells you if the key existed in the map.
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
