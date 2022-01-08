package main

type InputError struct{ Errstr string }

func (i *InputError) Error() string {
	return i.Errstr
}
