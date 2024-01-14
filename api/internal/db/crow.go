package db

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Crow struct {
	db *gorm.DB
	t  *testing.T
}

func NewCrow(t *testing.T, db *gorm.DB) Crow {
	return Crow{db: db, t: t}
}

func (c Crow) Seed(seed map[schema.Tabler][]any) {
	var models []any
	var records []any
	for model, recs := range seed {
		models = append(models, model)
		records = append(records, recs...)
	}

	if err := c.db.AutoMigrate(models...); err != nil {
		c.t.Fatalf("failed to migrate schema: %v", err)
	}

	for _, record := range records {
		if err := c.db.Create(record).Error; err != nil {
			c.t.Fatalf("failed to insert data: %v", err)
		}
	}
}

func (c Crow) Assert(want map[schema.Tabler][]any) {
	// TODO establish a consistent naming for the type of want's key and values

	for table, wantRecords := range want {
		gotRecordsInterface := reflect.New(reflect.SliceOf(reflect.TypeOf(table).Elem())).Interface()
		if err := c.db.Table(table.TableName()).Find(gotRecordsInterface).Error; err != nil {
			c.t.Fatalf("failed to get records from database\n%v", err)
		}

		// gotRecords is a pointer to a slice of the correct type
		// wantRecords is a slice of interface{} that are pointers to the correct type

		// convert gotRecords to a slice of the correct type
		gotSlice := reflect.ValueOf(gotRecordsInterface).Elem()
		var gotInterfaces []interface{}
		for i := 0; i < gotSlice.Len(); i++ {
			gotInterfaces = append(gotInterfaces, gotSlice.Index(i).Interface())
		}

		// convert wantRecords to a slice of the correct type
		var wantRecordsValues []interface{}
		for _, wantRecord := range wantRecords {
			// Dereference the pointer to get the value it points to
			wantRecordsValues = append(wantRecordsValues, reflect.ValueOf(wantRecord).Elem().Interface())
		}

		// TODO use generic soft unordered equality
		if !reflect.DeepEqual(gotInterfaces, wantRecordsValues) {
			c.t.Errorf("got records mismatch\nwant %+v\ngot %+v", wantRecordsValues, gotInterfaces)
		}
	}
}
