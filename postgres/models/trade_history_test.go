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

func testTradeHistories(t *testing.T) {
	t.Parallel()

	query := TradeHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTradeHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
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

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TradeHistories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TradeHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TradeHistoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TradeHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TradeHistoryExists to return true, but got false.")
	}
}

func testTradeHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tradeHistoryFound, err := FindTradeHistory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if tradeHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTradeHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TradeHistories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTradeHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TradeHistories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTradeHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tradeHistoryOne := &TradeHistory{}
	tradeHistoryTwo := &TradeHistory{}
	if err = randomize.Struct(seed, tradeHistoryOne, tradeHistoryDBTypes, false, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, tradeHistoryTwo, tradeHistoryDBTypes, false, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tradeHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tradeHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TradeHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTradeHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tradeHistoryOne := &TradeHistory{}
	tradeHistoryTwo := &TradeHistory{}
	if err = randomize.Struct(seed, tradeHistoryOne, tradeHistoryDBTypes, false, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, tradeHistoryTwo, tradeHistoryDBTypes, false, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tradeHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tradeHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTradeHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTradeHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tradeHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTradeHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
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

func testTradeHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TradeHistorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTradeHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TradeHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tradeHistoryDBTypes = map[string]string{`ID`: `uuid`, `AccountID`: `character varying`, `BaseCurrency`: `character varying`, `QouteCurrency`: `character varying`, `BaseAmount`: `bigint`, `QouteAmount`: `bigint`, `Date`: `bigint`}
	_                   = bytes.MinRead
)

func testTradeHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tradeHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tradeHistoryAllColumns) == len(tradeHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTradeHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tradeHistoryAllColumns) == len(tradeHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TradeHistory{}
	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tradeHistoryDBTypes, true, tradeHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tradeHistoryAllColumns, tradeHistoryPrimaryKeyColumns) {
		fields = tradeHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			tradeHistoryAllColumns,
			tradeHistoryPrimaryKeyColumns,
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

	slice := TradeHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTradeHistoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(tradeHistoryAllColumns) == len(tradeHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TradeHistory{}
	if err = randomize.Struct(seed, &o, tradeHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TradeHistory: %s", err)
	}

	count, err := TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tradeHistoryDBTypes, false, tradeHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeHistory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TradeHistory: %s", err)
	}

	count, err = TradeHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
