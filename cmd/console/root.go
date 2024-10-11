/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package console

import (
	"fmt"
	"os"

	"github.com/CESSProject/cess-miner/configs"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   configs.Name,
	Short: configs.Description,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("\x1b[%dm[err]\x1b[0m %v\n", 41, err)
		os.Exit(1)
	}
}

// init
func init() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "custom configuration file")
	rootCmd.PersistentFlags().StringSliceP("rpcs", "", nil, "rpc endpoint list")
	rootCmd.PersistentFlags().StringP("workspace", "", "", "workspace")
	rootCmd.PersistentFlags().StringP("staking", "", "", "staking account")
	rootCmd.PersistentFlags().StringP("earnings", "", "", "earnings account")
	rootCmd.PersistentFlags().Uint16P("port", "", 0, "listening port")
	rootCmd.PersistentFlags().IntP("cpu", "", 0, "number of cpus used, 0 means use all")
	rootCmd.PersistentFlags().Uint64P("space", "", 0, "maximum space used (TiB)")
	rootCmd.PersistentFlags().StringP("mnemonic", "", "", "signature account mnemonic")
	rootCmd.PersistentFlags().StringSliceP("tees", "", nil, "priority use of tee endpoint list")
	rootCmd.PersistentFlags().StringP("endpoint", "", "", "endpoint of miner communication")
}
