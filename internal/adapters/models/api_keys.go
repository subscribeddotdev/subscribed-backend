// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// APIKey is an object representing the database table.
type APIKey struct {
	SecretKey     string    `boil:"secret_key" json:"secret_key" toml:"secret_key" yaml:"secret_key"`
	Suffix        string    `boil:"suffix" json:"suffix" toml:"suffix" yaml:"suffix"`
	EnvironmentID string    `boil:"environment_id" json:"environment_id" toml:"environment_id" yaml:"environment_id"`
	Name          string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt     time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ExpiresAt     null.Time `boil:"expires_at" json:"expires_at,omitempty" toml:"expires_at" yaml:"expires_at,omitempty"`

	R *apiKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L apiKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var APIKeyColumns = struct {
	SecretKey     string
	Suffix        string
	EnvironmentID string
	Name          string
	CreatedAt     string
	ExpiresAt     string
}{
	SecretKey:     "secret_key",
	Suffix:        "suffix",
	EnvironmentID: "environment_id",
	Name:          "name",
	CreatedAt:     "created_at",
	ExpiresAt:     "expires_at",
}

var APIKeyTableColumns = struct {
	SecretKey     string
	Suffix        string
	EnvironmentID string
	Name          string
	CreatedAt     string
	ExpiresAt     string
}{
	SecretKey:     "api_keys.secret_key",
	Suffix:        "api_keys.suffix",
	EnvironmentID: "api_keys.environment_id",
	Name:          "api_keys.name",
	CreatedAt:     "api_keys.created_at",
	ExpiresAt:     "api_keys.expires_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var APIKeyWhere = struct {
	SecretKey     whereHelperstring
	Suffix        whereHelperstring
	EnvironmentID whereHelperstring
	Name          whereHelperstring
	CreatedAt     whereHelpertime_Time
	ExpiresAt     whereHelpernull_Time
}{
	SecretKey:     whereHelperstring{field: "\"api_keys\".\"secret_key\""},
	Suffix:        whereHelperstring{field: "\"api_keys\".\"suffix\""},
	EnvironmentID: whereHelperstring{field: "\"api_keys\".\"environment_id\""},
	Name:          whereHelperstring{field: "\"api_keys\".\"name\""},
	CreatedAt:     whereHelpertime_Time{field: "\"api_keys\".\"created_at\""},
	ExpiresAt:     whereHelpernull_Time{field: "\"api_keys\".\"expires_at\""},
}

// APIKeyRels is where relationship names are stored.
var APIKeyRels = struct {
	Environment string
}{
	Environment: "Environment",
}

// apiKeyR is where relationships are stored.
type apiKeyR struct {
	Environment *Environment `boil:"Environment" json:"Environment" toml:"Environment" yaml:"Environment"`
}

// NewStruct creates a new relationship struct
func (*apiKeyR) NewStruct() *apiKeyR {
	return &apiKeyR{}
}

func (r *apiKeyR) GetEnvironment() *Environment {
	if r == nil {
		return nil
	}
	return r.Environment
}

// apiKeyL is where Load methods for each relationship are stored.
type apiKeyL struct{}

var (
	apiKeyAllColumns            = []string{"secret_key", "suffix", "environment_id", "name", "created_at", "expires_at"}
	apiKeyColumnsWithoutDefault = []string{"secret_key", "suffix", "environment_id", "name"}
	apiKeyColumnsWithDefault    = []string{"created_at", "expires_at"}
	apiKeyPrimaryKeyColumns     = []string{"secret_key"}
	apiKeyGeneratedColumns      = []string{}
)

type (
	// APIKeySlice is an alias for a slice of pointers to APIKey.
	// This should almost always be used instead of []APIKey.
	APIKeySlice []*APIKey
	// APIKeyHook is the signature for custom APIKey hook methods
	APIKeyHook func(context.Context, boil.ContextExecutor, *APIKey) error

	apiKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	apiKeyType                 = reflect.TypeOf(&APIKey{})
	apiKeyMapping              = queries.MakeStructMapping(apiKeyType)
	apiKeyPrimaryKeyMapping, _ = queries.BindMapping(apiKeyType, apiKeyMapping, apiKeyPrimaryKeyColumns)
	apiKeyInsertCacheMut       sync.RWMutex
	apiKeyInsertCache          = make(map[string]insertCache)
	apiKeyUpdateCacheMut       sync.RWMutex
	apiKeyUpdateCache          = make(map[string]updateCache)
	apiKeyUpsertCacheMut       sync.RWMutex
	apiKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var apiKeyAfterSelectHooks []APIKeyHook

var apiKeyBeforeInsertHooks []APIKeyHook
var apiKeyAfterInsertHooks []APIKeyHook

var apiKeyBeforeUpdateHooks []APIKeyHook
var apiKeyAfterUpdateHooks []APIKeyHook

var apiKeyBeforeDeleteHooks []APIKeyHook
var apiKeyAfterDeleteHooks []APIKeyHook

var apiKeyBeforeUpsertHooks []APIKeyHook
var apiKeyAfterUpsertHooks []APIKeyHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *APIKey) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *APIKey) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *APIKey) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *APIKey) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *APIKey) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *APIKey) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *APIKey) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *APIKey) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *APIKey) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range apiKeyAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAPIKeyHook registers your hook function for all future operations.
func AddAPIKeyHook(hookPoint boil.HookPoint, apiKeyHook APIKeyHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		apiKeyAfterSelectHooks = append(apiKeyAfterSelectHooks, apiKeyHook)
	case boil.BeforeInsertHook:
		apiKeyBeforeInsertHooks = append(apiKeyBeforeInsertHooks, apiKeyHook)
	case boil.AfterInsertHook:
		apiKeyAfterInsertHooks = append(apiKeyAfterInsertHooks, apiKeyHook)
	case boil.BeforeUpdateHook:
		apiKeyBeforeUpdateHooks = append(apiKeyBeforeUpdateHooks, apiKeyHook)
	case boil.AfterUpdateHook:
		apiKeyAfterUpdateHooks = append(apiKeyAfterUpdateHooks, apiKeyHook)
	case boil.BeforeDeleteHook:
		apiKeyBeforeDeleteHooks = append(apiKeyBeforeDeleteHooks, apiKeyHook)
	case boil.AfterDeleteHook:
		apiKeyAfterDeleteHooks = append(apiKeyAfterDeleteHooks, apiKeyHook)
	case boil.BeforeUpsertHook:
		apiKeyBeforeUpsertHooks = append(apiKeyBeforeUpsertHooks, apiKeyHook)
	case boil.AfterUpsertHook:
		apiKeyAfterUpsertHooks = append(apiKeyAfterUpsertHooks, apiKeyHook)
	}
}

