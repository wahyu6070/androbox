package kopi

import (
"fmt"
*os/exec"
)
func CallClear() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}

