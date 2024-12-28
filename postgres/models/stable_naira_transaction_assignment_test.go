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

func testStableNairaTransactionAssignments(t *testing.T) {
	t.Parallel()

	query := StableNairaTransactionAssignments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStableNairaTransactionAssignmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
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

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStableNairaTransactionAssignmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StableNairaTransactionAssignments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStableNairaTransactionAssignmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StableNairaTransactionAssignmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStableNairaTransactionAssignmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StableNairaTransactionAssignmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if StableNairaTransactionAssignment exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StableNairaTransactionAssignmentExists to return true, but got false.")
	}
}

func testStableNairaTransactionAssignmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stableNairaTransactionAssignmentFound, err := FindStableNairaTransactionAssignment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if stableNairaTransactionAssignmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStableNairaTransactionAssignmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StableNairaTransactionAssignments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStableNairaTransactionAssignmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StableNairaTransactionAssignments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStableNairaTransactionAssignmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stableNairaTransactionAssignmentOne := &StableNairaTransactionAssignment{}
	stableNairaTransactionAssignmentTwo := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, stableNairaTransactionAssignmentOne, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}
	if err = randomize.Struct(seed, stableNairaTransactionAssignmentTwo, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stableNairaTransactionAssignmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stableNairaTransactionAssignmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StableNairaTransactionAssignments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStableNairaTransactionAssignmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stableNairaTransactionAssignmentOne := &StableNairaTransactionAssignment{}
	stableNairaTransactionAssignmentTwo := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, stableNairaTransactionAssignmentOne, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}
	if err = randomize.Struct(seed, stableNairaTransactionAssignmentTwo, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stableNairaTransactionAssignmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stableNairaTransactionAssignmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testStableNairaTransactionAssignmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStableNairaTransactionAssignmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stableNairaTransactionAssignmentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStableNairaTransactionAssignmentToOneAgentUsingAgent(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StableNairaTransactionAssignment
	var foreign Agent

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, agentDBTypes, false, agentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Agent struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AgentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Agent().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StableNairaTransactionAssignmentSlice{&local}
	if err = local.L.LoadAgent(ctx, tx, false, (*[]*StableNairaTransactionAssignment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Agent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Agent = nil
	if err = local.L.LoadAgent(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Agent == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStableNairaTransactionAssignmentToOneStableNairaTransactionUsingTransaction(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StableNairaTransactionAssignment
	var foreign StableNairaTransaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stableNairaTransactionDBTypes, false, stableNairaTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransaction struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TransactionID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Transaction().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StableNairaTransactionAssignmentSlice{&local}
	if err = local.L.LoadTransaction(ctx, tx, false, (*[]*StableNairaTransactionAssignment)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Transaction = nil
	if err = local.L.LoadTransaction(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStableNairaTransactionAssignmentToOneSetOpAgentUsingAgent(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StableNairaTransactionAssignment
	var b, c Agent

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stableNairaTransactionAssignmentDBTypes, false, strmangle.SetComplement(stableNairaTransactionAssignmentPrimaryKeyColumns, stableNairaTransactionAssignmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, agentDBTypes, false, strmangle.SetComplement(agentPrimaryKeyColumns, agentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, agentDBTypes, false, strmangle.SetComplement(agentPrimaryKeyColumns, agentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Agent{&b, &c} {
		err = a.SetAgent(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Agent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StableNairaTransactionAssignments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AgentID != x.ID {
			t.Error("foreign key was wrong value", a.AgentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AgentID))
		reflect.Indirect(reflect.ValueOf(&a.AgentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AgentID != x.ID {
			t.Error("foreign key was wrong value", a.AgentID, x.ID)
		}
	}
}
func testStableNairaTransactionAssignmentToOneSetOpStableNairaTransactionUsingTransaction(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StableNairaTransactionAssignment
	var b, c StableNairaTransaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stableNairaTransactionAssignmentDBTypes, false, strmangle.SetComplement(stableNairaTransactionAssignmentPrimaryKeyColumns, stableNairaTransactionAssignmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stableNairaTransactionDBTypes, false, strmangle.SetComplement(stableNairaTransactionPrimaryKeyColumns, stableNairaTransactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stableNairaTransactionDBTypes, false, strmangle.SetComplement(stableNairaTransactionPrimaryKeyColumns, stableNairaTransactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*StableNairaTransaction{&b, &c} {
		err = a.SetTransaction(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Transaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TransactionStableNairaTransactionAssignments[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID, x.ID)
		}
	}
}

func testStableNairaTransactionAssignmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
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

func testStableNairaTransactionAssignmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StableNairaTransactionAssignmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStableNairaTransactionAssignmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StableNairaTransactionAssignments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stableNairaTransactionAssignmentDBTypes = map[string]string{`ID`: `integer`, `AgentID`: `integer`, `TransactionID`: `uuid`, `Amount`: `bigint`, `Date`: `bigint`, `Status`: `integer`}
	_                                       = bytes.MinRead
)

func testStableNairaTransactionAssignmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stableNairaTransactionAssignmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stableNairaTransactionAssignmentAllColumns) == len(stableNairaTransactionAssignmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStableNairaTransactionAssignmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stableNairaTransactionAssignmentAllColumns) == len(stableNairaTransactionAssignmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stableNairaTransactionAssignmentDBTypes, true, stableNairaTransactionAssignmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stableNairaTransactionAssignmentAllColumns, stableNairaTransactionAssignmentPrimaryKeyColumns) {
		fields = stableNairaTransactionAssignmentAllColumns
	} else {
		fields = strmangle.SetComplement(
			stableNairaTransactionAssignmentAllColumns,
			stableNairaTransactionAssignmentPrimaryKeyColumns,
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

	slice := StableNairaTransactionAssignmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStableNairaTransactionAssignmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(stableNairaTransactionAssignmentAllColumns) == len(stableNairaTransactionAssignmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StableNairaTransactionAssignment{}
	if err = randomize.Struct(seed, &o, stableNairaTransactionAssignmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StableNairaTransactionAssignment: %s", err)
	}

	count, err := StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stableNairaTransactionAssignmentDBTypes, false, stableNairaTransactionAssignmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StableNairaTransactionAssignment struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StableNairaTransactionAssignment: %s", err)
	}

	count, err = StableNairaTransactionAssignments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}