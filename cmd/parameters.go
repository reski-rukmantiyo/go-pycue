package cmd

import (
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func getParameters(cmd *cobra.Command) map[string]string {
	flags := cmd.Flags()
	arrayFlags := make(map[string]string)

	flagNames := make([]string, 0, flags.NFlag())
	flags.Visit(func(flag *pflag.Flag) {
		flagNames = append(flagNames, flag.Name)
	})
	sort.Strings(flagNames)
	for _, name := range flagNames {
		value := flags.Lookup(name).Value.String()
		arrayFlags[name] = value
	}
	return arrayFlags
}
