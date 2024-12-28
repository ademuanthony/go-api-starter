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

// Transaction is an object representing the database table.
type Transaction struct {
	ID              string `boil:"id" json:"id" toml:"id" yaml:"id"`
	BankName        string `boil:"bank_name" json:"bank_name" toml:"bank_name" yaml:"bank_name"`
	AccountNumber   string `boil:"account_number" json:"account_number" toml:"account_number" yaml:"account_number"`
	AccountName     string `boil:"account_name" json:"account_name" toml:"account_name" yaml:"account_name"`
	Amount          int64  `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	TokenAmount     string `boil:"token_amount" json:"token_amount" toml:"token_amount" yaml:"token_amount"`
	AmountPaid      string `boil:"amount_paid" json:"amount_paid" toml:"amount_paid" yaml:"amount_paid"`
	Email           string `boil:"email" json:"email" toml:"email" yaml:"email"`
	Network         string `boil:"network" json:"network" toml:"network" yaml:"network"`
	Currency        string `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	WalletAddress   string `boil:"wallet_address" json:"wallet_address" toml:"wallet_address" yaml:"wallet_address"`
	PrivateKey      string `boil:"private_key" json:"private_key" toml:"private_key" yaml:"private_key"`
	PaymentLink     string `boil:"payment_link" json:"payment_link" toml:"payment_link" yaml:"payment_link"`
	Status          string `boil:"status" json:"status" toml:"status" yaml:"status"`
	NairaAmount     int64  `boil:"naira_amount" json:"naira_amount" toml:"naira_amount" yaml:"naira_amount"`
	PhoneNumber     string `boil:"phone_number" json:"phone_number" toml:"phone_number" yaml:"phone_number"`
	MobileNetwork   string `boil:"mobile_network" json:"mobile_network" toml:"mobile_network" yaml:"mobile_network"`
	SmartCardNumber string `boil:"smart_card_number" json:"smart_card_number" toml:"smart_card_number" yaml:"smart_card_number"`
	TVProvider      string `boil:"tv_provider" json:"tv_provider" toml:"tv_provider" yaml:"tv_provider"`
	MeterNumber     string `boil:"meter_number" json:"meter_number" toml:"meter_number" yaml:"meter_number"`
	PowerProvider   string `boil:"power_provider" json:"power_provider" toml:"power_provider" yaml:"power_provider"`
	Type            int    `boil:"type" json:"type" toml:"type" yaml:"type"`
	Date            int64  `boil:"date" json:"date" toml:"date" yaml:"date"`
	MerchantRef     string `boil:"merchant_ref" json:"merchant_ref" toml:"merchant_ref" yaml:"merchant_ref"`
	SuccessURL      string `boil:"success_url" json:"success_url" toml:"success_url" yaml:"success_url"`
	FailureURL      string `boil:"failure_url" json:"failure_url" toml:"failure_url" yaml:"failure_url"`
	RequestID       string `boil:"request_id" json:"request_id" toml:"request_id" yaml:"request_id"`
	PackageID       string `boil:"package_id" json:"package_id" toml:"package_id" yaml:"package_id"`
	PaymentMethod   string `boil:"payment_method" json:"payment_method" toml:"payment_method" yaml:"payment_method"`
	ElectricToken   string `boil:"electric_token" json:"electric_token" toml:"electric_token" yaml:"electric_token"`
	BillersCode     string `boil:"billers_code" json:"billers_code" toml:"billers_code" yaml:"billers_code"`
	VariationCode   string `boil:"variation_code" json:"variation_code" toml:"variation_code" yaml:"variation_code"`
	OperatorID      string `boil:"operator_id" json:"operator_id" toml:"operator_id" yaml:"operator_id"`
	CountryCode     string `boil:"country_code" json:"country_code" toml:"country_code" yaml:"country_code"`
	ProductTypeID   int    `boil:"product_type_id" json:"product_type_id" toml:"product_type_id" yaml:"product_type_id"`

	R *transactionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transactionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransactionColumns = struct {
	ID              string
	BankName        string
	AccountNumber   string
	AccountName     string
	Amount          string
	TokenAmount     string
	AmountPaid      string
	Email           string
	Network         string
	Currency        string
	WalletAddress   string
	PrivateKey      string
	PaymentLink     string
	Status          string
	NairaAmount     string
	PhoneNumber     string
	MobileNetwork   string
	SmartCardNumber string
	TVProvider      string
	MeterNumber     string
	PowerProvider   string
	Type            string
	Date            string
	MerchantRef     string
	SuccessURL      string
	FailureURL      string
	RequestID       string
	PackageID       string
	PaymentMethod   string
	ElectricToken   string
	BillersCode     string
	VariationCode   string
	OperatorID      string
	CountryCode     string
	ProductTypeID   string
}{
	ID:              "id",
	BankName:        "bank_name",
	AccountNumber:   "account_number",
	AccountName:     "account_name",
	Amount:          "amount",
	TokenAmount:     "token_amount",
	AmountPaid:      "amount_paid",
	Email:           "email",
	Network:         "network",
	Currency:        "currency",
	WalletAddress:   "wallet_address",
	PrivateKey:      "private_key",
	PaymentLink:     "payment_link",
	Status:          "status",
	NairaAmount:     "naira_amount",
	PhoneNumber:     "phone_number",
	MobileNetwork:   "mobile_network",
	SmartCardNumber: "smart_card_number",
	TVProvider:      "tv_provider",
	MeterNumber:     "meter_number",
	PowerProvider:   "power_provider",
	Type:            "type",
	Date:            "date",
	MerchantRef:     "merchant_ref",
	SuccessURL:      "success_url",
	FailureURL:      "failure_url",
	RequestID:       "request_id",
	PackageID:       "package_id",
	PaymentMethod:   "payment_method",
	ElectricToken:   "electric_token",
	BillersCode:     "billers_code",
	VariationCode:   "variation_code",
	OperatorID:      "operator_id",
	CountryCode:     "country_code",
	ProductTypeID:   "product_type_id",
}

var TransactionTableColumns = struct {
	ID              string
	BankName        string
	AccountNumber   string
	AccountName     string
	Amount          string
	TokenAmount     string
	AmountPaid      string
	Email           string
	Network         string
	Currency        string
	WalletAddress   string
	PrivateKey      string
	PaymentLink     string
	Status          string
	NairaAmount     string
	PhoneNumber     string
	MobileNetwork   string
	SmartCardNumber string
	TVProvider      string
	MeterNumber     string
	PowerProvider   string
	Type            string
	Date            string
	MerchantRef     string
	SuccessURL      string
	FailureURL      string
	RequestID       string
	PackageID       string
	PaymentMethod   string
	ElectricToken   string
	BillersCode     string
	VariationCode   string
	OperatorID      string
	CountryCode     string
	ProductTypeID   string
}{
	ID:              "transaction.id",
	BankName:        "transaction.bank_name",
	AccountNumber:   "transaction.account_number",
	AccountName:     "transaction.account_name",
	Amount:          "transaction.amount",
	TokenAmount:     "transaction.token_amount",
	AmountPaid:      "transaction.amount_paid",
	Email:           "transaction.email",
	Network:         "transaction.network",
	Currency:        "transaction.currency",
	WalletAddress:   "transaction.wallet_address",
	PrivateKey:      "transaction.private_key",
	PaymentLink:     "transaction.payment_link",
	Status:          "transaction.status",
	NairaAmount:     "transaction.naira_amount",
	PhoneNumber:     "transaction.phone_number",
	MobileNetwork:   "transaction.mobile_network",
	SmartCardNumber: "transaction.smart_card_number",
	TVProvider:      "transaction.tv_provider",
	MeterNumber:     "transaction.meter_number",
	PowerProvider:   "transaction.power_provider",
	Type:            "transaction.type",
	Date:            "transaction.date",
	MerchantRef:     "transaction.merchant_ref",
	SuccessURL:      "transaction.success_url",
	FailureURL:      "transaction.failure_url",
	RequestID:       "transaction.request_id",
	PackageID:       "transaction.package_id",
	PaymentMethod:   "transaction.payment_method",
	ElectricToken:   "transaction.electric_token",
	BillersCode:     "transaction.billers_code",
	VariationCode:   "transaction.variation_code",
	OperatorID:      "transaction.operator_id",
	CountryCode:     "transaction.country_code",
	ProductTypeID:   "transaction.product_type_id",
}

// Generated where

var TransactionWhere = struct {
	ID              whereHelperstring
	BankName        whereHelperstring
	AccountNumber   whereHelperstring
	AccountName     whereHelperstring
	Amount          whereHelperint64
	TokenAmount     whereHelperstring
	AmountPaid      whereHelperstring
	Email           whereHelperstring
	Network         whereHelperstring
	Currency        whereHelperstring
	WalletAddress   whereHelperstring
	PrivateKey      whereHelperstring
	PaymentLink     whereHelperstring
	Status          whereHelperstring
	NairaAmount     whereHelperint64
	PhoneNumber     whereHelperstring
	MobileNetwork   whereHelperstring
	SmartCardNumber whereHelperstring
	TVProvider      whereHelperstring
	MeterNumber     whereHelperstring
	PowerProvider   whereHelperstring
	Type            whereHelperint
	Date            whereHelperint64
	MerchantRef     whereHelperstring
	SuccessURL      whereHelperstring
	FailureURL      whereHelperstring
	RequestID       whereHelperstring
	PackageID       whereHelperstring
	PaymentMethod   whereHelperstring
	ElectricToken   whereHelperstring
	BillersCode     whereHelperstring
	VariationCode   whereHelperstring
	OperatorID      whereHelperstring
	CountryCode     whereHelperstring
	ProductTypeID   whereHelperint
}{
	ID:              whereHelperstring{field: "\"transaction\".\"id\""},
	BankName:        whereHelperstring{field: "\"transaction\".\"bank_name\""},
	AccountNumber:   whereHelperstring{field: "\"transaction\".\"account_number\""},
	AccountName:     whereHelperstring{field: "\"transaction\".\"account_name\""},
	Amount:          whereHelperint64{field: "\"transaction\".\"amount\""},
	TokenAmount:     whereHelperstring{field: "\"transaction\".\"token_amount\""},
	AmountPaid:      whereHelperstring{field: "\"transaction\".\"amount_paid\""},
	Email:           whereHelperstring{field: "\"transaction\".\"email\""},
	Network:         whereHelperstring{field: "\"transaction\".\"network\""},
	Currency:        whereHelperstring{field: "\"transaction\".\"currency\""},
	WalletAddress:   whereHelperstring{field: "\"transaction\".\"wallet_address\""},
	PrivateKey:      whereHelperstring{field: "\"transaction\".\"private_key\""},
	PaymentLink:     whereHelperstring{field: "\"transaction\".\"payment_link\""},
	Status:          whereHelperstring{field: "\"transaction\".\"status\""},
	NairaAmount:     whereHelperint64{field: "\"transaction\".\"naira_amount\""},
	PhoneNumber:     whereHelperstring{field: "\"transaction\".\"phone_number\""},
	MobileNetwork:   whereHelperstring{field: "\"transaction\".\"mobile_network\""},
	SmartCardNumber: whereHelperstring{field: "\"transaction\".\"smart_card_number\""},
	TVProvider:      whereHelperstring{field: "\"transaction\".\"tv_provider\""},
	MeterNumber:     whereHelperstring{field: "\"transaction\".\"meter_number\""},
	PowerProvider:   whereHelperstring{field: "\"transaction\".\"power_provider\""},
	Type:            whereHelperint{field: "\"transaction\".\"type\""},
	Date:            whereHelperint64{field: "\"transaction\".\"date\""},
	MerchantRef:     whereHelperstring{field: "\"transaction\".\"merchant_ref\""},
	SuccessURL:      whereHelperstring{field: "\"transaction\".\"success_url\""},
	FailureURL:      whereHelperstring{field: "\"transaction\".\"failure_url\""},
	RequestID:       whereHelperstring{field: "\"transaction\".\"request_id\""},
	PackageID:       whereHelperstring{field: "\"transaction\".\"package_id\""},
	PaymentMethod:   whereHelperstring{field: "\"transaction\".\"payment_method\""},
	ElectricToken:   whereHelperstring{field: "\"transaction\".\"electric_token\""},
	BillersCode:     whereHelperstring{field: "\"transaction\".\"billers_code\""},
	VariationCode:   whereHelperstring{field: "\"transaction\".\"variation_code\""},
	OperatorID:      whereHelperstring{field: "\"transaction\".\"operator_id\""},
	CountryCode:     whereHelperstring{field: "\"transaction\".\"country_code\""},
	ProductTypeID:   whereHelperint{field: "\"transaction\".\"product_type_id\""},
}

// TransactionRels is where relationship names are stored.
var TransactionRels = struct {
	TransactionAssignments string
}{
	TransactionAssignments: "TransactionAssignments",
}

// transactionR is where relationships are stored.
type transactionR struct {
	TransactionAssignments TransactionAssignmentSlice `boil:"TransactionAssignments" json:"TransactionAssignments" toml:"TransactionAssignments" yaml:"TransactionAssignments"`
}

// NewStruct creates a new relationship struct
func (*transactionR) NewStruct() *transactionR {
	return &transactionR{}
}

// transactionL is where Load methods for each relationship are stored.
type transactionL struct{}

var (
	transactionAllColumns            = []string{"id", "bank_name", "account_number", "account_name", "amount", "token_amount", "amount_paid", "email", "network", "currency", "wallet_address", "private_key", "payment_link", "status", "naira_amount", "phone_number", "mobile_network", "smart_card_number", "tv_provider", "meter_number", "power_provider", "type", "date", "merchant_ref", "success_url", "failure_url", "request_id", "package_id", "payment_method", "electric_token", "billers_code", "variation_code", "operator_id", "country_code", "product_type_id"}
	transactionColumnsWithoutDefault = []string{"bank_name", "account_number", "account_name", "amount", "token_amount", "amount_paid", "email", "network", "currency", "wallet_address", "private_key", "payment_link", "status"}
	transactionColumnsWithDefault    = []string{"id", "naira_amount", "phone_number", "mobile_network", "smart_card_number", "tv_provider", "meter_number", "power_provider", "type", "date", "merchant_ref", "success_url", "failure_url", "request_id", "package_id", "payment_method", "electric_token", "billers_code", "variation_code", "operator_id", "country_code", "product_type_id"}
	transactionPrimaryKeyColumns     = []string{"id"}
)

type (
	// TransactionSlice is an alias for a slice of pointers to Transaction.
	// This should almost always be used instead of []Transaction.
	TransactionSlice []*Transaction

	transactionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transactionType                 = reflect.TypeOf(&Transaction{})
	transactionMapping              = queries.MakeStructMapping(transactionType)
	transactionPrimaryKeyMapping, _ = queries.BindMapping(transactionType, transactionMapping, transactionPrimaryKeyColumns)
	transactionInsertCacheMut       sync.RWMutex
	transactionInsertCache          = make(map[string]insertCache)
	transactionUpdateCacheMut       sync.RWMutex
	transactionUpdateCache          = make(map[string]updateCache)
	transactionUpsertCacheMut       sync.RWMutex
	transactionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single transaction record from the query.
func (q transactionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Transaction, error) {
	o := &Transaction{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for transaction")
	}

	return o, nil
}

// All returns all Transaction records from the query.
func (q transactionQuery) All(ctx context.Context, exec boil.ContextExecutor) (TransactionSlice, error) {
	var o []*Transaction

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Transaction slice")
	}

	return o, nil
}

// Count returns the count of all Transaction records in the query.
func (q transactionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count transaction rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q transactionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if transaction exists")
	}

	return count > 0, nil
}

