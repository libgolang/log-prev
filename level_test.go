package logger

import (
	"encoding/json"
	"testing"
)

type tSruct struct {
	Level Level
}

func TestMarshalJSON(t *testing.T) {
	// given
	level := DEBUG
	expected := `"DEBUG"`

	// when
	b, _ := json.Marshal(&level)
	serializedLevel := string(b)

	//then
	if serializedLevel != expected {
		t.Errorf("expected `%s`, but got `%s`", expected, serializedLevel)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	// given
	var level Level
	expected := DEBUG

	// when
	_ = json.Unmarshal([]byte(`"DEBUG"`), &level)

	//then
	if level != expected {
		t.Errorf("expected `%s`, but got `%s`", expected, level)
	}
}

func TestLevelMarshalInStruct(t *testing.T) {
	// given
	tval := tSruct{WARN}
	expected := `{"Level":"WARN"}`

	// when
	b, _ := json.Marshal(&tval)
	result := string(b)

	//then
	if result != expected {
		t.Errorf("expected `%s`, but got `%s`", expected, result)
	}
}
