package main 

import  "fmt"

func main(){


	
	res := third()
	fmt.Println(res)

}

func third() []int{
	var repNum []int

	for i := 100; i < 1000; i++ {
		intString := fmt.Sprint(i)
		if intString[0] == intString[2]{
			repNum = append(repNum, i)
		}
	}
	return repNum
}