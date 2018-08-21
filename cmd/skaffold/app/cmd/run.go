/*
Copyright 2018 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NewCmdRun describes the CLI command to run a pipeline.
func NewCmdRun(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs a pipeline file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(out)
		},
	}
	AddRunDevFlags(cmd)
	AddRunFlags(cmd)

	cmd.Flags().StringVarP(&opts.CustomTag, "tag", "t", "", "The optional custom tag to use for images which overrides the current Tagger configuration")
	return cmd
}

func run(out io.Writer) error {
	ctx := context.Background()

	runner, config, err := newRunner(opts)
	if err != nil {
		return errors.Wrap(err, "creating runner")
	}

	return runner.Run(ctx, out, config.Build.Artifacts)
}
