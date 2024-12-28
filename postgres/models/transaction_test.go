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

func testTransactions(t *testing.T) {
	t.Parallel()

	query := Transactions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTransactionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
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

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Transactions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TransactionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Transaction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactionExists to return true, but got false.")
	}
}

func testTransactionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	transactionFound, err := FindTransaction(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if transactionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTransactionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Transactions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTransactionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Transactions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Transactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = transactionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTransactionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(transactionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionToManyTransactionAssignments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Transaction
	var b, c TransactionAssignment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, transactionAssignmentDBTypes, false, transactionAssignmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionAssignmentDBTypes, false, transactionAssignmentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.TransactionID = a.ID
	c.TransactionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.TransactionAssignments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.TransactionID == b.TransactionID {
			bFound = true
		}
		if v.TransactionID == c.TransactionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionSlice{&a}
	if err = a.L.LoadTransactionAssignments(ctx, tx, false, (*[]*Transaction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionAssignments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TransactionAssignments = nil
	if err = a.L.LoadTransactionAssignments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionAssignments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTransactionToManyAddOpTransactionAssignments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Transaction
	var b, c, d, e TransactionAssignment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TransactionAssignment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionAssignmentDBTypes, false, strmangle.SetComplement(transactionAssignmentPrimaryKeyColumns, transactionAssignmentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*TransactionAssignment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactionAssignments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionID)
		}
		if a.ID != second.TransactionID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionID)
		}

		if first.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TransactionAssignments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TransactionAssignments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TransactionAssignments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTransactionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
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

func testTransactionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTransactionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Transactions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactionDBTypes = map[string]string{`ID`: `uuid`, `BankName`: `character varying`, `AccountNumber`: `character varying`, `AccountName`: `character varying`, `Amount`: `bigint`, `TokenAmount`: `character varying`, `AmountPaid`: `character varying`, `Email`: `character varying`, `Network`: `character varying`, `Currency`: `character varying`, `WalletAddress`: `character varying`, `PrivateKey`: `character varying`, `PaymentLink`: `character varying`, `Status`: `character varying`, `NairaAmount`: `bigint`, `PhoneNumber`: `character varying`, `MobileNetwork`: `character varying`, `SmartCardNumber`: `character varying`, `TVProvider`: `character varying`, `MeterNumber`: `character varying`, `PowerProvider`: `character varying`, `Type`: `integer`, `Date`: `bigint`, `MerchantRef`: `character varying`, `SuccessURL`: `character varying`, `FailureURL`: `character varying`, `RequestID`: `character varying`, `PackageID`: `character varying`, `PaymentMethod`: `character varying`, `ElectricToken`: `character varying`, `BillersCode`: `character varying`, `VariationCode`: `character varying`, `OperatorID`: `character varying`, `CountryCode`: `character varying`, `ProductTypeID`: `integer`}
	_                  = bytes.MinRead
)

func testTransactionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTransactionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Transaction{}
	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, transactionDBTypes, true, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactionAllColumns, transactionPrimaryKeyColumns) {
		fields = transactionAllColumns
	} else {
		fields = strmangle.SetComplement(
			transactionAllColumns,
			transactionPrimaryKeyColumns,
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

	slice := TransactionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTransactionsUpsert(t *testing.T) {
	t.Parallel()

	if len(transactionAllColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Transaction{}
	if err = randomize.Struct(seed, &o, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err := Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, transactionDBTypes, false, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err = Transactions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
