// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/coze-dev/coze-studio/backend/domain/plugin/internal/dal/model"
)

func newToolVersion(db *gorm.DB, opts ...gen.DOOption) toolVersion {
	_toolVersion := toolVersion{}

	_toolVersion.toolVersionDo.UseDB(db, opts...)
	_toolVersion.toolVersionDo.UseModel(&model.ToolVersion{})

	tableName := _toolVersion.toolVersionDo.TableName()
	_toolVersion.ALL = field.NewAsterisk(tableName)
	_toolVersion.ID = field.NewInt64(tableName, "id")
	_toolVersion.ToolID = field.NewInt64(tableName, "tool_id")
	_toolVersion.PluginID = field.NewInt64(tableName, "plugin_id")
	_toolVersion.Version = field.NewString(tableName, "version")
	_toolVersion.SubURL = field.NewString(tableName, "sub_url")
	_toolVersion.Method = field.NewString(tableName, "method")
	_toolVersion.Operation = field.NewField(tableName, "operation")
	_toolVersion.CreatedAt = field.NewInt64(tableName, "created_at")
	_toolVersion.DeletedAt = field.NewField(tableName, "deleted_at")

	_toolVersion.fillFieldMap()

	return _toolVersion
}

// toolVersion Tool Version
type toolVersion struct {
	toolVersionDo

	ALL       field.Asterisk
	ID        field.Int64  // Primary Key ID
	ToolID    field.Int64  // Tool ID
	PluginID  field.Int64  // Plugin ID
	Version   field.String // Tool Version, e.g. v1.0.0
	SubURL    field.String // Sub URL Path
	Method    field.String // HTTP Request Method
	Operation field.Field  // Tool Openapi Operation Schema
	CreatedAt field.Int64  // Create Time in Milliseconds
	DeletedAt field.Field  // Delete Time

	fieldMap map[string]field.Expr
}

