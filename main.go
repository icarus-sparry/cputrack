package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("/proc/stat")
	check(err)
	ticker := time.NewTicker(1000 * time.Millisecond)
	b := make([]byte, 20000)
	var oldcpu [10]int
	var first bool = true
	for {
		_, err = f.Seek(0, 0)
		check(err)
		n, err := f.Read(b)
		check(err)
		str := string(b[:n])
		for _, l := range strings.Split(str, "\n") {
			l2 := strings.SplitN(l, " ", 2)
			switch l2[0] {
			case "cpu":
				if first {
					fmt.Print("time,type,user,nice,system,idle,iowait,irq,softirq,steal,guest,gnice")
				} else {
					fmt.Printf("%v,", time.Now().Unix())
					fmt.Print("cpu")
				}
				for i, v := range strings.Fields(l2[1]) {
					vi, err := strconv.Atoi(v)
					check(err)
					if !first {
						fmt.Printf(",%d", vi-oldcpu[i])
					}
					oldcpu[i] = vi
				}
				first = false
				fmt.Print("\n")
			// case "processes":
			// 	fmt.Printf("%s\n", l)
			// 	break
			default:
			}
		}
		<-ticker.C
	}
}
