package vault

import (
	"context"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Client interface {
	GetConfigValue(vaultKeyName string) any
}

const (
	_vaultRoleName  = "brahma-builder"
	_vaultMountPath = "apps"
	_serviceName    = "brahma-builder"
)

func New(ctx context.Context) (*Vault, error) {
	return NewVault(ctx, _vaultRoleName, _vaultMountPath, _serviceName+"/config")
}

// LoadConfig fills the input config object with vault data. It accepts a pointer
// to the config struct and iterate over its fields.
//
// NOTE-1: This method only fill one level of embedded and non embedded structs.
// NOTE-2: To fill the embedded structs, the config struct MUST have `mapstructure:",squash"` tag.
func LoadConfig(config any, vaultCli Client) error {
	keyVals := make(map[string]any)

	fields := jsonTagValues(config)
	for _, field := range fields {
		value := vaultCli.GetConfigValue(field)
		if value == nil {
			continue
		}
		keyVals[field] = value
	}

	findAndFillInnerStructs(vaultCli, config, keyVals)

	err := mapstructure.Decode(keyVals, config)
	if err != nil {
		return err
	}

	return nil
}

func jsonTagValues(in any) []string {
	var out []string
	var t reflect.Type

	refValue, ok := in.(reflect.Value)
	if ok {
		t = refValue.Type().Elem()
	} else {
		t = reflect.TypeOf(in)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// if it's an embedded struct get the fields
		if field.Type.Kind().String() == reflect.Struct.String() {
			if !field.Anonymous {
				continue
			}
			refVal := reflect.New(field.Type)
			out = append(out, jsonTagValues(refVal)...)
			continue
		}

		value := strings.Split(field.Tag.Get("json"), ",")[0] // first one is the json tag value
		if value == "" {
			value = field.Name
		}

		out = append(out, value)
	}

	return out
}

func findAndFillInnerStructs(vaultCli Client, in any, keyVals map[string]any) {
	values := structFields(in)
	for key, value := range values {
		keyVals[key] = getStructValues(vaultCli, value)
	}
}

func getStructValues(vaultCli Client, value reflect.Value) map[string]any {
	var t reflect.Type
	t = value.Type().Elem()

	out := make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := strings.Split(field.Tag.Get("json"), ",")[0]
		if key == "" {
			key = field.Name
		}

		out[key] = vaultCli.GetConfigValue(key)
	}

	return out
}

func structFields(in any) map[string]reflect.Value {
	var t reflect.Type
	t = reflect.TypeOf(in)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	out := make(map[string]reflect.Value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// if it's a struct and not embedded
		if field.Type.Kind().String() == reflect.Struct.String() && !field.Anonymous {
			key := strings.Split(field.Tag.Get("json"), ",")[0]
			if key == "" {
				key = field.Name
			}
			out[key] = reflect.New(field.Type)
		}
	}

	return out
}
