// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"pjm.dev/chin/internal/db/model"
)

func newInvoice(db *gorm.DB, opts ...gen.DOOption) invoice {
	_invoice := invoice{}

	_invoice.invoiceDo.UseDB(db, opts...)
	_invoice.invoiceDo.UseModel(&model.Invoice{})

	tableName := _invoice.invoiceDo.TableName()
	_invoice.ALL = field.NewAsterisk(tableName)
	_invoice.InvoiceID = field.NewInt32(tableName, "InvoiceId")
	_invoice.CustomerID = field.NewInt32(tableName, "CustomerId")
	_invoice.InvoiceDate = field.NewTime(tableName, "InvoiceDate")
	_invoice.BillingAddress = field.NewString(tableName, "BillingAddress")
	_invoice.BillingCity = field.NewString(tableName, "BillingCity")
	_invoice.BillingState = field.NewString(tableName, "BillingState")
	_invoice.BillingCountry = field.NewString(tableName, "BillingCountry")
	_invoice.BillingPostalCode = field.NewString(tableName, "BillingPostalCode")
	_invoice.Total = field.NewFloat64(tableName, "Total")

	_invoice.fillFieldMap()

	return _invoice
}

type invoice struct {
	invoiceDo

	ALL               field.Asterisk
	InvoiceID         field.Int32
	CustomerID        field.Int32
	InvoiceDate       field.Time
	BillingAddress    field.String
	BillingCity       field.String
	BillingState      field.String
	BillingCountry    field.String
	BillingPostalCode field.String
	Total             field.Float64

	fieldMap map[string]field.Expr
}

func (i invoice) Table(newTableName string) *invoice {
	i.invoiceDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i invoice) As(alias string) *invoice {
	i.invoiceDo.DO = *(i.invoiceDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *invoice) updateTableName(table string) *invoice {
	i.ALL = field.NewAsterisk(table)
	i.InvoiceID = field.NewInt32(table, "InvoiceId")
	i.CustomerID = field.NewInt32(table, "CustomerId")
	i.InvoiceDate = field.NewTime(table, "InvoiceDate")
	i.BillingAddress = field.NewString(table, "BillingAddress")
	i.BillingCity = field.NewString(table, "BillingCity")
	i.BillingState = field.NewString(table, "BillingState")
	i.BillingCountry = field.NewString(table, "BillingCountry")
	i.BillingPostalCode = field.NewString(table, "BillingPostalCode")
	i.Total = field.NewFloat64(table, "Total")

	i.fillFieldMap()

	return i
}

func (i *invoice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *invoice) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["InvoiceId"] = i.InvoiceID
	i.fieldMap["CustomerId"] = i.CustomerID
	i.fieldMap["InvoiceDate"] = i.InvoiceDate
	i.fieldMap["BillingAddress"] = i.BillingAddress
	i.fieldMap["BillingCity"] = i.BillingCity
	i.fieldMap["BillingState"] = i.BillingState
	i.fieldMap["BillingCountry"] = i.BillingCountry
	i.fieldMap["BillingPostalCode"] = i.BillingPostalCode
	i.fieldMap["Total"] = i.Total
}

func (i invoice) clone(db *gorm.DB) invoice {
	i.invoiceDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i invoice) replaceDB(db *gorm.DB) invoice {
	i.invoiceDo.ReplaceDB(db)
	return i
}

type invoiceDo struct{ gen.DO }

