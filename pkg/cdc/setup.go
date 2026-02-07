package cdc

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// EnsurePublication checks if a publication exists, and creates it if not.
// It accepts a list of tables to include in the publication.
func EnsurePublication(db *gorm.DB, publicationName string, tables ...string) error {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_publication WHERE pubname = ?)"

	if err := db.Raw(query, publicationName).Scan(&exists).Error; err != nil {
		return fmt.Errorf("failed to check publication existence: %w", err)
	}

	if len(tables) == 0 {
		return fmt.Errorf("no tables specified for publication %s", publicationName)
	}

	// Build table list string: "table1, table2, table3"
	tableList := ""
	for i, t := range tables {
		if i > 0 {
			tableList += ", "
		}
		tableList += t
	}

	if !exists {
		log.Printf("[CDC Setup] Publication '%s' for tables [%s] does not exist. Creating...", publicationName, tableList)
		createPubQuery := fmt.Sprintf("CREATE PUBLICATION %s FOR TABLE %s", publicationName, tableList)
		if err := db.Exec(createPubQuery).Error; err != nil {
			return fmt.Errorf("failed to create publication '%s': %w", publicationName, err)
		}
		log.Printf("[CDC Setup] Publication '%s' created successfully.", publicationName)
	} else {
		// Optional: We could check if tables are missing and ALTER PUBLICATION, but for now we assume it's correct or managed manually.
		// A robust implementation would query pg_publication_tables.
		log.Printf("[CDC Setup] Publication '%s' already exists.", publicationName)

		// For robustness, let's force an update of the tables in the publication to ensure nothing is missing
		// This handles the case where we add a new table to the code.
		log.Printf("[CDC Setup] Updating publication '%s' to ensure tables [%s] are included...", publicationName, tableList)
		alterPubQuery := fmt.Sprintf("ALTER PUBLICATION %s SET TABLE %s", publicationName, tableList)
		if err := db.Exec(alterPubQuery).Error; err != nil {
			// Don't fail hard on ALTER, just log warning as it might be a permissions issue or minor state mismatch
			log.Printf("[CDC Setup] WARN: Failed to update publication tables: %v", err)
		}
	}

	return nil
}
