package venom

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

type Flag struct {
	Name        string
	Short       string
	Description string
	Value       interface{}
}

type CommandFlag []Flag

func (flags CommandFlag) Register(cmd *cobra.Command) *cobra.Command {
	for _, i := range flags {
		viper.Set(i.Name, i.Value)
		switch i.Value.(type) {
		case int:
			cmd.Flags().IntP(i.Name, i.Short, i.Value.(int), i.Description)
		case bool:
			cmd.Flags().BoolP(i.Name, i.Short, i.Value.(bool), i.Description)
		case string:
			cmd.Flags().StringP(i.Name, i.Short, i.Value.(string), i.Description)
		}
		viper.BindPFlag(strings.ToUpper(i.Name), cmd.Flags().Lookup(i.Name))
		viper.BindEnv(strings.ToUpper(i.Name))
	}
	return cmd
}
