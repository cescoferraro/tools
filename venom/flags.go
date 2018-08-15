package venom

import (
	"sort"
	"strings"

	"github.com/cescoferraro/tools/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flag TODO: NEEDS COMMENT INFO
type Flag struct {
	Name        string
	Short       string
	Safe        bool
	Description string
	Value       interface{}
}

// CommandFlag TODO: NEEDS COMMENT INFO
type CommandFlag []Flag

// Register TODO: NEEDS COMMENT INFO
func (flags CommandFlag) Register(command *cobra.Command) *cobra.Command {
	for _, i := range flags {

		switch i.Value.(type) {
		case int:
			command.PersistentFlags().IntP(i.Name, i.Short, i.Value.(int), i.Description)
		case bool:
			command.PersistentFlags().BoolP(i.Name, i.Short, i.Value.(bool), i.Description)
		case string:
			command.PersistentFlags().StringP(i.Name, i.Short, i.Value.(string), i.Description)
		}
		viper.BindPFlag(i.Name, command.PersistentFlags().Lookup(i.Name))
		viper.BindEnv(i.Name)
	}
	return command
}

// VIPERLOGGER TODO: NEEDS COMMENT INFO
var VIPERLOGGER = logger.New("VIPER")

func flagByName(RunServerFlags CommandFlag, name string) Flag {
	for _, flag := range RunServerFlags {
		if flag.Name == name {
			return flag
		}
	}
	return Flag{}
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
