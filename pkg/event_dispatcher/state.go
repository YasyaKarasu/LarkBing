package eventDispatcher

// set of event ids
var eventIdList = make(map[string]bool)

var eventMap = make(map[string]CallbackType)
