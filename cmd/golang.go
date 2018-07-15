// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/tjololo/dunix/pkg/golang"
	"os"
	"github.com/tjololo/dunix/pkg/fileutils"
	"path/filepath"
)

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "Install Golang",
	Long: `Download and install golang`,
	Run: installGolang,
}

func init() {
	installCmd.AddCommand(golangCmd)
	golangCmd.Flags().StringVarP(&installPath, "install-path", "p", "/opt", "What dir to install Golang in.")
	golangCmd.Flags().BoolVarP(&createSymlink, "create-symlink", "s", true, "Create symlink in install folder.")
	golangCmd.Flags().StringVar(&symlinkName, "symlink-name", "golang", "Name of the symlink created with -s.")
	golangCmd.Flags().StringVarP(&version, "version", "v", golang.DefaultVersion, "What version of Golang to install.")
}

func installGolang(cmd *cobra.Command, args []string) {
	downloadURL := golang.GetDownloadURI(version)
	downloadTarTo := "/tmp/golang.tar.gz"
	defer os.Remove(downloadTarTo)
	fmt.Printf("Downloading golang version %s\n", version)
	if err := fileutils.DownloadFile(downloadTarTo, downloadURL); err != nil {
		fmt.Printf("failed to download golang from URL: %s to folder %s\n%v", downloadURL, installPath, err)
		return
	}
	fmt.Printf("Extracting to: %s\n", installPath)

	if err := fileutils.Untar(downloadTarTo, installPath); err != nil {
		fmt.Printf("failed to untar: %s\n%v", downloadTarTo, err)
		return
	}
	sourceSymlink := filepath.Join(installPath, fileutils.GetRootFolder(downloadTarTo))
	symlink := filepath.Join(installPath, symlinkName)
	fmt.Printf("Creating symlink %s -> %s\n", sourceSymlink, symlink)
	if err := fileutils.CreateUpdateSymlink(sourceSymlink, symlink); err != nil {
		fmt.Printf("failed to create symlink: %s\n%v", downloadTarTo, err)
	}
}