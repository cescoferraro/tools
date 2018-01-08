package venom

import (
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Flag struct {
	Name        string
	Short       string
	Safe        bool
	Description string
	Value       interface{}
}

type CommandFlag []Flag

func (flags CommandFlag) Register(command *cobra.Command) *cobra.Command {
	for _, i := range flags {

		switch i.Value.(type) {
		case int:
			command.Flags().IntP(i.Name, i.Short, i.Value.(int), i.Description)
		case bool:
			command.Flags().BoolP(i.Name, i.Short, i.Value.(bool), i.Description)
		case string:
			command.Flags().StringP(i.Name, i.Short, i.Value.(string), i.Description)
		}
		viper.BindPFlag(i.Name, command.Flags().Lookup(i.Name))
		viper.BindEnv(i.Name)
	}
	return command
}

// PrintViperConfig TODO: NEEDS COMMENT INFO
func PrintViperConfig(flags CommandFlag) {
	// TODO: HANDLE NESTED YAMLS BETTER
	keys := viper.AllKeys()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	sort.Strings(keys)
	for _, key := range keys {
		if flagByName(flags, key).Safe {
			VIPERLOGGER.Print(
				" " + red(strings.ToUpper(key)) +
					" " + yellow(key) +
					": " + viper.GetString(key)[0:len(viper.GetString(key))*4/10] + "...")
		} else {
			VIPERLOGGER.Print(
				" " + red(strings.ToUpper(key)) +
					" " + yellow(key) +
					": " + viper.GetString(key))
		}
	}
	return
}
