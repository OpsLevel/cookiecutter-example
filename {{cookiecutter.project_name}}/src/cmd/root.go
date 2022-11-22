package cmd

import (
	"strings"
	"github.com/spf13/cobra"
    {% if cookiecutter.use_viper == "y" %}"github.com/spf13/viper"{% endif %}
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.project_name }}",
	Short: "{{ cookiecutter.project_description }}",
	Long:  `{{ cookiecutter.project_description }}`,
}

func Execute(v string, c string) {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Add root flags here - IE
	rootCmd.PersistentFlags().String("log-format", "TEXT", "overrides environment variable 'OPSLEVEL_LOG_FORMAT' (options [\"JSON\", \"TEXT\"])")
	rootCmd.PersistentFlags().String("log-level", "INFO", "overrides environment variable 'OPSLEVEL_LOG_LEVEL' (options [\"ERROR\", \"WARN\", \"INFO\", \"DEBUG\"])")

	{% if cookiecutter.use_viper == "y" %}
	viper.BindPFlags(rootCmd.PersistentFlags())
	{% endif %}
	cobra.OnInitialize(setupLogging)

}

func setupLogging() {
	logFormat := strings.ToLower(viper.GetString("log-format"))
	logLevel := strings.ToLower(viper.GetString("log-level"))

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if logFormat == "text" {
		output := zerolog.ConsoleWriter{Out: os.Stderr}
		log.Logger = log.Output(output)
	}

	switch {
	case logLevel == "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case logLevel == "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case logLevel == "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case logLevel == "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}