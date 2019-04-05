package main 

import ( 
    "fmt" 
    "time" 
) 

// daysIn returns the number of days in a month for a given year. 
func daysIn(m time.Month, year int) int { 
    // This is equivalent to time.daysIn(m, year). 
    return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day() 
} 

func main() { 
    fmt.Println(daysIn(1, 2000)) 
    fmt.Println(daysIn(2, 2000)) 
    fmt.Println(daysIn(2, 2001)) 
    fmt.Println(daysIn(6, 2000)) 
    fmt.Println(daysIn(12, 2019)) 
fmt.Println(daysIn(2, 1900)) 
} 

/*
output
------
31
29
28
30
31
28
*/
