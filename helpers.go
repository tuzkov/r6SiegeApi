package r6

import (
	"encoding/json"
	"log"
)

// intFromJSONNumber возвращает int
func intFromJSONNumber(v json.Number) int {
	i, err := v.Int64()
	if err != nil {
		log.Println("Ошибка при intFromJSONNumber:", err)
		return 0
	}
	return int(i)
}
