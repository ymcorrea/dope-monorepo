// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/ent/asset"
	"github.com/dopedao/dope-monorepo/packages/api/ent/dope"
	"github.com/dopedao/dope-monorepo/packages/api/ent/listing"
	"github.com/dopedao/dope-monorepo/packages/api/ent/predicate"
)

// ListingUpdate is the builder for updating Listing entities.
type ListingUpdate struct {
	config
	hooks    []Hook
	mutation *ListingMutation
}

// Where appends a list predicates to the ListingUpdate builder.
func (lu *ListingUpdate) Where(ps ...predicate.Listing) *ListingUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetActive sets the "active" field.
func (lu *ListingUpdate) SetActive(b bool) *ListingUpdate {
	lu.mutation.SetActive(b)
	return lu
}

// SetDopeID sets the "dope" edge to the Dope entity by ID.
func (lu *ListingUpdate) SetDopeID(id string) *ListingUpdate {
	lu.mutation.SetDopeID(id)
	return lu
}

// SetNillableDopeID sets the "dope" edge to the Dope entity by ID if the given value is not nil.
func (lu *ListingUpdate) SetNillableDopeID(id *string) *ListingUpdate {
	if id != nil {
		lu = lu.SetDopeID(*id)
	}
	return lu
}

// SetDope sets the "dope" edge to the Dope entity.
func (lu *ListingUpdate) SetDope(d *Dope) *ListingUpdate {
	return lu.SetDopeID(d.ID)
}

// AddDopeLastsaleIDs adds the "dope_lastsales" edge to the Dope entity by IDs.
func (lu *ListingUpdate) AddDopeLastsaleIDs(ids ...string) *ListingUpdate {
	lu.mutation.AddDopeLastsaleIDs(ids...)
	return lu
}

// AddDopeLastsales adds the "dope_lastsales" edges to the Dope entity.
func (lu *ListingUpdate) AddDopeLastsales(d ...*Dope) *ListingUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return lu.AddDopeLastsaleIDs(ids...)
}

// AddInputIDs adds the "inputs" edge to the Asset entity by IDs.
func (lu *ListingUpdate) AddInputIDs(ids ...string) *ListingUpdate {
	lu.mutation.AddInputIDs(ids...)
	return lu
}

// AddInputs adds the "inputs" edges to the Asset entity.
func (lu *ListingUpdate) AddInputs(a ...*Asset) *ListingUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.AddInputIDs(ids...)
}

// AddOutputIDs adds the "outputs" edge to the Asset entity by IDs.
func (lu *ListingUpdate) AddOutputIDs(ids ...string) *ListingUpdate {
	lu.mutation.AddOutputIDs(ids...)
	return lu
}

// AddOutputs adds the "outputs" edges to the Asset entity.
func (lu *ListingUpdate) AddOutputs(a ...*Asset) *ListingUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.AddOutputIDs(ids...)
}

// Mutation returns the ListingMutation object of the builder.
func (lu *ListingUpdate) Mutation() *ListingMutation {
	return lu.mutation
}

// ClearDope clears the "dope" edge to the Dope entity.
func (lu *ListingUpdate) ClearDope() *ListingUpdate {
	lu.mutation.ClearDope()
	return lu
}

// ClearDopeLastsales clears all "dope_lastsales" edges to the Dope entity.
func (lu *ListingUpdate) ClearDopeLastsales() *ListingUpdate {
	lu.mutation.ClearDopeLastsales()
	return lu
}

// RemoveDopeLastsaleIDs removes the "dope_lastsales" edge to Dope entities by IDs.
func (lu *ListingUpdate) RemoveDopeLastsaleIDs(ids ...string) *ListingUpdate {
	lu.mutation.RemoveDopeLastsaleIDs(ids...)
	return lu
}

// RemoveDopeLastsales removes "dope_lastsales" edges to Dope entities.
func (lu *ListingUpdate) RemoveDopeLastsales(d ...*Dope) *ListingUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return lu.RemoveDopeLastsaleIDs(ids...)
}

