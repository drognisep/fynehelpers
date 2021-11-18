package cmd

import (
	"bytes"
	"log"
	"os"
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

	treeCmd.Flags().BoolVar(&eventTappedVal, eventTappedFlag, false, "Indicates that the tree node should be tappable")
	treeCmd.Flags().BoolVar(&eventDoubleTappedVal, eventDoubleTappedFlag, false, "Indicates that the tree node should be double-tappable")
	treeCmd.Flags().BoolVar(&eventSecondTappedVal, eventSecondTappedFlag, false, "Indicates that the tree node should be secondary-tappable")
}

func generateTree(*cobra.Command, []string) {
	normalization.TrimSpace(&packageVal, &fileVal, &typeVal)
	cobra.CheckErr(validation.NoneBlank(map[string]string{
		packageFlag: packageVal,
		fileFlag:    fileVal,
		typeFlag:    typeVal,
	}))
	fileVal = strings.TrimSuffix(fileVal, ".go")
	var nodeBuf bytes.Buffer
	var treeBuf bytes.Buffer
	params := &TreeGenParams{
		Package:         packageVal,
		File:            fileVal,
		TypeBase:        typeVal,
		GenTapped:       eventTappedVal,
		GenDoubleTapped: eventDoubleTappedVal,
		GenSecondTapped: eventSecondTappedVal,
	}
	cobra.CheckErr(treeNodeTemplate.Execute(&nodeBuf, params))
	nodeBuf.WriteString("\n\n")
	cobra.CheckErr(treeTemplate.Execute(&treeBuf, params))
	treeBuf.WriteString("\n\n")
	if err := os.WriteFile(fileVal + "Node.go", nodeBuf.Bytes(), 0766); err != nil {
		log.Printf("Error writing to '%s': %v\n", fileVal, err)
		return
	}
	if err := os.WriteFile(fileVal + "Tree.go", treeBuf.Bytes(), 0766); err != nil {
		log.Printf("Error writing to '%s': %v\n", fileVal, err)
		return
	}
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

var (
	treeNodeTemplate = template.Must(template.New("generateTreeNode").Parse(treeNodeText))
	treeTemplate = template.Must(template.New("generateTree").Parse(treeText))
)
