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
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustleritem"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustlerquest"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustlerrelation"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/schema"
)

// GameHustlerUpdate is the builder for updating GameHustler entities.
type GameHustlerUpdate struct {
	config
	hooks    []Hook
	mutation *GameHustlerMutation
}

// Where appends a list predicates to the GameHustlerUpdate builder.
func (ghu *GameHustlerUpdate) Where(ps ...predicate.GameHustler) *GameHustlerUpdate {
	ghu.mutation.Where(ps...)
	return ghu
}

// SetLastPosition sets the "last_position" field.
func (ghu *GameHustlerUpdate) SetLastPosition(s schema.Position) *GameHustlerUpdate {
	ghu.mutation.SetLastPosition(s)
	return ghu
}

// AddRelationIDs adds the "relations" edge to the GameHustlerRelation entity by IDs.
func (ghu *GameHustlerUpdate) AddRelationIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.AddRelationIDs(ids...)
	return ghu
}

// AddRelations adds the "relations" edges to the GameHustlerRelation entity.
func (ghu *GameHustlerUpdate) AddRelations(g ...*GameHustlerRelation) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.AddRelationIDs(ids...)
}

// AddItemIDs adds the "items" edge to the GameHustlerItem entity by IDs.
func (ghu *GameHustlerUpdate) AddItemIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.AddItemIDs(ids...)
	return ghu
}

// AddItems adds the "items" edges to the GameHustlerItem entity.
func (ghu *GameHustlerUpdate) AddItems(g ...*GameHustlerItem) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.AddItemIDs(ids...)
}

// AddQuestIDs adds the "quests" edge to the GameHustlerQuest entity by IDs.
func (ghu *GameHustlerUpdate) AddQuestIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.AddQuestIDs(ids...)
	return ghu
}

// AddQuests adds the "quests" edges to the GameHustlerQuest entity.
func (ghu *GameHustlerUpdate) AddQuests(g ...*GameHustlerQuest) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.AddQuestIDs(ids...)
}

// Mutation returns the GameHustlerMutation object of the builder.
func (ghu *GameHustlerUpdate) Mutation() *GameHustlerMutation {
	return ghu.mutation
}

// ClearRelations clears all "relations" edges to the GameHustlerRelation entity.
func (ghu *GameHustlerUpdate) ClearRelations() *GameHustlerUpdate {
	ghu.mutation.ClearRelations()
	return ghu
}

// RemoveRelationIDs removes the "relations" edge to GameHustlerRelation entities by IDs.
func (ghu *GameHustlerUpdate) RemoveRelationIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.RemoveRelationIDs(ids...)
	return ghu
}

// RemoveRelations removes "relations" edges to GameHustlerRelation entities.
func (ghu *GameHustlerUpdate) RemoveRelations(g ...*GameHustlerRelation) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.RemoveRelationIDs(ids...)
}

// ClearItems clears all "items" edges to the GameHustlerItem entity.
func (ghu *GameHustlerUpdate) ClearItems() *GameHustlerUpdate {
	ghu.mutation.ClearItems()
	return ghu
}

// RemoveItemIDs removes the "items" edge to GameHustlerItem entities by IDs.
func (ghu *GameHustlerUpdate) RemoveItemIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.RemoveItemIDs(ids...)
	return ghu
}

// RemoveItems removes "items" edges to GameHustlerItem entities.
func (ghu *GameHustlerUpdate) RemoveItems(g ...*GameHustlerItem) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.RemoveItemIDs(ids...)
}

// ClearQuests clears all "quests" edges to the GameHustlerQuest entity.
func (ghu *GameHustlerUpdate) ClearQuests() *GameHustlerUpdate {
	ghu.mutation.ClearQuests()
	return ghu
}

// RemoveQuestIDs removes the "quests" edge to GameHustlerQuest entities by IDs.
func (ghu *GameHustlerUpdate) RemoveQuestIDs(ids ...string) *GameHustlerUpdate {
	ghu.mutation.RemoveQuestIDs(ids...)
	return ghu
}

