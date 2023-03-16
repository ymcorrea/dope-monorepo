// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustler"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/gamehustlerrelation"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
)

// GameHustlerRelationQuery is the builder for querying GameHustlerRelation entities.
type GameHustlerRelationQuery struct {
	config
	ctx         *QueryContext
	order       []OrderFunc
	inters      []Interceptor
	predicates  []predicate.GameHustlerRelation
	withHustler *GameHustlerQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*GameHustlerRelation) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GameHustlerRelationQuery builder.
func (ghrq *GameHustlerRelationQuery) Where(ps ...predicate.GameHustlerRelation) *GameHustlerRelationQuery {
	ghrq.predicates = append(ghrq.predicates, ps...)
	return ghrq
}

// Limit the number of records to be returned by this query.
func (ghrq *GameHustlerRelationQuery) Limit(limit int) *GameHustlerRelationQuery {
	ghrq.ctx.Limit = &limit
	return ghrq
}

// Offset to start from.
func (ghrq *GameHustlerRelationQuery) Offset(offset int) *GameHustlerRelationQuery {
	ghrq.ctx.Offset = &offset
	return ghrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ghrq *GameHustlerRelationQuery) Unique(unique bool) *GameHustlerRelationQuery {
	ghrq.ctx.Unique = &unique
	return ghrq
}

// Order specifies how the records should be ordered.
func (ghrq *GameHustlerRelationQuery) Order(o ...OrderFunc) *GameHustlerRelationQuery {
	ghrq.order = append(ghrq.order, o...)
	return ghrq
}

