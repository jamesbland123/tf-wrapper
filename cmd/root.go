package cmd

import (
	"fmt"
	"os"
	"strings"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var environment string
var executionOrder []string
var containerImage string
var terraformTarget string
var backendConf string
var environmentConf string
var awsProfile string
var aws_region string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tf-wrapper",
	Short: "A wrapper application to simplify running terraform",
	Long: `A wrapper application to simplify running terraform 
	on different environments and configurations`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "", "The target environment ie prod/staging/dev (required)")
	rootCmd.MarkPersistentFlagRequired("environment")

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is tf-wrapper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	

	rootCmd.PersistentFlags().StringVarP(&containerImage, "image", "i", "", "The container image to be deployed (required)")
	rootCmd.MarkPersistentFlagRequired("image")

	rootCmd.PersistentFlags().StringVarP(&terraformTarget, "target", "t", "all", "The terraform target configuration/directory such as network or bastion")

	//rootCmd.PersistentFlags().StringSliceVar(&executionOrder, "executionOrder", []string{}, "Terraform execution ordering sequence")
	//executionOrder = viper.GetStringSlice("executionOrder")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		/* home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} */

		// Search config in home directory with name ".tf-wrapper" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("tf-wrapper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Config vars.  Root level
	executionOrder = viper.GetStringSlice("executionOrder")
	aws_region = strings.Trim(viper.GetString("aws_region"), "-") 

	// Config vars. Nested based on environment flag
	backendConf = viper.GetString(fmt.Sprintf("%s.backendConf", environment))
	environmentConf = viper.GetString(fmt.Sprintf("%s.environmentConf", environment))
	awsProfile = viper.GetString(fmt.Sprintf("%s.awsProfile", environment))
	
}
