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

const deviceTreeOverlaySummary = `allows access to system files or directories`

const deviceTreeOverlayBaseDeclarationPlugs = `
  device-tree-overlay-files:
    allow-installation: false
    deny-auto-connection: true
`

const deviceTreeOverlayConnectedPlugAppArmor = `
# Description: Can access specific "/configfs/device-tree/overlays/" files or directories.
# This is restricted because otherwise it would give arbitrary access to hardware parameters.
`

type deviceTreeOverlayInterface struct {
	commonFilesInterface
}

func validateSinglePathdeviceTreeOverlay(np string) error {
	if !strings.HasPrefix(np, "/configfs/device-tree/overlays/") {
		return fmt.Errorf(`%q must start with "/configfs/device-tree/overlays/"`, np)
	}

	return nil
}

func init() {
	registerIface(&deviceTreeOverlayInterface{
		commonFilesInterface{
			commonInterface: commonInterface{
				name:                 "device-tree-overlay-files",
				summary:              deviceTreeOverlaySummary,
				implicitOnCore:       true,
				implicitOnClassic:    true,
				baseDeclarationPlugs: deviceTreeOverlayBaseDeclarationPlugs,
			},
			apparmorHeader:    deviceTreeOverlayConnectedPlugAppArmor,
			extraPathValidate: validateSinglePathdeviceTreeOverlay,
		},
	})
}
