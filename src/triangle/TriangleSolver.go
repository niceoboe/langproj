package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	//Make a triangle which is actually a 2D slice of integers
	//The slice named values will be added to triangle
	//A temporary variable named val will be needed to convert values
	triangle := make([][]int, 0)
	line := make([]string, 0)
	values := make([]int, 0)
	var val int

	//Create a scanner to read in a text file
	file, err := os.Open("triangle.txt")
	if err != nil {
		fmt.Printf("File could not be opened: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	/* For every line, the scanner will clear the values slice
	 * and then split that line using whitespace as a delimiter.
	 * Each element in this new array will be converted to an int
	 * and added to the values slice. Afterwards, that values slice
	 * is added to the triangle.
	 */
	for scanner.Scan() {
		values := values[:0]
		line = strings.Split(scanner.Text(), " ")
    	for _, e := range line[:len(line)-1]{
    		val, err = strconv.Atoi(e)
    		if err != nil {
    			fmt.Printf("Error parsing file: %v\n", err)
    			os.Exit(2)
    		}
    		values = append(values, val)
    	}
    	triangle = append(triangle, values)
	}
	fmt.Printf("Largest sum of triangle: %v\n", collapseTriangle(triangle))
}

func collapseTriangle(triangle [][]int) int {
	//Starts at the second to last row and "collapses" the row below it
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			//Adds a child node to the parent node which will produce the largest sum
			v := triangle[i][j] + triangle[i+1][j]
			r := triangle[i][j] + triangle[i+1][j+1]
			if v >= r {
				triangle[i][j] = v
			} else {
				triangle[i][j] = r
			}
		}
	}
	return triangle[0][0]
}