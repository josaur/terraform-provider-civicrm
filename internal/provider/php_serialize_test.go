package provider

import (
	"reflect"
	"testing"
)

// TestPHPSerializeMatchesPHP pins phpSerialize's output against real PHP
// serialize() output, captured with `php -r 'echo serialize(...);'` — the
// exact bytes CiviRules' unserialize() must accept.
func TestPHPSerializeMatchesPHP(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"null", nil, "N;"},
		{"true", true, "b:1;"},
		{"false", false, "b:0;"},
		{"int", float64(2), "i:2;"},
		{"negative int", float64(-5), "i:-5;"},
		{"float", 1.5, "d:1.5;"},
		{"string", "foo", `s:3:"foo";`},
		{"empty string", "", `s:0:"";`},
		{"indexed array", []any{"a", "b"}, `a:2:{i:0;s:1:"a";i:1;s:1:"b";}`},
		{
			"assoc array single key",
			map[string]any{"status_id": float64(2)},
			`a:1:{s:9:"status_id";i:2;}`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := phpSerialize(tc.in)
			if got != tc.want {
				t.Errorf("phpSerialize(%#v) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

// TestPHPUnserializeMatchesPHP decodes the exact byte strings PHP's
// serialize() produces (see TestPHPSerializeMatchesPHP and the docker
// verification: `php -r 'echo serialize([...])`) and checks the resulting
// Go shape matches what encoding/json.Unmarshal would have produced for the
// equivalent JSON value.
func TestPHPUnserializeMatchesPHP(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want any
	}{
		{"null", "N;", nil},
		{"true", "b:1;", true},
		{"false", "b:0;", false},
		{"int", "i:2;", float64(2)},
		{"negative int", "i:-5;", float64(-5)},
		{"float", "d:1.5;", 1.5},
		{"string", `s:3:"foo";`, "foo"},
		{"empty string", `s:0:"";`, ""},
		{"indexed array", `a:2:{i:0;s:1:"a";i:1;s:1:"b";}`, []any{"a", "b"}},
		{
			"assoc array",
			`a:1:{s:9:"status_id";i:2;}`,
			map[string]any{"status_id": float64(2)},
		},
		{
			// Real serialize() output for
			// ["status_id"=>2,"name"=>"foo","flag"=>true,"nested"=>["a","b"],"assoc"=>["x"=>1,"y"=>2],"f"=>1.5,"n"=>null]
			// captured against a live CiviCRM container's php -r.
			"mixed nested",
			`a:7:{s:9:"status_id";i:2;s:4:"name";s:3:"foo";s:4:"flag";b:1;s:6:"nested";a:2:{i:0;s:1:"a";i:1;s:1:"b";}s:5:"assoc";a:2:{s:1:"x";i:1;s:1:"y";i:2;}s:1:"f";d:1.5;s:1:"n";N;}`,
			map[string]any{
				"status_id": float64(2),
				"name":      "foo",
				"flag":      true,
				"nested":    []any{"a", "b"},
				"assoc":     map[string]any{"x": float64(1), "y": float64(2)},
				"f":         1.5,
				"n":         nil,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := phpUnserialize(tc.in)
			if err != nil {
				t.Fatalf("phpUnserialize(%q) error: %s", tc.in, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("phpUnserialize(%q) = %#v, want %#v", tc.in, got, tc.want)
			}
		})
	}
}

// TestPHPSerializeRoundTrip checks phpSerialize -> phpUnserialize recovers
// the original JSON-shaped value, for values as they'd arrive after
// decodeJSONAttribute(jsonencode(...)).
func TestPHPSerializeRoundTrip(t *testing.T) {
	values := []any{
		nil,
		true,
		false,
		float64(42),
		3.25,
		"hello",
		[]any{float64(1), float64(2), float64(3)},
		map[string]any{"a": float64(1), "b": "two", "c": []any{"x", "y"}},
		map[string]any{
			"tag_id":  float64(46),
			"tag_ids": []any{float64(1), float64(2)},
		},
	}
	for _, v := range values {
		serialized := phpSerialize(v)
		got, err := phpUnserialize(serialized)
		if err != nil {
			t.Fatalf("phpUnserialize(phpSerialize(%#v)) = error: %s (serialized: %q)", v, err, serialized)
		}
		if !reflect.DeepEqual(got, v) {
			t.Errorf("round trip mismatch: in=%#v serialized=%q out=%#v", v, serialized, got)
		}
	}
}

// TestPHPUnserializeRejectsJSON is the regression test for the bug this
// feature fixes: a bare JSON string is not valid PHP serialize() data, and
// must fail loudly instead of being silently accepted (which is what made
// jsonencode(...)-written action_params/condition_params break CiviRules at
// runtime before this provider did the JSON<->serialize() conversion
// itself).
func TestPHPUnserializeRejectsJSON(t *testing.T) {
	_, err := phpUnserialize(`{"tag_id":46}`)
	if err == nil {
		t.Fatal("phpUnserialize of a JSON string should fail, got nil error")
	}
}
