// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// AllowList is an object representing the database table.
type AllowList struct {
	ID     int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	PubKey roomdb.DBFeedRef `boil:"pub_key" json:"pub_key" toml:"pub_key" yaml:"pub_key"`

	R *allowListR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L allowListL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AllowListColumns = struct {
	ID     string
	PubKey string
}{
	ID:     "id",
	PubKey: "pub_key",
}

// Generated where

type whereHelperroomdb_DBFeedRef struct{ field string }

func (w whereHelperroomdb_DBFeedRef) EQ(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperroomdb_DBFeedRef) NEQ(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperroomdb_DBFeedRef) LT(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperroomdb_DBFeedRef) LTE(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperroomdb_DBFeedRef) GT(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperroomdb_DBFeedRef) GTE(x roomdb.DBFeedRef) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AllowListWhere = struct {
	ID     whereHelperint64
	PubKey whereHelperroomdb_DBFeedRef
}{
	ID:     whereHelperint64{field: "\"allow_list\".\"id\""},
	PubKey: whereHelperroomdb_DBFeedRef{field: "\"allow_list\".\"pub_key\""},
}

// AllowListRels is where relationship names are stored.
var AllowListRels = struct {
	UserAliases string
}{
	UserAliases: "UserAliases",
}

// allowListR is where relationships are stored.
type allowListR struct {
	UserAliases AliasSlice `boil:"UserAliases" json:"UserAliases" toml:"UserAliases" yaml:"UserAliases"`
}

// NewStruct creates a new relationship struct
func (*allowListR) NewStruct() *allowListR {
	return &allowListR{}
}

// allowListL is where Load methods for each relationship are stored.
type allowListL struct{}

var (
	allowListAllColumns            = []string{"id", "pub_key"}
	allowListColumnsWithoutDefault = []string{}
	allowListColumnsWithDefault    = []string{"id", "pub_key"}
	allowListPrimaryKeyColumns     = []string{"id"}
)

type (
	// AllowListSlice is an alias for a slice of pointers to AllowList.
	// This should generally be used opposed to []AllowList.
	AllowListSlice []*AllowList
	// AllowListHook is the signature for custom AllowList hook methods
	AllowListHook func(context.Context, boil.ContextExecutor, *AllowList) error

	allowListQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	allowListType                 = reflect.TypeOf(&AllowList{})
	allowListMapping              = queries.MakeStructMapping(allowListType)
	allowListPrimaryKeyMapping, _ = queries.BindMapping(allowListType, allowListMapping, allowListPrimaryKeyColumns)
	allowListInsertCacheMut       sync.RWMutex
	allowListInsertCache          = make(map[string]insertCache)
	allowListUpdateCacheMut       sync.RWMutex
	allowListUpdateCache          = make(map[string]updateCache)
	allowListUpsertCacheMut       sync.RWMutex
	allowListUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var allowListBeforeInsertHooks []AllowListHook
var allowListBeforeUpdateHooks []AllowListHook
var allowListBeforeDeleteHooks []AllowListHook
var allowListBeforeUpsertHooks []AllowListHook

var allowListAfterInsertHooks []AllowListHook
var allowListAfterSelectHooks []AllowListHook
var allowListAfterUpdateHooks []AllowListHook
var allowListAfterDeleteHooks []AllowListHook
var allowListAfterUpsertHooks []AllowListHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AllowList) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AllowList) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AllowList) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AllowList) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AllowList) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AllowList) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AllowList) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AllowList) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AllowList) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range allowListAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAllowListHook registers your hook function for all future operations.
func AddAllowListHook(hookPoint boil.HookPoint, allowListHook AllowListHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		allowListBeforeInsertHooks = append(allowListBeforeInsertHooks, allowListHook)
	case boil.BeforeUpdateHook:
		allowListBeforeUpdateHooks = append(allowListBeforeUpdateHooks, allowListHook)
	case boil.BeforeDeleteHook:
		allowListBeforeDeleteHooks = append(allowListBeforeDeleteHooks, allowListHook)
	case boil.BeforeUpsertHook:
		allowListBeforeUpsertHooks = append(allowListBeforeUpsertHooks, allowListHook)
	case boil.AfterInsertHook:
		allowListAfterInsertHooks = append(allowListAfterInsertHooks, allowListHook)
	case boil.AfterSelectHook:
		allowListAfterSelectHooks = append(allowListAfterSelectHooks, allowListHook)
	case boil.AfterUpdateHook:
		allowListAfterUpdateHooks = append(allowListAfterUpdateHooks, allowListHook)
	case boil.AfterDeleteHook:
		allowListAfterDeleteHooks = append(allowListAfterDeleteHooks, allowListHook)
	case boil.AfterUpsertHook:
		allowListAfterUpsertHooks = append(allowListAfterUpsertHooks, allowListHook)
	}
}

