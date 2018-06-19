package main

import (
	"text/tabwriter"
	"os"
	"fmt"
)

func main(){
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "username \t firstname \t lastname")
	fmt.Fprintln(w, "chauhr \t ritesh \t chauhan")
	fmt.Fprintln(w, "asinghal \t ankur \t singhal")
	w.Flush()
}

