package main
import ("fmt";"os";"strconv")
func main() {
	total := 100
	if len(os.Args) > 1 { if n,err:=strconv.Atoi(os.Args[1]); err==nil && n>0 { total=n } }
	for i:=0; i<=total; i++ {
		pct := i*100/total
		bar := ""
		for j:=0; j<50; j++ {
			if j*100/50 <= pct { bar += "#" } else { bar += "." }
		}
		fmt.Printf("\r%s %3d%%", bar, pct)
	}
	fmt.Println()
}
