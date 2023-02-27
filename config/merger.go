package config

import (
	"fmt"
	"reflect"
	"strings"
)

type merger struct{}

func (m merger) Merge(src, tgt map[string]any) error {
	return m.mergeMaps(src, tgt, nil)
}

func (m merger) mergeMaps(src, tgt map[string]any, itgt map[any]any) error {
	for srcKey, srcVal := range src {
		// key ignore case, not exist, set
		tgtKey, ok := m.keyExists(srcKey, tgt)
		if !ok {
			m.setToTarget(srcKey, srcKey, tgt, itgt)
			continue
		}

		// care about case, not exist, set
		tgtVal, ok := tgt[tgtKey]
		if !ok {
			m.setToTarget(srcKey, srcKey, tgt, itgt)
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
			err := m.mergeMaps(m.anyMapToStringMap(tmpSrcVal), m.anyMapToStringMap(tmpTgtVal), tmpTgtVal)
			if err != nil {
				return err
			}
		case map[string]any:
			tmpSrcVal, ok := srcVal.(map[string]any)
			if !ok {
				return fmt.Errorf(`failed to cast src value to map[string]any, srcKey: %s, srcValType: %s, tgtValType: %s, srcVal: %s, tgtVal: %s`, srcKey, srcValType, tgtValType, srcVal, tgtVal)
			}
			err := m.mergeMaps(tmpSrcVal, tmpTgtVal, nil)
			if err != nil {
				return err
			}
		default:
			tgt[tgtKey] = srcVal
			if itgt != nil {
				itgt[tgtKey] = srcVal
			}
			m.setToTarget(tgtKey, srcVal, tgt, itgt)
		}
	}

	return nil
}

func (m merger) setToTarget(key string, val any, tgt map[string]any, itgt map[any]any) {
	tgt[key] = val
	if itgt != nil {
		itgt[key] = val
	}
}

func (m merger) anyMapToStringMap(src map[any]any) map[string]any {
	tgt := map[string]any{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

func (m merger) keyExists(key string, tgt map[string]any) (string, bool) {
	lowerKey := strings.ToLower(key)
	for configKey := range tgt {
		if strings.ToLower(configKey) == lowerKey {
			return configKey, true
		}
	}
	return "", false
}
