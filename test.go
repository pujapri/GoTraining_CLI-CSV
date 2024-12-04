package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	SiteID                string
	FixletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

// Load CSV file into a slice of CSV
func load_CSV_data(filePath string) ([]Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Not able to open a file  ", err.Error())
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic("Error while reading the file ")
	}

	var entries []Data
	for _, row := range rows[1:] {
		count, _ := strconv.Atoi(row[4])
		entries = append(entries, Data{row[0], row[1], row[2], row[3], count})
	}
	return entries, nil
}

// Save entries to a CSV file
func save_CSV_data(filePath string, entries []Data) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"SiteID", "FixletID", "Name", "Criticality", "RelevantComputerCount"})
	for _, e := range entries {
		writer.Write([]string{e.SiteID, e.FixletID, e.Name, e.Criticality, strconv.Itoa(e.RelevantComputerCount)})
	}
	return nil
}

func main() {
	filePath := "file.csv"
	entries, err := load_CSV_data(filePath)
	if err != nil {
		fmt.Println("Error loading file:", err)
		return
	}

	for {
		fmt.Print("Enter your choice!")
		fmt.Print("\n1. List  2. Query  3. Sort  4. Add  5. Delete  6. Exit\nChoose: ")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			listAllEntries(entries)
		} else if choice == 2 {
			var key, value string
			fmt.Print("Enter key (SiteID/FixletID/Name/Criticality): ")
			fmt.Scan(&key)
			fmt.Print("Enter value: ")
			fmt.Scan(&value)
			queryEntries(entries, key, value)
		} else if choice == 3 {
			sortEntries(entries)
			fmt.Println("Entries sorted.")
		} else if choice == 4 {
			entries = append(entries, addEntry())
		} else if choice == 5 {
			var fixletID string
			fmt.Print("Enter FixletID to delete: ")
			fmt.Scan(&fixletID)
			entries = deleteEntry(entries, fixletID)
		} else if choice == 6 {
			if err := save_CSV_data(filePath, entries); err != nil {
				fmt.Println("Error saving file:", err)
			}
			fmt.Println("Exiting the program!")
			break
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

// Display all entries
func listAllEntries(entries []Data) {
	for i, e := range entries {
		fmt.Printf("%d. %+v\n", i+1, e)
	}
}

// Query entries by a key-value pair
func queryEntries(entries []Data, key, value string) {
	for _, e := range entries {
		switch key {
		case "SiteID":
			if e.SiteID == value {
				fmt.Println(e)
			}
		case "FixletID":
			if e.FixletID == value {
				fmt.Println(e)
			}
		case "Name":
			if e.Name == value {
				fmt.Println(e)
			}
		case "Criticality":
			if e.Criticality == value {
				fmt.Println(e)
			}
		}
	}
}

// Sort entries by FixletID
func sortEntries(entries []Data) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].FixletID < entries[j].FixletID
	})
}

// Add a new entry
func addEntry() Data {
	var e Data
	fmt.Print("Enter the value of SiteID FixletID Name Criticality RelevantComputerCount: ")
	fmt.Scan(&e.SiteID, &e.FixletID, &e.Name, &e.Criticality, &e.RelevantComputerCount)
	return e
}

// Delete an entry by FixletID
func deleteEntry(entries []Data, fixletID string) []Data {
	var filtered []Data
	for _, e := range entries {
		if e.FixletID != fixletID {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
