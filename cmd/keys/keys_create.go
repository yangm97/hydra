/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package keys

import (
	"github.com/ory/hydra/cmd"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var keysCreateCmd = &cobra.Command{
	Use:   "create <set> <key>",
	Short: "Create a new JSON Web Key Set",
	Run:   cmd.cmdHandler.Keys.CreateKeys,
}

func init() {
	keysCreateCmd.Flags().StringP("alg", "a", "RS256", "The algorithm to be used to generated they key. Supports: RS256, ES512, HS256")
	keysCreateCmd.Flags().StringP("use", "u", "sig", "The intended use of this key")
}