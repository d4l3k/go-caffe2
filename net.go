package caffe2

// #include "caffe2.h"
import "C"

type Net struct {
	n C.C2Net

	// we hold on to the workspace pointer to ensure that Workspace doesn't get
	// freed before Net.
	w *Workspace
}