// One returns a single apiKey record from the query.
func (q apiKeyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*APIKey, error) {
	o := &APIKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for api_keys")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all APIKey records from the query.
func (q apiKeyQuery) All(ctx context.Context, exec boil.ContextExecutor) (APIKeySlice, error) {
	var o []*APIKey

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to APIKey slice")
	}

	if len(apiKeyAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all APIKey records in the query.
func (q apiKeyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count api_keys rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q apiKeyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if api_keys exists")
	}

	return count > 0, nil
}

// Environment pointed to by the foreign key.
func (o *APIKey) Environment(mods ...qm.QueryMod) environmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EnvironmentID),
	}

	queryMods = append(queryMods, mods...)

	return Environments(queryMods...)
}

// LoadEnvironment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (apiKeyL) LoadEnvironment(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAPIKey interface{}, mods queries.Applicator) error {
	var slice []*APIKey
	var object *APIKey

	if singular {
		var ok bool
		object, ok = maybeAPIKey.(*APIKey)
		if !ok {
			object = new(APIKey)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAPIKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAPIKey))
			}
		}
	} else {
		s, ok := maybeAPIKey.(*[]*APIKey)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAPIKey)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAPIKey))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &apiKeyR{}
		}
		args = append(args, object.EnvironmentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &apiKeyR{}
			}

			for _, a := range args {
				if a == obj.EnvironmentID {
					continue Outer
				}
			}

			args = append(args, obj.EnvironmentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`environments`),
		qm.WhereIn(`environments.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Environment")
	}

	var resultSlice []*Environment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Environment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for environments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for environments")
	}

	if len(environmentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Environment = foreign
		if foreign.R == nil {
			foreign.R = &environmentR{}
		}
		foreign.R.APIKeys = append(foreign.R.APIKeys, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EnvironmentID == foreign.ID {
				local.R.Environment = foreign
				if foreign.R == nil {
					foreign.R = &environmentR{}
				}
				foreign.R.APIKeys = append(foreign.R.APIKeys, local)
				break
			}
		}
	}

	return nil
}

// SetEnvironment of the apiKey to the related item.
// Sets o.R.Environment to related.
// Adds o to related.R.APIKeys.
func (o *APIKey) SetEnvironment(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Environment) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"api_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"environment_id"}),
		strmangle.WhereClause("\"", "\"", 2, apiKeyPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SecretKey}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EnvironmentID = related.ID
	if o.R == nil {
		o.R = &apiKeyR{
			Environment: related,
		}
	} else {
		o.R.Environment = related
	}

	if related.R == nil {
		related.R = &environmentR{
			APIKeys: APIKeySlice{o},
		}
	} else {
		related.R.APIKeys = append(related.R.APIKeys, o)
	}

	return nil
}

// APIKeys retrieves all the records using an executor.
func APIKeys(mods ...qm.QueryMod) apiKeyQuery {
	mods = append(mods, qm.From("\"api_keys\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"api_keys\".*"})
	}

	return apiKeyQuery{q}
}

// FindAPIKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAPIKey(ctx context.Context, exec boil.ContextExecutor, secretKey string, selectCols ...string) (*APIKey, error) {
	apiKeyObj := &APIKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"api_keys\" where \"secret_key\"=$1", sel,
	)

	q := queries.Raw(query, secretKey)

	err := q.Bind(ctx, exec, apiKeyObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from api_keys")
	}

	if err = apiKeyObj.doAfterSelectHooks(ctx, exec); err != nil {
		return apiKeyObj, err
	}

	return apiKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *APIKey) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no api_keys provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(apiKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	apiKeyInsertCacheMut.RLock()
	cache, cached := apiKeyInsertCache[key]
	apiKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			apiKeyAllColumns,
			apiKeyColumnsWithDefault,
			apiKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(apiKeyType, apiKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(apiKeyType, apiKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"api_keys\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"api_keys\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into api_keys")
	}

	if !cached {
		apiKeyInsertCacheMut.Lock()
		apiKeyInsertCache[key] = cache
		apiKeyInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the APIKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *APIKey) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	apiKeyUpdateCacheMut.RLock()
	cache, cached := apiKeyUpdateCache[key]
	apiKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			apiKeyAllColumns,
			apiKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update api_keys, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"api_keys\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, apiKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(apiKeyType, apiKeyMapping, append(wl, apiKeyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update api_keys row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for api_keys")
	}

	if !cached {
		apiKeyUpdateCacheMut.Lock()
		apiKeyUpdateCache[key] = cache
		apiKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q apiKeyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for api_keys")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o APIKeySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apiKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"api_keys\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, apiKeyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in apiKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all apiKey")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *APIKey) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no api_keys provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(apiKeyColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	apiKeyUpsertCacheMut.RLock()
	cache, cached := apiKeyUpsertCache[key]
	apiKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			apiKeyAllColumns,
			apiKeyColumnsWithDefault,
			apiKeyColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			apiKeyAllColumns,
			apiKeyPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert api_keys, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(apiKeyPrimaryKeyColumns))
			copy(conflict, apiKeyPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"api_keys\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(apiKeyType, apiKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(apiKeyType, apiKeyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert api_keys")
	}

	if !cached {
		apiKeyUpsertCacheMut.Lock()
		apiKeyUpsertCache[key] = cache
		apiKeyUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single APIKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *APIKey) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no APIKey provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), apiKeyPrimaryKeyMapping)
	sql := "DELETE FROM \"api_keys\" WHERE \"secret_key\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for api_keys")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q apiKeyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no apiKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from api_keys")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for api_keys")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o APIKeySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(apiKeyBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apiKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"api_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, apiKeyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from apiKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for api_keys")
	}

	if len(apiKeyAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *APIKey) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAPIKey(ctx, exec, o.SecretKey)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *APIKeySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := APIKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), apiKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"api_keys\".* FROM \"api_keys\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, apiKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in APIKeySlice")
	}

	*o = slice

	return nil
}

// APIKeyExists checks if the APIKey row exists.
func APIKeyExists(ctx context.Context, exec boil.ContextExecutor, secretKey string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"api_keys\" where \"secret_key\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, secretKey)
	}
	row := exec.QueryRowContext(ctx, sql, secretKey)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if api_keys exists")
	}

	return exists, nil
}

// Exists checks if the APIKey row exists.
func (o *APIKey) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return APIKeyExists(ctx, exec, o.SecretKey)
}