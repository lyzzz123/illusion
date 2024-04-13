package proxy

import "reflect"

type Proxy interface {
	SupportInterface() reflect.Type

	SetTarget(target interface{})
}
