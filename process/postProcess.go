// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package process

import (
	"os"
	"path/filepath"

	"../config"
)

// PostProcess moves the file(s) from working folder to the output folder
func PostProcess(config config.Config) {
	os.Rename(filepath.Join(config.Folders.Wrkg, config.UUID), filepath.Join(config.Folders.Out, config.UUID))
}
