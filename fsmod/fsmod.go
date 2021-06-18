// This file is part of fsmod.
//
// Copyright (C) 2021  David Gamba Rios
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package fsmod - Provides ways to query files to validate if they have been modified or if they differ from what it is expected.
package fsmod

import (
	"io/fs"
)

func Glob(fsys fs.FS, pattern ...string) (filepaths []string, err error) {
	for _, p := range pattern {
		matches, err := fs.Glob(fsys, p)
		if err != nil {
			return filepaths, err
		}
		filepaths = append(filepaths, matches...)
	}
	return
}
