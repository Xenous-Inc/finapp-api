package jsoner

import (
	"encoding/json"

	"github.com/Xenous-Inc/finapp-api/internal/utils/logger/log"
)

func Jsonify(obj interface{}) string {
	if obj == nil {
		log.Warn("Internal", "jsoner Jsonify")
		return ""
	}
	str, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		log.Warn("Internal", "jsoner Jsonify")
		return ""
	}

	return string(str)
}
