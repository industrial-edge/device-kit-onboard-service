/*
 * Copyright © Siemens 2025 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockFileSystem is a mock implementation of the FileSystem interface.
type MockFileSystem struct {
	mock.Mock
}

func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileSystem) Stat(path string) error {
	args := m.Called(path)
	return args.Error(0)
}
