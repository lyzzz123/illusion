package listener

import "sort"

var ServiceListenerArray = make([]Listener, 0)

func RegisterServiceListener(listener Listener) {
	ServiceListenerArray = append(ServiceListenerArray, listener)
	sort.SliceStable(ServiceListenerArray, func(i, j int) bool {
		return ServiceListenerArray[i].GetPriority() > ServiceListenerArray[j].GetPriority()
	})
}

func ExecuteHttpServerStartUpListener() error {

	if len(ServiceListenerArray) <= 0 {
		return nil
	}

	for i := 0; i < len(ServiceListenerArray); i++ {
		err := ServiceListenerArray[i].ServiceStartup()
		if err != nil {
			return err
		}
	}
	return nil
}

func ExecuteHttpServerShutdownListener() error {
	if len(ServiceListenerArray) <= 0 {
		return nil
	}

	for i := 0; i < len(ServiceListenerArray); i++ {
		err := ServiceListenerArray[i].ServiceShutdown()
		if err != nil {
			return err
		}
	}
	return nil
}

type Listener interface {
	ServiceStartup() error

	ServiceShutdown() error

	GetPriority() int
}