// ClearInputs clears all "inputs" edges to the Asset entity.
func (lu *ListingUpdate) ClearInputs() *ListingUpdate {
	lu.mutation.ClearInputs()
	return lu
}

// RemoveInputIDs removes the "inputs" edge to Asset entities by IDs.
func (lu *ListingUpdate) RemoveInputIDs(ids ...string) *ListingUpdate {
	lu.mutation.RemoveInputIDs(ids...)
	return lu
}

// RemoveInputs removes "inputs" edges to Asset entities.
func (lu *ListingUpdate) RemoveInputs(a ...*Asset) *ListingUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.RemoveInputIDs(ids...)
}

// ClearOutputs clears all "outputs" edges to the Asset entity.
func (lu *ListingUpdate) ClearOutputs() *ListingUpdate {
	lu.mutation.ClearOutputs()
	return lu
}

// RemoveOutputIDs removes the "outputs" edge to Asset entities by IDs.
func (lu *ListingUpdate) RemoveOutputIDs(ids ...string) *ListingUpdate {
	lu.mutation.RemoveOutputIDs(ids...)
	return lu
}

// RemoveOutputs removes "outputs" edges to Asset entities.
func (lu *ListingUpdate) RemoveOutputs(a ...*Asset) *ListingUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return lu.RemoveOutputIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *ListingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *ListingUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *ListingUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *ListingUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *ListingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   listing.Table,
			Columns: listing.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: listing.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: listing.FieldActive,
		})
	}
	if lu.mutation.DopeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   listing.DopeTable,
			Columns: []string{listing.DopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.DopeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   listing.DopeTable,
			Columns: []string{listing.DopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.DopeLastsalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedDopeLastsalesIDs(); len(nodes) > 0 && !lu.mutation.DopeLastsalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.DopeLastsalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.InputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedInputsIDs(); len(nodes) > 0 && !lu.mutation.InputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.InputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.OutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedOutputsIDs(); len(nodes) > 0 && !lu.mutation.OutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{listing.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ListingUpdateOne is the builder for updating a single Listing entity.
type ListingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ListingMutation
}

// SetActive sets the "active" field.
func (luo *ListingUpdateOne) SetActive(b bool) *ListingUpdateOne {
	luo.mutation.SetActive(b)
	return luo
}

// SetDopeID sets the "dope" edge to the Dope entity by ID.
func (luo *ListingUpdateOne) SetDopeID(id string) *ListingUpdateOne {
	luo.mutation.SetDopeID(id)
	return luo
}

// SetNillableDopeID sets the "dope" edge to the Dope entity by ID if the given value is not nil.
func (luo *ListingUpdateOne) SetNillableDopeID(id *string) *ListingUpdateOne {
	if id != nil {
		luo = luo.SetDopeID(*id)
	}
	return luo
}

// SetDope sets the "dope" edge to the Dope entity.
func (luo *ListingUpdateOne) SetDope(d *Dope) *ListingUpdateOne {
	return luo.SetDopeID(d.ID)
}

// AddDopeLastsaleIDs adds the "dope_lastsales" edge to the Dope entity by IDs.
func (luo *ListingUpdateOne) AddDopeLastsaleIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.AddDopeLastsaleIDs(ids...)
	return luo
}

// AddDopeLastsales adds the "dope_lastsales" edges to the Dope entity.
func (luo *ListingUpdateOne) AddDopeLastsales(d ...*Dope) *ListingUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return luo.AddDopeLastsaleIDs(ids...)
}

// AddInputIDs adds the "inputs" edge to the Asset entity by IDs.
func (luo *ListingUpdateOne) AddInputIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.AddInputIDs(ids...)
	return luo
}

// AddInputs adds the "inputs" edges to the Asset entity.
func (luo *ListingUpdateOne) AddInputs(a ...*Asset) *ListingUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.AddInputIDs(ids...)
}

// AddOutputIDs adds the "outputs" edge to the Asset entity by IDs.
func (luo *ListingUpdateOne) AddOutputIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.AddOutputIDs(ids...)
	return luo
}

