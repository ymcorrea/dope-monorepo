// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/bodypart"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
)

// BodyPartQuery is the builder for querying BodyPart entities.
type BodyPartQuery struct {
	config
	ctx                    *QueryContext
	order                  []OrderFunc
	inters                 []Interceptor
	predicates             []predicate.BodyPart
	withHustlerBodies      *HustlerQuery
	withHustlerHairs       *HustlerQuery
	withHustlerBeards      *HustlerQuery
	modifiers              []func(*sql.Selector)
	loadTotal              []func(context.Context, []*BodyPart) error
	withNamedHustlerBodies map[string]*HustlerQuery
	withNamedHustlerHairs  map[string]*HustlerQuery
	withNamedHustlerBeards map[string]*HustlerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BodyPartQuery builder.
func (bpq *BodyPartQuery) Where(ps ...predicate.BodyPart) *BodyPartQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit the number of records to be returned by this query.
func (bpq *BodyPartQuery) Limit(limit int) *BodyPartQuery {
	bpq.ctx.Limit = &limit
	return bpq
}

// Offset to start from.
func (bpq *BodyPartQuery) Offset(offset int) *BodyPartQuery {
	bpq.ctx.Offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BodyPartQuery) Unique(unique bool) *BodyPartQuery {
	bpq.ctx.Unique = &unique
	return bpq
}

// Order specifies how the records should be ordered.
func (bpq *BodyPartQuery) Order(o ...OrderFunc) *BodyPartQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryHustlerBodies chains the current query on the "hustler_bodies" edge.
func (bpq *BodyPartQuery) QueryHustlerBodies() *HustlerQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bodypart.Table, bodypart.FieldID, selector),
			sqlgraph.To(hustler.Table, hustler.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bodypart.HustlerBodiesTable, bodypart.HustlerBodiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHustlerHairs chains the current query on the "hustler_hairs" edge.
func (bpq *BodyPartQuery) QueryHustlerHairs() *HustlerQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bodypart.Table, bodypart.FieldID, selector),
			sqlgraph.To(hustler.Table, hustler.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bodypart.HustlerHairsTable, bodypart.HustlerHairsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHustlerBeards chains the current query on the "hustler_beards" edge.
func (bpq *BodyPartQuery) QueryHustlerBeards() *HustlerQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bodypart.Table, bodypart.FieldID, selector),
			sqlgraph.To(hustler.Table, hustler.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bodypart.HustlerBeardsTable, bodypart.HustlerBeardsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BodyPart entity from the query.
// Returns a *NotFoundError when no BodyPart was found.
func (bpq *BodyPartQuery) First(ctx context.Context) (*BodyPart, error) {
	nodes, err := bpq.Limit(1).All(setContextOp(ctx, bpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bodypart.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BodyPartQuery) FirstX(ctx context.Context) *BodyPart {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BodyPart ID from the query.
// Returns a *NotFoundError when no BodyPart ID was found.
func (bpq *BodyPartQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bpq.Limit(1).IDs(setContextOp(ctx, bpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bodypart.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BodyPartQuery) FirstIDX(ctx context.Context) string {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BodyPart entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BodyPart entity is found.
// Returns a *NotFoundError when no BodyPart entities are found.
func (bpq *BodyPartQuery) Only(ctx context.Context) (*BodyPart, error) {
	nodes, err := bpq.Limit(2).All(setContextOp(ctx, bpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bodypart.Label}
	default:
		return nil, &NotSingularError{bodypart.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BodyPartQuery) OnlyX(ctx context.Context) *BodyPart {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BodyPart ID in the query.
// Returns a *NotSingularError when more than one BodyPart ID is found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BodyPartQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bpq.Limit(2).IDs(setContextOp(ctx, bpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bodypart.Label}
	default:
		err = &NotSingularError{bodypart.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BodyPartQuery) OnlyIDX(ctx context.Context) string {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BodyParts.
func (bpq *BodyPartQuery) All(ctx context.Context) ([]*BodyPart, error) {
	ctx = setContextOp(ctx, bpq.ctx, "All")
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BodyPart, *BodyPartQuery]()
	return withInterceptors[[]*BodyPart](ctx, bpq, qr, bpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BodyPartQuery) AllX(ctx context.Context) []*BodyPart {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BodyPart IDs.
func (bpq *BodyPartQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bpq.ctx.Unique == nil && bpq.path != nil {
		bpq.Unique(true)
	}
	ctx = setContextOp(ctx, bpq.ctx, "IDs")
	if err = bpq.Select(bodypart.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BodyPartQuery) IDsX(ctx context.Context) []string {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BodyPartQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bpq.ctx, "Count")
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bpq, querierCount[*BodyPartQuery](), bpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BodyPartQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BodyPartQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bpq.ctx, "Exist")
	switch _, err := bpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BodyPartQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BodyPartQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BodyPartQuery) Clone() *BodyPartQuery {
	if bpq == nil {
		return nil
	}
	return &BodyPartQuery{
		config:            bpq.config,
		ctx:               bpq.ctx.Clone(),
		order:             append([]OrderFunc{}, bpq.order...),
		inters:            append([]Interceptor{}, bpq.inters...),
		predicates:        append([]predicate.BodyPart{}, bpq.predicates...),
		withHustlerBodies: bpq.withHustlerBodies.Clone(),
		withHustlerHairs:  bpq.withHustlerHairs.Clone(),
		withHustlerBeards: bpq.withHustlerBeards.Clone(),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// WithHustlerBodies tells the query-builder to eager-load the nodes that are connected to
// the "hustler_bodies" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithHustlerBodies(opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withHustlerBodies = query
	return bpq
}

// WithHustlerHairs tells the query-builder to eager-load the nodes that are connected to
// the "hustler_hairs" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithHustlerHairs(opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withHustlerHairs = query
	return bpq
}

// WithHustlerBeards tells the query-builder to eager-load the nodes that are connected to
// the "hustler_beards" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithHustlerBeards(opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withHustlerBeards = query
	return bpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type bodypart.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BodyPart.Query().
//		GroupBy(bodypart.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bpq *BodyPartQuery) GroupBy(field string, fields ...string) *BodyPartGroupBy {
	bpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BodyPartGroupBy{build: bpq}
	grbuild.flds = &bpq.ctx.Fields
	grbuild.label = bodypart.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type bodypart.Type `json:"type,omitempty"`
//	}
//
//	client.BodyPart.Query().
//		Select(bodypart.FieldType).
//		Scan(ctx, &v)
func (bpq *BodyPartQuery) Select(fields ...string) *BodyPartSelect {
	bpq.ctx.Fields = append(bpq.ctx.Fields, fields...)
	sbuild := &BodyPartSelect{BodyPartQuery: bpq}
	sbuild.label = bodypart.Label
	sbuild.flds, sbuild.scan = &bpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BodyPartSelect configured with the given aggregations.
func (bpq *BodyPartQuery) Aggregate(fns ...AggregateFunc) *BodyPartSelect {
	return bpq.Select().Aggregate(fns...)
}

func (bpq *BodyPartQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bpq); err != nil {
				return err
			}
		}
	}
	for _, f := range bpq.ctx.Fields {
		if !bodypart.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BodyPartQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BodyPart, error) {
	var (
		nodes       = []*BodyPart{}
		_spec       = bpq.querySpec()
		loadedTypes = [3]bool{
			bpq.withHustlerBodies != nil,
			bpq.withHustlerHairs != nil,
			bpq.withHustlerBeards != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BodyPart).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BodyPart{config: bpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bpq.modifiers) > 0 {
		_spec.Modifiers = bpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bpq.withHustlerBodies; query != nil {
		if err := bpq.loadHustlerBodies(ctx, query, nodes,
			func(n *BodyPart) { n.Edges.HustlerBodies = []*Hustler{} },
			func(n *BodyPart, e *Hustler) { n.Edges.HustlerBodies = append(n.Edges.HustlerBodies, e) }); err != nil {
			return nil, err
		}
	}
	if query := bpq.withHustlerHairs; query != nil {
		if err := bpq.loadHustlerHairs(ctx, query, nodes,
			func(n *BodyPart) { n.Edges.HustlerHairs = []*Hustler{} },
			func(n *BodyPart, e *Hustler) { n.Edges.HustlerHairs = append(n.Edges.HustlerHairs, e) }); err != nil {
			return nil, err
		}
	}
	if query := bpq.withHustlerBeards; query != nil {
		if err := bpq.loadHustlerBeards(ctx, query, nodes,
			func(n *BodyPart) { n.Edges.HustlerBeards = []*Hustler{} },
			func(n *BodyPart, e *Hustler) { n.Edges.HustlerBeards = append(n.Edges.HustlerBeards, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bpq.withNamedHustlerBodies {
		if err := bpq.loadHustlerBodies(ctx, query, nodes,
			func(n *BodyPart) { n.appendNamedHustlerBodies(name) },
			func(n *BodyPart, e *Hustler) { n.appendNamedHustlerBodies(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bpq.withNamedHustlerHairs {
		if err := bpq.loadHustlerHairs(ctx, query, nodes,
			func(n *BodyPart) { n.appendNamedHustlerHairs(name) },
			func(n *BodyPart, e *Hustler) { n.appendNamedHustlerHairs(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bpq.withNamedHustlerBeards {
		if err := bpq.loadHustlerBeards(ctx, query, nodes,
			func(n *BodyPart) { n.appendNamedHustlerBeards(name) },
			func(n *BodyPart, e *Hustler) { n.appendNamedHustlerBeards(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bpq.loadTotal {
		if err := bpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bpq *BodyPartQuery) loadHustlerBodies(ctx context.Context, query *HustlerQuery, nodes []*BodyPart, init func(*BodyPart), assign func(*BodyPart, *Hustler)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*BodyPart)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Hustler(func(s *sql.Selector) {
		s.Where(sql.InValues(bodypart.HustlerBodiesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.body_part_hustler_bodies
		if fk == nil {
			return fmt.Errorf(`foreign-key "body_part_hustler_bodies" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "body_part_hustler_bodies" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bpq *BodyPartQuery) loadHustlerHairs(ctx context.Context, query *HustlerQuery, nodes []*BodyPart, init func(*BodyPart), assign func(*BodyPart, *Hustler)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*BodyPart)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Hustler(func(s *sql.Selector) {
		s.Where(sql.InValues(bodypart.HustlerHairsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.body_part_hustler_hairs
		if fk == nil {
			return fmt.Errorf(`foreign-key "body_part_hustler_hairs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "body_part_hustler_hairs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bpq *BodyPartQuery) loadHustlerBeards(ctx context.Context, query *HustlerQuery, nodes []*BodyPart, init func(*BodyPart), assign func(*BodyPart, *Hustler)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*BodyPart)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Hustler(func(s *sql.Selector) {
		s.Where(sql.InValues(bodypart.HustlerBeardsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.body_part_hustler_beards
		if fk == nil {
			return fmt.Errorf(`foreign-key "body_part_hustler_beards" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "body_part_hustler_beards" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bpq *BodyPartQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	if len(bpq.modifiers) > 0 {
		_spec.Modifiers = bpq.modifiers
	}
	_spec.Node.Columns = bpq.ctx.Fields
	if len(bpq.ctx.Fields) > 0 {
		_spec.Unique = bpq.ctx.Unique != nil && *bpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BodyPartQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bodypart.Table, bodypart.Columns, sqlgraph.NewFieldSpec(bodypart.FieldID, field.TypeString))
	_spec.From = bpq.sql
	if unique := bpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bpq.path != nil {
		_spec.Unique = true
	}
	if fields := bpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bodypart.FieldID)
		for i := range fields {
			if fields[i] != bodypart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BodyPartQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(bodypart.Table)
	columns := bpq.ctx.Fields
	if len(columns) == 0 {
		columns = bodypart.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.ctx.Unique != nil && *bpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedHustlerBodies tells the query-builder to eager-load the nodes that are connected to the "hustler_bodies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithNamedHustlerBodies(name string, opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bpq.withNamedHustlerBodies == nil {
		bpq.withNamedHustlerBodies = make(map[string]*HustlerQuery)
	}
	bpq.withNamedHustlerBodies[name] = query
	return bpq
}

// WithNamedHustlerHairs tells the query-builder to eager-load the nodes that are connected to the "hustler_hairs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithNamedHustlerHairs(name string, opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bpq.withNamedHustlerHairs == nil {
		bpq.withNamedHustlerHairs = make(map[string]*HustlerQuery)
	}
	bpq.withNamedHustlerHairs[name] = query
	return bpq
}

// WithNamedHustlerBeards tells the query-builder to eager-load the nodes that are connected to the "hustler_beards"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bpq *BodyPartQuery) WithNamedHustlerBeards(name string, opts ...func(*HustlerQuery)) *BodyPartQuery {
	query := (&HustlerClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bpq.withNamedHustlerBeards == nil {
		bpq.withNamedHustlerBeards = make(map[string]*HustlerQuery)
	}
	bpq.withNamedHustlerBeards[name] = query
	return bpq
}

// BodyPartGroupBy is the group-by builder for BodyPart entities.
type BodyPartGroupBy struct {
	selector
	build *BodyPartQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BodyPartGroupBy) Aggregate(fns ...AggregateFunc) *BodyPartGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the selector query and scans the result into the given value.
func (bpgb *BodyPartGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bpgb.build.ctx, "GroupBy")
	if err := bpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BodyPartQuery, *BodyPartGroupBy](ctx, bpgb.build, bpgb, bpgb.build.inters, v)
}

func (bpgb *BodyPartGroupBy) sqlScan(ctx context.Context, root *BodyPartQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bpgb.flds)+len(bpgb.fns))
		for _, f := range *bpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BodyPartSelect is the builder for selecting fields of BodyPart entities.
type BodyPartSelect struct {
	*BodyPartQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bps *BodyPartSelect) Aggregate(fns ...AggregateFunc) *BodyPartSelect {
	bps.fns = append(bps.fns, fns...)
	return bps
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BodyPartSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bps.ctx, "Select")
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BodyPartQuery, *BodyPartSelect](ctx, bps.BodyPartQuery, bps, bps.inters, v)
}

func (bps *BodyPartSelect) sqlScan(ctx context.Context, root *BodyPartQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bps.fns))
	for _, fn := range bps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
