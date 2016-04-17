package main

import "testing"

func shouldTrimStringCorrectly(t *testing.T) {
	testCase0 := "    Taki tam string\n"
	expected0 := "Taki tam string"

	testCase1 := " string inny       "
	expected1 := "string inny"

	expected := [2]string{expected0, expected1}
	toBeTested := [2]string{testCase0, testCase1}

	for index, element := range toBeTested {
		trimmed := trimString(element)
		assertEquals(trimmed, expected[index], t)
	}
}

func shouldCorrectlyDetectThatTwoStringsHaveSomeStringInCommon(t *testing.T) {
	testCase0 := [2]string{"hellow", "world"}
	testCase1 := [2]string{"hi", "world"}

	toBeTested := [2][2]string{testCase0, testCase1}
	expected := [2]string{"YES", "NO"}

	for index, element := range toBeTested {
		result := detectIfCommonSubStringExists(element[0], element[1])
		assertEquals(result, expected[index], t)
	}
}

func assertEquals(actual string, expected string, t *testing.T) {
	if (actual != expected) {
		t.Fatalf("Expected: %s Returned: %s", expected, actual)
	}
}

