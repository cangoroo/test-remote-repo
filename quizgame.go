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


func getUserAnswer() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = strings.Replace(num, "\n", "", -1)
	ans, err := strconv.Atoi(num)
	return ans, err
}

	
func displayCounter(ctr int) {
        fmt.Printf("No. of correct answers: %d \n", ctr)
}

var limit  = flag.Int("limit", 10, "The time limit for solving the problems") 

func main() {
     /* parse command line flag limit, setup "done" timer channel, create csv reader */	
     flag.Parse()
     done := time.After(time.Duration(*limit) * time.Second)
     csvFile, _ := os.Open("qanda.csv")
     reader := csv.NewReader(bufio.NewReader(csvFile))

     counter := 0  // counter for correct no. of answers
     
     for {

         select {
	      case <-done:
		   displayCounter(counter)
	      	   fmt.Println("Timed out")
		   os.Exit(0)
              default:
		   break 
	 }
 
	 /* read each line of the csv file,
            line[0] contains "addend1 + addend2"
            line[1] contains the sum */

     	 line, error := reader.Read()
	 if error == io.EOF {
	    break
	 } else if error != nil {
	    log.Fatal(error)
         }
	 
	 fmt.Printf("%s = ", line[0])
	 /* enter user answer */
	 ans, err := getUserAnswer() 
	 if err != nil {
	    fmt.Printf("There is an issue converting the user response %s to integer.\n", ans) 
         }
	 
	 /* get the sum from the line */
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
     displayCounter(counter)	
}   
