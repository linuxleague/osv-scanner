package testutility

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

// LoadJSONFixture loads a JSON fixture file and returns the decoded version.
func LoadJSONFixture[V any](t *testing.T, path string) V {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open fixture: %s", err)
	}
	var value V
	err = json.NewDecoder(file).Decode(&value)
	if err != nil {
		t.Fatalf("Failed to parse fixture: %s", err)
	}

	return value
}

// AssertMatchFixtureJSON matches the JSON at path with the value val, failing if not equal, printing out the difference.
func AssertMatchFixtureJSON[V any](t *testing.T, path string, val V) {
	t.Helper()
	fileA, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to open fixture: %s", err)
	}

	var elem V
	err = json.Unmarshal(fileA, &elem)
	if err != nil {
		t.Fatalf("Failed to unmarshal val: %s", err)
	}

	if !reflect.DeepEqual(val, elem) {
		t.Errorf("Not equal: \n%s", strings.Join(pretty.Diff(val, elem), "\n"))
	}
}

// CreateJSONFixture creates a JSON file at path of the value val,
// can be used with AssertMatchFixtureJSON to compare against future values.
func CreateJSONFixture[V any](t *testing.T, path string, val V) {
	t.Helper()
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to open file to write: %s", err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(val)
	if err != nil {
		t.Fatalf("Failed to encode val: %s", err)
	}
}

// CreateTextFixture creates a text file at path of the value val,
// can be used with AssertMatchFixtureJSON to compare against future values.
func CreateTextFixture(t *testing.T, path string, val string) {
	t.Helper()
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to open file to write: %s", err)
	}

	_, err = file.WriteString(val)
	if err != nil {
		t.Fatalf("Failed to write string to file: %s", err)
	}
}

// AssertMatchFixtureText matches the Text file at path with val
func AssertMatchFixtureText(t *testing.T, path string, val string) {
	t.Helper()
	fileA, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to open fixture: %s", err)
	}

	if val != string(fileA) {
		t.Errorf("Not equal: \ngot: \n%s\n\n---\nwant:\n%s", val, string(fileA))
	}
}
