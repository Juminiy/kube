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
		name += "create" // Create
	case 'Q': // query
		name += "query" // Query
	case 'U': // update
		name += "update" // Update
	case 'D': // delete
		name += "delete" // Delete
	case 'E': // raw
		name += "raw" // Exec
	case 'R': // row
		name += "row" // Row
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
