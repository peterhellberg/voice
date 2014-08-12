package voice

import (
	"fmt"
	"os/exec"
)

func Say(d ...interface{}) {
	exec.Command("say", fmt.Sprintf("%v", d), "-v", "Vicki").Run()
}
