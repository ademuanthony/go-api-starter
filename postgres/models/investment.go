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

// Investment is an object representing the database table.
type Investment struct {
	ID             string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountID      string `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	Amount         int64  `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Date           int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	ActivationDate int64  `boil:"activation_date" json:"activation_date" toml:"activation_date" yaml:"activation_date"`
	Status         int    `boil:"status" json:"status" toml:"status" yaml:"status"`

	R *investmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L investmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvestmentColumns = struct {
	ID             string
	AccountID      string
	Amount         string
	Date           string
	ActivationDate string
	Status         string
}{
	ID:             "id",
	AccountID:      "account_id",
	Amount:         "amount",
	Date:           "date",
	ActivationDate: "activation_date",
	Status:         "status",
}

var InvestmentTableColumns = struct {
	ID             string
	AccountID      string
	Amount         string
	Date           string
	ActivationDate string
	Status         string
}{
	ID:             "investment.id",
	AccountID:      "investment.account_id",
	Amount:         "investment.amount",
	Date:           "investment.date",
	ActivationDate: "investment.activation_date",
	Status:         "investment.status",
}

// Generated where

var InvestmentWhere = struct {
	ID             whereHelperstring
	AccountID      whereHelperstring
	Amount         whereHelperint64
	Date           whereHelperint64
	ActivationDate whereHelperint64
	Status         whereHelperint
}{
	ID:             whereHelperstring{field: "\"investment\".\"id\""},
	AccountID:      whereHelperstring{field: "\"investment\".\"account_id\""},
	Amount:         whereHelperint64{field: "\"investment\".\"amount\""},
	Date:           whereHelperint64{field: "\"investment\".\"date\""},
	ActivationDate: whereHelperint64{field: "\"investment\".\"activation_date\""},
	Status:         whereHelperint{field: "\"investment\".\"status\""},
}

// InvestmentRels is where relationship names are stored.
var InvestmentRels = struct {
	Account string
}{
	Account: "Account",
}

// investmentR is where relationships are stored.
type investmentR struct {
	Account *Account `boil:"Account" json:"Account" toml:"Account" yaml:"Account"`
}

// NewStruct creates a new relationship struct
func (*investmentR) NewStruct() *investmentR {
	return &investmentR{}
}

// investmentL is where Load methods for each relationship are stored.
type investmentL struct{}

var (
	investmentAllColumns            = []string{"id", "account_id", "amount", "date", "activation_date", "status"}
	investmentColumnsWithoutDefault = []string{"id", "account_id", "amount", "date", "activation_date"}
	investmentColumnsWithDefault    = []string{"status"}
	investmentPrimaryKeyColumns     = []string{"id"}
)

type (
	// InvestmentSlice is an alias for a slice of pointers to Investment.
	// This should almost always be used instead of []Investment.
	InvestmentSlice []*Investment

	investmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	investmentType                 = reflect.TypeOf(&Investment{})
	investmentMapping              = queries.MakeStructMapping(investmentType)
	investmentPrimaryKeyMapping, _ = queries.BindMapping(investmentType, investmentMapping, investmentPrimaryKeyColumns)
	investmentInsertCacheMut       sync.RWMutex
	investmentInsertCache          = make(map[string]insertCache)
	investmentUpdateCacheMut       sync.RWMutex
	investmentUpdateCache          = make(map[string]updateCache)
	investmentUpsertCacheMut       sync.RWMutex
	investmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single investment record from the query.
func (q investmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Investment, error) {
	o := &Investment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for investment")
	}

	return o, nil
}

// All returns all Investment records from the query.
func (q investmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvestmentSlice, error) {
	var o []*Investment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Investment slice")
	}

	return o, nil
}

// Count returns the count of all Investment records in the query.
func (q investmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count investment rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q investmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if investment exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *Investment) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"account\"")

	return query
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (investmentL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvestment interface{}, mods queries.Applicator) error {
	var slice []*Investment
	var object *Investment

	if singular {
		object = maybeInvestment.(*Investment)
	} else {
		slice = *maybeInvestment.(*[]*Investment)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &investmentR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &investmentR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`account`),
		qm.WhereIn(`account.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for account")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Account = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.Investments = append(foreign.R.Investments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.ID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.Investments = append(foreign.R.Investments, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the investment to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.Investments.
func (o *Investment) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"investment\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 2, investmentPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.ID
	if o.R == nil {
		o.R = &investmentR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			Investments: InvestmentSlice{o},
		}
	} else {
		related.R.Investments = append(related.R.Investments, o)
	}

	return nil
}

// Investments retrieves all the records using an executor.
func Investments(mods ...qm.QueryMod) investmentQuery {
	mods = append(mods, qm.From("\"investment\""))
	return investmentQuery{NewQuery(mods...)}
}

// FindInvestment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvestment(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Investment, error) {
	investmentObj := &Investment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"investment\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, investmentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from investment")
	}

	return investmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Investment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no investment provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(investmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	investmentInsertCacheMut.RLock()
	cache, cached := investmentInsertCache[key]
	investmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			investmentAllColumns,
			investmentColumnsWithDefault,
			investmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(investmentType, investmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(investmentType, investmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"investment\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"investment\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into investment")
	}

	if !cached {
		investmentInsertCacheMut.Lock()
		investmentInsertCache[key] = cache
		investmentInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Investment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Investment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	investmentUpdateCacheMut.RLock()
	cache, cached := investmentUpdateCache[key]
	investmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			investmentAllColumns,
			investmentPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update investment, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"investment\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, investmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(investmentType, investmentMapping, append(wl, investmentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update investment row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for investment")
	}

	if !cached {
		investmentUpdateCacheMut.Lock()
		investmentUpdateCache[key] = cache
		investmentUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q investmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for investment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for investment")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvestmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"investment\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, investmentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in investment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all investment")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Investment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no investment provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(investmentColumnsWithDefault, o)

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

	investmentUpsertCacheMut.RLock()
	cache, cached := investmentUpsertCache[key]
	investmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			investmentAllColumns,
			investmentColumnsWithDefault,
			investmentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			investmentAllColumns,
			investmentPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert investment, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(investmentPrimaryKeyColumns))
			copy(conflict, investmentPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"investment\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(investmentType, investmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(investmentType, investmentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert investment")
	}

	if !cached {
		investmentUpsertCacheMut.Lock()
		investmentUpsertCache[key] = cache
		investmentUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Investment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Investment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Investment provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), investmentPrimaryKeyMapping)
	sql := "DELETE FROM \"investment\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from investment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for investment")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q investmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no investmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from investment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for investment")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvestmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"investment\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, investmentPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from investment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for investment")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Investment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvestment(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvestmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvestmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"investment\".* FROM \"investment\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, investmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvestmentSlice")
	}

	*o = slice

	return nil
}

// InvestmentExists checks if the Investment row exists.
func InvestmentExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"investment\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if investment exists")
	}

	return exists, nil
}