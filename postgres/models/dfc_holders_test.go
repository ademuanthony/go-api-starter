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

func testDFCHolders(t *testing.T) {
	t.Parallel()

	query := DFCHolders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDFCHoldersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
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

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDFCHoldersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DFCHolders().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDFCHoldersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DFCHolderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDFCHoldersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DFCHolderExists(ctx, tx, o.WalletAddress)
	if err != nil {
		t.Errorf("Unable to check if DFCHolder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DFCHolderExists to return true, but got false.")
	}
}

func testDFCHoldersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dfcHolderFound, err := FindDFCHolder(ctx, tx, o.WalletAddress)
	if err != nil {
		t.Error(err)
	}

	if dfcHolderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDFCHoldersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DFCHolders().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDFCHoldersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DFCHolders().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDFCHoldersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dfcHolderOne := &DFCHolder{}
	dfcHolderTwo := &DFCHolder{}
	if err = randomize.Struct(seed, dfcHolderOne, dfcHolderDBTypes, false, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}
	if err = randomize.Struct(seed, dfcHolderTwo, dfcHolderDBTypes, false, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dfcHolderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dfcHolderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DFCHolders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDFCHoldersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dfcHolderOne := &DFCHolder{}
	dfcHolderTwo := &DFCHolder{}
	if err = randomize.Struct(seed, dfcHolderOne, dfcHolderDBTypes, false, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}
	if err = randomize.Struct(seed, dfcHolderTwo, dfcHolderDBTypes, false, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dfcHolderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dfcHolderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDFCHoldersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDFCHoldersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dfcHolderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDFCHoldersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
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

func testDFCHoldersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DFCHolderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDFCHoldersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DFCHolders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dfcHolderDBTypes = map[string]string{`WalletAddress`: `character varying`, `Balance`: `character varying`, `Sent`: `integer`, `Balance2`: `bigint`, `Hash`: `character varying`}
	_                = bytes.MinRead
)

func testDFCHoldersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dfcHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dfcHolderAllColumns) == len(dfcHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDFCHoldersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dfcHolderAllColumns) == len(dfcHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DFCHolder{}
	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dfcHolderDBTypes, true, dfcHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dfcHolderAllColumns, dfcHolderPrimaryKeyColumns) {
		fields = dfcHolderAllColumns
	} else {
		fields = strmangle.SetComplement(
			dfcHolderAllColumns,
			dfcHolderPrimaryKeyColumns,
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

	slice := DFCHolderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDFCHoldersUpsert(t *testing.T) {
	t.Parallel()

	if len(dfcHolderAllColumns) == len(dfcHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DFCHolder{}
	if err = randomize.Struct(seed, &o, dfcHolderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DFCHolder: %s", err)
	}

	count, err := DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dfcHolderDBTypes, false, dfcHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DFCHolder struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DFCHolder: %s", err)
	}

	count, err = DFCHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
