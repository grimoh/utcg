package utcg

import (
	"fmt"
	"strconv"
	"time"
	"errors"

	"github.com/spf13/cobra"
)

var jstSubCmd = &cobra.Command{
	Use: "jst",
	Short: "Convert unixtime to jst.",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected 1 arg.")
		}
		return convert(args[0], c)
	},
}

func init() {
	rootCmd.AddCommand(jstSubCmd)
}

func convert(ut string, c *cobra.Command) error {
	i, err := strconv.ParseInt(ut, 10, 64)
	if err != nil {
		return errors.New("invalid syntax")
	}
	nano, _ := c.Flags().GetString("nano")
	if nano != "" {
		fmt.Println(time.Unix(i/1000000000, i%1000000000))
	} else {
		fmt.Println(time.Unix(i, 0))
	}
	return nil
}
