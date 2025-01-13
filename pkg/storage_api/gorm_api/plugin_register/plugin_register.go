package plugin_register

import "errors"

func CallbackName(pluginName string, before bool, do byte) string {
	name := pluginName + ":"
	if before {
		name += "before_"
	} else {
		name += "after_"
	}
	switch do {
	case 'C': // create
		name += "create"
	case 'Q': // query
		name += "query"
	case 'U': // update
		name += "update"
	case 'D': // delete
		name += "delete"
	case 'E': // raw
		name += "raw"
	case 'R': // row
		name += "row"
	default:
		panic(do)
	}
	return name
}

func OneError(err ...error) error {
	for _, errElem := range err {
		if errElem != nil {
			return errElem
		}
	}
	return nil
}

var NoPluginName = errors.New("external gorm plugin registered no name")
