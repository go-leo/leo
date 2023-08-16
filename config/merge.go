package config

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

func merge(src, tgt map[string]any) error {
	return mergeMaps(src, tgt, nil)
}

func mergeMaps(src, tgt map[string]any, itgt map[any]any) error {
	for srcKey, srcVal := range src {
		// key ignore case, not exist, set
		tgtKey, ok := keyExists(srcKey, tgt)
		if !ok {
			setToTarget(srcKey, srcVal, tgt, itgt)
			continue
		}

		// care about case, not exist, set
		tgtVal, ok := tgt[tgtKey]
		if !ok {
			setToTarget(srcKey, srcVal, tgt, itgt)
			continue
		}

		srcValType := reflect.TypeOf(srcVal)
		tgtValType := reflect.TypeOf(tgtVal)
		switch tmpTgtVal := tgtVal.(type) {
		case map[any]any:
			tmpSrcVal, ok := srcVal.(map[any]any)
			if !ok {
				return fmt.Errorf(`failed to cast src value to map[any]any, srcKey: %s, srcValType: %s, tgtValType: %s, srcVal: %s, tgtVal: %s`, srcKey, srcValType, tgtValType, srcVal, tgtVal)
			}
			err := mergeMaps(anyMapToStringMap(tmpSrcVal), anyMapToStringMap(tmpTgtVal), tmpTgtVal)
			if err != nil {
				return err
			}
		case map[string]any:
			tmpSrcVal, ok := srcVal.(map[string]any)
			if !ok {
				return fmt.Errorf(`failed to cast src value to map[string]any, srcKey: %s, srcValType: %s, tgtValType: %s, srcVal: %s, tgtVal: %s`, srcKey, srcValType, tgtValType, srcVal, tgtVal)
			}
			err := mergeMaps(tmpSrcVal, tmpTgtVal, nil)
			if err != nil {
				return err
			}
		default:
			tgt[tgtKey] = srcVal
			if itgt != nil {
				itgt[tgtKey] = srcVal
			}
			setToTarget(tgtKey, srcVal, tgt, itgt)
		}
	}

	return nil
}

func setToTarget(key string, val any, tgt map[string]any, itgt map[any]any) {
	tgt[key] = val
	if itgt != nil {
		itgt[key] = val
	}
}

func anyMapToStringMap(src map[any]any) map[string]any {
	tgt := map[string]any{}
	for k, v := range src {
		tgt[cast.ToString(k)] = v
	}
	return tgt
}

func keyExists(key string, tgt map[string]any) (string, bool) {
	lowerKey := strings.ToLower(key)
	for configKey := range tgt {
		if strings.ToLower(configKey) == lowerKey {
			return configKey, true
		}
	}
	return "", false
}
