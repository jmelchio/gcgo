package main

import "testing"

func TestExerciseOne(t *testing.T) {
	t.Log("testing is fun")
}

func TestExerciseThree(t *testing.T) {
	someError := customErr{
		Stuff:         "stuff",
		Morestuff:     "morestuff",
		Evenmorestuff: "evenmorestuff",
	}
	err := exerciseThree(someError)
	if err != nil {
		t.Log("Did not expect a non nil error return here:", err)
		t.Fail()
	}

	someError.Erroneousstuff = "It's just wrong"
	err = exerciseThree(someError)
	if err == nil {
		t.Log("Did not expect nil error return here")
		t.Fail()
	}
}
