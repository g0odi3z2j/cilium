// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var cmdRefDir string

var cmdRef = &cobra.Command{
	Use:   "cmdref",
	Short: "Generate Cilium command reference",
	Run: func(cmd *cobra.Command, args []string) {
		genCmdRef()
	},
	Hidden: true,
}

func genCmdRef() {
	// Remove the line 'Auto generated by spf13/cobra on ...'
	rootCmd.DisableAutoGenTag = true
	if err := doc.GenMarkdownTreeCustom(rootCmd, cmdRefDir, filePrepend, linkHandler); err != nil {
		log.Fatal(err)
	}
}

func linkHandler(s string) string {
	// The generated files have a 'See also' section but the URL's are
	// hardcoded to use Markdown but we only want / have them in HTML
	// later.
	return "../" + strings.Replace(s, ".md", "", 1)
}

func filePrepend(s string) string {
	// Prepend a HTML comment that this file is autogenerated. So that
	// users are warned before fixing issues in the Markdown files.  Should
	// never show up on the web.
	return fmt.Sprintf("%s\n\n", "<!-- This file was autogenerated via cilium cmdref, do not edit manually-->")
}

func init() {
	cmdRef.Flags().StringVarP(&cmdRefDir, "directory", "d", "./", "Path to the output directory")
	rootCmd.AddCommand(cmdRef)
}
