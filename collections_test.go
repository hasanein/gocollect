package gocollect

import (
	"strings"
	"testing"
)

var planets = []interface{}{
	"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
}

var allCapsPlanets = []interface{}{
	"MERCURY", "VENUS", "EARTH", "MARS", "JUPITER", "SATURN", "URANUS", "NEPTUNE", "PLUTO",
}

// This tests against the zero value of an array
func TestCollectionCreationFromNilProducesEmptyArray(t *testing.T)  {
	if len(CollectionFrom(nil).Unwrap()) != 0{
		t.Errorf("Expecting an empty array created from \"nil\"")
	}
}

// this tests against an array of 0 elements, rather than the zero value of an array
func TestCollectionCreationFromEmptyArrayProducesEmptyArray(t *testing.T)  {
	var emptyArray = [...]interface{}{}

	if len(CollectionFrom(emptyArray[0:0]).Unwrap()) != 0{
		t.Errorf("Expecting an empty array created from \"nil\"")
	}
}

func TestFilteringProducesTheExpectedOutcome(t *testing.T) {
	// Given
	expectedArrayLength := 1
	expectedFilteredEntry := "Mars"

	isMarsPredicate := func(value interface{}) bool {
		return value.(string) == "Mars"
	}

	// When
	underlyingArray := CollectionFrom(planets).Filter(isMarsPredicate).Unwrap()

	// Then
	arrayLength := len(underlyingArray)
	if arrayLength != expectedArrayLength{
		t.Errorf("Expecting a single element in the array but found %d", arrayLength)
	}

	actualArrayElement := underlyingArray[0]
	if actualArrayElement != expectedFilteredEntry{
		t.Errorf("Expected: %s, Actual: %s", expectedFilteredEntry, actualArrayElement)
	}
}

func TestMappingProducesTheExpectedOutcome(t *testing.T)  {
	// Given
	expectedArrayLength := len(planets)
	toAllCaps := func(v interface{}) interface{} {return strings.ToUpper(v.(string))}

	// When
	underlyingArray := CollectionFrom(planets).Map(toAllCaps).Unwrap()

	// Then
	if len(underlyingArray) != expectedArrayLength{
		t.Fatalf("The length of the mapped array does not match the actual, expected %d, actual: %d", expectedArrayLength, len(underlyingArray))
	}

	// TODO Is there a better way to perform array equality in Go?
	// my initial attempt to use "==" operator on []interface{} resulted in a compilation error
	// Go compiler reported that "==" is not implemented on []interface{}
	// hence I am doing this in a rather manual fashion
	for i := 0; i < len(underlyingArray); i++{
		if underlyingArray[i] != allCapsPlanets[i]{
			t.Errorf("Expected %s, Actual %s", underlyingArray[i], allCapsPlanets[i])
		}
	}
}

func TestPrinting(t *testing.T)  {
	CollectionFrom(planets).Print()
}

func TestThatForEachPerformsTheActionOnEachElement(t *testing.T)  {
	// Given
	var passedElements []interface{}
	consumer := func(v interface{}) {
		passedElements = append(passedElements, v)
	}

	// When
	CollectionFrom(planets).ForEach(consumer)

	// Then
	if len(passedElements) != len(planets){
		t.Fatalf("Not all array elements are passed to the consumer, expecting %d, actual %d", len(planets), len(passedElements))
	}

	for i := 0; i < len(planets); i++{
		if passedElements[i] != passedElements[i]{
			t.Errorf("Expected %s, Actual %s", planets[i], passedElements[i])
		}
	}
}

func TestThatPeekPerformsTheActionOnEachElement(t *testing.T)  {
	// Given
	var passedElements []interface{}
	consumer := func(v interface{}) {
		passedElements = append(passedElements, v.(string) + "b")
	}

	// When
	underlyingArray := CollectionFrom(planets).Peek(consumer).Unwrap()

	// Then

	// test that all elements have been peeked on
	if len(passedElements) != len(planets){
		t.Fatalf("Not all array elements are passed to the consumer, expecting %d, actual %d", len(planets), len(passedElements))
	}

	for i := 0; i < len(planets); i++{
		if passedElements[i] != passedElements[i]{
			t.Errorf("Expected %s, Actual %s", planets[i], passedElements[i])
		}
	}

	// Test that the original elements of the collection are not altered during peeking
	for i := 0; i < len(planets); i++{
		if underlyingArray[i] != planets[i]{
			t.Errorf("Expected %s, Actual %s", underlyingArray[i], planets[i])
		}
	}
}

func TestThatCountWorksAsExpected(t *testing.T)  {
	// When
	count := CollectionFrom(planets).Count()

	// Then
	if count != len(planets) {
		t.Fatalf("Incorrect count, expecting %d, actual %d", len(planets), count)
	}
}