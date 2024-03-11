package goption

import (
	"encoding/json"
	"testing"
	"time"
)

type testStruct struct {
	Name string
	Date time.Time
	Ptr  *int
}

func TestSome(t *testing.T) {
	t.Run("should create Some", func(t *testing.T) {
		s := Some(testStruct{
			Name: "test",
			Date: time.Now(),
		})
		if !s.defined {
			t.Errorf("expected Some to be defined")
		}
	})
}

func TestNone(t *testing.T) {
	t.Run("should create None", func(t *testing.T) {
		s := None[testStruct]()
		if s.defined {
			t.Errorf("expected None to be not defined")
		}
	})
}

func TestIsDefined(t *testing.T) {
	t.Run("should not be defined", func(t *testing.T) {
		var s GOption[testStruct]
		if s.IsDefined() {
			t.Errorf("expected Some to be not defined")
		}
	})

	t.Run("should be defined", func(t *testing.T) {
		s := Some(testStruct{
			Name: "test",
			Date: time.Now(),
		})
		if !s.IsDefined() {
			t.Errorf("expected Some to be defined")
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("should be empty", func(t *testing.T) {
		var s GOption[testStruct]
		if !s.IsEmpty() {
			t.Errorf("expected Some to be empty")
		}
	})

	t.Run("should not be empty", func(t *testing.T) {
		s := Some(testStruct{
			Name: "test",
			Date: time.Now(),
		})
		if s.IsEmpty() {
			t.Errorf("expected Some to be not empty")
		}
	})
}

func TestGetOrElse(t *testing.T) {
	t.Run("should return default value", func(t *testing.T) {
		var s GOption[testStruct]
		defaultValue := testStruct{
			Name: "default",
			Date: time.Now(),
		}
		if s.GetOrElse(defaultValue) != defaultValue {
			t.Errorf("expected Some to return default value")
		}
	})

	t.Run("should return value", func(t *testing.T) {
		value := testStruct{
			Name: "test",
			Date: time.Now(),
		}
		s := Some(value)
		defaultValue := testStruct{
			Name: "default",
			Date: time.Now(),
		}
		if s.GetOrElse(defaultValue) != value {
			t.Errorf("expected Some to return value")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("should return value and true", func(t *testing.T) {
		value := testStruct{
			Name: "test",
			Date: time.Now(),
		}
		s := Some(value)
		v, ok := s.Get()
		if !ok || v != value {
			t.Errorf("expected Some to return value and true")
		}
	})

	t.Run("should return zero value and false", func(t *testing.T) {
		var s GOption[testStruct]
		v, ok := s.Get()
		if ok || v != (testStruct{}) {
			t.Errorf("expected Some to return zero value and false")
		}
	})
}

func TestMarshalJSON(t *testing.T) {
	t.Run("should marshal to JSON", func(t *testing.T) {
		value := testStruct{
			Name: "test",
			Date: time.Now(),
		}
		s := Some(value)
		res, err := json.Marshal(s)
		if err != nil {
			t.Errorf("expected Some to marshal to JSON without error")
		}

		expected := `{"Name":"test","Date":"` + value.Date.Format(time.RFC3339Nano) + `","Ptr":null}`
		if string(res) != expected {
			t.Errorf("expected Some to marshal to JSON %s, %s", string(res), expected)
		}
	})

	t.Run("should marshal to null", func(t *testing.T) {
		var s GOption[testStruct]
		res, err := json.Marshal(s)
		if err != nil {
			t.Errorf("expected Some to marshal to JSON without error")
		}
		if string(res) != "null" {
			t.Errorf("expected Some to marshal to null")
		}
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("should unmarshal from JSON", func(t *testing.T) {
		value := testStruct{
			Name: "test",
			Date: time.Now(),
		}
		expected := Some(value)
		data, err := json.Marshal(expected)
		if err != nil {
			t.Errorf("expected Some to marshal to JSON without error")
		}

		var s GOption[testStruct]
		err = json.Unmarshal(data, &s)
		if err != nil {
			t.Errorf("expected Some to unmarshal from JSON without error")
		}
		if s.defined != expected.defined {
			t.Errorf("Defined different after unmarshal")
		}
		if s.value.Name != expected.value.Name {
			t.Errorf("Name different after unmarshal")
		}
		if s.value.Date.Compare(expected.value.Date) != 0 {
			t.Errorf("Date different after unmarshal")
		}
	})

	t.Run("should unmarshal from null", func(t *testing.T) {
		data := []byte("null")
		var s GOption[testStruct]
		err := json.Unmarshal(data, &s)
		if err != nil {
			t.Errorf("expected Some to unmarshal from JSON without error")
		}
	})
}
