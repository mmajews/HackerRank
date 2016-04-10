package main

import "testing"

var pangram = "We promptly judged antique ivory buckles for the next prize"
var notPangram = "We promptly judged antique ivory buckles for the prize"

func shouldReturnThatIsPangram(t *testing.T){
	toBeTested := detectIfPangramOrNot(pangram)
	if toBeTested != "pangram" {
		t.Fatalf("Expected: %s Returned: %s","pangram",toBeTested)
	}
}

func shouldReturnThatIsNotPangram(t *testing.T){
	toBeTested := detectIfPangramOrNot(notPangram)
	if toBeTested != "not pangram" {
		t.Fatalf("Expected: %s Returned: %s","not pangram",toBeTested)
	}
}