// RemoveQuests removes "quests" edges to GameHustlerQuest entities.
func (ghu *GameHustlerUpdate) RemoveQuests(g ...*GameHustlerQuest) *GameHustlerUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghu.RemoveQuestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ghu *GameHustlerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GameHustlerMutation](ctx, ghu.sqlSave, ghu.mutation, ghu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghu *GameHustlerUpdate) SaveX(ctx context.Context) int {
	affected, err := ghu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ghu *GameHustlerUpdate) Exec(ctx context.Context) error {
	_, err := ghu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghu *GameHustlerUpdate) ExecX(ctx context.Context) {
	if err := ghu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghu *GameHustlerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamehustler.Table, gamehustler.Columns, sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString))
	if ps := ghu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ghu.mutation.LastPosition(); ok {
		_spec.SetField(gamehustler.FieldLastPosition, field.TypeJSON, value)
	}
	if ghu.mutation.RelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.RemovedRelationsIDs(); len(nodes) > 0 && !ghu.mutation.RelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.RelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ghu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !ghu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ghu.mutation.QuestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.RemovedQuestsIDs(); len(nodes) > 0 && !ghu.mutation.QuestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghu.mutation.QuestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ghu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamehustler.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ghu.mutation.done = true
	return n, nil
}

// GameHustlerUpdateOne is the builder for updating a single GameHustler entity.
type GameHustlerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameHustlerMutation
}

// SetLastPosition sets the "last_position" field.
func (ghuo *GameHustlerUpdateOne) SetLastPosition(s schema.Position) *GameHustlerUpdateOne {
	ghuo.mutation.SetLastPosition(s)
	return ghuo
}

// AddRelationIDs adds the "relations" edge to the GameHustlerRelation entity by IDs.
func (ghuo *GameHustlerUpdateOne) AddRelationIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.AddRelationIDs(ids...)
	return ghuo
}

// AddRelations adds the "relations" edges to the GameHustlerRelation entity.
func (ghuo *GameHustlerUpdateOne) AddRelations(g ...*GameHustlerRelation) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.AddRelationIDs(ids...)
}

// AddItemIDs adds the "items" edge to the GameHustlerItem entity by IDs.
func (ghuo *GameHustlerUpdateOne) AddItemIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.AddItemIDs(ids...)
	return ghuo
}

// AddItems adds the "items" edges to the GameHustlerItem entity.
func (ghuo *GameHustlerUpdateOne) AddItems(g ...*GameHustlerItem) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.AddItemIDs(ids...)
}

// AddQuestIDs adds the "quests" edge to the GameHustlerQuest entity by IDs.
func (ghuo *GameHustlerUpdateOne) AddQuestIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.AddQuestIDs(ids...)
	return ghuo
}

// AddQuests adds the "quests" edges to the GameHustlerQuest entity.
func (ghuo *GameHustlerUpdateOne) AddQuests(g ...*GameHustlerQuest) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.AddQuestIDs(ids...)
}

// Mutation returns the GameHustlerMutation object of the builder.
func (ghuo *GameHustlerUpdateOne) Mutation() *GameHustlerMutation {
	return ghuo.mutation
}

// ClearRelations clears all "relations" edges to the GameHustlerRelation entity.
func (ghuo *GameHustlerUpdateOne) ClearRelations() *GameHustlerUpdateOne {
	ghuo.mutation.ClearRelations()
	return ghuo
}

// RemoveRelationIDs removes the "relations" edge to GameHustlerRelation entities by IDs.
func (ghuo *GameHustlerUpdateOne) RemoveRelationIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.RemoveRelationIDs(ids...)
	return ghuo
}

// RemoveRelations removes "relations" edges to GameHustlerRelation entities.
func (ghuo *GameHustlerUpdateOne) RemoveRelations(g ...*GameHustlerRelation) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.RemoveRelationIDs(ids...)
}

