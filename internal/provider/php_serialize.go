package provider

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// phpSerialize encodes a Go value (as produced by encoding/json.Unmarshal
// into `any`, i.e. only map[string]any, []any, string, float64, bool, nil)
// into PHP's serialize() wire format. CiviRules stores action_params /
// condition_params this way and decodes them with unserialize(); see
// resource_civirules_rule_action.go / resource_civirules_rule_condition.go.
//
// JSON objects become PHP associative arrays (a:N:{...} with string keys);
// JSON arrays become PHP indexed arrays (a:N:{...} with 0-based int keys) —
// both are the same PHP "array" type, so this round-trips through
// unserialize() correctly for either shape.
func phpSerialize(v any) string {
	var b strings.Builder
	phpSerializeInto(&b, v)
	return b.String()
}

func phpSerializeInto(b *strings.Builder, v any) {
	switch val := v.(type) {
	case nil:
		b.WriteString("N;")
	case bool:
		if val {
			b.WriteString("b:1;")
		} else {
			b.WriteString("b:0;")
		}
	case float64:
		if val == float64(int64(val)) {
			b.WriteString("i:")
			b.WriteString(strconv.FormatInt(int64(val), 10))
			b.WriteString(";")
		} else {
			b.WriteString("d:")
			b.WriteString(strconv.FormatFloat(val, 'G', -1, 64))
			b.WriteString(";")
		}
	case string:
		phpSerializeString(b, val)
	case []any:
		fmt.Fprintf(b, "a:%d:{", len(val))
		for i, item := range val {
			b.WriteString("i:")
			b.WriteString(strconv.Itoa(i))
			b.WriteString(";")
			phpSerializeInto(b, item)
		}
		b.WriteString("}")
	case map[string]any:
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		fmt.Fprintf(b, "a:%d:{", len(val))
		for _, k := range keys {
			phpSerializeString(b, k)
			phpSerializeInto(b, val[k])
		}
		b.WriteString("}")
	default:
		// Unreachable for values decoded from encoding/json, but fall back
		// to a string representation rather than silently dropping data.
		phpSerializeString(b, fmt.Sprintf("%v", val))
	}
}

func phpSerializeString(b *strings.Builder, s string) {
	fmt.Fprintf(b, "s:%d:\"%s\";", len(s), s)
}

// phpUnserialize decodes PHP's serialize() wire format into the same Go
// shapes encoding/json.Unmarshal would produce (map[string]any, []any,
// string, float64, bool, nil), so the result can be passed straight into
// encodeJSONAttribute for state. PHP arrays with purely sequential 0-based
// integer keys decode to []any; any other key shape (string keys, sparse or
// non-zero-based integer keys) decodes to map[string]any so no data is lost.
func phpUnserialize(s string) (any, error) {
	v, rest, err := phpUnserializeValue(s)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(rest) != "" {
		return nil, fmt.Errorf("trailing data after PHP-serialized value: %q", rest)
	}
	return v, nil
}

func phpUnserializeValue(s string) (any, string, error) {
	if len(s) == 0 {
		return nil, "", fmt.Errorf("unexpected end of input")
	}
	switch s[0] {
	case 'N':
		// N;
		if !strings.HasPrefix(s, "N;") {
			return nil, "", fmt.Errorf("malformed null at %q", s)
		}
		return nil, s[2:], nil
	case 'b':
		// b:0; or b:1;
		if len(s) < 4 || s[1] != ':' || s[3] != ';' {
			return nil, "", fmt.Errorf("malformed bool at %q", s)
		}
		return s[2] == '1', s[4:], nil
	case 'i':
		// i:123;
		end := strings.IndexByte(s, ';')
		if end < 0 || len(s) < 3 || s[1] != ':' {
			return nil, "", fmt.Errorf("malformed int at %q", s)
		}
		n, err := strconv.ParseInt(s[2:end], 10, 64)
		if err != nil {
			return nil, "", fmt.Errorf("malformed int at %q: %w", s, err)
		}
		return float64(n), s[end+1:], nil
	case 'd':
		// d:1.5;
		end := strings.IndexByte(s, ';')
		if end < 0 || len(s) < 3 || s[1] != ':' {
			return nil, "", fmt.Errorf("malformed float at %q", s)
		}
		f, err := strconv.ParseFloat(s[2:end], 64)
		if err != nil {
			return nil, "", fmt.Errorf("malformed float at %q: %w", s, err)
		}
		return f, s[end+1:], nil
	case 's':
		str, rest, err := phpUnserializeStringValue(s)
		if err != nil {
			return nil, "", err
		}
		return str, rest, nil
	case 'a':
		return phpUnserializeArray(s)
	default:
		return nil, "", fmt.Errorf("unsupported PHP serialize type marker %q in %q", s[0:1], s)
	}
}

