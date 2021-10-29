package rune_by_index

import "fmt"

func RuneByIndex(s *string, i *int) (r rune, err error) {
	runes := []rune(*s)
	err = nil

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recove: %v", r)
		}
	}()
	return runes[*i], err
}
