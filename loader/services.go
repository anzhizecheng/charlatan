// Copyright (C) 2019 Guillaume CAMUS
//
// This file is part of Charlatan.
//
// Charlatan is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Charlatan is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Charlatan.  If not, see <http://www.gnu.org/licenses/>.

package loader

import (
	"fmt"

	"github.com/sarulabs/di"
	"github.com/spf13/afero"
)

var Services []di.Def = []di.Def{
	{
		Name:  "app.locator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			if fs, ok := ctn.Get("app.fs").(afero.Fs); ok {
				return NewFixtureLocator(fs), nil
			}
			return nil, fmt.Errorf("Could not cast 'app.fs' in afero.Fs")
		},
	},
	{
		Name:  "app.fileloader",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			if fs, ok := ctn.Get("app.fs").(afero.Fs); ok {
				return NewFileLoader(fs), nil
			}
			return nil, fmt.Errorf("Could not cast 'app.fs' in afero.Fs")
		},
	},
}