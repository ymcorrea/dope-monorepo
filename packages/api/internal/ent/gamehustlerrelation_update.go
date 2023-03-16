// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustler"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustlerrelation"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
)

// GameHustlerRelationUpdate is the builder for updating GameHustlerRelation entities.
type GameHustlerRelationUpdate struct {
	config
	hooks    []Hook
	mutation *GameHustlerRelationMutation
}

// Where appends a list predicates to the GameHustlerRelationUpdate builder.
func (ghru *GameHustlerRelationUpdate) Where(ps ...predicate.GameHustlerRelation) *GameHustlerRelationUpdate {
	ghru.mutation.Where(ps...)
	return ghru
}

// SetCitizen sets the "citizen" field.
func (ghru *GameHustlerRelationUpdate) SetCitizen(s string) *GameHustlerRelationUpdate {
	ghru.mutation.SetCitizen(s)
	return ghru
}

// SetConversation sets the "conversation" field.
func (ghru *GameHustlerRelationUpdate) SetConversation(s string) *GameHustlerRelationUpdate {
	ghru.mutation.SetConversation(s)
	return ghru
}

// SetText sets the "text" field.
func (ghru *GameHustlerRelationUpdate) SetText(u uint) *GameHustlerRelationUpdate {
	ghru.mutation.ResetText()
	ghru.mutation.SetText(u)
	return ghru
}

// AddText adds u to the "text" field.
func (ghru *GameHustlerRelationUpdate) AddText(u int) *GameHustlerRelationUpdate {
	ghru.mutation.AddText(u)
	return ghru
}

// SetHustlerID sets the "hustler" edge to the GameHustler entity by ID.
func (ghru *GameHustlerRelationUpdate) SetHustlerID(id string) *GameHustlerRelationUpdate {
	ghru.mutation.SetHustlerID(id)
	return ghru
}

// SetNillableHustlerID sets the "hustler" edge to the GameHustler entity by ID if the given value is not nil.
func (ghru *GameHustlerRelationUpdate) SetNillableHustlerID(id *string) *GameHustlerRelationUpdate {
	if id != nil {
		ghru = ghru.SetHustlerID(*id)
	}
	return ghru
}

// SetHustler sets the "hustler" edge to the GameHustler entity.
func (ghru *GameHustlerRelationUpdate) SetHustler(g *GameHustler) *GameHustlerRelationUpdate {
	return ghru.SetHustlerID(g.ID)
}

// Mutation returns the GameHustlerRelationMutation object of the builder.
func (ghru *GameHustlerRelationUpdate) Mutation() *GameHustlerRelationMutation {
	return ghru.mutation
}

// ClearHustler clears the "hustler" edge to the GameHustler entity.
func (ghru *GameHustlerRelationUpdate) ClearHustler() *GameHustlerRelationUpdate {
	ghru.mutation.ClearHustler()
	return ghru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ghru *GameHustlerRelationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GameHustlerRelationMutation](ctx, ghru.sqlSave, ghru.mutation, ghru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghru *GameHustlerRelationUpdate) SaveX(ctx context.Context) int {
	affected, err := ghru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ghru *GameHustlerRelationUpdate) Exec(ctx context.Context) error {
	_, err := ghru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghru *GameHustlerRelationUpdate) ExecX(ctx context.Context) {
	if err := ghru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghru *GameHustlerRelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamehustlerrelation.Table, gamehustlerrelation.Columns, sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString))
	if ps := ghru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ghru.mutation.Citizen(); ok {
		_spec.SetField(gamehustlerrelation.FieldCitizen, field.TypeString, value)
	}
	if value, ok := ghru.mutation.Conversation(); ok {
		_spec.SetField(gamehustlerrelation.FieldConversation, field.TypeString, value)
	}
	if value, ok := ghru.mutation.Text(); ok {
		_spec.SetField(gamehustlerrelation.FieldText, field.TypeUint, value)
	}
	if value, ok := ghru.mutation.AddedText(); ok {
		_spec.AddField(gamehustlerrelation.FieldText, field.TypeUint, value)
	}
	if ghru.mutation.HustlerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamehustlerrelation.HustlerTable,
			Columns: []string{gamehustlerrelation.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghru.mutation.HustlerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamehustlerrelation.HustlerTable,
			Columns: []string{gamehustlerrelation.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ghru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamehustlerrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ghru.mutation.done = true
	return n, nil
}

// GameHustlerRelationUpdateOne is the builder for updating a single GameHustlerRelation entity.
type GameHustlerRelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameHustlerRelationMutation
}

// SetCitizen sets the "citizen" field.
func (ghruo *GameHustlerRelationUpdateOne) SetCitizen(s string) *GameHustlerRelationUpdateOne {
	ghruo.mutation.SetCitizen(s)
	return ghruo
}

