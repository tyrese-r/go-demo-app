package db

import (
	"fmt"
	"go-demo-app/internal/utils/logger"
	"log"
	"sync"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Database struct to hold the GORM instance
type Database struct {
	DB *gorm.DB
}

var (
	database *Database
	once     sync.Once
)

// GetDB returns the singleton GORM database instance
func GetDB() *Database {
	return database
}

// ConnectToDatabase initializes and connects to the SQLite database
func ConnectToDatabase() (*Database, error) {
	var err error

	once.Do(func() {
		var db *gorm.DB

		// Connection string with optimized SQLite pragmas for performance and memory usage:
		// - _journal=WAL: Use Write-Ahead Logging for better concurrency and less I/O blocking
		// - _cache_size=-8000: Use 8MB of RAM for DB cache (negative means kilobytes)
		// - _page_size=4096: Optimal page size for most file systems
		// - _busy_timeout=5000: Wait 5s on busy DB instead of returning error immediately
		// - _synchronous=NORMAL: Reduced I/O synchronization without risking corruption
		// - _temp_store=MEMORY: Store temp tables in memory instead of on disk
		// - _mmap_size=67108864: Use memory mapping for up to 64MB to reduce I/O
		// - _foreign_keys=ON: Enforce data integrity with foreign key constraints
		dsn := "app.db?_journal=WAL&_cache_size=-8000&_page_size=4096&_busy_timeout=5000&" +
			"_synchronous=NORMAL&_temp_store=MEMORY&_mmap_size=67108864&_foreign_keys=ON"

		// Configure GORM logger to reduce verbosity
		gormLogger := gormlogger.New(
			log.New(log.Writer(), "\r\n", log.LstdFlags),
			gormlogger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  gormlogger.Error, // Only log errors
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)

		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: gormLogger,
			// Speed up queries by disabling transaction wrapping for single create/update
			SkipDefaultTransaction: true,
		})
		if err != nil {
			logger.Error.Fatalf("Failed to open db: %v", err)
		}

		// Configure connection pool
		sqlDB, err := db.DB()
		if err != nil {
			logger.Error.Fatalf("Failed to get DB connection: %v", err)
		}

		// Limit connections to prevent memory bloat
		sqlDB.SetMaxOpenConns(10)           // Maximum 10 open connections
		sqlDB.SetMaxIdleConns(5)            // Keep up to 5 idle connections
		sqlDB.SetConnMaxLifetime(time.Hour) // Recycle connections after 1 hour

		database = &Database{DB: db}
		log.Println("Connected to SQLite database with optimized settings!")
	})

	return database, err
}

// CloseDatabase closes the database connection (GORM handles this automatically, but included for completeness)
func CloseDatabase() {
	sqlDB, err := database.DB.DB()
	if err != nil {
		logger.Error.Fatalf("Error getting SQL DB instance: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		logger.Error.Fatalf("Error closing database: %v", err)
	}

	fmt.Println("Database connection closed.")
}
