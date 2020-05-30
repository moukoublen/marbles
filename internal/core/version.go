package core

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

// Version struct represent a version
// hidding the inner implementation which is hashicorp go-version
type Version struct {
	inner *version.Version
}

// Compare compares this version to another version.
// Returns hashicorp/go-version compare which is:
//   -1 this is smalle
//    0 equals
//    1 this is larger
func (v *Version) Compare(other *Version) int {
	return v.inner.Compare(other.inner)
}

// Equal tests if two versions are equal.
func (v *Version) Equal(o *Version) bool {
	return v.Compare(o) == 0
}

// GreaterThan tests if this version is greater than another version.
func (v *Version) GreaterThan(o *Version) bool {
	return v.Compare(o) > 0
}

// GreaterThanOrEqual tests if this version is greater than or equal to another version.
func (v *Version) GreaterThanOrEqual(o *Version) bool {
	return v.Compare(o) >= 0
}

// LessThan tests if this version is less than another version.
func (v *Version) LessThan(o *Version) bool {
	return v.Compare(o) < 0
}

// LessThanOrEqual tests if this version is less than or equal to another version.
func (v *Version) LessThanOrEqual(o *Version) bool {
	return v.Compare(o) <= 0
}

// ParseVersion fubction
func ParseVersion(ver string) (*Version, error) {
	v, err := version.NewVersion(ver)
	if err != nil {
		return nil, fmt.Errorf("Version string \"%s\" cannot be parsed", ver)
	}

	return &Version{v}, nil
}
