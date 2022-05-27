package ptype

import (
	"time"
)

// PTime converts a time.Time to *time.Time
func PTime(v time.Time) *time.Time {
	return &v
}

// TimeP converts a *time.Time to time.Time
func TimeP(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

// PFloat64 converts a *float64 to float64
func PFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// Float64P converts a float64 to float
func Float64P(v float64) *float64 {
	return &v
}

// PFloat32 converts a *float32 to float32
func PFloat32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

// Float32P converts a float to *float32
func Float32P(v float32) *float32 {
	return &v
}

// IntP converts an int to *int
func IntP(v int) *int {
	return &v
}

// PInt converts a pointer *int to int
func PInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// Int32P converts an int32 to *int32
func Int32P(v int32) *int32 {
	return &v
}

// PInt32 converts a *int32 to int32
func PInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// Int64P converts an int64 to *int64
func Int64P(v int64) *int64 {
	return &v
}

// PInt64 converts a *int64 to int64
func PInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

// StringP converts a string to *string
func StringP(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// PString converts a *string to string
func PString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// PBool converts *bool to bool
func PBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// BoolP converts bool to *bool
func BoolP(v bool) *bool {
	return &v
}

func PIntToPInt32(in *int) *int32 {
	if in == nil {
		return nil
	}

	out := int32(*in)
	return &out
}
func PInt32ToPInt(in *int32) *int {
	if in == nil {
		return nil
	}

	out := int(*in)
	return &out
}
func IntToPInt32(in int) *int32 {
	out := int32(in)
	return &out
}

func Int32ToPInt(in int32) *int {
	out := int(in)
	return &out
}
