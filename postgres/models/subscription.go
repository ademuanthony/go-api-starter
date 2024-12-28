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

// Subscription is an object representing the database table.
type Subscription struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	AccountID string `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	PackageID string `boil:"package_id" json:"package_id" toml:"package_id" yaml:"package_id"`
	StartDate int64  `boil:"start_date" json:"start_date" toml:"start_date" yaml:"start_date"`
	EndDate   int64  `boil:"end_date" json:"end_date" toml:"end_date" yaml:"end_date"`

	R *subscriptionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L subscriptionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SubscriptionColumns = struct {
	ID        string
	AccountID string
	PackageID string
	StartDate string
	EndDate   string
}{
	ID:        "id",
	AccountID: "account_id",
	PackageID: "package_id",
	StartDate: "start_date",
	EndDate:   "end_date",
}

var SubscriptionTableColumns = struct {
	ID        string
	AccountID string
	PackageID string
	StartDate string
	EndDate   string
}{
	ID:        "subscription.id",
	AccountID: "subscription.account_id",
	PackageID: "subscription.package_id",
	StartDate: "subscription.start_date",
	EndDate:   "subscription.end_date",
}

// Generated where

var SubscriptionWhere = struct {
	ID        whereHelperstring
	AccountID whereHelperstring
	PackageID whereHelperstring
	StartDate whereHelperint64
	EndDate   whereHelperint64
}{
	ID:        whereHelperstring{field: "\"subscription\".\"id\""},
	AccountID: whereHelperstring{field: "\"subscription\".\"account_id\""},
	PackageID: whereHelperstring{field: "\"subscription\".\"package_id\""},
	StartDate: whereHelperint64{field: "\"subscription\".\"start_date\""},
	EndDate:   whereHelperint64{field: "\"subscription\".\"end_date\""},
}

// SubscriptionRels is where relationship names are stored.
var SubscriptionRels = struct {
	Account         string
	Package         string
	ReferralPayouts string
}{
	Account:         "Account",
	Package:         "Package",
	ReferralPayouts: "ReferralPayouts",
}

// subscriptionR is where relationships are stored.
type subscriptionR struct {
	Account         *Account            `boil:"Account" json:"Account" toml:"Account" yaml:"Account"`
	Package         *Package            `boil:"Package" json:"Package" toml:"Package" yaml:"Package"`
	ReferralPayouts ReferralPayoutSlice `boil:"ReferralPayouts" json:"ReferralPayouts" toml:"ReferralPayouts" yaml:"ReferralPayouts"`
}

// NewStruct creates a new relationship struct
func (*subscriptionR) NewStruct() *subscriptionR {
	return &subscriptionR{}
}

// subscriptionL is where Load methods for each relationship are stored.
type subscriptionL struct{}

var (
	subscriptionAllColumns            = []string{"id", "account_id", "package_id", "start_date", "end_date"}
	subscriptionColumnsWithoutDefault = []string{"id", "account_id", "package_id", "start_date", "end_date"}
	subscriptionColumnsWithDefault    = []string{}
	subscriptionPrimaryKeyColumns     = []string{"id"}
)

type (
	// SubscriptionSlice is an alias for a slice of pointers to Subscription.
	// This should almost always be used instead of []Subscription.
	SubscriptionSlice []*Subscription

	subscriptionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	subscriptionType                 = reflect.TypeOf(&Subscription{})
	subscriptionMapping              = queries.MakeStructMapping(subscriptionType)
	subscriptionPrimaryKeyMapping, _ = queries.BindMapping(subscriptionType, subscriptionMapping, subscriptionPrimaryKeyColumns)
	subscriptionInsertCacheMut       sync.RWMutex
	subscriptionInsertCache          = make(map[string]insertCache)
	subscriptionUpdateCacheMut       sync.RWMutex
	subscriptionUpdateCache          = make(map[string]updateCache)
	subscriptionUpsertCacheMut       sync.RWMutex
	subscriptionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single subscription record from the query.
func (q subscriptionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Subscription, error) {
	o := &Subscription{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for subscription")
	}

	return o, nil
}

// All returns all Subscription records from the query.
func (q subscriptionQuery) All(ctx context.Context, exec boil.ContextExecutor) (SubscriptionSlice, error) {
	var o []*Subscription

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Subscription slice")
	}

	return o, nil
}

// Count returns the count of all Subscription records in the query.
func (q subscriptionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count subscription rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q subscriptionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if subscription exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *Subscription) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "\"account\"")

	return query
}

// Package pointed to by the foreign key.
func (o *Subscription) Package(mods ...qm.QueryMod) packageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PackageID),
	}

	queryMods = append(queryMods, mods...)

	query := Packages(queryMods...)
	queries.SetFrom(query.Query, "\"package\"")

	return query
}

// ReferralPayouts retrieves all the referral_payout's ReferralPayouts with an executor.
func (o *Subscription) ReferralPayouts(mods ...qm.QueryMod) referralPayoutQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"referral_payout\".\"subscription_id\"=?", o.ID),
	)

	query := ReferralPayouts(queryMods...)
	queries.SetFrom(query.Query, "\"referral_payout\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"referral_payout\".*"})
	}

	return query
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (subscriptionL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubscription interface{}, mods queries.Applicator) error {
	var slice []*Subscription
	var object *Subscription

	if singular {
		object = maybeSubscription.(*Subscription)
	} else {
		slice = *maybeSubscription.(*[]*Subscription)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subscriptionR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subscriptionR{}
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
		foreign.R.Subscriptions = append(foreign.R.Subscriptions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.ID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.Subscriptions = append(foreign.R.Subscriptions, local)
				break
			}
		}
	}

	return nil
}

// LoadPackage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (subscriptionL) LoadPackage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubscription interface{}, mods queries.Applicator) error {
	var slice []*Subscription
	var object *Subscription

	if singular {
		object = maybeSubscription.(*Subscription)
	} else {
		slice = *maybeSubscription.(*[]*Subscription)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subscriptionR{}
		}
		args = append(args, object.PackageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subscriptionR{}
			}

			for _, a := range args {
				if a == obj.PackageID {
					continue Outer
				}
			}

			args = append(args, obj.PackageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`package`),
		qm.WhereIn(`package.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Package")
	}

	var resultSlice []*Package
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Package")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for package")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for package")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Package = foreign
		if foreign.R == nil {
			foreign.R = &packageR{}
		}
		foreign.R.Subscriptions = append(foreign.R.Subscriptions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PackageID == foreign.ID {
				local.R.Package = foreign
				if foreign.R == nil {
					foreign.R = &packageR{}
				}
				foreign.R.Subscriptions = append(foreign.R.Subscriptions, local)
				break
			}
		}
	}

	return nil
}

