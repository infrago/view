package view

import (
	. "github.com/infrago/base"
)

func (this *Module) Parse(body Body) (string, error) {
	if this.instance == nil {
		return "", ErrInvalidConnection
	}

	inst := this.instance

	//view层的helper
	if body.Helpers == nil {
		body.Helpers = Map{}
	}

	for key, helper := range this.helperActions {
		//默认不替换，因为http层携带context的方法，更好用一些
		if _, ok := body.Helpers[key]; !ok {
			body.Helpers[key] = helper
		}
	}

	return inst.connect.Parse(body)
}
