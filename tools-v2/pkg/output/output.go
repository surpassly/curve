/*
 *  Copyright (c) 2022 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveCli
 * Created Date: 2022-05-11
 * Author: chengyi (Cyber-SiKu)
 */

package output

import (
	"encoding/json"
	"fmt"

	cmderror "github.com/opencurve/curve/tools-v2/internal/error"
	basecmd "github.com/opencurve/curve/tools-v2/pkg/cli/command"
	"github.com/opencurve/curve/tools-v2/pkg/config"
	"github.com/spf13/viper"
)

const (
	FORMAT_JSON  = "json"
	FORMAT_PLAIN = "plain"
)

func FinalCmdOutputJson(finalCmd *basecmd.FinalCurveCmd) error {
	output, err := json.MarshalIndent(finalCmd, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func FinalCmdOutput(finalCmd *basecmd.FinalCurveCmd,
	funcs basecmd.FinalCurveCmdFunc) error {
	format := viper.GetString("format")
	finalCmd.Error = cmderror.MostImportantCmdError(finalCmd.AllError)
	var err error
	switch format {
	case FORMAT_JSON:
		err = FinalCmdOutputJson(finalCmd)
	case FORMAT_PLAIN:
		err = funcs.ResultPlainOutput()
		if viper.GetBool(config.VIPER_GLOBALE_SHOWERROR) {
			for _, output := range finalCmd.AllError {
				fmt.Printf("%+v\n", output)
			}
		}
	default:
		err = nil
	}
	return err
}
