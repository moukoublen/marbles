package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVersion(t *testing.T) {
	var testcases = map[string]struct {
		ver          string
		parseSuccess bool
		printedVer   string
	}{
		"parse error":   {"hello", false, ""},
		"parse success": {"v1.13.0", true, "1.13.0"},
	}

	for name, testcase := range testcases {
		tc := testcase
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			v, err := ParseVersion(tc.ver)
			parseSuccess := (err == nil) && (v != nil)
			if tc.parseSuccess != parseSuccess {
				tt.Errorf("Parse status should be %v and is %v", tc.parseSuccess, parseSuccess)
			}

			if tc.parseSuccess {
				s := v.String()
				if s != tc.printedVer {
					tt.Errorf("Error in version string. Expected %s got %s", tc.printedVer, s)
				}
			}
		})
	}
}

func TestVersionComparison(t *testing.T) {
	var testcases = map[string]struct {
		verA               string
		verB               string
		Equal              bool
		GreaterThan        bool
		GreaterThanOrEqual bool
		LessThan           bool
		LessThanOrEqual    bool
	}{
		"test a": {"v1.13.0", "v1.13.1", false, false, false, true, true},
		"test b": {"v1.13.1", "v1.13.1", true, false, true, false, true},
		"test c": {"v1.13.1", "v1.13.0", false, true, true, false, false},
	}

	for name, testcase := range testcases {
		tc := testcase
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			va, err := ParseVersion(tc.verA)
			assert.NoError(t, err)
			assert.NotNil(t, va)
			vb, err := ParseVersion(tc.verB)
			assert.NoError(t, err)
			assert.NotNil(t, vb)

			boolCheck(tt, tc.Equal, va.Equal(vb), "Equal")
			boolCheck(tt, tc.GreaterThan, va.GreaterThan(vb), "GreaterThan")
			boolCheck(tt, tc.GreaterThanOrEqual, va.GreaterThanOrEqual(vb), "GreaterThanOrEqual")
			boolCheck(tt, tc.LessThan, va.LessThan(vb), "LessThan")
			boolCheck(tt, tc.LessThanOrEqual, va.LessThanOrEqual(vb), "LessThanOrEqual")
		})
	}
}

func boolCheck(t *testing.T, expected, outcome bool, name string) {
	if expected != outcome {
		t.Errorf("%s expected to be %v", name, expected)
	}
}
