package voice

import (
	"fmt"
	"os/exec"
)

// Say executes the say command with the
// provided arguments as the string being said
func Say(d ...interface{}) {
	exec.Command("say", fmt.Sprintf("%v", d), "-v", "Vicki").Run()
}