// ClearItems clears all "items" edges to the GameHustlerItem entity.
func (ghuo *GameHustlerUpdateOne) ClearItems() *GameHustlerUpdateOne {
	ghuo.mutation.ClearItems()
	return ghuo
}

// RemoveItemIDs removes the "items" edge to GameHustlerItem entities by IDs.
func (ghuo *GameHustlerUpdateOne) RemoveItemIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.RemoveItemIDs(ids...)
	return ghuo
}

// RemoveItems removes "items" edges to GameHustlerItem entities.
func (ghuo *GameHustlerUpdateOne) RemoveItems(g ...*GameHustlerItem) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.RemoveItemIDs(ids...)
}

// ClearQuests clears all "quests" edges to the GameHustlerQuest entity.
func (ghuo *GameHustlerUpdateOne) ClearQuests() *GameHustlerUpdateOne {
	ghuo.mutation.ClearQuests()
	return ghuo
}

// RemoveQuestIDs removes the "quests" edge to GameHustlerQuest entities by IDs.
func (ghuo *GameHustlerUpdateOne) RemoveQuestIDs(ids ...string) *GameHustlerUpdateOne {
	ghuo.mutation.RemoveQuestIDs(ids...)
	return ghuo
}

// RemoveQuests removes "quests" edges to GameHustlerQuest entities.
func (ghuo *GameHustlerUpdateOne) RemoveQuests(g ...*GameHustlerQuest) *GameHustlerUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return ghuo.RemoveQuestIDs(ids...)
}

// Where appends a list predicates to the GameHustlerUpdate builder.
func (ghuo *GameHustlerUpdateOne) Where(ps ...predicate.GameHustler) *GameHustlerUpdateOne {
	ghuo.mutation.Where(ps...)
	return ghuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ghuo *GameHustlerUpdateOne) Select(field string, fields ...string) *GameHustlerUpdateOne {
	ghuo.fields = append([]string{field}, fields...)
	return ghuo
}

// Save executes the query and returns the updated GameHustler entity.
func (ghuo *GameHustlerUpdateOne) Save(ctx context.Context) (*GameHustler, error) {
	return withHooks[*GameHustler, GameHustlerMutation](ctx, ghuo.sqlSave, ghuo.mutation, ghuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ghuo *GameHustlerUpdateOne) SaveX(ctx context.Context) *GameHustler {
	node, err := ghuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ghuo *GameHustlerUpdateOne) Exec(ctx context.Context) error {
	_, err := ghuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ghuo *GameHustlerUpdateOne) ExecX(ctx context.Context) {
	if err := ghuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ghuo *GameHustlerUpdateOne) sqlSave(ctx context.Context) (_node *GameHustler, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamehustler.Table, gamehustler.Columns, sqlgraph.NewFieldSpec(gamehustler.FieldID, field.TypeString))
	id, ok := ghuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GameHustler.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ghuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gamehustler.FieldID)
		for _, f := range fields {
			if !gamehustler.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gamehustler.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ghuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ghuo.mutation.LastPosition(); ok {
		_spec.SetField(gamehustler.FieldLastPosition, field.TypeJSON, value)
	}
	if ghuo.mutation.RelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.RemovedRelationsIDs(); len(nodes) > 0 && !ghuo.mutation.RelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.RelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.RelationsTable,
			Columns: []string{gamehustler.RelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ghuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !ghuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.ItemsTable,
			Columns: []string{gamehustler.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustleritem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ghuo.mutation.QuestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.RemovedQuestsIDs(); len(nodes) > 0 && !ghuo.mutation.QuestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ghuo.mutation.QuestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gamehustler.QuestsTable,
			Columns: []string{gamehustler.QuestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gamehustlerquest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GameHustler{config: ghuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ghuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamehustler.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ghuo.mutation.done = true
	return _node, nil
}
