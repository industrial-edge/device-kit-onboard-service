/*
 * Copyright © Siemens 2025 - 2026. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package utils

import (
	"os"
)

// FileSystemOperations interface defines file system operations
type FileSystemOperations interface {

	// ReadFile reads the content of the file at the given path
	ReadFile(path string) ([]byte, error)

	// Stat checks the status of the file at the given path
	Stat(path string) error
}

// OsFileSystemOperations FileSystemOperations implementation, wrapper for functionalities from `os` package
type OsFileSystemOperations struct{}

func (f *OsFileSystemOperations) ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (f *OsFileSystemOperations) Stat(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}
