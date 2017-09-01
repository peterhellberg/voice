package voice

import (
	"fmt"
	"os/exec"
)

// Name of the voice use by the Say function.
// List all voices by running: say -v ?
var Name = "Samantha"

// Say executes the say command with the
// provided arguments as the string being said
func Say(d ...interface{}) {
	exec.Command("say", fmt.Sprintf("%v", d), "-v", Name).Run()
}
