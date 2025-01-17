// Copyright © 2021 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/labring/sealos/pkg/utils/logger"

	"github.com/labring/sealos/pkg/apply"
	"github.com/labring/sealos/pkg/apply/processor"
)

var exampleReset = `
reset you current cluster:
	sealos reset --name xxx [--force]
`

func newResetCmd() *cobra.Command {
	resetArgs := &apply.ResetArgs{
		Cluster: &apply.Cluster{},
		SSH:     &apply.SSH{},
	}

	var resetCmd = &cobra.Command{
		Use:     "reset",
		Short:   "simplest way to reset your cluster",
		Long:    `sealos reset --name [arg]`,
		Example: exampleReset,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := processor.ConfirmDeleteNodes(); err != nil {
				return err
			}
			applier, err := apply.NewApplierFromResetArgs(resetArgs)
			if err != nil {
				return err
			}
			return applier.Delete()
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			logger.Info(contact)
		},
	}
	resetArgs.RegisterFlags(resetCmd.Flags())
	resetCmd.Flags().BoolVar(&processor.ForceDelete, "force", false, "we also can input an --force flag to reset cluster by force")
	return resetCmd
}

func init() {
	rootCmd.AddCommand(newResetCmd())
}
