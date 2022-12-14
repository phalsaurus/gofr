package datastore

import (
	"database/sql"
	"io"
	"testing"

	"github.com/jmoiron/sqlx"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"developer.zopsmart.com/go/gofr/pkg"
	"developer.zopsmart.com/go/gofr/pkg/gofr/config"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"
	"developer.zopsmart.com/go/gofr/pkg/log"
)

type mockPubSub struct {
	Param string
}

func (m *mockPubSub) HealthCheck() types.Health {
	if m.Param == "kafka" {
		return types.Health{
			Status:   pkg.StatusUp,
			Database: pkg.Kafka,
		}
	}

	if m.Param == "eventhub" {
		return types.Health{
			Status:   pkg.StatusUp,
			Database: pkg.Kafka,
		}
	}

	return types.Health{
		Status: pkg.StatusDown,
	}
}

func TestDataStore_GORM(t *testing.T) {
	{
		ds := new(DataStore)
		ds.gorm.DB = new(gorm.DB)

		db := ds.GORM()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)
		c := config.NewGoDotEnvProvider(log.NewMockLogger(io.Discard), "../../configs")
		cfg := &DBConfig{
			HostName: c.Get("DB_HOST"),
			Username: c.Get("DB_USER"),
			Password: c.Get("DB_PASSWORD"),
			Database: c.Get("DB_NAME"),
			Port:     c.Get("DB_PORT"),
			Dialect:  c.Get("DB_DIALECT"),
		}

		client, _ := NewORM(cfg)
		ds.SetORM(client)

		db := ds.GORM()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)

		db := ds.GORM()
		if db != nil {
			t.Errorf("FAILED, Expected: nil, Got: %v", db)
		}
	}

	{
		ds := new(DataStore)
		if db, ok := ds.ORM.(*gorm.DB); ok {
			if db != nil {
				t.Errorf("FAILED, expected nil, Got: %v", db)
			}
		}
	}
}

func TestDataStore_SQLX(t *testing.T) {
	{
		ds := new(DataStore)
		ds.sqlx = SQLXClient{DB: new(sqlx.DB)}

		db := ds.SQLX()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)
		ds.SetORM(SQLXClient{DB: new(sqlx.DB)})

		db := ds.SQLX()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)

		db := ds.SQLX()
		if db != nil {
			t.Errorf("FAILED, Expected: nil, Got: %v", db)
		}
	}
}

// TestDataStore_DB tests the behavior of ds.DB() when DB connection is not established.
// It tests, whether it will panic or throw error. For example when /.well-known/health-check api pings DB for its status
// it shouldn't panic if the DB connection is not established.
func TestDataStore_DB(t *testing.T) {
	{
		ds := new(DataStore)
		ds.rdb.DB = new(sql.DB)

		db := ds.DB()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)

		// passing incorrect dsn will not establish a db connection. But gorm.ConnPool will not be nil.
		// passing new(gorm.DB) panics, as gorm.ConnPool will be nil.
		dsn := "incorrect DSN"

		gdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

		ds.SetORM(GORMClient{DB: gdb})

		db := ds.GORM()
		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}

	{
		ds := new(DataStore)
		ds.SetORM(SQLXClient{DB: new(sqlx.DB)})
		db := ds.DB()

		if db.DB != nil {
			t.Errorf("FAILED, Expected the db object to be nil, Got: %v", db)
		}
	}

	{
		ds := new(DataStore)
		db := ds.DB()

		if db != nil {
			t.Errorf("FAILED, Expected the db object to be nil, Got: %v", db)
		}
	}

	{
		ds := new(DataStore)
		ds.ORM = new(sql.DB)
		db := ds.DB()

		if db == nil {
			t.Error("FAILED, Expected the db object to be initialized, Got: nil")
		}
	}
}

func TestDataStore_KafkaHealthCheck(t *testing.T) {
	{
		kafkaClient := &mockPubSub{}
		healthCheck := kafkaClient.HealthCheck()

		if healthCheck.Status != pkg.StatusDown {
			t.Errorf("Failed: Expected kafka health check status as DOWN Got %v", healthCheck.Status)
		}
	}
	{
		kafkaClient := &mockPubSub{"kafka"}
		healthCheck := kafkaClient.HealthCheck()

		if healthCheck.Status != pkg.StatusUp {
			t.Errorf("Failed: Expected kafka health check status as UP Got %v", healthCheck.Status)
		}
	}
}

func TestDataStore_EventHubHealthCheck(t *testing.T) {
	{
		eventhubClient := &mockPubSub{}
		healthCheck := eventhubClient.HealthCheck()
		if healthCheck.Status != pkg.StatusDown {
			t.Errorf("Failed: Expected EventHub health check status as DOWN Got %v", healthCheck.Status)
		}
	}
	{
		kafkaClient := &mockPubSub{"eventhub"}
		healthCheck := kafkaClient.HealthCheck()

		if healthCheck.Status != pkg.StatusUp {
			t.Errorf("Failed: Expected EventHub health check status as UP Got %v", healthCheck.Status)
		}
	}
}

// TestSQLX_ORM tests when sqlx instance is initialized to ORM
func TestSQLX_ORM(t *testing.T) {
	c := config.NewGoDotEnvProvider(log.NewMockLogger(io.Discard), "../../configs")

	db, _ := NewSQLX(&DBConfig{
		HostName: c.Get("DB_HOST"),
		Username: c.Get("DB_USER"),
		Password: c.Get("DB_PASSWORD"),
		Database: "mysql",
		Port:     c.Get("DB_PORT"),
		Dialect:  c.Get("DB_DIALECT"),
	})

	ds := &DataStore{ORM: db.DB}
	if ds.SQLX() == nil {
		t.Errorf("Not got sqxl.DB")
	}
}
