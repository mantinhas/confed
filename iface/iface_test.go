package iface_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mantinhas/confed/iface"
	"github.com/mantinhas/confed/utils"
)

func TestGetStringAttributionTokensTypes(t *testing.T) {
	testdata := utils.ReadFile("../scanner/testdata/string_attribution_test.toml")

	var value string
	var ok bool

	value, ok = iface.Get("name", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "john smith", value)

	value, ok = iface.Get("nationality", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "usa", value)

	_, ok = iface.Get("not-included", testdata)
	assert.Equal(t, false, ok)
}

func TestGetNumberAttributionTokensTypes(t *testing.T) {
	testdata := utils.ReadFile("../scanner/testdata/number_attribution_test.toml")

	var value string
	var ok bool

	value, ok = iface.Get("age", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "16", value)

	value, ok = iface.Get("grade", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "20.170", value)

	_, ok = iface.Get("not-included", testdata)
	assert.Equal(t, false, ok)
}

func TestGetBoolAttributionTokensTypes(t *testing.T) {
	testdata := utils.ReadFile("../scanner/testdata/bool_attribution_test.toml")

	var value string
	var ok bool

	value, ok = iface.Get("isUp_", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "true", value)

	value, ok = iface.Get("ISDOWN", testdata)
	assert.Equal(t, true, ok)
	assert.Equal(t, "false", value)

	_, ok = iface.Get("not-included", testdata)
	assert.Equal(t, false, ok)
}
