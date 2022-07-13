package main

import (
	"log"
	"os/exec"
)

func exCmd(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err, string(out))
	}
	log.Print(string(out))
}

/* Strips JSON Schema to the subset that is JSON Typedef, for easier code generation (better ecosystem) */
func genTypedef(path string) {
	
}

func main() {
	names := []string{"capability", "permission", "url", "manifest"}
	list := ""
	for _, name := range names {
		list = list + "schemas/" + name + ".schema.json "
	}
	println("LIST: ", "gojsonschema -p main " + list + "> src/types.go")
	cmd := exec.Command("bash", "-c", "gojsonschema -p main " + list + "> src/types.go")
	cmd.Dir = "../"
	exCmd(cmd)

}
