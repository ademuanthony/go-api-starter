// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testNetworkCurrencies(t *testing.T) {
	t.Parallel()

	query := NetworkCurrencies()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNetworkCurrenciesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkCurrenciesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NetworkCurrencies().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkCurrenciesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkCurrencySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNetworkCurrenciesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NetworkCurrencyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if NetworkCurrency exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NetworkCurrencyExists to return true, but got false.")
	}
}

func testNetworkCurrenciesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	networkCurrencyFound, err := FindNetworkCurrency(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if networkCurrencyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNetworkCurrenciesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NetworkCurrencies().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNetworkCurrenciesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NetworkCurrencies().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNetworkCurrenciesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	networkCurrencyOne := &NetworkCurrency{}
	networkCurrencyTwo := &NetworkCurrency{}
	if err = randomize.Struct(seed, networkCurrencyOne, networkCurrencyDBTypes, false, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}
	if err = randomize.Struct(seed, networkCurrencyTwo, networkCurrencyDBTypes, false, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkCurrencyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkCurrencyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkCurrencies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNetworkCurrenciesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	networkCurrencyOne := &NetworkCurrency{}
	networkCurrencyTwo := &NetworkCurrency{}
	if err = randomize.Struct(seed, networkCurrencyOne, networkCurrencyDBTypes, false, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}
	if err = randomize.Struct(seed, networkCurrencyTwo, networkCurrencyDBTypes, false, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = networkCurrencyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = networkCurrencyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testNetworkCurrenciesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkCurrenciesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(networkCurrencyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNetworkCurrencyToOneCurrencyUsingSymbolCurrency(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local NetworkCurrency
	var foreign Currency

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, networkCurrencyDBTypes, false, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, currencyDBTypes, false, currencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Currency struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.Symbol = foreign.Symbol
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SymbolCurrency().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Symbol != foreign.Symbol {
		t.Errorf("want: %v, got %v", foreign.Symbol, check.Symbol)
	}

	slice := NetworkCurrencySlice{&local}
	if err = local.L.LoadSymbolCurrency(ctx, tx, false, (*[]*NetworkCurrency)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SymbolCurrency == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SymbolCurrency = nil
	if err = local.L.LoadSymbolCurrency(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SymbolCurrency == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNetworkCurrencyToOneSetOpCurrencyUsingSymbolCurrency(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NetworkCurrency
	var b, c Currency

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, networkCurrencyDBTypes, false, strmangle.SetComplement(networkCurrencyPrimaryKeyColumns, networkCurrencyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, currencyDBTypes, false, strmangle.SetComplement(currencyPrimaryKeyColumns, currencyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, currencyDBTypes, false, strmangle.SetComplement(currencyPrimaryKeyColumns, currencyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Currency{&b, &c} {
		err = a.SetSymbolCurrency(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SymbolCurrency != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SymbolNetworkCurrencies[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Symbol != x.Symbol {
			t.Error("foreign key was wrong value", a.Symbol)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Symbol))
		reflect.Indirect(reflect.ValueOf(&a.Symbol)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Symbol != x.Symbol {
			t.Error("foreign key was wrong value", a.Symbol, x.Symbol)
		}
	}
}

func testNetworkCurrenciesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNetworkCurrenciesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NetworkCurrencySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNetworkCurrenciesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NetworkCurrencies().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	networkCurrencyDBTypes = map[string]string{`ID`: `uuid`, `Symbol`: `character varying`, `Network`: `character varying`, `ContractAddress`: `character varying`, `Decimals`: `integer`, `Active`: `boolean`}
	_                      = bytes.MinRead
)

func testNetworkCurrenciesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(networkCurrencyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(networkCurrencyAllColumns) == len(networkCurrencyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNetworkCurrenciesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(networkCurrencyAllColumns) == len(networkCurrencyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NetworkCurrency{}
	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, networkCurrencyDBTypes, true, networkCurrencyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(networkCurrencyAllColumns, networkCurrencyPrimaryKeyColumns) {
		fields = networkCurrencyAllColumns
	} else {
		fields = strmangle.SetComplement(
			networkCurrencyAllColumns,
			networkCurrencyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := NetworkCurrencySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNetworkCurrenciesUpsert(t *testing.T) {
	t.Parallel()

	if len(networkCurrencyAllColumns) == len(networkCurrencyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NetworkCurrency{}
	if err = randomize.Struct(seed, &o, networkCurrencyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkCurrency: %s", err)
	}

	count, err := NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, networkCurrencyDBTypes, false, networkCurrencyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NetworkCurrency struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NetworkCurrency: %s", err)
	}

	count, err = NetworkCurrencies().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
