package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Word not found...")
var errWordExists = errors.New("That word already exists...")
var errUnavailableUpdate = errors.New("Cannot update non-existing word...")
var errDeleteFail = errors.New("Cannot delete non-existing word...")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

	// if err == errNotFound{
	// 	d[word] = def
	// }else if err == nil{
	// 	return errWordExists
	// }
	// return nil
}

// Update the definition of a word in the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def

	case errNotFound:
		return errUnavailableUpdate

	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errDeleteFail
	}
	return nil
}