// QueryHustler chains the current query on the "hustler" edge.
func (ghrq *GameHustlerRelationQuery) QueryHustler() *GameHustlerQuery {
	query := (&GameHustlerClient{config: ghrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ghrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ghrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(gamehustlerrelation.Table, gamehustlerrelation.FieldID, selector),
			sqlgraph.To(gamehustler.Table, gamehustler.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, gamehustlerrelation.HustlerTable, gamehustlerrelation.HustlerColumn),
		)
		fromU = sqlgraph.SetNeighbors(ghrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GameHustlerRelation entity from the query.
// Returns a *NotFoundError when no GameHustlerRelation was found.
func (ghrq *GameHustlerRelationQuery) First(ctx context.Context) (*GameHustlerRelation, error) {
	nodes, err := ghrq.Limit(1).All(setContextOp(ctx, ghrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gamehustlerrelation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) FirstX(ctx context.Context) *GameHustlerRelation {
	node, err := ghrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GameHustlerRelation ID from the query.
// Returns a *NotFoundError when no GameHustlerRelation ID was found.
func (ghrq *GameHustlerRelationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ghrq.Limit(1).IDs(setContextOp(ctx, ghrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gamehustlerrelation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) FirstIDX(ctx context.Context) string {
	id, err := ghrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GameHustlerRelation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GameHustlerRelation entity is found.
// Returns a *NotFoundError when no GameHustlerRelation entities are found.
func (ghrq *GameHustlerRelationQuery) Only(ctx context.Context) (*GameHustlerRelation, error) {
	nodes, err := ghrq.Limit(2).All(setContextOp(ctx, ghrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gamehustlerrelation.Label}
	default:
		return nil, &NotSingularError{gamehustlerrelation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) OnlyX(ctx context.Context) *GameHustlerRelation {
	node, err := ghrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GameHustlerRelation ID in the query.
// Returns a *NotSingularError when more than one GameHustlerRelation ID is found.
// Returns a *NotFoundError when no entities are found.
func (ghrq *GameHustlerRelationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ghrq.Limit(2).IDs(setContextOp(ctx, ghrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gamehustlerrelation.Label}
	default:
		err = &NotSingularError{gamehustlerrelation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) OnlyIDX(ctx context.Context) string {
	id, err := ghrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GameHustlerRelations.
func (ghrq *GameHustlerRelationQuery) All(ctx context.Context) ([]*GameHustlerRelation, error) {
	ctx = setContextOp(ctx, ghrq.ctx, "All")
	if err := ghrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GameHustlerRelation, *GameHustlerRelationQuery]()
	return withInterceptors[[]*GameHustlerRelation](ctx, ghrq, qr, ghrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) AllX(ctx context.Context) []*GameHustlerRelation {
	nodes, err := ghrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GameHustlerRelation IDs.
func (ghrq *GameHustlerRelationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ghrq.ctx.Unique == nil && ghrq.path != nil {
		ghrq.Unique(true)
	}
	ctx = setContextOp(ctx, ghrq.ctx, "IDs")
	if err = ghrq.Select(gamehustlerrelation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) IDsX(ctx context.Context) []string {
	ids, err := ghrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ghrq *GameHustlerRelationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ghrq.ctx, "Count")
	if err := ghrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ghrq, querierCount[*GameHustlerRelationQuery](), ghrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) CountX(ctx context.Context) int {
	count, err := ghrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ghrq *GameHustlerRelationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ghrq.ctx, "Exist")
	switch _, err := ghrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ghrq *GameHustlerRelationQuery) ExistX(ctx context.Context) bool {
	exist, err := ghrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GameHustlerRelationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ghrq *GameHustlerRelationQuery) Clone() *GameHustlerRelationQuery {
	if ghrq == nil {
		return nil
	}
	return &GameHustlerRelationQuery{
		config:      ghrq.config,
		ctx:         ghrq.ctx.Clone(),
		order:       append([]OrderFunc{}, ghrq.order...),
		inters:      append([]Interceptor{}, ghrq.inters...),
		predicates:  append([]predicate.GameHustlerRelation{}, ghrq.predicates...),
		withHustler: ghrq.withHustler.Clone(),
		// clone intermediate query.
		sql:  ghrq.sql.Clone(),
		path: ghrq.path,
	}
}

// WithHustler tells the query-builder to eager-load the nodes that are connected to
// the "hustler" edge. The optional arguments are used to configure the query builder of the edge.
func (ghrq *GameHustlerRelationQuery) WithHustler(opts ...func(*GameHustlerQuery)) *GameHustlerRelationQuery {
	query := (&GameHustlerClient{config: ghrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ghrq.withHustler = query
	return ghrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Citizen string `json:"citizen,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GameHustlerRelation.Query().
//		GroupBy(gamehustlerrelation.FieldCitizen).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ghrq *GameHustlerRelationQuery) GroupBy(field string, fields ...string) *GameHustlerRelationGroupBy {
	ghrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GameHustlerRelationGroupBy{build: ghrq}
	grbuild.flds = &ghrq.ctx.Fields
	grbuild.label = gamehustlerrelation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Citizen string `json:"citizen,omitempty"`
//	}
//
//	client.GameHustlerRelation.Query().
//		Select(gamehustlerrelation.FieldCitizen).
//		Scan(ctx, &v)
func (ghrq *GameHustlerRelationQuery) Select(fields ...string) *GameHustlerRelationSelect {
	ghrq.ctx.Fields = append(ghrq.ctx.Fields, fields...)
	sbuild := &GameHustlerRelationSelect{GameHustlerRelationQuery: ghrq}
	sbuild.label = gamehustlerrelation.Label
	sbuild.flds, sbuild.scan = &ghrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GameHustlerRelationSelect configured with the given aggregations.
func (ghrq *GameHustlerRelationQuery) Aggregate(fns ...AggregateFunc) *GameHustlerRelationSelect {
	return ghrq.Select().Aggregate(fns...)
}

func (ghrq *GameHustlerRelationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ghrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ghrq); err != nil {
				return err
			}
		}
	}
	for _, f := range ghrq.ctx.Fields {
		if !gamehustlerrelation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ghrq.path != nil {
		prev, err := ghrq.path(ctx)
		if err != nil {
			return err
		}
		ghrq.sql = prev
	}
	return nil
}

func (ghrq *GameHustlerRelationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GameHustlerRelation, error) {
	var (
		nodes       = []*GameHustlerRelation{}
		withFKs     = ghrq.withFKs
		_spec       = ghrq.querySpec()
		loadedTypes = [1]bool{
			ghrq.withHustler != nil,
		}
	)
	if ghrq.withHustler != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, gamehustlerrelation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GameHustlerRelation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GameHustlerRelation{config: ghrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ghrq.modifiers) > 0 {
		_spec.Modifiers = ghrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ghrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ghrq.withHustler; query != nil {
		if err := ghrq.loadHustler(ctx, query, nodes, nil,
			func(n *GameHustlerRelation, e *GameHustler) { n.Edges.Hustler = e }); err != nil {
			return nil, err
		}
	}
	for i := range ghrq.loadTotal {
		if err := ghrq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ghrq *GameHustlerRelationQuery) loadHustler(ctx context.Context, query *GameHustlerQuery, nodes []*GameHustlerRelation, init func(*GameHustlerRelation), assign func(*GameHustlerRelation, *GameHustler)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*GameHustlerRelation)
	for i := range nodes {
		if nodes[i].game_hustler_relations == nil {
			continue
		}
		fk := *nodes[i].game_hustler_relations
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(gamehustler.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "game_hustler_relations" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ghrq *GameHustlerRelationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ghrq.querySpec()
	if len(ghrq.modifiers) > 0 {
		_spec.Modifiers = ghrq.modifiers
	}
	_spec.Node.Columns = ghrq.ctx.Fields
	if len(ghrq.ctx.Fields) > 0 {
		_spec.Unique = ghrq.ctx.Unique != nil && *ghrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ghrq.driver, _spec)
}

func (ghrq *GameHustlerRelationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gamehustlerrelation.Table, gamehustlerrelation.Columns, sqlgraph.NewFieldSpec(gamehustlerrelation.FieldID, field.TypeString))
	_spec.From = ghrq.sql
	if unique := ghrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ghrq.path != nil {
		_spec.Unique = true
	}
	if fields := ghrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gamehustlerrelation.FieldID)
		for i := range fields {
			if fields[i] != gamehustlerrelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ghrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ghrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ghrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ghrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ghrq *GameHustlerRelationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ghrq.driver.Dialect())
	t1 := builder.Table(gamehustlerrelation.Table)
	columns := ghrq.ctx.Fields
	if len(columns) == 0 {
		columns = gamehustlerrelation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ghrq.sql != nil {
		selector = ghrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ghrq.ctx.Unique != nil && *ghrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ghrq.predicates {
		p(selector)
	}
	for _, p := range ghrq.order {
		p(selector)
	}
	if offset := ghrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ghrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GameHustlerRelationGroupBy is the group-by builder for GameHustlerRelation entities.
type GameHustlerRelationGroupBy struct {
	selector
	build *GameHustlerRelationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ghrgb *GameHustlerRelationGroupBy) Aggregate(fns ...AggregateFunc) *GameHustlerRelationGroupBy {
	ghrgb.fns = append(ghrgb.fns, fns...)
	return ghrgb
}

// Scan applies the selector query and scans the result into the given value.
func (ghrgb *GameHustlerRelationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ghrgb.build.ctx, "GroupBy")
	if err := ghrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameHustlerRelationQuery, *GameHustlerRelationGroupBy](ctx, ghrgb.build, ghrgb, ghrgb.build.inters, v)
}

func (ghrgb *GameHustlerRelationGroupBy) sqlScan(ctx context.Context, root *GameHustlerRelationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ghrgb.fns))
	for _, fn := range ghrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ghrgb.flds)+len(ghrgb.fns))
		for _, f := range *ghrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ghrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ghrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GameHustlerRelationSelect is the builder for selecting fields of GameHustlerRelation entities.
type GameHustlerRelationSelect struct {
	*GameHustlerRelationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ghrs *GameHustlerRelationSelect) Aggregate(fns ...AggregateFunc) *GameHustlerRelationSelect {
	ghrs.fns = append(ghrs.fns, fns...)
	return ghrs
}

// Scan applies the selector query and scans the result into the given value.
func (ghrs *GameHustlerRelationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ghrs.ctx, "Select")
	if err := ghrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameHustlerRelationQuery, *GameHustlerRelationSelect](ctx, ghrs.GameHustlerRelationQuery, ghrs, ghrs.inters, v)
}

func (ghrs *GameHustlerRelationSelect) sqlScan(ctx context.Context, root *GameHustlerRelationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ghrs.fns))
	for _, fn := range ghrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ghrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ghrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
