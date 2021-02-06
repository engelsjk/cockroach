// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/typeconv"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecbase"
	"github.com/cockroachdb/cockroach/pkg/sql/colmem"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/duration"
	"github.com/cockroachdb/errors"
)

// Workaround for bazel auto-generated code. goimports does not automatically
// pick up the right packages when run within the bazel sandbox.
var (
	_ apd.Context
	_ duration.Duration
)

// NewConstOp creates a new operator that produces a constant value constVal of
// type t at index outputIdx.
func NewConstOp(
	allocator *colmem.Allocator,
	input colexecbase.Operator,
	t *types.T,
	constVal interface{},
	outputIdx int,
) (colexecbase.Operator, error) {
	input = newVectorTypeEnforcer(allocator, input, t, outputIdx)
	switch typeconv.TypeFamilyToCanonicalTypeFamily(t.Family()) {
	case types.BoolFamily:
		switch t.Width() {
		case -1:
		default:
			return &constBoolOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(bool),
			}, nil
		}
	case types.BytesFamily:
		switch t.Width() {
		case -1:
		default:
			return &constBytesOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.([]byte),
			}, nil
		}
	case types.DecimalFamily:
		switch t.Width() {
		case -1:
		default:
			return &constDecimalOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(apd.Decimal),
			}, nil
		}
	case types.IntFamily:
		switch t.Width() {
		case 16:
			return &constInt16Op{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(int16),
			}, nil
		case 32:
			return &constInt32Op{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(int32),
			}, nil
		case -1:
		default:
			return &constInt64Op{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(int64),
			}, nil
		}
	case types.FloatFamily:
		switch t.Width() {
		case -1:
		default:
			return &constFloat64Op{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(float64),
			}, nil
		}
	case types.TimestampTZFamily:
		switch t.Width() {
		case -1:
		default:
			return &constTimestampOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(time.Time),
			}, nil
		}
	case types.IntervalFamily:
		switch t.Width() {
		case -1:
		default:
			return &constIntervalOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(duration.Duration),
			}, nil
		}
	case typeconv.DatumVecCanonicalTypeFamily:
		switch t.Width() {
		case -1:
		default:
			return &constDatumOp{
				OneInputNode: NewOneInputNode(input),
				allocator:    allocator,
				outputIdx:    outputIdx,
				constVal:     constVal.(interface{}),
			}, nil
		}
	}
	return nil, errors.Errorf("unsupported const type %s", t.Name())
}

type constBoolOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  bool
}

func (c constBoolOp) Init() {
	c.input.Init()
}

func (c constBoolOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Bool()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constBytesOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  []byte
}

func (c constBytesOp) Init() {
	c.input.Init()
}

func (c constBytesOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Bytes()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col.Set(i, c.constVal)
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					col.Set(i, c.constVal)
				}
			}
		},
	)
	return batch
}

type constDecimalOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  apd.Decimal
}

func (c constDecimalOp) Init() {
	c.input.Init()
}

func (c constDecimalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Decimal()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i].Set(&c.constVal)
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i].Set(&c.constVal)
				}
			}
		},
	)
	return batch
}

type constInt16Op struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  int16
}

func (c constInt16Op) Init() {
	c.input.Init()
}

func (c constInt16Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int16()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constInt32Op struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  int32
}

func (c constInt32Op) Init() {
	c.input.Init()
}

func (c constInt32Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int32()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constInt64Op struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  int64
}

func (c constInt64Op) Init() {
	c.input.Init()
}

func (c constInt64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Int64()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constFloat64Op struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  float64
}

func (c constFloat64Op) Init() {
	c.input.Init()
}

func (c constFloat64Op) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Float64()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constTimestampOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  time.Time
}

func (c constTimestampOp) Init() {
	c.input.Init()
}

func (c constTimestampOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Timestamp()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constIntervalOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  duration.Duration
}

func (c constIntervalOp) Init() {
	c.input.Init()
}

func (c constIntervalOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Interval()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col[i] = c.constVal
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					//gcassert:bce
					col[i] = c.constVal
				}
			}
		},
	)
	return batch
}

type constDatumOp struct {
	OneInputNode

	allocator *colmem.Allocator
	outputIdx int
	constVal  interface{}
}

func (c constDatumOp) Init() {
	c.input.Init()
}

func (c constDatumOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}
	vec := batch.ColVec(c.outputIdx)
	col := vec.Datum()
	if vec.MaybeHasNulls() {
		// We need to make sure that there are no left over null values in the
		// output vector.
		vec.Nulls().UnsetNulls()
	}
	c.allocator.PerformOperation(
		[]coldata.Vec{vec},
		func() {
			// Shallow copy col to work around Go issue
			// https://github.com/golang/go/issues/39756 which prevents bound check
			// elimination from working in this case.
			col := col
			if sel := batch.Selection(); sel != nil {
				for _, i := range sel[:n] {
					col.Set(i, c.constVal)
				}
			} else {
				_ = col.Get(n - 1)
				for i := 0; i < n; i++ {
					col.Set(i, c.constVal)
				}
			}
		},
	)
	return batch
}

// NewConstNullOp creates a new operator that produces a constant (untyped) NULL
// value at index outputIdx.
func NewConstNullOp(
	allocator *colmem.Allocator, input colexecbase.Operator, outputIdx int,
) colexecbase.Operator {
	input = newVectorTypeEnforcer(allocator, input, types.Unknown, outputIdx)
	return &constNullOp{
		OneInputNode: NewOneInputNode(input),
		outputIdx:    outputIdx,
	}
}

type constNullOp struct {
	OneInputNode
	outputIdx int
}

var _ colexecbase.Operator = &constNullOp{}

func (c constNullOp) Init() {
	c.input.Init()
}

func (c constNullOp) Next(ctx context.Context) coldata.Batch {
	batch := c.input.Next(ctx)
	n := batch.Length()
	if n == 0 {
		return coldata.ZeroBatch
	}

	batch.ColVec(c.outputIdx).Nulls().SetNulls()
	return batch
}
