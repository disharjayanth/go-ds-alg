package main

import "fmt"

type Table struct {
	Name        string
	ColumnNames []string
	Rows        []Row
}

type Row struct {
	Id      int
	Columns []Column
}

type Column struct {
	Id    int
	Value string
}

func printTable(table *Table) {
	fmt.Println("Table name:", table.Name)
	for _, row := range table.Rows {
		fmt.Println("----------------------------------")
		for i, col := range row.Columns {
			fmt.Println(table.ColumnNames[i], col.Id, col.Value)
		}
	}
}

func main() {
	table := &Table{}
	table.Name = "Customer"
	table.ColumnNames = []string{"Id", "Name", "SSN"}
	rows := make([]Row, 2)

	rows[0].Id = 1
	columns1 := make([]Column, 3)
	columns1[0] = Column{1, "323"}
	columns1[1] = Column{1, "John"}
	columns1[2] = Column{1, "392529922"}
	rows[0].Columns = columns1

	rows[1].Id = 2
	columns2 := make([]Column, 3)
	columns2[0] = Column{2, "120"}
	columns2[1] = Column{2, "Smith"}
	columns2[2] = Column{2, "719333"}
	rows[1].Columns = columns2

	table.Rows = rows
	fmt.Println(table)

	printTable(table)
}
