// Copyright 2019 The cleanup authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	cleanup := &cobra.Command{
		Use:   "cleanup",
		Short: `💫 Remove gone Git branches periodically.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cleanup.AddCommand(&cobra.Command{
		Use:   "branches <PATH>",
		Short: `Delete local branches that are gone on the remote`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	})

	if err := cleanup.Execute(); err != nil {
		log.Fatal(err)
	}
}
