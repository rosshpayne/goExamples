package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Pagesize = ", os.Getpagesize())
	cmd := exec.Command("sleep", "3")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	//
	cmd = exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe() // connects cmd.Stderr to a pipe which is read laster using ioutil.ReadAll
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil { // safe to cmd.Wait() as ReadAll has read all from pipe
		log.Fatal(err)
	}
	//
	cmd = exec.Command("cat")
	stdin, err := cmd.StdinPipe() // connect cmd.Stdin to local var
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, " ***** values written to stdin are passed to cmd's standard input") // write to cmd.Stdin as a goroutine
	}()

	out2, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out2)
	//
	cmd = exec.Command("echo", "-n", `{"Name": "Bob", "Age": 32}`)
	_, err = cmd.StdoutPipe() // hookup cmd.Stdout to var stdout
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil { // run command but do not wait...use wait() instead
		log.Fatal(err)
	}
	/*if err := json.NewDecoder(stdout).Decode(&person); err != nil { // hookup var stdout to json decorder (io.Reader is stdout where stdout is *os.File)
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil { // wait for cmd to finish. /
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
*/
	//
	//
	cmd = exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input") //
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	//
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		log.Fatal(err)
	}
}
