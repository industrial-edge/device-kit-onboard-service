/*
 * Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
 * Licensed under the MIT license
 * See LICENSE file in the top-level directory
 */

package main

import (
	"fmt"
	onboardapp "onboardservice/app"
	"os"
)

func main() {
	onboardserviceApp := onboardapp.CreateServiceApp(onboardapp.ClientFactoryImpl{})
	onboardserviceApp.StartApp()
	fmt.Println(onboardserviceApp.StartGRPC(os.Args))
}
