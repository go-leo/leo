package config

import (
	"encoding/json"
	"github.com/derekparker/trie"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonContent = `{"age":35,"author":{"bio":"Gopher.\nCoding addict.\nGood man.\n","e-mail":"fake@localhost","github":"https://github.com/Unknown","name":"Unknown"},"batters":{"batter":[{"type":"Regular"},{"type":"Chocolate"},{"type":"Blueberry"},{"type":"Devil's Food"}]},"beard":true,"clothing":{"jacket":"leather","pants":{"size":"large"},"trousers":"denim"},"default":{"import_path":"gopkg.in/ini.v1","name":"ini","version":"v1"},"eyes":"brown","foos":[{"foo":[{"key":1},{"key":2},{"key":3},{"key":4}]}],"hacker":true,"hobbies":["skateboarding","snowboarding","go"],"id":"0001","name":"Cake","name_dotenv":"Cake","newkey":"remote","owner":{"bio":"MongoDB Chief Developer Advocate \u0026 Hacker at Large","dob":"1979-05-27T07:32:00Z","organization":"MongoDB","likes":["apple","orange"]},"p_batters":{"batter":{"type":"Regular"}},"p_id":"0001","p_name":"Cake","p_ppu":"0.55","p_type":"donut","ppu":0.55,"super":{"deep":{"nested":"value"}},"title":"TOML Example","title_dotenv":"DotEnv Example","type":"donut","type_dotenv":"donut"}`

func TestTree(t *testing.T) {
	var configMap map[string]any
	err := json.Unmarshal([]byte(jsonContent), &configMap)
	assert.NoError(t, err)
	configTree := trie.New()
	tree(configMap, configTree)

	dob := "1979-05-27T07:32:00Z"

	expected := map[string]interface{}{
		"super": map[string]interface{}{
			"deep": map[string]interface{}{
				"nested": "value",
			},
		},
		"super.deep": map[string]interface{}{
			"nested": "value",
		},
		"super.deep.nested":  "value",
		"owner.organization": "MongoDB",
		"batters.batter": []interface{}{
			map[string]interface{}{
				"type": "Regular",
			},
			map[string]interface{}{
				"type": "Chocolate",
			},
			map[string]interface{}{
				"type": "Blueberry",
			},
			map[string]interface{}{
				"type": "Devil's Food",
			},
		},
		"hobbies": []interface{}{
			"skateboarding", "snowboarding", "go",
		},
		"title":  "TOML Example",
		"newkey": "remote",
		"batters": map[string]interface{}{
			"batter": []interface{}{
				map[string]interface{}{
					"type": "Regular",
				},
				map[string]interface{}{
					"type": "Chocolate",
				},
				map[string]interface{}{
					"type": "Blueberry",
				},
				map[string]interface{}{
					"type": "Devil's Food",
				},
			},
		},
		"eyes": "brown",
		"age":  float64(35),
		"owner": map[string]interface{}{
			"organization": "MongoDB",
			"bio":          "MongoDB Chief Developer Advocate & Hacker at Large",
			"dob":          dob,
			"likes":        []interface{}{"apple", "orange"},
		},
		"owner.bio":     "MongoDB Chief Developer Advocate & Hacker at Large",
		"owner.likes":   []interface{}{"apple", "orange"},
		"owner.likes.0": "apple",
		"owner.likes.1": "orange",
		"type":          "donut",
		"id":            "0001",
		"name":          "Cake",
		"hacker":        true,
		"ppu":           0.55,
		"clothing": map[string]interface{}{
			"jacket":   "leather",
			"trousers": "denim",
			"pants": map[string]interface{}{
				"size": "large",
			},
		},
		"clothing.jacket":     "leather",
		"clothing.pants.size": "large",
		"clothing.trousers":   "denim",
		"owner.dob":           dob,
		"beard":               true,
		"foos": []interface{}{
			map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"key": float64(1),
					},
					map[string]interface{}{
						"key": float64(2),
					},
					map[string]interface{}{
						"key": float64(3),
					},
					map[string]interface{}{
						"key": float64(4),
					},
				},
			},
		},
	}

	for key, expectedValue := range expected {
		node, ok := configTree.Find(key)
		assert.Equal(t, true, ok, "key: "+key)
		assert.Equal(t, expectedValue, node.Meta(), "key: "+key)
	}
}
