// This file is part of fsmod.
//
// Copyright (C) 2021  David Gamba Rios
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package fsmod

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGlob(t *testing.T) {
	mfs := make(fstest.MapFS)
	mfs["a"] = &fstest.MapFile{Mode: 0777 | fs.ModeDir}
	mfs["a/file.txt"] = &fstest.MapFile{}
	mfs["a/file_test.txt"] = &fstest.MapFile{}
	mfs["a/file.go"] = &fstest.MapFile{}
	mfs["a/file_test.go"] = &fstest.MapFile{}
	mfs["b"] = &fstest.MapFile{Mode: 0777 | fs.ModeDir}
	mfs["b/file.txt"] = &fstest.MapFile{}
	mfs["b/file_test.txt"] = &fstest.MapFile{}
	mfs["b/file.go"] = &fstest.MapFile{}
	mfs["b/file_test.go"] = &fstest.MapFile{}

	tests := []struct {
		name     string
		glob     string
		expFiles []string
		expErr   error
	}{
		{"all", "*", []string{"a/file.txt"}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filepaths, err := Glob(mfs, test.glob)
			if err != test.expErr {
				t.Errorf("Unexpected error. Exp: %v, Got: %v\n", test.expErr, err)
			}
			if !reflect.DeepEqual(filepaths, test.expFiles) {
				t.Errorf("exp: %v\ngot: %v\n", test.expFiles, filepaths)
			}
		})
	}
}
