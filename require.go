// Copyright 2020 The Godror Authors
//
// SPDX-License-Identifier: UPL-1.0 OR Apache-2.0

//go:build require
// +build require

package godror

import (
	_ "github.com/devMake-a11y/godror/odpi/embed"   // ODPI-C
	_ "github.com/devMake-a11y/godror/odpi/include" // ODPI-C
	_ "github.com/devMake-a11y/godror/odpi/src"     // ODPI-C
)
