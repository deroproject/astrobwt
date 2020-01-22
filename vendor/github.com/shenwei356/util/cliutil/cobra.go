package cliutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shenwei356/util/stringutil"
	"github.com/spf13/cobra"
)

func isStdin(file string) bool {
	return file == "-"
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func GetFileList(args []string) []string {
	files := []string{}
	if len(args) == 0 {
		files = append(files, "-")
	} else {
		for _, file := range files {
			if isStdin(file) {
				continue
			}
			if _, err := os.Stat(file); os.IsNotExist(err) {
				CheckError(err)
			}
		}
		files = args
	}
	return files
}

func GetFlagInt(cmd *cobra.Command, flag string) int {
	value, err := cmd.Flags().GetInt(flag)
	CheckError(err)
	return value
}

func GetFlagPositiveInt(cmd *cobra.Command, flag string) int {
	value, err := cmd.Flags().GetInt(flag)
	CheckError(err)
	if value <= 0 {
		CheckError(fmt.Errorf("value of flag --%s should be greater than 0", flag))
	}
	return value
}

func GetFlagPositiveFloat64(cmd *cobra.Command, flag string) float64 {
	value, err := cmd.Flags().GetFloat64(flag)
	CheckError(err)
	if value <= 0 {
		CheckError(fmt.Errorf("value of flag --%s should be greater than 0", flag))
	}
	return value
}

func GetFlagNonNegativeInt(cmd *cobra.Command, flag string) int {
	value, err := cmd.Flags().GetInt(flag)
	CheckError(err)
	if value < 0 {
		CheckError(fmt.Errorf("value of flag --%s should be greater than or equal to 0", flag))
	}
	return value
}

func GetFlagNonNegativeFloat64(cmd *cobra.Command, flag string) float64 {
	value, err := cmd.Flags().GetFloat64(flag)
	CheckError(err)
	if value < 0 {
		CheckError(fmt.Errorf("value of flag --%s should be greater than or equal to ", flag))
	}
	return value
}

func GetFlagBool(cmd *cobra.Command, flag string) bool {
	value, err := cmd.Flags().GetBool(flag)
	CheckError(err)
	return value
}

func GetFlagString(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)
	CheckError(err)
	return value
}

func GetFlagNonEmptyString(cmd *cobra.Command, flag string) string {
	value, err := cmd.Flags().GetString(flag)
	CheckError(err)
	if value == "" {
		CheckError(fmt.Errorf("value of flag --%s should not be empty", flag))

	}
	return value
}

func GetFlagCommaSeparatedStrings(cmd *cobra.Command, flag string) []string {
	value, err := cmd.Flags().GetString(flag)
	CheckError(err)
	return stringutil.Split(value, ",")
}

func GetFlagSemicolonSeparatedStrings(cmd *cobra.Command, flag string) []string {
	value, err := cmd.Flags().GetString(flag)
	CheckError(err)
	return stringutil.Split(value, ";")
}

func GetFlagCommaSeparatedInts(cmd *cobra.Command, flag string) []int {
	filedsStrList := GetFlagCommaSeparatedStrings(cmd, flag)
	fields := make([]int, len(filedsStrList))
	for i, value := range filedsStrList {
		v, err := strconv.Atoi(value)
		if err != nil {
			CheckError(fmt.Errorf("value of flag --%s should be comma separated integers", flag))
		}
		fields[i] = v
	}
	return fields
}

func GetFlagRune(cmd *cobra.Command, flag string) rune {
	value, err := cmd.Flags().GetString(flag)
	CheckError(err)
	if len(value) > 1 {
		CheckError(fmt.Errorf("value of flag --%s should has length of 1", flag))
	}
	var v rune
	for _, r := range value {
		v = r
		break
	}
	return v
}

func GetFlagFloat64(cmd *cobra.Command, flag string) float64 {
	value, err := cmd.Flags().GetFloat64(flag)
	CheckError(err)
	return value
}

func GetFlagInt64(cmd *cobra.Command, flag string) int64 {
	value, err := cmd.Flags().GetInt64(flag)
	CheckError(err)
	return value
}

func GetFlagStringSlice(cmd *cobra.Command, flag string) []string {
	value, err := cmd.Flags().GetStringSlice(flag)
	CheckError(err)
	return value
}
