// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CgoldHolder is an object representing the database table.
type CgoldHolder struct {
	WalletAddress string `boil:"wallet_address" json:"wallet_address" toml:"wallet_address" yaml:"wallet_address"`
	Balance       string `boil:"balance" json:"balance" toml:"balance" yaml:"balance"`

	R *cgoldHolderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cgoldHolderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CgoldHolderColumns = struct {
	WalletAddress string
	Balance       string
}{
	WalletAddress: "wallet_address",
	Balance:       "balance",
}

var CgoldHolderTableColumns = struct {
	WalletAddress string
	Balance       string
}{
	WalletAddress: "cgold_holders.wallet_address",
	Balance:       "cgold_holders.balance",
}

// Generated where

var CgoldHolderWhere = struct {
	WalletAddress whereHelperstring
	Balance       whereHelperstring
}{
	WalletAddress: whereHelperstring{field: "\"cgold_holders\".\"wallet_address\""},
	Balance:       whereHelperstring{field: "\"cgold_holders\".\"balance\""},
}

// CgoldHolderRels is where relationship names are stored.
var CgoldHolderRels = struct {
}{}

// cgoldHolderR is where relationships are stored.
type cgoldHolderR struct {
}

// NewStruct creates a new relationship struct
func (*cgoldHolderR) NewStruct() *cgoldHolderR {
	return &cgoldHolderR{}
}

// cgoldHolderL is where Load methods for each relationship are stored.
type cgoldHolderL struct{}

var (
	cgoldHolderAllColumns            = []string{"wallet_address", "balance"}
	cgoldHolderColumnsWithoutDefault = []string{"wallet_address", "balance"}
	cgoldHolderColumnsWithDefault    = []string{}
	cgoldHolderPrimaryKeyColumns     = []string{"wallet_address"}
)

type (
	// CgoldHolderSlice is an alias for a slice of pointers to CgoldHolder.
	// This should almost always be used instead of []CgoldHolder.
	CgoldHolderSlice []*CgoldHolder

	cgoldHolderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cgoldHolderType                 = reflect.TypeOf(&CgoldHolder{})
	cgoldHolderMapping              = queries.MakeStructMapping(cgoldHolderType)
	cgoldHolderPrimaryKeyMapping, _ = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, cgoldHolderPrimaryKeyColumns)
	cgoldHolderInsertCacheMut       sync.RWMutex
	cgoldHolderInsertCache          = make(map[string]insertCache)
	cgoldHolderUpdateCacheMut       sync.RWMutex
	cgoldHolderUpdateCache          = make(map[string]updateCache)
	cgoldHolderUpsertCacheMut       sync.RWMutex
	cgoldHolderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single cgoldHolder record from the query.
func (q cgoldHolderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CgoldHolder, error) {
	o := &CgoldHolder{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cgold_holders")
	}

	return o, nil
}

// All returns all CgoldHolder records from the query.
func (q cgoldHolderQuery) All(ctx context.Context, exec boil.ContextExecutor) (CgoldHolderSlice, error) {
	var o []*CgoldHolder

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CgoldHolder slice")
	}

	return o, nil
}

// Count returns the count of all CgoldHolder records in the query.
func (q cgoldHolderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cgold_holders rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q cgoldHolderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cgold_holders exists")
	}

	return count > 0, nil
}

// CgoldHolders retrieves all the records using an executor.
func CgoldHolders(mods ...qm.QueryMod) cgoldHolderQuery {
	mods = append(mods, qm.From("\"cgold_holders\""))
	return cgoldHolderQuery{NewQuery(mods...)}
}

// FindCgoldHolder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCgoldHolder(ctx context.Context, exec boil.ContextExecutor, walletAddress string, selectCols ...string) (*CgoldHolder, error) {
	cgoldHolderObj := &CgoldHolder{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cgold_holders\" where \"wallet_address\"=$1", sel,
	)

	q := queries.Raw(query, walletAddress)

	err := q.Bind(ctx, exec, cgoldHolderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cgold_holders")
	}

	return cgoldHolderObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CgoldHolder) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cgold_holders provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(cgoldHolderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	cgoldHolderInsertCacheMut.RLock()
	cache, cached := cgoldHolderInsertCache[key]
	cgoldHolderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			cgoldHolderAllColumns,
			cgoldHolderColumnsWithDefault,
			cgoldHolderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"cgold_holders\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"cgold_holders\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into cgold_holders")
	}

	if !cached {
		cgoldHolderInsertCacheMut.Lock()
		cgoldHolderInsertCache[key] = cache
		cgoldHolderInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CgoldHolder.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CgoldHolder) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	cgoldHolderUpdateCacheMut.RLock()
	cache, cached := cgoldHolderUpdateCache[key]
	cgoldHolderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			cgoldHolderAllColumns,
			cgoldHolderPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update cgold_holders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cgold_holders\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cgoldHolderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, append(wl, cgoldHolderPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update cgold_holders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for cgold_holders")
	}

	if !cached {
		cgoldHolderUpdateCacheMut.Lock()
		cgoldHolderUpdateCache[key] = cache
		cgoldHolderUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q cgoldHolderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for cgold_holders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for cgold_holders")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CgoldHolderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cgoldHolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"cgold_holders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, cgoldHolderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in cgoldHolder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all cgoldHolder")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CgoldHolder) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no cgold_holders provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(cgoldHolderColumnsWithDefault, o)

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

	cgoldHolderUpsertCacheMut.RLock()
	cache, cached := cgoldHolderUpsertCache[key]
	cgoldHolderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			cgoldHolderAllColumns,
			cgoldHolderColumnsWithDefault,
			cgoldHolderColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			cgoldHolderAllColumns,
			cgoldHolderPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert cgold_holders, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cgoldHolderPrimaryKeyColumns))
			copy(conflict, cgoldHolderPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"cgold_holders\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cgoldHolderType, cgoldHolderMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert cgold_holders")
	}

	if !cached {
		cgoldHolderUpsertCacheMut.Lock()
		cgoldHolderUpsertCache[key] = cache
		cgoldHolderUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CgoldHolder record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CgoldHolder) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CgoldHolder provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cgoldHolderPrimaryKeyMapping)
	sql := "DELETE FROM \"cgold_holders\" WHERE \"wallet_address\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from cgold_holders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for cgold_holders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q cgoldHolderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no cgoldHolderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cgold_holders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cgold_holders")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CgoldHolderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cgoldHolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"cgold_holders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cgoldHolderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from cgoldHolder slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for cgold_holders")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CgoldHolder) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCgoldHolder(ctx, exec, o.WalletAddress)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CgoldHolderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CgoldHolderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cgoldHolderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"cgold_holders\".* FROM \"cgold_holders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cgoldHolderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CgoldHolderSlice")
	}

	*o = slice

	return nil
}

// CgoldHolderExists checks if the CgoldHolder row exists.
func CgoldHolderExists(ctx context.Context, exec boil.ContextExecutor, walletAddress string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"cgold_holders\" where \"wallet_address\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, walletAddress)
	}
	row := exec.QueryRowContext(ctx, sql, walletAddress)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cgold_holders exists")
	}

	return exists, nil
}
