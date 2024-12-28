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

func testKycForms(t *testing.T) {
	t.Parallel()

	query := KycForms()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testKycFormsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
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

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKycFormsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := KycForms().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKycFormsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KycFormSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKycFormsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := KycFormExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if KycForm exists: %s", err)
	}
	if !e {
		t.Errorf("Expected KycFormExists to return true, but got false.")
	}
}

func testKycFormsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	kycFormFound, err := FindKycForm(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if kycFormFound == nil {
		t.Error("want a record, got nil")
	}
}

func testKycFormsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = KycForms().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testKycFormsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := KycForms().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testKycFormsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	kycFormOne := &KycForm{}
	kycFormTwo := &KycForm{}
	if err = randomize.Struct(seed, kycFormOne, kycFormDBTypes, false, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}
	if err = randomize.Struct(seed, kycFormTwo, kycFormDBTypes, false, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = kycFormOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = kycFormTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KycForms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testKycFormsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	kycFormOne := &KycForm{}
	kycFormTwo := &KycForm{}
	if err = randomize.Struct(seed, kycFormOne, kycFormDBTypes, false, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}
	if err = randomize.Struct(seed, kycFormTwo, kycFormDBTypes, false, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = kycFormOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = kycFormTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testKycFormsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKycFormsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(kycFormColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKycFormsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
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

func testKycFormsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KycFormSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testKycFormsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KycForms().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	kycFormDBTypes = map[string]string{`ID`: `uuid`, `AccountID`: `character varying`, `FirstName`: `character varying`, `LastName`: `character varying`, `MiddleName`: `character varying`, `ApplicationDate`: `bigint`, `DateOfBirth`: `bigint`, `IdentityType`: `character varying`, `IdentityIsseuDate`: `bigint`, `IdentityExpiryDate`: `bigint`, `IdentityCountry`: `character varying`, `Country`: `character varying`, `State`: `character varying`, `Address`: `character varying`, `IdentityFile`: `character varying`, `PhotoFile`: `character varying`, `VerificationStatus`: `integer`}
	_              = bytes.MinRead
)

func testKycFormsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(kycFormPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(kycFormAllColumns) == len(kycFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testKycFormsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(kycFormAllColumns) == len(kycFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KycForm{}
	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, kycFormDBTypes, true, kycFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(kycFormAllColumns, kycFormPrimaryKeyColumns) {
		fields = kycFormAllColumns
	} else {
		fields = strmangle.SetComplement(
			kycFormAllColumns,
			kycFormPrimaryKeyColumns,
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

	slice := KycFormSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testKycFormsUpsert(t *testing.T) {
	t.Parallel()

	if len(kycFormAllColumns) == len(kycFormPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := KycForm{}
	if err = randomize.Struct(seed, &o, kycFormDBTypes, true); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KycForm: %s", err)
	}

	count, err := KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, kycFormDBTypes, false, kycFormPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KycForm struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KycForm: %s", err)
	}

	count, err = KycForms().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}