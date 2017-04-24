/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// set-gen is an example usage of go2idl.
//
// Structs in the input directories with the below line in their comments will
// have sets generated for them.
// // +genset
//
// Any builtin type referenced anywhere in the input directories will have a
// set generated for it.
package main

import (
	"os"
	"path/filepath"

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/set-gen/generators"

	"github.com/golang/glog"
)

func main() {
	arguments := args.Default()

	// Override defaults. These are Kubernetes specific input and output
	// locations.
	arguments.InputDirs = []string{"github.com/sourcegraph/monorepo-test-1/kubernetes-3/pkg/util/sets/types"}
	arguments.OutputPackagePath = "k8s.io/apimachinery/pkg/util/sets"
	arguments.GoHeaderFilePath = filepath.Join(args.DefaultSourceTree(), "github.com/sourcegraph/monorepo-test-1/kubernetes-3/hack/boilerplate/boilerplate.go.txt")

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		generators.Packages,
	); err != nil {
		glog.Errorf("Error: %v", err)
		os.Exit(1)
	}
	glog.V(2).Info("Completed successfully.")
}