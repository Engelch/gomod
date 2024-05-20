package debugerrorce

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type structA struct {
	A string `json:"A"`
	B string `json:"B"`
}

type structB struct {
	structA
	C string `json:"C"`
}

var a structA
var b structB

func TestMain(m *testing.M) {
	a.A = "A string"
	a.B = "B string"
	b.A = "A2 string"
	b.B = "B2 string"
	b.C = "C string"
	os.Exit(m.Run())
}

func TestFormatString2String(t *testing.T) {
	demo := "blabla"
	res, err := Format(FormatText, demo)
	assert.Nil(t, err, "Format returned err")
	assert.Equal(t, "blabla", res)
	// fmt.Println("text is:", res)
}

func TestFormatString2JSON(t *testing.T) {
	demo := "blabla"
	res, err := Format(FormatJSON, demo)
	assert.Nil(t, err, "Format returned err")
	assert.Equal(t, "{\"blabla\"}", res)
	// fmt.Println("json is:", res)
}

func TestFormatString2PrettyJSON(t *testing.T) {
	demo := "blabla"
	res, err := Format(FormatPrettyJson, demo)
	assert.Nil(t, err, "Format returned err")
	assert.Equal(t, "{ \"blabla\" }", res)
	// fmt.Println("json is:", res)
}

func TestFormatStruct2Text(t *testing.T) {
	dest := FormatText
	_, err := Format(dest, a)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("a struct is", res)
}

func TestFormatStruct2JSON(t *testing.T) {
	dest := FormatJSON
	_, err := Format(dest, a)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("a struct is", res)
}

func TestFormatStruct2PrettyJSON(t *testing.T) {
	dest := FormatPrettyJson
	_, err := Format(dest, a)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("a struct is", res)
}

// nested struct ===============

func TestFormatStructB2Text(t *testing.T) {
	dest := FormatText
	_, err := Format(dest, b)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("b struct is", res)
}

func TestFormatStructB2JSON(t *testing.T) {
	dest := FormatJSON
	_, err := Format(dest, b)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("b struct is", res)
}

func TestFormatStructB2PrettyJSON(t *testing.T) {
	dest := FormatPrettyJson
	_, err := Format(dest, b)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("b struct is", res)
}

// maps ===============

func TestFormatMap2String(t *testing.T) {
	c := map[string]int{"key1": 1, "key2": 42}
	dest := FormatText
	_, err := Format(dest, c)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("c map is", res)
}

func TestFormatMap2PrettyJSON(t *testing.T) {
	c := map[string]int{"key1": 1, "key2": 42}
	dest := FormatPrettyJson
	_, err := Format(dest, c)
	assert.Nil(t, err, "Format returned err")
	// fmt.Println("c map is", res)
}

// EOF
