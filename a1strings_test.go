package a1strings

import "testing"
import "../golib"

func TestRpad(t *testing.T) {
	actual := a1strings.Rpad("abc", 5)
	expected := "abc  "
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRpadMuliBite(t *testing.T) {
	actual := a1strings.Rpad("あいう", 5)
	expected := "あいう  "
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLpad(t *testing.T) {
	actual := a1strings.Lpad("abc", 5)
	expected := "  abc"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLpadMultiByte(t *testing.T) {
	actual := a1strings.Lpad("あいう", 5)
	expected := "  あいう"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSnake2Camel(t *testing.T) {
	actual := a1strings.Snake2Camel("spring_has_come")
	expected := "springHasCome"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSnake2CamelUcase(t *testing.T) {
	actual := a1strings.Snake2Camel("SPRING_HAS_COME")
	expected := "springHasCome"
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLenFontWidth(t *testing.T) {
	actual := a1strings.LenFontWidth("123456")
	expected := 6
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestLenFontWidthMultiByte(t *testing.T) {
	actual := a1strings.LenFontWidth("あいうabc")
	expected := 9
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
