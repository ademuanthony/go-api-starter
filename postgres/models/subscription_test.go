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

func testSubscriptions(t *testing.T) {
	t.Parallel()

	query := Subscriptions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSubscriptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
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

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Subscriptions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubscriptionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSubscriptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SubscriptionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Subscription exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SubscriptionExists to return true, but got false.")
	}
}

func testSubscriptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	subscriptionFound, err := FindSubscription(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if subscriptionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSubscriptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Subscriptions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSubscriptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Subscriptions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSubscriptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	subscriptionOne := &Subscription{}
	subscriptionTwo := &Subscription{}
	if err = randomize.Struct(seed, subscriptionOne, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}
	if err = randomize.Struct(seed, subscriptionTwo, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subscriptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subscriptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subscriptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSubscriptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	subscriptionOne := &Subscription{}
	subscriptionTwo := &Subscription{}
	if err = randomize.Struct(seed, subscriptionOne, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}
	if err = randomize.Struct(seed, subscriptionTwo, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = subscriptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = subscriptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSubscriptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubscriptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(subscriptionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSubscriptionToManyReferralPayouts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c ReferralPayout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, referralPayoutDBTypes, false, referralPayoutColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SubscriptionID = a.ID
	c.SubscriptionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ReferralPayouts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SubscriptionID == b.SubscriptionID {
			bFound = true
		}
		if v.SubscriptionID == c.SubscriptionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SubscriptionSlice{&a}
	if err = a.L.LoadReferralPayouts(ctx, tx, false, (*[]*Subscription)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ReferralPayouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ReferralPayouts = nil
	if err = a.L.LoadReferralPayouts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ReferralPayouts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSubscriptionToManyAddOpReferralPayouts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c, d, e ReferralPayout

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, false, strmangle.SetComplement(subscriptionPrimaryKeyColumns, subscriptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ReferralPayout{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, referralPayoutDBTypes, false, strmangle.SetComplement(referralPayoutPrimaryKeyColumns, referralPayoutColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ReferralPayout{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddReferralPayouts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.SubscriptionID {
			t.Error("foreign key was wrong value", a.ID, first.SubscriptionID)
		}
		if a.ID != second.SubscriptionID {
			t.Error("foreign key was wrong value", a.ID, second.SubscriptionID)
		}

		if first.R.Subscription != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Subscription != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ReferralPayouts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ReferralPayouts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ReferralPayouts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSubscriptionToOneAccountUsingAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Subscription
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
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

	slice := SubscriptionSlice{&local}
	if err = local.L.LoadAccount(ctx, tx, false, (*[]*Subscription)(&slice), nil); err != nil {
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

func testSubscriptionToOnePackageUsingPackage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Subscription
	var foreign Package

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, subscriptionDBTypes, false, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PackageID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Package().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SubscriptionSlice{&local}
	if err = local.L.LoadPackage(ctx, tx, false, (*[]*Subscription)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Package == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Package = nil
	if err = local.L.LoadPackage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Package == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSubscriptionToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, false, strmangle.SetComplement(subscriptionPrimaryKeyColumns, subscriptionColumnsWithoutDefault)...); err != nil {
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

		if x.R.Subscriptions[0] != &a {
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
func testSubscriptionToOneSetOpPackageUsingPackage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Subscription
	var b, c Package

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, subscriptionDBTypes, false, strmangle.SetComplement(subscriptionPrimaryKeyColumns, subscriptionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, packageDBTypes, false, strmangle.SetComplement(packagePrimaryKeyColumns, packageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, packageDBTypes, false, strmangle.SetComplement(packagePrimaryKeyColumns, packageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Package{&b, &c} {
		err = a.SetPackage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Package != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Subscriptions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PackageID != x.ID {
			t.Error("foreign key was wrong value", a.PackageID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PackageID))
		reflect.Indirect(reflect.ValueOf(&a.PackageID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PackageID != x.ID {
			t.Error("foreign key was wrong value", a.PackageID, x.ID)
		}
	}
}

func testSubscriptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
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

func testSubscriptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SubscriptionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSubscriptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Subscriptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	subscriptionDBTypes = map[string]string{`ID`: `character varying`, `AccountID`: `character varying`, `PackageID`: `character varying`, `StartDate`: `bigint`, `EndDate`: `bigint`}
	_                   = bytes.MinRead
)

func testSubscriptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(subscriptionAllColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSubscriptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(subscriptionAllColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Subscription{}
	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, subscriptionDBTypes, true, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(subscriptionAllColumns, subscriptionPrimaryKeyColumns) {
		fields = subscriptionAllColumns
	} else {
		fields = strmangle.SetComplement(
			subscriptionAllColumns,
			subscriptionPrimaryKeyColumns,
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

	slice := SubscriptionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSubscriptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(subscriptionAllColumns) == len(subscriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Subscription{}
	if err = randomize.Struct(seed, &o, subscriptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subscription: %s", err)
	}

	count, err := Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, subscriptionDBTypes, false, subscriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Subscription struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Subscription: %s", err)
	}

	count, err = Subscriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