// TransactionAssignments retrieves all the transaction_assignment's TransactionAssignments with an executor.
func (o *Transaction) TransactionAssignments(mods ...qm.QueryMod) transactionAssignmentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"transaction_assignment\".\"transaction_id\"=?", o.ID),
	)

	query := TransactionAssignments(queryMods...)
	queries.SetFrom(query.Query, "\"transaction_assignment\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"transaction_assignment\".*"})
	}

	return query
}

// LoadTransactionAssignments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (transactionL) LoadTransactionAssignments(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTransaction interface{}, mods queries.Applicator) error {
	var slice []*Transaction
	var object *Transaction

	if singular {
		object = maybeTransaction.(*Transaction)
	} else {
		slice = *maybeTransaction.(*[]*Transaction)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &transactionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &transactionR{}
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
		qm.From(`transaction_assignment`),
		qm.WhereIn(`transaction_assignment.transaction_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load transaction_assignment")
	}

	var resultSlice []*TransactionAssignment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice transaction_assignment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on transaction_assignment")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transaction_assignment")
	}

	if singular {
		object.R.TransactionAssignments = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &transactionAssignmentR{}
			}
			foreign.R.Transaction = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.TransactionID {
				local.R.TransactionAssignments = append(local.R.TransactionAssignments, foreign)
				if foreign.R == nil {
					foreign.R = &transactionAssignmentR{}
				}
				foreign.R.Transaction = local
				break
			}
		}
	}

	return nil
}

// AddTransactionAssignments adds the given related objects to the existing relationships
// of the transaction, optionally inserting them as new records.
// Appends related to o.R.TransactionAssignments.
// Sets related.R.Transaction appropriately.
func (o *Transaction) AddTransactionAssignments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TransactionAssignment) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TransactionID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"transaction_assignment\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"transaction_id"}),
				strmangle.WhereClause("\"", "\"", 2, transactionAssignmentPrimaryKeyColumns),
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

			rel.TransactionID = o.ID
		}
	}

	if o.R == nil {
		o.R = &transactionR{
			TransactionAssignments: related,
		}
	} else {
		o.R.TransactionAssignments = append(o.R.TransactionAssignments, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &transactionAssignmentR{
				Transaction: o,
			}
		} else {
			rel.R.Transaction = o
		}
	}
	return nil
}

// Transactions retrieves all the records using an executor.
func Transactions(mods ...qm.QueryMod) transactionQuery {
	mods = append(mods, qm.From("\"transaction\""))
	return transactionQuery{NewQuery(mods...)}
}

// FindTransaction retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransaction(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Transaction, error) {
	transactionObj := &Transaction{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"transaction\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, transactionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from transaction")
	}

	return transactionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Transaction) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(transactionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	transactionInsertCacheMut.RLock()
	cache, cached := transactionInsertCache[key]
	transactionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			transactionAllColumns,
			transactionColumnsWithDefault,
			transactionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transactionType, transactionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"transaction\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"transaction\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into transaction")
	}

	if !cached {
		transactionInsertCacheMut.Lock()
		transactionInsertCache[key] = cache
		transactionInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Transaction.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Transaction) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	transactionUpdateCacheMut.RLock()
	cache, cached := transactionUpdateCache[key]
	transactionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			transactionAllColumns,
			transactionPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update transaction, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"transaction\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, transactionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, append(wl, transactionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update transaction row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for transaction")
	}

	if !cached {
		transactionUpdateCacheMut.Lock()
		transactionUpdateCache[key] = cache
		transactionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q transactionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for transaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for transaction")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransactionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"transaction\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, transactionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in transaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all transaction")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Transaction) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no transaction provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(transactionColumnsWithDefault, o)

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

	transactionUpsertCacheMut.RLock()
	cache, cached := transactionUpsertCache[key]
	transactionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			transactionAllColumns,
			transactionColumnsWithDefault,
			transactionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			transactionAllColumns,
			transactionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert transaction, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(transactionPrimaryKeyColumns))
			copy(conflict, transactionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"transaction\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(transactionType, transactionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transactionType, transactionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert transaction")
	}

	if !cached {
		transactionUpsertCacheMut.Lock()
		transactionUpsertCache[key] = cache
		transactionUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Transaction record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Transaction) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Transaction provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transactionPrimaryKeyMapping)
	sql := "DELETE FROM \"transaction\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from transaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for transaction")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q transactionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no transactionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transaction")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransactionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"transaction\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transaction slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for transaction")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Transaction) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTransaction(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransactionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TransactionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transactionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"transaction\".* FROM \"transaction\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, transactionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransactionSlice")
	}

	*o = slice

	return nil
}

// TransactionExists checks if the Transaction row exists.
func TransactionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"transaction\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if transaction exists")
	}

	return exists, nil
}
