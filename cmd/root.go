package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "gazer",
	Long: "Gazer, an open source log collection tool.",
}

var rootTemplate = `{{if (ne .Long "")}}{{.Long | trimTrailingWhitespaces}}{{end}}
Usage:  gazer [add]

These are common Gazer commands:

operate for gazer:{{range .Commands}}{{if (not .Hidden)}}
  {{rpad .Name .NamePadding}} {{.Short}}{{end}}{{end}}

Use "gazer [command] --help" for more information about a command.
`

func Execute() {
	// 隐藏help命令
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	// 隐藏completion命令
	rootCmd.AddCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})
	// 隐藏help标志
	rootCmd.Flags().BoolP("help", "h", false, "help")
	rootCmd.Flag("help").Hidden = true
	// 重新设置usage
	rootCmd.SetHelpTemplate(rootTemplate)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
