package infrastructure

import "golang-project/pkg/helper"

type Helper struct {
	ICommon helper.ICommon
}

func CreateHelper() Helper {
	helperCommon := Helper{
		ICommon: helper.MakeICommon(),
	}

	return helperCommon
}
