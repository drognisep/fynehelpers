package cmd

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Used to generate code",
	Long:  `Use the generate sub-command to create code in the current directory.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("generate called")
	//},
}

const (
	packageFlag = "pkg"
	fileFlag    = "file"
	typeFlag    = "type"
)

var (
	packageVal string
	fileVal    string
	typeVal    string
)

func init() {
	rootCmd.AddCommand(generateCmd)
	cur := currentDirectory()
	generateCmd.PersistentFlags().StringVar(&packageVal, packageFlag, cur, "Sets the generated file's package name. Defaults to the current directory's name")
	generateCmd.PersistentFlags().StringVar(&fileVal, fileFlag, "", "Path to the file to be generated")
	generateCmd.PersistentFlags().StringVar(&typeVal, typeFlag, "", "The base name of generated types")
	cobra.CheckErr(generateCmd.MarkPersistentFlagRequired(fileFlag))
	cobra.CheckErr(generateCmd.MarkPersistentFlagRequired(typeFlag))
}

func currentDirectory() (cur string) {
	cur, err := filepath.Abs(".")
	cobra.CheckErr(err)
	curSegments := strings.Split(filepath.ToSlash(cur), "/")
	if len(curSegments) == 0 {
		cobra.CheckErr(errors.New("unable to determine current directory"))
	}
	return curSegments[len(curSegments)-1]
}
