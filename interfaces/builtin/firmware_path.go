// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

import (
	"fmt"
	"strings"
)

const firmwarePathSummary = `allows access to system files or directories`

const firmwarePathBaseDeclarationPlugs = `
  firmware-path-file:
    allow-installation: false
    deny-auto-connection: true
`

const firmwarePathConnectedPlugAppArmor = `
# Description: Can access specific "/sys/module/firmware_class/parameters/path" to control the firmware lookup 
#  directory.
`

type firmwarePathInterface struct {
	commonFilesInterface
}

func validateSinglePathFirmwarePath(np string) error {
	if !(strings.Compare(np, "/sys/module/firmware_class/parameters/path") == 0) {
		return fmt.Errorf(`%q must be "/sys/module/firmware_class/parameters/path"`, np)
	}

	return nil
}

func init() {
	registerIface(&firmwarePathInterface{
		commonFilesInterface{
			commonInterface: commonInterface{
				name:                 "firmware-path-files",
				summary:              firmwarePathSummary,
				implicitOnCore:       true,
				implicitOnClassic:    true,
				baseDeclarationPlugs: firmwarePathBaseDeclarationPlugs,
			},
			apparmorHeader:    firmwarePathConnectedPlugAppArmor,
			extraPathValidate: validateSinglePathFirmwarePath,
		},
	})
}
