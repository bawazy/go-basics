package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rows := readOrders("orders.csv")
	rows = calculate(rows)
	writeOrders("ordersReport.csv", rows)
}
func readOrders(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	return rows
}

func calculate(rows [][]string) [][]string {

	sum := 0
	nb := 0

	for i := range rows {

		if i == 0 { // the first row
			rows[0] = append(rows[0], "TOTAL")
			continue // start the next loop iteration
			// fmt.Println(rows[0][2])
		}

		// All Subsequent rows:

		item := rows[i][2]

		price, err := strconv.Atoi(strings.Replace(rows[i][3], ".", "", -1))
		if err != nil {
			log.Fatalf("Cannot retrieve price of %s: %s\n", item, err)
		}
		qty, err := strconv.Atoi(rows[i][4])
		if err != nil {
			log.Fatalf("Cannot retrieve Quantity of %s: %s\n", item, err)
		}

		total := price * qty
		rows[i] = append(rows[i], intToFloatString(total))
		sum += total

		if item == "Ball Pen" {
			nb += qty
		}
	}
	rows = append(rows, []string{"", "", "", "Sum", "", intToFloatString(sum)})
	rows = append(rows, []string{"", "", "", "Ball Pens", "", fmt.Sprint(nb)})

	return rows
}

func intToFloatString(n int) string {
	intgr := n / 100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d", intgr, frac)
}

func writeOrders(name string, rows [][]string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal("Cannot create %s: %s\n", name, err)
	}
	// we are going to write to a file , so any errors are relevant and need to be logged. Hence the anonymous instead of a one liner

	defer func() {
		e := f.Close()
		if e != nil {
			log.Fatalf("Cannot close %s: %s\n", name, err.Error())
		}
	}()

	w := csv.NewWriter(f)
	err = w.WriteAll(rows)
}
