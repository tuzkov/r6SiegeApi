package r6

import (
	"log"
)

// intFromInterface кастует interface{} в int. Если не инт - то 0
func intFromInterface(v interface{}) int {
	switch v.(type) {
	case float32:
		if i, ok := v.(float32); ok {
			return int(i)
		}
	case float64:
		if i, ok := v.(float64); ok {
			return int(i)
		}
	case int:
		if i, ok := v.(int); ok {
			return i
		}
	case int32:
		if i, ok := v.(int32); ok {
			return int(i)
		}
	case int64:
		if i, ok := v.(int64); ok {
			return int(i)
		}
	}
	log.Println("interface is not digit", v)
	return 0
}
