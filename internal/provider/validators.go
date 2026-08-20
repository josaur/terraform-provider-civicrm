package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// stringOneOf returns a String validator that accepts only the given values.
func stringOneOf(values ...string) validator.String {
	return &oneOfStringValidator{values: values}
}

type oneOfStringValidator struct {
	values []string
}

func (v *oneOfStringValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Must be one of: %s", strings.Join(v.values, ", "))
}

func (v *oneOfStringValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *oneOfStringValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	for _, allowed := range v.values {
		if val == allowed {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid value",
		fmt.Sprintf("%q is not a valid value. Allowed values: %s", val, strings.Join(v.values, ", ")),
	)
}

// stringLengthAtLeast returns a String validator that rejects empty or too-short strings.
func stringLengthAtLeast(n int) validator.String {
	return &stringMinLenValidator{min: n}
}

type stringMinLenValidator struct {
	min int
}

func (v *stringMinLenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Must be at least %d character(s) long", v.min)
}

func (v *stringMinLenValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *stringMinLenValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if len(req.ConfigValue.ValueString()) < v.min {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"String too short",
			fmt.Sprintf("Value must be at least %d character(s) long, got %q", v.min, req.ConfigValue.ValueString()),
		)
	}
}

// int64OneOf returns an Int64 validator that accepts only the given values.
func int64OneOf(values ...int64) validator.Int64 {
	return &oneOfInt64Validator{values: values}
}

type oneOfInt64Validator struct {
	values []int64
}

func (v *oneOfInt64Validator) Description(_ context.Context) string {
	parts := make([]string, len(v.values))
	for i, val := range v.values {
		parts[i] = strconv.FormatInt(val, 10)
	}
	return fmt.Sprintf("Must be one of: %s", strings.Join(parts, ", "))
}

func (v *oneOfInt64Validator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *oneOfInt64Validator) ValidateInt64(_ context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueInt64()
	for _, allowed := range v.values {
		if val == allowed {
			return
		}
	}
	parts := make([]string, len(v.values))
	for i, allowed := range v.values {
		parts[i] = strconv.FormatInt(allowed, 10)
	}
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid value",
		fmt.Sprintf("%d is not a valid value. Allowed values: %s", val, strings.Join(parts, ", ")),
	)
}

// int64AtLeast returns an Int64 validator that rejects values below min.
func int64AtLeast(min int64) validator.Int64 {
	return &int64MinValidator{min: min}
}

type int64MinValidator struct {
	min int64
}

func (v *int64MinValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Must be at least %d", v.min)
}

func (v *int64MinValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *int64MinValidator) ValidateInt64(_ context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.ConfigValue.ValueInt64() < v.min {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Value too small",
			fmt.Sprintf("Value must be at least %d, got %d", v.min, req.ConfigValue.ValueInt64()),
		)
	}
}
