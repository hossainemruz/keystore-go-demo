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
	"flag"
	"github.com/hossainemruz/keytool-go/pkg"
	"github.com/spf13/cobra"
	"log"
)

var(
	cert string
	keystore string
	alias string
	pass string
)
var importcertCmd = &cobra.Command{
	Use:   "importcert",
	Short: "import certificate to keystore",
	Long: `Import certificate to keystore`,
	Run: func(cmd *cobra.Command, args []string) {
		err:=pkg.ImportCert(cert,keystore,pass,alias)
		if err!=nil{
			log.Fatal(err)
		}
		log.Println("Certificate imported successfully")
	},
}

func init() {

	importcertCmd.Flags().StringVar(&cert, "cert", "", "Target certificate which will be imported")
	importcertCmd.Flags().StringVar(&keystore, "keystore", "keystore.jks", "Name of the keystore file. Default is keystore.jks")
	importcertCmd.Flags().StringVar(&alias, "alias", "alias", "Alias to use")
	importcertCmd.Flags().StringVar(&pass, "pass", "", "Password for keystore")

	flag.Parse()

	RootCmd.AddCommand(importcertCmd)
}
