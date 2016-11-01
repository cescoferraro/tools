package venom

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
	"github.com/cescoferraro/tools/logger"
	"github.com/fatih/color"
	"sort"
)

type Flag struct {
	Name        string
	Short       string
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
		viper.BindEnv( i.Name)
	}
	return command
}


var VIPERLOGGER = logger.New("VIPER")
func PrintViperConfig() {
	// TODO: HANDLE NESTED YAMLS BETTER
	keys := viper.AllKeys()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	sort.Strings(keys)
	for _, key := range keys {
		VIPERLOGGER.Print(" "+ red(strings.ToUpper(key)) + " " + yellow(key) + ": " + viper.GetString(key))
	}
	return
}