// phpUnserializeStringValue parses a s:<byte-len>:"<bytes>"; token. PHP
// measures the length in bytes, not runes, so slicing must happen on the
// byte representation.
func phpUnserializeStringValue(s string) (string, string, error) {
	if len(s) < 2 || s[1] != ':' {
		return "", "", fmt.Errorf("malformed string at %q", s)
	}
	colon2 := strings.IndexByte(s[2:], ':')
	if colon2 < 0 {
		return "", "", fmt.Errorf("malformed string length at %q", s)
	}
	colon2 += 2
	length, err := strconv.Atoi(s[2:colon2])
	if err != nil {
		return "", "", fmt.Errorf("malformed string length at %q: %w", s, err)
	}
	if len(s) < colon2+2 || s[colon2] != ':' || s[colon2+1] != '"' {
		return "", "", fmt.Errorf("malformed string opening quote at %q", s)
	}
	start := colon2 + 2
	end := start + length
	if len(s) < end+2 || s[end] != '"' || s[end+1] != ';' {
		return "", "", fmt.Errorf("malformed string body/terminator at %q", s)
	}
	return s[start:end], s[end+2:], nil
}

func phpUnserializeArray(s string) (any, string, error) {
	if len(s) < 2 || s[1] != ':' {
		return nil, "", fmt.Errorf("malformed array at %q", s)
	}
	colon2 := strings.IndexByte(s[2:], ':')
	if colon2 < 0 {
		return nil, "", fmt.Errorf("malformed array length at %q", s)
	}
	colon2 += 2
	count, err := strconv.Atoi(s[2:colon2])
	if err != nil {
		return nil, "", fmt.Errorf("malformed array length at %q: %w", s, err)
	}
	if len(s) < colon2+2 || s[colon2] != ':' || s[colon2+1] != '{' {
		return nil, "", fmt.Errorf("malformed array opening brace at %q", s)
	}
	rest := s[colon2+2:]

	type kv struct {
		key any
		val any
	}
	entries := make([]kv, 0, count)
	for i := 0; i < count; i++ {
		var key any
		var err error
		key, rest, err = phpUnserializeValue(rest)
		if err != nil {
			return nil, "", fmt.Errorf("array key %d: %w", i, err)
		}
		var val any
		val, rest, err = phpUnserializeValue(rest)
		if err != nil {
			return nil, "", fmt.Errorf("array value %d: %w", i, err)
		}
		entries = append(entries, kv{key: key, val: val})
	}
	if len(rest) < 1 || rest[0] != '}' {
		return nil, "", fmt.Errorf("malformed array terminator at %q", rest)
	}
	rest = rest[1:]

	isSequential := true
	for i, e := range entries {
		n, ok := e.key.(float64)
		if !ok || n != float64(i) {
			isSequential = false
			break
		}
	}

	if isSequential {
		list := make([]any, len(entries))
		for i, e := range entries {
			list[i] = e.val
		}
		return list, rest, nil
	}

	obj := make(map[string]any, len(entries))
	for _, e := range entries {
		var k string
		switch kt := e.key.(type) {
		case string:
			k = kt
		case float64:
			k = strconv.FormatInt(int64(kt), 10)
		default:
			k = fmt.Sprintf("%v", kt)
		}
		obj[k] = e.val
	}
	return obj, rest, nil
}
