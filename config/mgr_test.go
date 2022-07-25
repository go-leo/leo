package config

/*
package config_test


import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/go-leo/leo/config/mgr"
	"github.com/go-leo/leo/config/mgr/medium/file"
	"github.com/go-leo/leo/config/mgr/parser"
	"github.com/go-leo/leo/config/mgr/valuer"
)

type Config struct {
	String               string              `mapstructure:"string" json:"string" yaml:"string"`
	Bool                 bool                `mapstructure:"bool" json:"bool" yaml:"bool"`
	Int                  int                 `mapstructure:"int" json:"int" yaml:"int"`
	Int32                int32               `mapstructure:"int_32" json:"int_32" yaml:"int_32"`
	Int64                int64               `mapstructure:"int_64" json:"int_64" yaml:"int_64"`
	UInt                 uint                `mapstructure:"u_int" json:"u_int" yaml:"u_int"`
	UInt32               uint32              `mapstructure:"u_int_32" json:"u_int_32" yaml:"u_int_32"`
	UInt64               uint64              `mapstructure:"u_int_64" json:"u_int_64" yaml:"u_int_64"`
	Float64              float64             `mapstructure:"float_64" json:"float_64" yaml:"float_64"`
	Time                 time.Time           `mapstructure:"time" json:"time" yaml:"time"`
	Duration             time.Duration       `mapstructure:"duration" json:"duration" yaml:"duration"`
	IntSlice             []int               `mapstructure:"int_slice" json:"int_slice" yaml:"int_slice"`
	StringSlice          []string            `mapstructure:"string_slice" json:"string_slice" yaml:"string_slice"`
	StringMapString      map[string]string   `mapstructure:"string_map_string" json:"string_map_string" yaml:"string_map_string"`
	StringMapStringSlice map[string][]string `mapstructure:"string_map_string_slice" json:"string_map_string_slice" yaml:"string_map_string_slice"`
	Map                  Map                 `mapstructure:"map" json:"map" yaml:"map"`
}

type Map struct {
	String               string              `mapstructure:"string" json:"string" yaml:"string"`
	Bool                 bool                `mapstructure:"bool" json:"bool" yaml:"bool"`
	Int                  int                 `mapstructure:"int" json:"int" yaml:"int"`
	Int32                int32               `mapstructure:"int_32" json:"int_32" yaml:"int_32"`
	Int64                int64               `mapstructure:"int_64" json:"int_64" yaml:"int_64"`
	UInt                 uint                `mapstructure:"u_int" json:"u_int" yaml:"u_int"`
	UInt32               uint32              `mapstructure:"u_int_32" json:"u_int_32" yaml:"u_int_32"`
	UInt64               uint64              `mapstructure:"u_int_64" json:"u_int_64" yaml:"u_int_64"`
	Float64              float64             `mapstructure:"float_64" json:"float_64" yaml:"float_64"`
	Time                 time.Time           `mapstructure:"time" json:"time" yaml:"time"`
	Duration             time.Duration       `mapstructure:"duration" json:"duration" yaml:"duration"`
	IntSlice             []int               `mapstructure:"int_slice" json:"int_slice" yaml:"int_slice"`
	StringSlice          []string            `mapstructure:"string_slice" json:"string_slice" yaml:"string_slice"`
	StringMapString      map[string]string   `mapstructure:"string_map_string" json:"string_map_string" yaml:"string_map_string"`
	StringMapStringSlice map[string][]string `mapstructure:"string_map_string_slice" json:"string_map_string_slice" yaml:"string_map_string_slice"`
}

var now = time.Now()

var configMap = Map{
	String:               "subString",
	Bool:                 true,
	Int:                  -10,
	Int32:                200,
	Int64:                3000,
	UInt:                 1333,
	UInt32:               4133,
	UInt64:               5643,
	Float64:              4.15,
	Time:                 time.Now(),
	Duration:             time.Hour,
	IntSlice:             []int{1, 2},
	StringSlice:          []string{"a", "b"},
	StringMapString:      map[string]string{"k1": "v1"},
	StringMapStringSlice: map[string][]string{"k1": {"c", "d"}},
}

var configuration = Config{
	String:               "String",
	Bool:                 true,
	Int:                  10,
	Int32:                -200,
	Int64:                -3000,
	UInt:                 133,
	UInt32:               413,
	UInt64:               564,
	Float64:              1.3,
	Time:                 now,
	Duration:             time.Second,
	IntSlice:             []int{1, 2, 3, 4, 5, 6},
	StringSlice:          []string{"a", "b", "c", "d"},
	StringMapString:      map[string]string{"k1": "v1"},
	StringMapStringSlice: map[string][]string{"k1": {"a", "b", "c", "d"}},
	Map:                  configMap,
}

var configFile = "/tmp/mgr.yaml"

var contentType = "yaml"

func TestMain(m *testing.M) {
	data, _ := yaml.Marshal(&configuration)
	err := ioutil.WriteFile(configFile, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	m.Run()
	err = os.Remove(configFile)
	if err != nil {
		panic(err)
	}
}

func TestManager(t *testing.T) {
	manager := mgr.NewManager(
		mgr.WithLoader(file.NewLoader(configFile)),
		mgr.WithParser(parser.NewYamlParser()),
		mgr.WithValuer(valuer.NewTrieTreeValuer()),
	)
	err := manager.ReadConfig()
	if err != nil {
		t.Fatal(err)
	}

	stringVal, err := manager.GetString("string")
	assert.Nil(t, err)
	assert.Equal(t, configuration.String, stringVal)

	boolVal, err := manager.GetBool("bool")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Bool, boolVal)

	intVal, err := manager.GetInt("int")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Int, intVal)

	int32Val, err := manager.GetInt32("int_32")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Int32, int32Val)

	int64Val, err := manager.GetInt64("int_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Int64, int64Val)

	uintVal, err := manager.GetUint("u_int")
	assert.Nil(t, err)
	assert.Equal(t, configuration.UInt, uintVal)

	uint32Val, err := manager.GetUint32("u_int_32")
	assert.Nil(t, err)
	assert.Equal(t, configuration.UInt32, uint32Val)

	uint64Val, err := manager.GetUint64("u_int_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.UInt64, uint64Val)

	float64Val, err := manager.GetFloat64("float_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Float64, float64Val)

	timeVal, err := manager.GetTime("time")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Time.UnixNano(), timeVal.UnixNano())

	durationVal, err := manager.GetDuration("duration")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Duration, durationVal)

	intSliceVal, err := manager.GetIntSlice("int_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.IntSlice, intSliceVal)

	stringSliceVal, err := manager.GetStringSlice("string_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.StringSlice, stringSliceVal)

	StringMapStringVal, err := manager.GetStringMapString("string_map_string")
	assert.Nil(t, err)
	assert.Equal(t, configuration.StringMapString, StringMapStringVal)

	StringMapStringSliceVal, err := manager.GetStringMapStringSlice("string_map_string_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.StringMapStringSlice, StringMapStringSliceVal)

	MapVal, err := manager.Get("map")
	assert.Nil(t, err)
	assert.NotNil(t, MapVal)

	mapString, err := manager.GetString("map.string")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.String, mapString)

	mapBool, err := manager.GetBool("map.bool")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Bool, mapBool)

	mapInt, err := manager.GetInt("map.int")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Int, mapInt)

	mapInt32, err := manager.GetInt32("map.int_32")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Int32, mapInt32)

	mapInt64, err := manager.GetInt64("map.int_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Int64, mapInt64)

	mapUInt, err := manager.GetUint("map.u_int")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.UInt, mapUInt)

	mapUInt32, err := manager.GetUint32("map.u_int_32")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.UInt32, mapUInt32)

	mapUInt64, err := manager.GetUint64("map.u_int_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.UInt64, mapUInt64)

	mapFloat64, err := manager.GetFloat64("map.float_64")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Float64, mapFloat64)

	mapTime, err := manager.GetTime("map.time")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Time.UnixNano(), mapTime.UnixNano())

	mapDuration, err := manager.GetDuration("map.duration")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.Duration, mapDuration)

	mapIntSlice, err := manager.GetIntSlice("map.int_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.IntSlice, mapIntSlice)

	mapStringSlice, err := manager.GetStringSlice("map.string_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.StringSlice, mapStringSlice)

	mapStringMapString, err := manager.GetStringMapString("map.string_map_string")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.StringMapString, mapStringMapString)

	mapStringMapStringSlice, err := manager.GetStringMapStringSlice("map.string_map_string_slice")
	assert.Nil(t, err)
	assert.Equal(t, configuration.Map.StringMapStringSlice, mapStringMapStringSlice)

	var conf Config
	err = manager.Unmarshal(&conf)
	assert.Nil(t, err)
	actualConfData, err := json.Marshal(&conf)
	assert.Nil(t, err)
	expectedConfData, err := json.Marshal(&configuration)
	assert.Nil(t, err)
	assert.Equal(t, string(expectedConfData), string(actualConfData))

	var confMap Map
	err = manager.UnmarshalKey("map", &confMap)
	assert.Nil(t, err)
	actualConfMapData, err := json.Marshal(&confMap)
	assert.Nil(t, err)
	expectedConfigMapData, err := json.Marshal(&configMap)
	assert.Nil(t, err)
	assert.Equal(t, string(expectedConfigMapData), string(actualConfMapData))

}

func TestWatcher(t *testing.T) {
	confFilename := "filewatcher.yaml"
	fp := filepath.Join(os.TempDir(), confFilename)
	var confContent = `
bool: true
int: 10
int_32: -200
int_64: -3000
u_int: 133
u_int_32: 413
u_int_64: 564
float_64: 1.3
time: 2021-07-07T17:16:12.361234+08:00
duration: 1s
`
	err := ioutil.WriteFile(fp, []byte(confContent), os.ModePerm)
	assert.Nil(t, err)

	watcher := file.NewWatcher(fp)

	manager := mgr.NewManager(mgr.WithWatcher(watcher))

	err = manager.StartWatch()
	assert.Nil(t, err)

	go func() {
		file, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, os.ModePerm)
		defer file.Close()
		assert.Nil(t, err, "failed open file")

		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			n, err := file.WriteString("float_32: 0.3\n")
			assert.Nil(t, err, "failed write string")
			assert.Greater(t, n, 0, "failed write string")
		}
		time.Sleep(2 * time.Second)
		err = manager.StopWatch()
		assert.Nil(t, err)
	}()

	events := manager.Events()
	tmpConfContent := confContent
	for event := range events {
		tmpConfContent += "float_32: 0.3\n"
		assert.Equal(t, tmpConfContent, string(event.Data()))
	}

	err = manager.StopWatch()
	assert.Nil(t, err)

	err = manager.StopWatch()
	assert.Nil(t, err)
	err = os.Remove(fp)
	assert.Nil(t, err)
}
*/
