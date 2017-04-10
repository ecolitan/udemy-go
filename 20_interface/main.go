package main

import "fmt"
import "sort"

type People []string

func (a People) Len() int           { return len(a) }
func (a People) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a People) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	studyGroup := People{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
