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

func testCgoldHolders(t *testing.T) {
	t.Parallel()

	query := CgoldHolders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCgoldHoldersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
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

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCgoldHoldersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CgoldHolders().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCgoldHoldersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CgoldHolderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCgoldHoldersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CgoldHolderExists(ctx, tx, o.WalletAddress)
	if err != nil {
		t.Errorf("Unable to check if CgoldHolder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CgoldHolderExists to return true, but got false.")
	}
}

func testCgoldHoldersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	cgoldHolderFound, err := FindCgoldHolder(ctx, tx, o.WalletAddress)
	if err != nil {
		t.Error(err)
	}

	if cgoldHolderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCgoldHoldersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CgoldHolders().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCgoldHoldersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CgoldHolders().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCgoldHoldersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cgoldHolderOne := &CgoldHolder{}
	cgoldHolderTwo := &CgoldHolder{}
	if err = randomize.Struct(seed, cgoldHolderOne, cgoldHolderDBTypes, false, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}
	if err = randomize.Struct(seed, cgoldHolderTwo, cgoldHolderDBTypes, false, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cgoldHolderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cgoldHolderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CgoldHolders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCgoldHoldersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cgoldHolderOne := &CgoldHolder{}
	cgoldHolderTwo := &CgoldHolder{}
	if err = randomize.Struct(seed, cgoldHolderOne, cgoldHolderDBTypes, false, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}
	if err = randomize.Struct(seed, cgoldHolderTwo, cgoldHolderDBTypes, false, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cgoldHolderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cgoldHolderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testCgoldHoldersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCgoldHoldersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(cgoldHolderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCgoldHoldersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
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

func testCgoldHoldersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CgoldHolderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCgoldHoldersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CgoldHolders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cgoldHolderDBTypes = map[string]string{`WalletAddress`: `character varying`, `Balance`: `character varying`}
	_                  = bytes.MinRead
)

func testCgoldHoldersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(cgoldHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(cgoldHolderAllColumns) == len(cgoldHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCgoldHoldersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cgoldHolderAllColumns) == len(cgoldHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CgoldHolder{}
	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cgoldHolderDBTypes, true, cgoldHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cgoldHolderAllColumns, cgoldHolderPrimaryKeyColumns) {
		fields = cgoldHolderAllColumns
	} else {
		fields = strmangle.SetComplement(
			cgoldHolderAllColumns,
			cgoldHolderPrimaryKeyColumns,
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

	slice := CgoldHolderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCgoldHoldersUpsert(t *testing.T) {
	t.Parallel()

	if len(cgoldHolderAllColumns) == len(cgoldHolderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CgoldHolder{}
	if err = randomize.Struct(seed, &o, cgoldHolderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CgoldHolder: %s", err)
	}

	count, err := CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, cgoldHolderDBTypes, false, cgoldHolderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CgoldHolder struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CgoldHolder: %s", err)
	}

	count, err = CgoldHolders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
