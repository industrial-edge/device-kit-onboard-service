/*
 * Copyright (c) Siemens 2021
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package system

import "github.com/stretchr/testify/mock"

type MockHelper struct {
	mock.Mock
}

func (m *MockHelper) DaemonReload(query string) bool {
	return true
}
