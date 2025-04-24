#!/bin/bash
# Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
# Licensed under the MIT license
# See LICENSE file in the top-level directory

#
# Code coverage generation
set -ex
COVERAGE_DIR="${COVERAGE_DIR:-coverage}"
PKG_LIST=$(go list ../../... | grep -v /cmd/ | grep -v /api/)

# Create the coverage files directory
mkdir -p "$COVERAGE_DIR";

# Create a coverage file for each package
for package in ${PKG_LIST}; do
    echo "Generating global code coverage report for $package"
    go test -gcflags=all=-l -covermode=count -coverprofile "${COVERAGE_DIR}/${package##*/}.cov" "$package" ;
done ;

# Merge the coverage profile files
tail -q -n +2 "${COVERAGE_DIR}"/*.cov >> "${COVERAGE_DIR}"/coverage_all.cov ;
echo 'mode: count' > "${COVERAGE_DIR}"/coverage.cov ;
cat "${COVERAGE_DIR}"/coverage_all.cov | grep -v "mocks" >> "${COVERAGE_DIR}"/coverage.cov;

# Display the global code coverage
go tool cover -func="${COVERAGE_DIR}"/coverage.cov | tee coverage_global.txt;

go tool cover -html="${COVERAGE_DIR}"/coverage.cov -o coverage.html ;
 

# Remove the coverage files directory
cp "$COVERAGE_DIR"/coverage.cov .;
rm -rf "$COVERAGE_DIR";
