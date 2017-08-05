package cmd

import (
	"github.com/exercism/cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	viperUserConfig *viper.Viper
	viperAPIConfig  *viper.Viper
)

// configureCmd configures the command-line client with user-specific settings.
var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"c"},
	Short:   "Configure the command-line client.",
	Long: `Configure the command-line client to customize it to your needs.

This lets you set up the CLI to talk to the API on your behalf,
and tells the CLI about your setup so it puts things in the right
places.

You can also override certain default settings to suit your preferences.
`,
	Run: func(cmd *cobra.Command, args []string) {
		usrCfg := config.NewEmptyUserConfig()
		if err := usrCfg.Load(viperUserConfig); err != nil {
			Bail(err)
		}
		if err := usrCfg.Write(); err != nil {
			Bail(err)
		}

		apiCfg := config.NewEmptyAPIConfig()
		if err := apiCfg.Load(viperAPIConfig); err != nil {
			Bail(err)
		}
		if err := apiCfg.Write(); err != nil {
			Bail(err)
		}

		return
	},
}

func initConfigureCfg() {
	configureCmd.Flags().StringP("token", "t", "", "authentication token used to connect to exercism.io")
	configureCmd.Flags().StringP("workspace", "w", "", "directory for exercism exercises")
	configureCmd.Flags().StringP("api", "a", "", "API base url")

	viperUserConfig = viper.New()
	viperUserConfig.BindPFlag("token", configureCmd.Flags().Lookup("token"))
	viperUserConfig.BindPFlag("workspace", configureCmd.Flags().Lookup("workspace"))

	viperAPIConfig = viper.New()
	viperAPIConfig.BindPFlag("baseurl", configureCmd.Flags().Lookup("api"))
}

func init() {
	RootCmd.AddCommand(configureCmd)

	initConfigureCfg()
}
