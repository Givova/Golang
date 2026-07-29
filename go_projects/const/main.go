package main

import (
	"fmt"
	"strings"
)

type Season int

const (
	_      Season = iota
	Spring        //1
	Summer        //2
	Autumn        //3
	Winter        //4
)



type FileMode int
 
const (
    ModeRead    FileMode = 1 << iota
    ModeWrite    
    ModeExecute  
)

func Describe(m FileMode) string{
	 var parts []string
	if m&ModeRead != 0 {
		parts = append(parts, "read")
	}
	if m&ModeWrite != 0 {
        parts = append(parts, "write")
  }
  if m&ModeExecute != 0 {
        parts = append(parts, "execute")
  }
  
	return strings.Join(parts, " ")
}

func (s Season) String() string {
	switch s {
	case 1:
		return "Весна"
	case 2:
		return "Лето"
	case 3:
		return "Осень"
	case 4:
		return "Зима"
	default:
		return "Uncnown seson"
	}
}


const (
    _  = iota
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
    TB                     // 1099511627776
)
 
func FormatSize(bytes int) string {
    switch {
    case bytes >= TB:
        return fmt.Sprintf("%.2f TB", float64(bytes)/float64(TB))
    case bytes >= GB:
        return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
    case bytes >= MB:
        return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
    case bytes >= KB:
        return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
    default:
        return fmt.Sprintf("%d B", bytes)
    }
}

func main() {
	fmt.Println(int(Spring)) // 1
	fmt.Println(Spring)      // "весна"

	perms := ModeRead | ModeWrite
fmt.Println(Describe(perms)) // "read write"
}