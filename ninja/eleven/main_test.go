package main

import "testing"

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

func TestExerciseOne(t *testing.T) {
	stuff, err := exerciseOne()
	if err != nil {
		t.Log("Did not expect error to be returned:", err)
		t.Fail()
	} else {
		if stuff != `{"First":"James","Last":"Bond","Sayings":["Shaken, not stirred","Any last wishes?","Never say never"]}` {
			t.Log("Did not get expected string:", stuff)
			t.Fail()
		}
	}
}
