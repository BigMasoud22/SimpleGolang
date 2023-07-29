package main

import (
	"os"
	"testing"
)

func TestReadAndWriteFilesData(t *testing.T) {
	//in order to test the file operation functions
	//we must at first check that there is no file with name of Test.txt
	if _, err := os.Stat("Test.txt"); err == nil {
		os.Remove("Test.txt")
	}
	mydeck := newDeck()
	filename, err := mydeck.writeOnFile("Test.txt")
	if err != nil {
		t.Errorf("an error occurred with fabricating the file%v", err.Error())
	}
	if filename == "" {
		t.Error("The string data is empty")
	}
	//if we got here it means that writing operation is working correctly
	//so we start to test reading operation
	returndDeck, err := readFromFile("Test.txt")
	if err != nil {
		t.Errorf("%v", err.Error())
	}
	if len(returndDeck) == 0 {
		t.Error("There is an issue with amount of returned deck")
	}

	os.Remove("Test.txt")
}
