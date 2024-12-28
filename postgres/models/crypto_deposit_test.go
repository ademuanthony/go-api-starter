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

func testCryptoDeposits(t *testing.T) {
	t.Parallel()

	query := CryptoDeposits()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCryptoDepositsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
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

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCryptoDepositsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CryptoDeposits().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCryptoDepositsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CryptoDepositSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCryptoDepositsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CryptoDepositExists(ctx, tx, o.TransactionHash)
	if err != nil {
		t.Errorf("Unable to check if CryptoDeposit exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CryptoDepositExists to return true, but got false.")
	}
}

func testCryptoDepositsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	cryptoDepositFound, err := FindCryptoDeposit(ctx, tx, o.TransactionHash)
	if err != nil {
		t.Error(err)
	}

	if cryptoDepositFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCryptoDepositsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CryptoDeposits().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCryptoDepositsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CryptoDeposits().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCryptoDepositsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cryptoDepositOne := &CryptoDeposit{}
	cryptoDepositTwo := &CryptoDeposit{}
	if err = randomize.Struct(seed, cryptoDepositOne, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}
	if err = randomize.Struct(seed, cryptoDepositTwo, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cryptoDepositOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cryptoDepositTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CryptoDeposits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCryptoDepositsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cryptoDepositOne := &CryptoDeposit{}
	cryptoDepositTwo := &CryptoDeposit{}
	if err = randomize.Struct(seed, cryptoDepositOne, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}
	if err = randomize.Struct(seed, cryptoDepositTwo, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cryptoDepositOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cryptoDepositTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testCryptoDepositsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCryptoDepositsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(cryptoDepositColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCryptoDepositToOneAccountUsingAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CryptoDeposit
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
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

	slice := CryptoDepositSlice{&local}
	if err = local.L.LoadAccount(ctx, tx, false, (*[]*CryptoDeposit)(&slice), nil); err != nil {
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

func testCryptoDepositToOneDepositWalletUsingWalletAddressDepositWallet(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CryptoDeposit
	var foreign DepositWallet

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, cryptoDepositDBTypes, false, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, depositWalletDBTypes, false, depositWalletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DepositWallet struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.WalletAddress = foreign.Address
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WalletAddressDepositWallet().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Address != foreign.Address {
		t.Errorf("want: %v, got %v", foreign.Address, check.Address)
	}

	slice := CryptoDepositSlice{&local}
	if err = local.L.LoadWalletAddressDepositWallet(ctx, tx, false, (*[]*CryptoDeposit)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WalletAddressDepositWallet == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WalletAddressDepositWallet = nil
	if err = local.L.LoadWalletAddressDepositWallet(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WalletAddressDepositWallet == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCryptoDepositToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CryptoDeposit
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cryptoDepositDBTypes, false, strmangle.SetComplement(cryptoDepositPrimaryKeyColumns, cryptoDepositColumnsWithoutDefault)...); err != nil {
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

		if x.R.CryptoDeposits[0] != &a {
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
func testCryptoDepositToOneSetOpDepositWalletUsingWalletAddressDepositWallet(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CryptoDeposit
	var b, c DepositWallet

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cryptoDepositDBTypes, false, strmangle.SetComplement(cryptoDepositPrimaryKeyColumns, cryptoDepositColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, depositWalletDBTypes, false, strmangle.SetComplement(depositWalletPrimaryKeyColumns, depositWalletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, depositWalletDBTypes, false, strmangle.SetComplement(depositWalletPrimaryKeyColumns, depositWalletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DepositWallet{&b, &c} {
		err = a.SetWalletAddressDepositWallet(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WalletAddressDepositWallet != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WalletAddressCryptoDeposits[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.WalletAddress != x.Address {
			t.Error("foreign key was wrong value", a.WalletAddress)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WalletAddress))
		reflect.Indirect(reflect.ValueOf(&a.WalletAddress)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.WalletAddress != x.Address {
			t.Error("foreign key was wrong value", a.WalletAddress, x.Address)
		}
	}
}

func testCryptoDepositsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
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

func testCryptoDepositsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CryptoDepositSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCryptoDepositsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CryptoDeposits().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cryptoDepositDBTypes = map[string]string{`TransactionHash`: `character varying`, `WalletAddress`: `character varying`, `TokenAmount`: `character varying`, `BlockNumber`: `bigint`, `Currency`: `character varying`, `AccountID`: `character varying`, `Date`: `bigint`, `Network`: `character varying`, `Status`: `character varying`}
	_                    = bytes.MinRead
)

func testCryptoDepositsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(cryptoDepositPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(cryptoDepositAllColumns) == len(cryptoDepositPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCryptoDepositsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cryptoDepositAllColumns) == len(cryptoDepositPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CryptoDeposit{}
	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cryptoDepositDBTypes, true, cryptoDepositPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cryptoDepositAllColumns, cryptoDepositPrimaryKeyColumns) {
		fields = cryptoDepositAllColumns
	} else {
		fields = strmangle.SetComplement(
			cryptoDepositAllColumns,
			cryptoDepositPrimaryKeyColumns,
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

	slice := CryptoDepositSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCryptoDepositsUpsert(t *testing.T) {
	t.Parallel()

	if len(cryptoDepositAllColumns) == len(cryptoDepositPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CryptoDeposit{}
	if err = randomize.Struct(seed, &o, cryptoDepositDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CryptoDeposit: %s", err)
	}

	count, err := CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, cryptoDepositDBTypes, false, cryptoDepositPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CryptoDeposit struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CryptoDeposit: %s", err)
	}

	count, err = CryptoDeposits().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
