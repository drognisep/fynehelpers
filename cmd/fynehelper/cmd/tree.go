package cmd

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/drognisep/fynehelpers/cmd/fynehelper/normalization"
	"github.com/drognisep/fynehelpers/cmd/fynehelper/validation"
	"github.com/spf13/cobra"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Generates tree boilerplate",
	Long: `Generates a tree based on a user model adapted to fit a Fyne tree.
The user model may be any type that implements generation.TreeModel.`,
	Run: generateTree,
}

const (
	eventTappedFlag       = "event-tapped"
	eventDoubleTappedFlag = "event-double-tapped"
	eventSecondTappedFlag = "event-secondary-tapped"
)

var (
	eventTappedVal       bool
	eventDoubleTappedVal bool
	eventSecondTappedVal bool
)

func init() {
	generateCmd.AddCommand(treeCmd)

	treeCmd.Flags().Bool(eventTappedFlag, false, "Indicates that the tree node should be tappable")
	treeCmd.Flags().Bool(eventDoubleTappedFlag, false, "Indicates that the tree node should be double-tappable")
	treeCmd.Flags().Bool(eventSecondTappedFlag, false, "Indicates that the tree node should be secondary-tappable")
}

func generateTree(cmd *cobra.Command, args []string) {
	normalization.TrimSpace(&packageVal, &fileVal, &typeVal)
	cobra.CheckErr(validation.NoneBlank(map[string]string{
		packageFlag: packageVal,
		fileFlag:    fileVal,
		typeFlag:    typeVal,
	}))
	var buf bytes.Buffer
	cobra.CheckErr(treeTemplate.Execute(&buf, TreeGenParams{
		Package:         packageVal,
		File:            fileVal,
		TypeBase:        typeVal,
		GenTapped:       eventTappedVal,
		GenDoubleTapped: eventDoubleTappedVal,
		GenSecondTapped: eventSecondTappedVal,
	}))

}

type TreeGenParams struct {
	Package         string
	File            string
	TypeBase        string
	GenTapped       bool
	GenDoubleTapped bool
	GenSecondTapped bool
}

func (p *TreeGenParams) TypeBaseTitle() string {
	return strings.Title(p.TypeBase)
}

func (p *TreeGenParams) TypeBaseHidden() string {
	splits := strings.SplitN(p.TypeBase, "", 2)
	switch len(splits) {
	case 2:
		return strings.ToLower(splits[0]) + splits[1]
	case 1:
		return strings.ToLower(splits[0])
	default:
		panic("Unexpectedly blank type base")
	}
}

const (
	treeText = `
package {{ .Package }}



`
	treeImplText = `
package {{ .Package }}
`
)

var (
	treeTemplate = template.Must(template.New("generateTree").Parse(treeText))
)
