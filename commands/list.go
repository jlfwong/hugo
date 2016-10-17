// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/jlfwong/hugo/hugolib"
	"github.com/spf13/viper"
)

func init() {
	listCmd.AddCommand(listDraftsCmd)
	listCmd.AddCommand(listFutureCmd)
	listCmd.AddCommand(listExpiredCmd)
	listCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "filesystem path to read files relative from")
	listCmd.PersistentFlags().SetAnnotation("source", cobra.BashCompSubdirsInDir, []string{})
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing out various types of content",
	Long: `Listing out various types of content.

List requires a subcommand, e.g. ` + "`hugo list drafts`.",
	RunE: nil,
}

var listDraftsCmd = &cobra.Command{
	Use:   "drafts",
	Short: "List all drafts",
	Long:  `List all of the drafts in your content directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := InitializeConfig(); err != nil {
			return err
		}

		viper.Set("BuildDrafts", true)

		sites, err := hugolib.NewHugoSitesFromConfiguration()

		if err != nil {
			return newSystemError("Error creating sites", err)
		}

		if err := sites.Build(hugolib.BuildCfg{SkipRender: true}); err != nil {
			return newSystemError("Error Processing Source Content", err)
		}

		for _, p := range sites.Pages() {
			if p.IsDraft() {
				fmt.Println(filepath.Join(p.File.Dir(), p.File.LogicalName()))
			}

		}

		return nil

	},
}

var listFutureCmd = &cobra.Command{
	Use:   "future",
	Short: "List all posts dated in the future",
	Long: `List all of the posts in your content directory which will be
posted in the future.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := InitializeConfig(); err != nil {
			return err
		}

		viper.Set("BuildFuture", true)

		sites, err := hugolib.NewHugoSitesFromConfiguration()

		if err != nil {
			return newSystemError("Error creating sites", err)
		}

		if err := sites.Build(hugolib.BuildCfg{SkipRender: true}); err != nil {
			return newSystemError("Error Processing Source Content", err)
		}

		for _, p := range sites.Pages() {
			if p.IsFuture() {
				fmt.Println(filepath.Join(p.File.Dir(), p.File.LogicalName()))
			}

		}

		return nil

	},
}

var listExpiredCmd = &cobra.Command{
	Use:   "expired",
	Short: "List all posts already expired",
	Long: `List all of the posts in your content directory which has already
expired.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := InitializeConfig(); err != nil {
			return err
		}

		viper.Set("BuildExpired", true)

		sites, err := hugolib.NewHugoSitesFromConfiguration()

		if err != nil {
			return newSystemError("Error creating sites", err)
		}

		if err := sites.Build(hugolib.BuildCfg{SkipRender: true}); err != nil {
			return newSystemError("Error Processing Source Content", err)
		}

		for _, p := range sites.Pages() {
			if p.IsExpired() {
				fmt.Println(filepath.Join(p.File.Dir(), p.File.LogicalName()))
			}

		}

		return nil

	},
}
