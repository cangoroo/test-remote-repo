package main

import (
     "bufio"
     "encoding/csv"
     "fmt"
     "io"
     "log"
     "os"
     "strings"
     "strconv"
     "time"
     "flag"
)

var limit  = flag.Int("limit", 30, "Default time limit is 30 seconds.") 

func main() {
     flag.Parse() // limit no. of seconds allowed before terminating the app
     done := make(chan struct{}) 
     csvFile, _ := os.Open("qanda.csv")
     reader := csv.NewReader(bufio.NewReader(csvFile))

     counter := 0 // counter variable for correct answers

     go func() {
        defer close(done)
	for {

     	    line, error := reader.Read()
	    if error == io.EOF {
	       break
	    } else if error != nil {
	      log.Fatal(error)
            }

	    // line[0] is "a+b ="	 
	    fmt.Printf("%s = ", line[0])
	    
	    /* enter user answer */
	    reader := bufio.NewReader(os.Stdin)
	    num, _ := reader.ReadString('\n')
	    num = strings.Replace(num, "\n", "", -1)
	    ans, err := strconv.Atoi(num)
	    if err != nil {
	       fmt.Printf("There is an issue converting the user response %s to integer.\n", num)
            }
	 
	    /* get the sum from the line[1], line[1] is c from "a+b, c" line */
	    sum, err := strconv.Atoi(line[1]) 
	    if err != nil {
	       fmt.Printf("There is an issue with converting %s to integer.\n", line[1])
	       os.Exit(1)
            }

	    /* compare user response with sum from line */
	    if ans == sum {
	       counter += 1
            }

     	 }
     }()

     select {
     	    case <-done:
     	    	 fmt.Println("completed processing sums.")
     	    case <-time.After(time.Duration(*limit) * time.Second):
     	    	 fmt.Println("sorry, took to long add numbers.")
     }
     fmt.Printf("\n num of correct answers: %d \n", counter)

}   
