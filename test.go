package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Define a struct for the CSV data
type Entry struct {
	SiteID                string
	FxiletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

// Function to load CSV into a slice of Entry
func loadCSV(filePath string) ([]Entry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for i, row := range rows {
		if i == 0 {
			continue // Skip header
		}
		count, _ := strconv.Atoi(row[4]) // Convert RelevantComputerCount to int
		entries = append(entries, Entry{
			SiteID:                row[0],
			FxiletID:              row[1],
			Name:                  row[2],
			Criticality:           row[3],
			RelevantComputerCount: count,
		})
	}
	return entries, nil
}

// Function to save entries to CSV
func saveCSV(filePath string, entries []Entry) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"SiteID", "FxiletID", "Name", "Criticality", "RelevantComputerCount"})

	// Write data
	for _, entry := range entries {
		writer.Write([]string{
			entry.SiteID,
			entry.FxiletID,
			entry.Name,
			entry.Criticality,
			strconv.Itoa(entry.RelevantComputerCount),
		})
	}
	return nil
}

// Function to list entries
func listEntries(entries []Entry) {
	fmt.Println("Listing all entries:")
	for i, entry := range entries {
		fmt.Printf("%d. %+v\n", i+1, entry)
	}
}

// Function to query entries
func queryEntries(entries []Entry, key, value string) {
	fmt.Printf("Querying entries where %s = %s:\n", key, value)
	for _, entry := range entries {
		switch key {
		case "SiteID":
			if entry.SiteID == value {
				fmt.Println(entry)
			}
		case "FxiletID":
			if entry.FxiletID == value {
				fmt.Println(entry)
			}
		case "Name":
			if entry.Name == value {
				fmt.Println(entry)
			}
		case "Criticality":
			if entry.Criticality == value {
				fmt.Println(entry)
			}
		}
	}
}

// Function to sort entries by FxiletID
func sortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].FxiletID < entries[j].FxiletID
	})
	fmt.Println("Entries sorted by FxiletID.")
	listEntries(entries)
}

// Function to add an entry
func addEntry(entries []Entry) []Entry {
	var entry Entry
	fmt.Println("Enter new entry details:")
	fmt.Print("SiteID: ")
	fmt.Scan(&entry.SiteID)
	fmt.Print("FxiletID: ")
	fmt.Scan(&entry.FxiletID)
	fmt.Print("Name: ")
	fmt.Scan(&entry.Name)
	fmt.Print("Criticality: ")
	fmt.Scan(&entry.Criticality)
	fmt.Print("RelevantComputerCount: ")
	fmt.Scan(&entry.RelevantComputerCount)

	entries = append(entries, entry)
	fmt.Println("Entry added.")
	return entries
}

// Function to delete an entry
func deleteEntry(entries []Entry, fxiletID string) []Entry {
	var updatedEntries []Entry
	for _, entry := range entries {
		if entry.FxiletID != fxiletID {
			updatedEntries = append(updatedEntries, entry)
		}
	}
	fmt.Printf("Deleted entries with FxiletID = %s.\n", fxiletID)
	return updatedEntries
}

func main() {
	const filePath = "file.csv"
	entries, err := loadCSV(filePath)
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}

	for {
		fmt.Println("\nChoose an operation:")
		fmt.Println("1. List entries")
		fmt.Println("2. Query entries")
		fmt.Println("3. Sort entries by FxiletID")
		fmt.Println("4. Add entry")
		fmt.Println("5. Delete entry")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			listEntries(entries)
		case 2:
			fmt.Print("Enter column to query (SiteID, FxiletID, Name, Criticality): ")
			var key string
			fmt.Scan(&key)
			fmt.Print("Enter value to query: ")
			var value string
			fmt.Scan(&value)
			queryEntries(entries, key, value)
		case 3:
			sortEntries(entries)
		case 4:
			entries = addEntry(entries)
		case 5:
			fmt.Print("Enter FxiletID of the entry to delete: ")
			var fxiletID string
			fmt.Scan(&fxiletID)
			entries = deleteEntry(entries, fxiletID)
		case 6:
			if err := saveCSV(filePath, entries); err != nil {
				fmt.Println("Error saving CSV:", err)
			}
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
