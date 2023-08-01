package common

import "reflect"

const (
	RouteMethod = "Routes"
)

func RegisterRoutes(handlers ...any) {
	for _, handler := range handlers {
		refType := reflect.TypeOf(handler)
		_, ok := refType.MethodByName(RouteMethod)
		if ok {
			reflect.ValueOf(handler).MethodByName(RouteMethod).Call([]reflect.Value{})
		}
	}
}
