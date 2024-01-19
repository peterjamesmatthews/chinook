package crow

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"pjm.dev/chinook/util"
)

// Crow is a *gorm.DB seeding and dumping tool.
//
// It's used to quickly seed a database with data for testing, and to quickly
// dump that data back out of the database.
type Crow struct {
	db *gorm.DB
	s  map[any][]any
	d  map[any][]any
}

func (c *Crow) seed(db *gorm.DB) error {
	c.db = db

	var models []any
	var records []any
	for model, recs := range c.s {
		models = append(models, model)
		records = append(records, recs...)
	}

	if err := c.db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to automigrate models: %v\n%w", models, err)
	}

	for _, record := range records {
		if err := c.db.Create(util.GetPointerTo(record)).Error; err != nil {
			return fmt.Errorf("failed to create record %v\n%w", record, err)
		}
	}

	return nil
}

func (c *Crow) dump() error {
	d := make(map[any][]any)

	// for each model that was seeded
	for model := range c.s {
		// model is a concrete type whose pointer is a schema.Tabler
		tabler, ok := util.GetPointerTo(model).(schema.Tabler)
		if !ok {
			return fmt.Errorf("model %v does not implement schema.Tabler", model)
		}

		// create pointer to slice on model's concrete type
		m := reflect.TypeOf(util.GetPointedTo(model)) // m is the concrete type of model (not a pointer)
		ms := reflect.SliceOf(m)                      // ms is []m
		records := reflect.New(ms).Interface()        // [records] = any(*[]m)

		// find records of model from database
		if err := c.db.Table(tabler.TableName()).Find(records).Error; err != nil {
			return fmt.Errorf("failed to find records of model %v\n%w", model, err)
		}

		// convert [records] to []any(m)
		models := reflect.ValueOf(records).Elem() // [models] = []m
		anys := make([]any, models.Len())
		for i := range anys {
			anys[i] = models.Index(i).Interface() // [anys[i]] = any(m)
		}
		d[util.GetPointedTo(model)] = anys
	}

	c.d = d
	return nil
}

// Seed maps models to the records that should be seeded for that model.
//
// Usage:
//
//	c.Seed(map[any][]any{
//	  model1: []any{
//	    model1{...},
//	    model1{...},
//	  },
//	  model2: []any{
//	    model2{...},
//	    model2{...},
//	  },
//	})
//
// You may also seed the models as pointer or mix and match pointers & values.
func (c *Crow) Seed(s map[any][]any) {
	c.s = s
}

// Dump returns a map of models to the records that were dumped from the
// database.
//
// Usage:
//
//	d := c.Dump()
//	model1s := d[model1].([]model1)
//	model2s := d[model2].([]model2)
//
// Dump()'s keys and values will always be concrete types, not pointers.
func (c *Crow) Dump() map[any][]any {
	return c.d
}