// AddOutputs adds the "outputs" edges to the Asset entity.
func (luo *ListingUpdateOne) AddOutputs(a ...*Asset) *ListingUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.AddOutputIDs(ids...)
}

// Mutation returns the ListingMutation object of the builder.
func (luo *ListingUpdateOne) Mutation() *ListingMutation {
	return luo.mutation
}

// ClearDope clears the "dope" edge to the Dope entity.
func (luo *ListingUpdateOne) ClearDope() *ListingUpdateOne {
	luo.mutation.ClearDope()
	return luo
}

// ClearDopeLastsales clears all "dope_lastsales" edges to the Dope entity.
func (luo *ListingUpdateOne) ClearDopeLastsales() *ListingUpdateOne {
	luo.mutation.ClearDopeLastsales()
	return luo
}

// RemoveDopeLastsaleIDs removes the "dope_lastsales" edge to Dope entities by IDs.
func (luo *ListingUpdateOne) RemoveDopeLastsaleIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.RemoveDopeLastsaleIDs(ids...)
	return luo
}

// RemoveDopeLastsales removes "dope_lastsales" edges to Dope entities.
func (luo *ListingUpdateOne) RemoveDopeLastsales(d ...*Dope) *ListingUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return luo.RemoveDopeLastsaleIDs(ids...)
}

// ClearInputs clears all "inputs" edges to the Asset entity.
func (luo *ListingUpdateOne) ClearInputs() *ListingUpdateOne {
	luo.mutation.ClearInputs()
	return luo
}

// RemoveInputIDs removes the "inputs" edge to Asset entities by IDs.
func (luo *ListingUpdateOne) RemoveInputIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.RemoveInputIDs(ids...)
	return luo
}

// RemoveInputs removes "inputs" edges to Asset entities.
func (luo *ListingUpdateOne) RemoveInputs(a ...*Asset) *ListingUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.RemoveInputIDs(ids...)
}

// ClearOutputs clears all "outputs" edges to the Asset entity.
func (luo *ListingUpdateOne) ClearOutputs() *ListingUpdateOne {
	luo.mutation.ClearOutputs()
	return luo
}

// RemoveOutputIDs removes the "outputs" edge to Asset entities by IDs.
func (luo *ListingUpdateOne) RemoveOutputIDs(ids ...string) *ListingUpdateOne {
	luo.mutation.RemoveOutputIDs(ids...)
	return luo
}

// RemoveOutputs removes "outputs" edges to Asset entities.
func (luo *ListingUpdateOne) RemoveOutputs(a ...*Asset) *ListingUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return luo.RemoveOutputIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *ListingUpdateOne) Select(field string, fields ...string) *ListingUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Listing entity.
func (luo *ListingUpdateOne) Save(ctx context.Context) (*Listing, error) {
	var (
		err  error
		node *Listing
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *ListingUpdateOne) SaveX(ctx context.Context) *Listing {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *ListingUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *ListingUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *ListingUpdateOne) sqlSave(ctx context.Context) (_node *Listing, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   listing.Table,
			Columns: listing.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: listing.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Listing.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, listing.FieldID)
		for _, f := range fields {
			if !listing.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != listing.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: listing.FieldActive,
		})
	}
	if luo.mutation.DopeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   listing.DopeTable,
			Columns: []string{listing.DopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.DopeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   listing.DopeTable,
			Columns: []string{listing.DopeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.DopeLastsalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedDopeLastsalesIDs(); len(nodes) > 0 && !luo.mutation.DopeLastsalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.DopeLastsalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.DopeLastsalesTable,
			Columns: []string{listing.DopeLastsalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dope.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.InputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedInputsIDs(); len(nodes) > 0 && !luo.mutation.InputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.InputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.InputsTable,
			Columns: []string{listing.InputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.OutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedOutputsIDs(); len(nodes) > 0 && !luo.mutation.OutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   listing.OutputsTable,
			Columns: []string{listing.OutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: asset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Listing{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{listing.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