// One returns a single allowList record from the query.
func (q allowListQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AllowList, error) {
	o := &AllowList{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for allow_list")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AllowList records from the query.
func (q allowListQuery) All(ctx context.Context, exec boil.ContextExecutor) (AllowListSlice, error) {
	var o []*AllowList

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AllowList slice")
	}

	if len(allowListAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AllowList records in the query.
func (q allowListQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count allow_list rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q allowListQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if allow_list exists")
	}

	return count > 0, nil
}

// UserAliases retrieves all the alias's Aliases with an executor via user_id column.
func (o *AllowList) UserAliases(mods ...qm.QueryMod) aliasQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"aliases\".\"user_id\"=?", o.ID),
	)

	query := Aliases(queryMods...)
	queries.SetFrom(query.Query, "\"aliases\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"aliases\".*"})
	}

	return query
}

// LoadUserAliases allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (allowListL) LoadUserAliases(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAllowList interface{}, mods queries.Applicator) error {
	var slice []*AllowList
	var object *AllowList

	if singular {
		object = maybeAllowList.(*AllowList)
	} else {
		slice = *maybeAllowList.(*[]*AllowList)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &allowListR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &allowListR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`aliases`),
		qm.WhereIn(`aliases.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load aliases")
	}

	var resultSlice []*Alias
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice aliases")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on aliases")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for aliases")
	}

	if len(aliasAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserAliases = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &aliasR{}
			}
			foreign.R.User = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserID {
				local.R.UserAliases = append(local.R.UserAliases, foreign)
				if foreign.R == nil {
					foreign.R = &aliasR{}
				}
				foreign.R.User = local
				break
			}
		}
	}

	return nil
}

// AddUserAliases adds the given related objects to the existing relationships
// of the allow_list, optionally inserting them as new records.
// Appends related to o.R.UserAliases.
// Sets related.R.User appropriately.
func (o *AllowList) AddUserAliases(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Alias) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"aliases\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"user_id"}),
				strmangle.WhereClause("\"", "\"", 0, aliasPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.UserID = o.ID
		}
	}

	if o.R == nil {
		o.R = &allowListR{
			UserAliases: related,
		}
	} else {
		o.R.UserAliases = append(o.R.UserAliases, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &aliasR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// AllowLists retrieves all the records using an executor.
func AllowLists(mods ...qm.QueryMod) allowListQuery {
	mods = append(mods, qm.From("\"allow_list\""))
	return allowListQuery{NewQuery(mods...)}
}

// FindAllowList retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAllowList(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AllowList, error) {
	allowListObj := &AllowList{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"allow_list\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, allowListObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from allow_list")
	}

	return allowListObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AllowList) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no allow_list provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(allowListColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	allowListInsertCacheMut.RLock()
	cache, cached := allowListInsertCache[key]
	allowListInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			allowListAllColumns,
			allowListColumnsWithDefault,
			allowListColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(allowListType, allowListMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(allowListType, allowListMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"allow_list\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"allow_list\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"allow_list\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, allowListPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into allow_list")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == allowListMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for allow_list")
	}

CacheNoHooks:
	if !cached {
		allowListInsertCacheMut.Lock()
		allowListInsertCache[key] = cache
		allowListInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AllowList.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AllowList) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	allowListUpdateCacheMut.RLock()
	cache, cached := allowListUpdateCache[key]
	allowListUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			allowListAllColumns,
			allowListPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update allow_list, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"allow_list\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, allowListPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(allowListType, allowListMapping, append(wl, allowListPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update allow_list row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for allow_list")
	}

	if !cached {
		allowListUpdateCacheMut.Lock()
		allowListUpdateCache[key] = cache
		allowListUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q allowListQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for allow_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for allow_list")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AllowListSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allowListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"allow_list\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, allowListPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in allowList slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all allowList")
	}
	return rowsAff, nil
}

// Delete deletes a single AllowList record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AllowList) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AllowList provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), allowListPrimaryKeyMapping)
	sql := "DELETE FROM \"allow_list\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from allow_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for allow_list")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q allowListQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no allowListQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from allow_list")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for allow_list")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AllowListSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(allowListBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allowListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"allow_list\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, allowListPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from allowList slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for allow_list")
	}

	if len(allowListAfterDeleteHooks) != 0 {
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
func (o *AllowList) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAllowList(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AllowListSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AllowListSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), allowListPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"allow_list\".* FROM \"allow_list\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, allowListPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AllowListSlice")
	}

	*o = slice

	return nil
}

// AllowListExists checks if the AllowList row exists.
func AllowListExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"allow_list\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if allow_list exists")
	}

	return exists, nil
}
