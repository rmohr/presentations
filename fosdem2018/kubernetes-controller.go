// LISTWATCHER BEGIN OMIT
type ListerWatcher interface { // HL
    // List should return a list type object; the Items field will be extracted, and the
    // ResourceVersion field will be used to start the watch in the right place.
    List(options v1.ListOptions) (runtime.Object, error) // HL
    // Watch should begin a watch at the specified version.
    Watch(options v1.ListOptions) (watch.Interface, error) // HL
} // HL
// LISTWATCHER END OMIT

// STORE BEGIN OMIT
type Store interface { // HL
    // Add a new runtime object
    Add(obj interface{}) error // HL
    // Update a new runtime object
    Update(obj interface{}) error // HL
    // Delete a runtime object
    Delete(obj interface{}) error // HL
    // List the store content
    List() []interface{} // HL
    // Get keys
    ListKeys() []string // HL
    // Check if an object exists in a store and retrieve it
    Get(obj interface{}) (item interface{}, exists bool, err error) // HL
    // Check if an object exists in a store and retrieve it via a key
    GetByKey(key string) (item interface{}, exists bool, err error) // HL
} // HL

// STORE END OMIT
// LISTWATCHERIMPL BEGIN OMIT
func NewListWatchFromClient(restcli Getter, resource string, namespace string, fieldSelector fields.Selector) *ListWatch {
	listFunc := func(options v1.ListOptions) (runtime.Object, error) {
		return restcli.Get().
			Namespace(namespace).
			Resource(resource).
			VersionedParams(&options, api.ParameterCodec).
			FieldsSelectorParam(fieldSelector).
			Do().
			Get()
	}
	watchFunc := func(options v1.ListOptions) (watch.Interface, error) {
		return restcli.Get().
			Prefix("watch").
			Namespace(namespace).
			Resource(resource).
			VersionedParams(&options, api.ParameterCodec).
			FieldsSelectorParam(fieldSelector).
			Watch()
	}
	return &ListWatch{ListFunc: listFunc, WatchFunc: watchFunc}
}
// LISTWATCHERIMPL END OMIT


// INFORMER BEGIN OMIT
callbacks := cache.ResourceEventHandlerFuncs {
		AddFunc: func(obj interface{}) {
				fmt.Println(obj)
		},
		UpdateFunc: func(old interface{}, new interface{}) {
				fmt.Println(obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println(obj)
		},
	}
}
// INFORMER END OMIT
// INFORMERCALLBACKS BEGIN OMIT
store, informer := cache.NewIndexerInformer(lw, objType, 0, callbacks, cache.Indexers{})
cache.WaitForCacheSync(stopCh, informer.HasSynced)
// INFORMERCALLBACKS END OMIT
// SHAREDINFORMER BEGIN OMIT
callbacks := cache.ResourceEventHandlerFuncs {
		AddFunc: func(obj interface{}) {
				fmt.Println(obj)
		},
		UpdateFunc: func(old interface{}, new interface{}) {
				fmt.Println(obj)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println(obj)
		},
	}
}
cache, informer := cache.NewSharedIndexInformer(lw, objType, 0, cache.Indexers{})
informer.AddEventHandler(callbacks)
// SHAREDINFORMER END OMIT

// WORKQUEUE BEGIN OMIT
type Interface interface {
	Add(item interface{})
	Len() int
	Get() (item interface{}, shutdown bool)
	Done(item interface{})
	AddRateLimited(item interface{})
	Forget(item interface{})
	NumRequeues(item interface{}) int
}
// WORKQUEUE END OMIT
