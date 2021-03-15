package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)
func addEven(args []string){
	var sum int
	for _, strVal := range args{
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Println(err)
		}
		if intVal % 2 == 0 {
			sum += intVal
		}
	}
	fmt.Printf("The even addition of numbers %s is %d", args, sum)
}
// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addEven(args)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)
}
