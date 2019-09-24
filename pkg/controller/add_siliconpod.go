package controller

import (
	"github.com/aneeshkp/silicon-pod-operator/pkg/controller/siliconpod"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, siliconpod.Add)
}