// LoadReferralPayouts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (subscriptionL) LoadReferralPayouts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSubscription interface{}, mods queries.Applicator) error {
	var slice []*Subscription
	var object *Subscription

	if singular {
		object = maybeSubscription.(*Subscription)
	} else {
		slice = *maybeSubscription.(*[]*Subscription)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &subscriptionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &subscriptionR{}
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
		qm.From(`referral_payout`),
		qm.WhereIn(`referral_payout.subscription_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load referral_payout")
	}

	var resultSlice []*ReferralPayout
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice referral_payout")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on referral_payout")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for referral_payout")
	}

	if singular {
		object.R.ReferralPayouts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &referralPayoutR{}
			}
			foreign.R.Subscription = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SubscriptionID {
				local.R.ReferralPayouts = append(local.R.ReferralPayouts, foreign)
				if foreign.R == nil {
					foreign.R = &referralPayoutR{}
				}
				foreign.R.Subscription = local
				break
			}
		}
	}

	return nil
}

// SetAccount of the subscription to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.Subscriptions.
func (o *Subscription) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"subscription\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"account_id"}),
		strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
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
		o.R = &subscriptionR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			Subscriptions: SubscriptionSlice{o},
		}
	} else {
		related.R.Subscriptions = append(related.R.Subscriptions, o)
	}

	return nil
}

// SetPackage of the subscription to the related item.
// Sets o.R.Package to related.
// Adds o to related.R.Subscriptions.
func (o *Subscription) SetPackage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Package) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"subscription\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"package_id"}),
		strmangle.WhereClause("\"", "\"", 2, subscriptionPrimaryKeyColumns),
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

	o.PackageID = related.ID
	if o.R == nil {
		o.R = &subscriptionR{
			Package: related,
		}
	} else {
		o.R.Package = related
	}

	if related.R == nil {
		related.R = &packageR{
			Subscriptions: SubscriptionSlice{o},
		}
	} else {
		related.R.Subscriptions = append(related.R.Subscriptions, o)
	}

	return nil
}

// AddReferralPayouts adds the given related objects to the existing relationships
// of the subscription, optionally inserting them as new records.
// Appends related to o.R.ReferralPayouts.
// Sets related.R.Subscription appropriately.
func (o *Subscription) AddReferralPayouts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ReferralPayout) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SubscriptionID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"referral_payout\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"subscription_id"}),
				strmangle.WhereClause("\"", "\"", 2, referralPayoutPrimaryKeyColumns),
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

			rel.SubscriptionID = o.ID
		}
	}

	if o.R == nil {
		o.R = &subscriptionR{
			ReferralPayouts: related,
		}
	} else {
		o.R.ReferralPayouts = append(o.R.ReferralPayouts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &referralPayoutR{
				Subscription: o,
			}
		} else {
			rel.R.Subscription = o
		}
	}
	return nil
}

// Subscriptions retrieves all the records using an executor.
func Subscriptions(mods ...qm.QueryMod) subscriptionQuery {
	mods = append(mods, qm.From("\"subscription\""))
	return subscriptionQuery{NewQuery(mods...)}
}

// FindSubscription retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSubscription(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Subscription, error) {
	subscriptionObj := &Subscription{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"subscription\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, subscriptionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from subscription")
	}

	return subscriptionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Subscription) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subscription provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(subscriptionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	subscriptionInsertCacheMut.RLock()
	cache, cached := subscriptionInsertCache[key]
	subscriptionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			subscriptionAllColumns,
			subscriptionColumnsWithDefault,
			subscriptionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"subscription\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"subscription\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into subscription")
	}

	if !cached {
		subscriptionInsertCacheMut.Lock()
		subscriptionInsertCache[key] = cache
		subscriptionInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Subscription.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Subscription) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	subscriptionUpdateCacheMut.RLock()
	cache, cached := subscriptionUpdateCache[key]
	subscriptionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			subscriptionAllColumns,
			subscriptionPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update subscription, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"subscription\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, subscriptionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, append(wl, subscriptionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update subscription row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for subscription")
	}

	if !cached {
		subscriptionUpdateCacheMut.Lock()
		subscriptionUpdateCache[key] = cache
		subscriptionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q subscriptionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for subscription")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for subscription")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SubscriptionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"subscription\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, subscriptionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in subscription slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all subscription")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Subscription) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no subscription provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(subscriptionColumnsWithDefault, o)

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

	subscriptionUpsertCacheMut.RLock()
	cache, cached := subscriptionUpsertCache[key]
	subscriptionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			subscriptionAllColumns,
			subscriptionColumnsWithDefault,
			subscriptionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			subscriptionAllColumns,
			subscriptionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert subscription, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(subscriptionPrimaryKeyColumns))
			copy(conflict, subscriptionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"subscription\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(subscriptionType, subscriptionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert subscription")
	}

	if !cached {
		subscriptionUpsertCacheMut.Lock()
		subscriptionUpsertCache[key] = cache
		subscriptionUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Subscription record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Subscription) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Subscription provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), subscriptionPrimaryKeyMapping)
	sql := "DELETE FROM \"subscription\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from subscription")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for subscription")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q subscriptionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no subscriptionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subscription")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subscription")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SubscriptionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"subscription\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subscriptionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from subscription slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for subscription")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Subscription) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSubscription(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubscriptionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SubscriptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), subscriptionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"subscription\".* FROM \"subscription\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, subscriptionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SubscriptionSlice")
	}

	*o = slice

	return nil
}

// SubscriptionExists checks if the Subscription row exists.
func SubscriptionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"subscription\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if subscription exists")
	}

	return exists, nil
}
