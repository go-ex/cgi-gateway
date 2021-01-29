package broadcast

import "gateway/constants"

type broadcast struct {
}

func NewController() {
	controller := &broadcast{}

	constants.Router.POST("/broadcast", controller.index)
}
