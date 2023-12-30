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

	"pjm.dev/chinook/internal/db/model"
)

func newMediaType(db *gorm.DB, opts ...gen.DOOption) mediaType {
	_mediaType := mediaType{}

	_mediaType.mediaTypeDo.UseDB(db, opts...)
	_mediaType.mediaTypeDo.UseModel(&model.MediaType{})

	tableName := _mediaType.mediaTypeDo.TableName()
	_mediaType.ALL = field.NewAsterisk(tableName)
	_mediaType.MediaTypeID = field.NewInt32(tableName, "MediaTypeId")
	_mediaType.Name = field.NewString(tableName, "Name")

	_mediaType.fillFieldMap()

	return _mediaType
}

type mediaType struct {
	mediaTypeDo

	ALL         field.Asterisk
	MediaTypeID field.Int32
	Name        field.String

	fieldMap map[string]field.Expr
}

func (m mediaType) Table(newTableName string) *mediaType {
	m.mediaTypeDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m mediaType) As(alias string) *mediaType {
	m.mediaTypeDo.DO = *(m.mediaTypeDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *mediaType) updateTableName(table string) *mediaType {
	m.ALL = field.NewAsterisk(table)
	m.MediaTypeID = field.NewInt32(table, "MediaTypeId")
	m.Name = field.NewString(table, "Name")

	m.fillFieldMap()

	return m
}

func (m *mediaType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *mediaType) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 2)
	m.fieldMap["MediaTypeId"] = m.MediaTypeID
	m.fieldMap["Name"] = m.Name
}

func (m mediaType) clone(db *gorm.DB) mediaType {
	m.mediaTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m mediaType) replaceDB(db *gorm.DB) mediaType {
	m.mediaTypeDo.ReplaceDB(db)
	return m
}

type mediaTypeDo struct{ gen.DO }

type IMediaTypeDo interface {
	gen.SubQuery
	Debug() IMediaTypeDo
	WithContext(ctx context.Context) IMediaTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMediaTypeDo
	WriteDB() IMediaTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMediaTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMediaTypeDo
	Not(conds ...gen.Condition) IMediaTypeDo
	Or(conds ...gen.Condition) IMediaTypeDo
	Select(conds ...field.Expr) IMediaTypeDo
	Where(conds ...gen.Condition) IMediaTypeDo
	Order(conds ...field.Expr) IMediaTypeDo
	Distinct(cols ...field.Expr) IMediaTypeDo
	Omit(cols ...field.Expr) IMediaTypeDo
	Join(table schema.Tabler, on ...field.Expr) IMediaTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMediaTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMediaTypeDo
	Group(cols ...field.Expr) IMediaTypeDo
	Having(conds ...gen.Condition) IMediaTypeDo
	Limit(limit int) IMediaTypeDo
	Offset(offset int) IMediaTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMediaTypeDo
	Unscoped() IMediaTypeDo
	Create(values ...*model.MediaType) error
	CreateInBatches(values []*model.MediaType, batchSize int) error
	Save(values ...*model.MediaType) error
	First() (*model.MediaType, error)
	Take() (*model.MediaType, error)
	Last() (*model.MediaType, error)
	Find() ([]*model.MediaType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MediaType, err error)
	FindInBatches(result *[]*model.MediaType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MediaType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMediaTypeDo
	Assign(attrs ...field.AssignExpr) IMediaTypeDo
	Joins(fields ...field.RelationField) IMediaTypeDo
	Preload(fields ...field.RelationField) IMediaTypeDo
	FirstOrInit() (*model.MediaType, error)
	FirstOrCreate() (*model.MediaType, error)
	FindByPage(offset int, limit int) (result []*model.MediaType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMediaTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m mediaTypeDo) Debug() IMediaTypeDo {
	return m.withDO(m.DO.Debug())
}

func (m mediaTypeDo) WithContext(ctx context.Context) IMediaTypeDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m mediaTypeDo) ReadDB() IMediaTypeDo {
	return m.Clauses(dbresolver.Read)
}

func (m mediaTypeDo) WriteDB() IMediaTypeDo {
	return m.Clauses(dbresolver.Write)
}

func (m mediaTypeDo) Session(config *gorm.Session) IMediaTypeDo {
	return m.withDO(m.DO.Session(config))
}

func (m mediaTypeDo) Clauses(conds ...clause.Expression) IMediaTypeDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m mediaTypeDo) Returning(value interface{}, columns ...string) IMediaTypeDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m mediaTypeDo) Not(conds ...gen.Condition) IMediaTypeDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m mediaTypeDo) Or(conds ...gen.Condition) IMediaTypeDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m mediaTypeDo) Select(conds ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m mediaTypeDo) Where(conds ...gen.Condition) IMediaTypeDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m mediaTypeDo) Order(conds ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m mediaTypeDo) Distinct(cols ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m mediaTypeDo) Omit(cols ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m mediaTypeDo) Join(table schema.Tabler, on ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m mediaTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m mediaTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m mediaTypeDo) Group(cols ...field.Expr) IMediaTypeDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m mediaTypeDo) Having(conds ...gen.Condition) IMediaTypeDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m mediaTypeDo) Limit(limit int) IMediaTypeDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m mediaTypeDo) Offset(offset int) IMediaTypeDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m mediaTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMediaTypeDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m mediaTypeDo) Unscoped() IMediaTypeDo {
	return m.withDO(m.DO.Unscoped())
}

func (m mediaTypeDo) Create(values ...*model.MediaType) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m mediaTypeDo) CreateInBatches(values []*model.MediaType, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m mediaTypeDo) Save(values ...*model.MediaType) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m mediaTypeDo) First() (*model.MediaType, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MediaType), nil
	}
}

func (m mediaTypeDo) Take() (*model.MediaType, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MediaType), nil
	}
}

func (m mediaTypeDo) Last() (*model.MediaType, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MediaType), nil
	}
}

func (m mediaTypeDo) Find() ([]*model.MediaType, error) {
	result, err := m.DO.Find()
	return result.([]*model.MediaType), err
}

func (m mediaTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MediaType, err error) {
	buf := make([]*model.MediaType, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m mediaTypeDo) FindInBatches(result *[]*model.MediaType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m mediaTypeDo) Attrs(attrs ...field.AssignExpr) IMediaTypeDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m mediaTypeDo) Assign(attrs ...field.AssignExpr) IMediaTypeDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m mediaTypeDo) Joins(fields ...field.RelationField) IMediaTypeDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m mediaTypeDo) Preload(fields ...field.RelationField) IMediaTypeDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m mediaTypeDo) FirstOrInit() (*model.MediaType, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MediaType), nil
	}
}

func (m mediaTypeDo) FirstOrCreate() (*model.MediaType, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MediaType), nil
	}
}

func (m mediaTypeDo) FindByPage(offset int, limit int) (result []*model.MediaType, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m mediaTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m mediaTypeDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m mediaTypeDo) Delete(models ...*model.MediaType) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *mediaTypeDo) withDO(do gen.Dao) *mediaTypeDo {
	m.DO = *do.(*gen.DO)
	return m
}
