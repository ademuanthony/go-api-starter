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

func testReferralPayouts(t *testing.T) {
	t.Parallel()

	query := ReferralPayouts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testReferralPayoutsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
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

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReferralPayoutsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ReferralPayouts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReferralPayoutsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReferralPayoutSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReferralPayoutsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ReferralPayoutExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ReferralPayout exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReferralPayoutExists to return true, but got false.")
	}
}

func testReferralPayoutsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	referralPayoutFound, err := FindReferralPayout(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if referralPayoutFound == nil {
		t.Error("want a record, got nil")
	}
}

func testReferralPayoutsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ReferralPayouts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testReferralPayoutsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ReferralPayouts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReferralPayoutsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	referralPayoutOne := &ReferralPayout{}
	referralPayoutTwo := &ReferralPayout{}
	if err = randomize.Struct(seed, referralPayoutOne, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}
	if err = randomize.Struct(seed, referralPayoutTwo, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = referralPayoutOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = referralPayoutTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReferralPayouts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReferralPayoutsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	referralPayoutOne := &ReferralPayout{}
	referralPayoutTwo := &ReferralPayout{}
	if err = randomize.Struct(seed, referralPayoutOne, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}
	if err = randomize.Struct(seed, referralPayoutTwo, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = referralPayoutOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = referralPayoutTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testReferralPayoutsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReferralPayoutsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(referralPayoutColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReferralPayoutToOneAccountUsingAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ReferralPayout
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
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

	slice := ReferralPayoutSlice{&local}
	if err = local.L.LoadAccount(ctx, tx, false, (*[]*ReferralPayout)(&slice), nil); err != nil {
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

func testReferralPayoutToOneDepositUsingDeposit(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ReferralPayout
	var foreign Deposit

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, depositDBTypes, false, depositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deposit struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DepositID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Deposit().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ReferralPayoutSlice{&local}
	if err = local.L.LoadDeposit(ctx, tx, false, (*[]*ReferralPayout)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Deposit == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Deposit = nil
	if err = local.L.LoadDeposit(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Deposit == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testReferralPayoutToOneAccountUsingFromAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ReferralPayout
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FromAccountID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FromAccount().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ReferralPayoutSlice{&local}
	if err = local.L.LoadFromAccount(ctx, tx, false, (*[]*ReferralPayout)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromAccount == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FromAccount = nil
	if err = local.L.LoadFromAccount(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FromAccount == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testReferralPayoutToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ReferralPayout
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, referralPayoutDBTypes, false, strmangle.SetComplement(referralPayoutPrimaryKeyColumns, referralPayoutColumnsWithoutDefault)...); err != nil {
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

		if x.R.ReferralPayouts[0] != &a {
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
func testReferralPayoutToOneSetOpDepositUsingDeposit(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ReferralPayout
	var b, c Deposit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, referralPayoutDBTypes, false, strmangle.SetComplement(referralPayoutPrimaryKeyColumns, referralPayoutColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, depositDBTypes, false, strmangle.SetComplement(depositPrimaryKeyColumns, depositColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, depositDBTypes, false, strmangle.SetComplement(depositPrimaryKeyColumns, depositColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Deposit{&b, &c} {
		err = a.SetDeposit(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Deposit != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ReferralPayouts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DepositID != x.ID {
			t.Error("foreign key was wrong value", a.DepositID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DepositID))
		reflect.Indirect(reflect.ValueOf(&a.DepositID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DepositID != x.ID {
			t.Error("foreign key was wrong value", a.DepositID, x.ID)
		}
	}
}
func testReferralPayoutToOneSetOpAccountUsingFromAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ReferralPayout
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, referralPayoutDBTypes, false, strmangle.SetComplement(referralPayoutPrimaryKeyColumns, referralPayoutColumnsWithoutDefault)...); err != nil {
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
		err = a.SetFromAccount(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FromAccount != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FromAccountReferralPayouts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FromAccountID != x.ID {
			t.Error("foreign key was wrong value", a.FromAccountID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FromAccountID))
		reflect.Indirect(reflect.ValueOf(&a.FromAccountID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FromAccountID != x.ID {
			t.Error("foreign key was wrong value", a.FromAccountID, x.ID)
		}
	}
}

func testReferralPayoutsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
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

func testReferralPayoutsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReferralPayoutSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReferralPayoutsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ReferralPayouts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	referralPayoutDBTypes = map[string]string{`ID`: `character varying`, `AccountID`: `character varying`, `FromAccountID`: `character varying`, `DepositID`: `character varying`, `Generation`: `integer`, `Amount`: `bigint`, `Date`: `bigint`, `PaymentMethod`: `integer`, `PaymentStatus`: `integer`, `PaymentRef`: `text`}
	_                     = bytes.MinRead
)

func testReferralPayoutsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(referralPayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(referralPayoutAllColumns) == len(referralPayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testReferralPayoutsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(referralPayoutAllColumns) == len(referralPayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ReferralPayout{}
	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, referralPayoutDBTypes, true, referralPayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(referralPayoutAllColumns, referralPayoutPrimaryKeyColumns) {
		fields = referralPayoutAllColumns
	} else {
		fields = strmangle.SetComplement(
			referralPayoutAllColumns,
			referralPayoutPrimaryKeyColumns,
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

	slice := ReferralPayoutSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testReferralPayoutsUpsert(t *testing.T) {
	t.Parallel()

	if len(referralPayoutAllColumns) == len(referralPayoutPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ReferralPayout{}
	if err = randomize.Struct(seed, &o, referralPayoutDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReferralPayout: %s", err)
	}

	count, err := ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, referralPayoutDBTypes, false, referralPayoutPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ReferralPayout struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ReferralPayout: %s", err)
	}

	count, err = ReferralPayouts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
