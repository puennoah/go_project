//Program for finding lowest highest and average salary
package main 
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func main(){
	var fileName string = "salary.txt"
	salary, err := readFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", fileName, err)
		return
	}
	minName, minSalary := min(salary)
	maxName, maxSalary := max(salary)
	avgSalary := average(salary)
	fmt.Println("The person with lowest salary is", minName, "with", minSalary,"$")
	fmt.Println("The person with highest salary is", maxName, "with", maxSalary,"$")
	fmt.Println("The company's average salary is", avgSalary,"$")
}

func readFile(fileName string) (map[string]uint64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	salary := make(map[string]uint64)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		name := strings.Join([]string{fields[0],fields[1]},".")
		var wage uint64
		wage,err = strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			fmt.Println("Converting error")
			continue
		}
		salary[name] = wage
	}
	if err := scanner.Err(); err != nil {
    	return nil, err  // Return nil map and the error from scanning
	}
	return salary, nil
}

func min(salary map[string]uint64) (string, uint64) {//For finding person the earn lowest salry
	if len(salary) == 0 {
		return "Empty map", 0
	}
	var minName string
	var min uint64 = math.MaxUint64
	for name, wage := range salary {
		if min > wage {
			min = wage
			minName = name
		}
	}
	return minName, min
}

func max(salary map[string]uint64) (string, uint64) {//For finding person the earn highest salry
	if len(salary) == 0 {
		return "Empty map", 0
	}
	var maxName string
	var max uint64 = 0
	for name, wage := range salary {
		if max < wage {
			max = wage
			maxName = name
		}
	}
	return maxName, max
}

func average(salary map[string]uint64) uint64 { //For finding average
	if len(salary) == 0 {
		return 0
	}
	var result, accum uint64
	for _, wage := range salary {
		accum += wage
	}
	result = accum/uint64(len(salary))
	return result
	
}