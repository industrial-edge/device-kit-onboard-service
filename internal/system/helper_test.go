/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package system

import (
	"testing"
)

func TestHelperMethods(t *testing.T) {

	var sut = Helper{}

	if sut.DaemonReload("true") != true {
		t.FailNow()
	}
	if sut.DaemonReload("false") != false {
		t.FailNow()
	}
}
