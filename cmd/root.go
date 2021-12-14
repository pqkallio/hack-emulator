package cmd

import (
	"github.com/pqkallio/hack-emulator/hack"
	"github.com/pqkallio/hack-emulator/util"
	"github.com/spf13/cobra"
)

var (
	romFile string

	rootCmd = &cobra.Command{
		Use:   "run",
		Short: "Hack system emulator",
		Long:  "Run Hack machine language ROM files",
		Run: func(cmd *cobra.Command, args []string) {
			if romFile == "" {
				cmd.Help()
				return
			}

			rom, err := util.ReadRomFile(romFile)
			if err != nil {
				panic(err)
			}

			hack.Run(rom)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&romFile, "rom", "r", "", "ROM file to load and run")
}
