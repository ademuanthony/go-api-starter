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

func testTradeSchedules(t *testing.T) {
	t.Parallel()

	query := TradeSchedules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTradeSchedulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
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

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeSchedulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TradeSchedules().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeSchedulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TradeScheduleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTradeSchedulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TradeScheduleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TradeSchedule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TradeScheduleExists to return true, but got false.")
	}
}

func testTradeSchedulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tradeScheduleFound, err := FindTradeSchedule(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if tradeScheduleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTradeSchedulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TradeSchedules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTradeSchedulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TradeSchedules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTradeSchedulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tradeScheduleOne := &TradeSchedule{}
	tradeScheduleTwo := &TradeSchedule{}
	if err = randomize.Struct(seed, tradeScheduleOne, tradeScheduleDBTypes, false, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}
	if err = randomize.Struct(seed, tradeScheduleTwo, tradeScheduleDBTypes, false, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tradeScheduleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tradeScheduleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TradeSchedules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTradeSchedulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tradeScheduleOne := &TradeSchedule{}
	tradeScheduleTwo := &TradeSchedule{}
	if err = randomize.Struct(seed, tradeScheduleOne, tradeScheduleDBTypes, false, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}
	if err = randomize.Struct(seed, tradeScheduleTwo, tradeScheduleDBTypes, false, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tradeScheduleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tradeScheduleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTradeSchedulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTradeSchedulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tradeScheduleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTradeScheduleToOneAccountUsingAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TradeSchedule
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, tradeScheduleDBTypes, false, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AccountID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Account().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TradeScheduleSlice{&local}
	if err = local.L.LoadAccount(ctx, tx, false, (*[]*TradeSchedule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Account = nil
	if err = local.L.LoadAccount(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTradeScheduleToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TradeSchedule
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tradeScheduleDBTypes, false, strmangle.SetComplement(tradeSchedulePrimaryKeyColumns, tradeScheduleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Account{&b, &c} {
		err = a.SetAccount(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Account != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TradeSchedules[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AccountID != x.ID {
			t.Error("foreign key was wrong value", a.AccountID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AccountID))
		reflect.Indirect(reflect.ValueOf(&a.AccountID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AccountID != x.ID {
			t.Error("foreign key was wrong value", a.AccountID, x.ID)
		}
	}
}

func testTradeSchedulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
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

func testTradeSchedulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TradeScheduleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTradeSchedulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TradeSchedules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tradeScheduleDBTypes = map[string]string{`ID`: `uuid`, `AccountID`: `character varying`, `TradeNo`: `integer`, `TotalTrades`: `integer`, `Date`: `bigint`, `TargetProfitPercentage`: `integer`, `StartDate`: `integer`}
	_                    = bytes.MinRead
)

func testTradeSchedulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tradeSchedulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tradeScheduleAllColumns) == len(tradeSchedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeSchedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTradeSchedulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tradeScheduleAllColumns) == len(tradeSchedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TradeSchedule{}
	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeScheduleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tradeScheduleDBTypes, true, tradeSchedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tradeScheduleAllColumns, tradeSchedulePrimaryKeyColumns) {
		fields = tradeScheduleAllColumns
	} else {
		fields = strmangle.SetComplement(
			tradeScheduleAllColumns,
			tradeSchedulePrimaryKeyColumns,
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

	slice := TradeScheduleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTradeSchedulesUpsert(t *testing.T) {
	t.Parallel()

	if len(tradeScheduleAllColumns) == len(tradeSchedulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TradeSchedule{}
	if err = randomize.Struct(seed, &o, tradeScheduleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TradeSchedule: %s", err)
	}

	count, err := TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tradeScheduleDBTypes, false, tradeSchedulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TradeSchedule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TradeSchedule: %s", err)
	}

	count, err = TradeSchedules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
