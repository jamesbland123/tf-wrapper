package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Run a terraform plan on a pipeline",
	Long: `Run a terraform plan on a specific environment
	and pipeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		executor(terraformTarget, "plan")
	},
}

func init() {
	rootCmd.AddCommand(planCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