func (t toolVersion) Table(newTableName string) *toolVersion {
	t.toolVersionDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t toolVersion) As(alias string) *toolVersion {
	t.toolVersionDo.DO = *(t.toolVersionDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *toolVersion) updateTableName(table string) *toolVersion {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.ToolID = field.NewInt64(table, "tool_id")
	t.PluginID = field.NewInt64(table, "plugin_id")
	t.Version = field.NewString(table, "version")
	t.SubURL = field.NewString(table, "sub_url")
	t.Method = field.NewString(table, "method")
	t.Operation = field.NewField(table, "operation")
	t.CreatedAt = field.NewInt64(table, "created_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *toolVersion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *toolVersion) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.ID
	t.fieldMap["tool_id"] = t.ToolID
	t.fieldMap["plugin_id"] = t.PluginID
	t.fieldMap["version"] = t.Version
	t.fieldMap["sub_url"] = t.SubURL
	t.fieldMap["method"] = t.Method
	t.fieldMap["operation"] = t.Operation
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t toolVersion) clone(db *gorm.DB) toolVersion {
	t.toolVersionDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t toolVersion) replaceDB(db *gorm.DB) toolVersion {
	t.toolVersionDo.ReplaceDB(db)
	return t
}

type toolVersionDo struct{ gen.DO }

type IToolVersionDo interface {
	gen.SubQuery
	Debug() IToolVersionDo
	WithContext(ctx context.Context) IToolVersionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IToolVersionDo
	WriteDB() IToolVersionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IToolVersionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IToolVersionDo
	Not(conds ...gen.Condition) IToolVersionDo
	Or(conds ...gen.Condition) IToolVersionDo
	Select(conds ...field.Expr) IToolVersionDo
	Where(conds ...gen.Condition) IToolVersionDo
	Order(conds ...field.Expr) IToolVersionDo
	Distinct(cols ...field.Expr) IToolVersionDo
	Omit(cols ...field.Expr) IToolVersionDo
	Join(table schema.Tabler, on ...field.Expr) IToolVersionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IToolVersionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IToolVersionDo
	Group(cols ...field.Expr) IToolVersionDo
	Having(conds ...gen.Condition) IToolVersionDo
	Limit(limit int) IToolVersionDo
	Offset(offset int) IToolVersionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IToolVersionDo
	Unscoped() IToolVersionDo
	Create(values ...*model.ToolVersion) error
	CreateInBatches(values []*model.ToolVersion, batchSize int) error
	Save(values ...*model.ToolVersion) error
	First() (*model.ToolVersion, error)
	Take() (*model.ToolVersion, error)
	Last() (*model.ToolVersion, error)
	Find() ([]*model.ToolVersion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ToolVersion, err error)
	FindInBatches(result *[]*model.ToolVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ToolVersion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IToolVersionDo
	Assign(attrs ...field.AssignExpr) IToolVersionDo
	Joins(fields ...field.RelationField) IToolVersionDo
	Preload(fields ...field.RelationField) IToolVersionDo
	FirstOrInit() (*model.ToolVersion, error)
	FirstOrCreate() (*model.ToolVersion, error)
	FindByPage(offset int, limit int) (result []*model.ToolVersion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IToolVersionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t toolVersionDo) Debug() IToolVersionDo {
	return t.withDO(t.DO.Debug())
}

func (t toolVersionDo) WithContext(ctx context.Context) IToolVersionDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t toolVersionDo) ReadDB() IToolVersionDo {
	return t.Clauses(dbresolver.Read)
}

func (t toolVersionDo) WriteDB() IToolVersionDo {
	return t.Clauses(dbresolver.Write)
}

func (t toolVersionDo) Session(config *gorm.Session) IToolVersionDo {
	return t.withDO(t.DO.Session(config))
}

func (t toolVersionDo) Clauses(conds ...clause.Expression) IToolVersionDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t toolVersionDo) Returning(value interface{}, columns ...string) IToolVersionDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t toolVersionDo) Not(conds ...gen.Condition) IToolVersionDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t toolVersionDo) Or(conds ...gen.Condition) IToolVersionDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t toolVersionDo) Select(conds ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t toolVersionDo) Where(conds ...gen.Condition) IToolVersionDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t toolVersionDo) Order(conds ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t toolVersionDo) Distinct(cols ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t toolVersionDo) Omit(cols ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t toolVersionDo) Join(table schema.Tabler, on ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t toolVersionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t toolVersionDo) RightJoin(table schema.Tabler, on ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t toolVersionDo) Group(cols ...field.Expr) IToolVersionDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t toolVersionDo) Having(conds ...gen.Condition) IToolVersionDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t toolVersionDo) Limit(limit int) IToolVersionDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t toolVersionDo) Offset(offset int) IToolVersionDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t toolVersionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IToolVersionDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t toolVersionDo) Unscoped() IToolVersionDo {
	return t.withDO(t.DO.Unscoped())
}

func (t toolVersionDo) Create(values ...*model.ToolVersion) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t toolVersionDo) CreateInBatches(values []*model.ToolVersion, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t toolVersionDo) Save(values ...*model.ToolVersion) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t toolVersionDo) First() (*model.ToolVersion, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ToolVersion), nil
	}
}

func (t toolVersionDo) Take() (*model.ToolVersion, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ToolVersion), nil
	}
}

func (t toolVersionDo) Last() (*model.ToolVersion, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ToolVersion), nil
	}
}

func (t toolVersionDo) Find() ([]*model.ToolVersion, error) {
	result, err := t.DO.Find()
	return result.([]*model.ToolVersion), err
}

func (t toolVersionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ToolVersion, err error) {
	buf := make([]*model.ToolVersion, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t toolVersionDo) FindInBatches(result *[]*model.ToolVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t toolVersionDo) Attrs(attrs ...field.AssignExpr) IToolVersionDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t toolVersionDo) Assign(attrs ...field.AssignExpr) IToolVersionDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t toolVersionDo) Joins(fields ...field.RelationField) IToolVersionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t toolVersionDo) Preload(fields ...field.RelationField) IToolVersionDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t toolVersionDo) FirstOrInit() (*model.ToolVersion, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ToolVersion), nil
	}
}

func (t toolVersionDo) FirstOrCreate() (*model.ToolVersion, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ToolVersion), nil
	}
}

func (t toolVersionDo) FindByPage(offset int, limit int) (result []*model.ToolVersion, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t toolVersionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t toolVersionDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t toolVersionDo) Delete(models ...*model.ToolVersion) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *toolVersionDo) withDO(do gen.Dao) *toolVersionDo {
	t.DO = *do.(*gen.DO)
	return t
}