type IInvoiceDo interface {
	gen.SubQuery
	Debug() IInvoiceDo
	WithContext(ctx context.Context) IInvoiceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInvoiceDo
	WriteDB() IInvoiceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInvoiceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInvoiceDo
	Not(conds ...gen.Condition) IInvoiceDo
	Or(conds ...gen.Condition) IInvoiceDo
	Select(conds ...field.Expr) IInvoiceDo
	Where(conds ...gen.Condition) IInvoiceDo
	Order(conds ...field.Expr) IInvoiceDo
	Distinct(cols ...field.Expr) IInvoiceDo
	Omit(cols ...field.Expr) IInvoiceDo
	Join(table schema.Tabler, on ...field.Expr) IInvoiceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInvoiceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInvoiceDo
	Group(cols ...field.Expr) IInvoiceDo
	Having(conds ...gen.Condition) IInvoiceDo
	Limit(limit int) IInvoiceDo
	Offset(offset int) IInvoiceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInvoiceDo
	Unscoped() IInvoiceDo
	Create(values ...*model.Invoice) error
	CreateInBatches(values []*model.Invoice, batchSize int) error
	Save(values ...*model.Invoice) error
	First() (*model.Invoice, error)
	Take() (*model.Invoice, error)
	Last() (*model.Invoice, error)
	Find() ([]*model.Invoice, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Invoice, err error)
	FindInBatches(result *[]*model.Invoice, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Invoice) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInvoiceDo
	Assign(attrs ...field.AssignExpr) IInvoiceDo
	Joins(fields ...field.RelationField) IInvoiceDo
	Preload(fields ...field.RelationField) IInvoiceDo
	FirstOrInit() (*model.Invoice, error)
	FirstOrCreate() (*model.Invoice, error)
	FindByPage(offset int, limit int) (result []*model.Invoice, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInvoiceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i invoiceDo) Debug() IInvoiceDo {
	return i.withDO(i.DO.Debug())
}

func (i invoiceDo) WithContext(ctx context.Context) IInvoiceDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i invoiceDo) ReadDB() IInvoiceDo {
	return i.Clauses(dbresolver.Read)
}

func (i invoiceDo) WriteDB() IInvoiceDo {
	return i.Clauses(dbresolver.Write)
}

func (i invoiceDo) Session(config *gorm.Session) IInvoiceDo {
	return i.withDO(i.DO.Session(config))
}

func (i invoiceDo) Clauses(conds ...clause.Expression) IInvoiceDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i invoiceDo) Returning(value interface{}, columns ...string) IInvoiceDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i invoiceDo) Not(conds ...gen.Condition) IInvoiceDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i invoiceDo) Or(conds ...gen.Condition) IInvoiceDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i invoiceDo) Select(conds ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i invoiceDo) Where(conds ...gen.Condition) IInvoiceDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i invoiceDo) Order(conds ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i invoiceDo) Distinct(cols ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i invoiceDo) Omit(cols ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i invoiceDo) Join(table schema.Tabler, on ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i invoiceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i invoiceDo) RightJoin(table schema.Tabler, on ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i invoiceDo) Group(cols ...field.Expr) IInvoiceDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i invoiceDo) Having(conds ...gen.Condition) IInvoiceDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i invoiceDo) Limit(limit int) IInvoiceDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i invoiceDo) Offset(offset int) IInvoiceDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i invoiceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInvoiceDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i invoiceDo) Unscoped() IInvoiceDo {
	return i.withDO(i.DO.Unscoped())
}

func (i invoiceDo) Create(values ...*model.Invoice) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i invoiceDo) CreateInBatches(values []*model.Invoice, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i invoiceDo) Save(values ...*model.Invoice) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i invoiceDo) First() (*model.Invoice, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invoice), nil
	}
}

func (i invoiceDo) Take() (*model.Invoice, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invoice), nil
	}
}

func (i invoiceDo) Last() (*model.Invoice, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invoice), nil
	}
}

func (i invoiceDo) Find() ([]*model.Invoice, error) {
	result, err := i.DO.Find()
	return result.([]*model.Invoice), err
}

func (i invoiceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Invoice, err error) {
	buf := make([]*model.Invoice, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i invoiceDo) FindInBatches(result *[]*model.Invoice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i invoiceDo) Attrs(attrs ...field.AssignExpr) IInvoiceDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i invoiceDo) Assign(attrs ...field.AssignExpr) IInvoiceDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i invoiceDo) Joins(fields ...field.RelationField) IInvoiceDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i invoiceDo) Preload(fields ...field.RelationField) IInvoiceDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i invoiceDo) FirstOrInit() (*model.Invoice, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invoice), nil
	}
}

func (i invoiceDo) FirstOrCreate() (*model.Invoice, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Invoice), nil
	}
}

func (i invoiceDo) FindByPage(offset int, limit int) (result []*model.Invoice, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i invoiceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i invoiceDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i invoiceDo) Delete(models ...*model.Invoice) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *invoiceDo) withDO(do gen.Dao) *invoiceDo {
	i.DO = *do.(*gen.DO)
	return i
}
