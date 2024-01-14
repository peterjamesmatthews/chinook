package db

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pjm.dev/chinook/util"
)

type Crow struct {
	t    *testing.T
	db   *gorm.DB
	seed map[any][]any
}

func NewCrow(t *testing.T, db *gorm.DB) Crow {
	return Crow{t: t, db: db}
}

func (c *Crow) Seed(seed map[any][]any) {
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
		if err := c.db.Create(util.GetPointerTo(record)).Error; err != nil {
			c.t.Fatalf("failed to insert data: %v", err)
		}
	}

	c.seed = seed
}

// Dump returns a map of models to their records in the database.
func (c Crow) Dump() map[any][]any {
	dump := make(map[any][]any)

	// for each model that was seeded
	for model := range c.seed {
		// model is a concrete type whose pointer is a schema.Tabler
		tabler, ok := util.GetPointerTo(model).(schema.Tabler)
		if !ok {
			c.t.Fatalf("model %v does not implement schema.Tabler", model)
		}

		// create pointer to slice on model's concrete type
		m := reflect.TypeOf(util.GetPointedTo(model)) // m is the concrete type of model (not a pointer)
		ms := reflect.SliceOf(m)                      // ms is []m
		records := reflect.New(ms).Interface()        // [records] = any(*[]m)

		// find records of model from database
		if err := c.db.Table(tabler.TableName()).Find(records).Error; err != nil {
			c.t.Fatalf("failed to find records of %v from database\n%v", tabler.TableName(), err)
		}

		// convert [records] to []any(m)
		models := reflect.ValueOf(records).Elem() // [models] = []m
		anys := make([]any, models.Len())
		for i := range anys {
			anys[i] = models.Index(i).Interface() // [anys[i]] = any(m)
		}
		dump[util.GetPointedTo(model)] = anys
	}

	return dump
}
