// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_resyncPeriod", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("resyncPeriod"); err == nil {
				assert.Equal(t, string("30s"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "30s"

			cmdFlags.Set("resyncPeriod", testValue)
			if vString, err := cmdFlags.GetString("resyncPeriod"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ResyncPeriod)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_limitNamespace", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("limitNamespace"); err == nil {
				assert.Equal(t, string(""), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("limitNamespace", testValue)
			if vString, err := cmdFlags.GetString("limitNamespace"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.LimitNamespace)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_metricsPrefix", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("metricsPrefix"); err == nil {
				assert.Equal(t, string("flinkk8soperator"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metricsPrefix", testValue)
			if vString, err := cmdFlags.GetString("metricsPrefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MetricsPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_prof-port", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("prof-port"); err == nil {
				assert.Equal(t, string("10254"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "10254"

			cmdFlags.Set("prof-port", testValue)
			if vString, err := cmdFlags.GetString("prof-port"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ProfilerPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_ingressUrlFormat", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("ingressUrlFormat"); err == nil {
				assert.Equal(t, string(*new(string)), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("ingressUrlFormat", testValue)
			if vString, err := cmdFlags.GetString("ingressUrlFormat"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.FlinkIngressURLFormat)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_useKubectlProxy", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vBool, err := cmdFlags.GetBool("useKubectlProxy"); err == nil {
				assert.Equal(t, bool(*new(bool)), vBool)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("useKubectlProxy", testValue)
			if vBool, err := cmdFlags.GetBool("useKubectlProxy"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.UseProxy)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_ProxyPort", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("ProxyPort"); err == nil {
				assert.Equal(t, string("8001"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "8001"

			cmdFlags.Set("ProxyPort", testValue)
			if vString, err := cmdFlags.GetString("ProxyPort"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ProxyPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_containerNameFormat", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("containerNameFormat"); err == nil {
				assert.Equal(t, string(*new(string)), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("containerNameFormat", testValue)
			if vString, err := cmdFlags.GetString("containerNameFormat"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ContainerNameFormat)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_workers", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				assert.Equal(t, int(4), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("workers", testValue)
			if vInt, err := cmdFlags.GetInt("workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_baseBackoffDuration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("baseBackoffDuration"); err == nil {
				assert.Equal(t, string("100ms"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "100ms"

			cmdFlags.Set("baseBackoffDuration", testValue)
			if vString, err := cmdFlags.GetString("baseBackoffDuration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.BaseBackoffDuration)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxBackoffDuration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("maxBackoffDuration"); err == nil {
				assert.Equal(t, string("30s"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "30s"

			cmdFlags.Set("maxBackoffDuration", testValue)
			if vString, err := cmdFlags.GetString("maxBackoffDuration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MaxBackoffDuration)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_maxErrDuration", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("maxErrDuration"); err == nil {
				assert.Equal(t, string("5m"), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "5m"

			cmdFlags.Set("maxErrDuration", testValue)
			if vString, err := cmdFlags.GetString("maxErrDuration"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MaxErrDuration)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
