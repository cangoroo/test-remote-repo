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


func convertToInt(x string) int {
     i, err := strconv.Atoi(x)
     if err != nil {
     	return -1
     }
     return i
}


var limit  = flag.Int("limit", 10, "The time limit for solving the problems") 

func main() {
     flag.Parse()
     done := time.After(time.Duration(*limit) * time.Second)
     
     csvFile, _ := os.Open("/Users/virgiliodevera/go/src/readingcsv/qanda.csv")
     reader := csv.NewReader(bufio.NewReader(csvFile))

     counter := 0
     
     for {

         select {
	      case <-done:
                   fmt.Printf("No. of correct answers: %d \n", counter)
	      	   fmt.Println("Timed out")
		   os.Exit(0)
              default:
		   break 
	 }

     	 line, error := reader.Read()
	 if error == io.EOF {
	    break
	 } else if error != nil {
	    log.Fatal(error)
         }
	 
	 fmt.Printf("%s = ", line[0])
	 /* enter user answer */
	 reader := bufio.NewReader(os.Stdin)
	 num, _ := reader.ReadString('\n')
	 num = strings.Replace(num, "\n", "", -1)
	 ans, err := strconv.Atoi(num)
	 if err != nil {
	    fmt.Printf("There is an issue converting the user response %s to integer.\n", num)
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
     fmt.Printf("\n num of correct answers: %d \n", counter)

}   