// SetConversation sets the "conversation" field.
func (ghruo *GameHustlerRelationUpdateOne) SetConversation(s string) *GameHustlerRelationUpdateOne {
	ghruo.mutation.SetConversation(s)
	return ghruo
}

// SetText sets the "text" field.
func (ghruo *GameHustlerRelationUpdateOne) SetText(u uint) *GameHustlerRelationUpdateOne {
	ghruo.mutation.ResetText()
	ghruo.mutation.SetText(u)
	return ghruo
}

// AddText adds u to the "text" field.
func (ghruo *GameHustlerRelationUpdateOne) AddText(u int) *GameHustlerRelationUpdateOne {
	ghruo.mutation.AddText(u)
	return ghruo
}

// SetHustlerID sets the "hustler" edge to the GameHustler entity by ID.
func (ghruo *GameHustlerRelationUpdateOne) SetHustlerID(id string) *GameHustlerRelationUpdateOne {
	ghruo.mutation.SetHustlerID(id)
	return ghruo
}

// SetNillableHustlerID sets the "hustler" edge to the GameHustler entity by ID if the given value is not nil.
func (ghruo *GameHustlerRelationUpdateOne) SetNillableHustlerID(id *string) *GameHustlerRelationUpdateOne {
	if id != nil {
		ghruo = ghruo.SetHustlerID(*id)
	}
	return ghruo
}

// SetHustler sets the "hustler" edge to the GameHustler entity.
func (ghruo *GameHustlerRelationUpdateOne) SetHustler(g *GameHustler) *GameHustlerRelationUpdateOne {
	return ghruo.SetHustlerID(g.ID)
}

// Mutation returns the GameHustlerRelationMutation object of the builder.
func (ghruo *GameHustlerRelationUpdateOne) Mutation() *GameHustlerRelationMutation {
	return ghruo.mutation
}

// ClearHustler clears the "hustler" edge to the GameHustler entity.
func (ghruo *GameHustlerRelationUpdateOne) ClearHustler() *GameHustlerRelationUpdateOne {
	ghruo.mutation.ClearHustler()
	return ghruo
}

// Where appends a list predicates to the GameHustlerRelationUpdate builder.
func (ghruo *GameHustlerRelationUpdateOne) Where(ps ...predicate.GameHustlerRelation) *GameHustlerRelationUpdateOne {
	ghruo.mutation.Where(ps...)
	return ghruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ghruo *GameHustlerRelationUpdateOne) Select(field string, fields ...string) *GameHustlerRelationUpdateOne {
	ghruo.fields = append([]string{field}, fields...)
	return ghruo
}

// Save executes the query and returns the updated GameHustlerRelation entity.
func (ghruo *GameHustlerRelationUpdateOne) Save(ctx context.Context) (*GameHustlerRelation, error) {
	return withHooks[*GameHustlerRelation, GameHustlerRelationMutation](ctx, ghruo.sqlSave, ghruo.mutation, ghruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghruo *GameHustlerRelationUpdateOne) SaveX(ctx context.Context) *GameHustlerRelation {
	node, err := ghruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ghruo *GameHustlerRelationUpdateOne) Exec(ctx context.Context) error {
	_, err := ghruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghruo *GameHustlerRelationUpdateOne) ExecX(ctx context.Context) {
	if err := ghruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghruo *GameHustlerRelationUpdateOne) sqlSave(ctx context.Context) (_node *GameHustlerRelation, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamehustlerrelation.Table, gamehustlerrelation.Columns, sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString))
	id, ok := ghruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GameHustlerRelation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ghruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gamehustlerrelation.FieldID)
		for _, f := range fields {
			if !gamehustlerrelation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gamehustlerrelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ghruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ghruo.mutation.Citizen(); ok {
		_spec.SetField(gamehustlerrelation.FieldCitizen, field.TypeString, value)
	}
	if value, ok := ghruo.mutation.Conversation(); ok {
		_spec.SetField(gamehustlerrelation.FieldConversation, field.TypeString, value)
	}
	if value, ok := ghruo.mutation.Text(); ok {
		_spec.SetField(gamehustlerrelation.FieldText, field.TypeUint, value)
	}
	if value, ok := ghruo.mutation.AddedText(); ok {
		_spec.AddField(gamehustlerrelation.FieldText, field.TypeUint, value)
	}
	if ghruo.mutation.HustlerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamehustlerrelation.HustlerTable,
			Columns: []string{gamehustlerrelation.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghruo.mutation.HustlerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gamehustlerrelation.HustlerTable,
			Columns: []string{gamehustlerrelation.HustlerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GameHustlerRelation{config: ghruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ghruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamehustlerrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ghruo.mutation.done = true
	return _node, nil
}
