package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVersion(t *testing.T) {
	var testcases = map[string]struct {
		ver          string
		parseSuccess bool
	}{
		"parse error":   {"hello", false},
		"parse success": {"v1.13.0", true},
	}

	for name, testcase := range testcases {
		tc := testcase
		t.Run(name, func(tt *testing.T) {
			tt.Parallel()
			v, err := ParseVersion(tc.ver)
			parseSuccess := (err == nil) && (v != nil)
			if tc.parseSuccess != parseSuccess {
				tt.Errorf("Parse status sould be %v and is %v", tc.parseSuccess, parseSuccess)
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

			assert.Equal(t, tc.Equal, va.Equal(vb), "Equal expected to be %v", tc.Equal)
			assert.Equal(t, tc.GreaterThan, va.GreaterThan(vb), "GreaterThan expected to be %v", tc.GreaterThan)
			assert.Equal(t, tc.GreaterThanOrEqual, va.GreaterThanOrEqual(vb), "GreaterThanOrEqual expected to be %v", tc.GreaterThanOrEqual)
			assert.Equal(t, tc.LessThan, va.LessThan(vb), "LessThan expected to be %v", tc.LessThan)
			assert.Equal(t, tc.LessThanOrEqual, va.LessThanOrEqual(vb), "LessThanOrEqual expected to be %v", tc.LessThanOrEqual)
		})
	}
}
