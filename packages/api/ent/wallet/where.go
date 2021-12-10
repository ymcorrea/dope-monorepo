// Code generated by entc, DO NOT EDIT.

package wallet

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dopedao/dope-monorepo/packages/api/ent/predicate"
	"github.com/dopedao/dope-monorepo/packages/api/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Paper applies equality check predicate on the "paper" field. It's identical to PaperEQ.
func Paper(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaper), v))
	})
}

// PaperEQ applies the EQ predicate on the "paper" field.
func PaperEQ(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaper), v))
	})
}

// PaperNEQ applies the NEQ predicate on the "paper" field.
func PaperNEQ(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaper), v))
	})
}

// PaperIn applies the In predicate on the "paper" field.
func PaperIn(vs ...schema.BigInt) predicate.Wallet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wallet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaper), v...))
	})
}

// PaperNotIn applies the NotIn predicate on the "paper" field.
func PaperNotIn(vs ...schema.BigInt) predicate.Wallet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wallet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaper), v...))
	})
}

// PaperGT applies the GT predicate on the "paper" field.
func PaperGT(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaper), v))
	})
}

// PaperGTE applies the GTE predicate on the "paper" field.
func PaperGTE(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaper), v))
	})
}

// PaperLT applies the LT predicate on the "paper" field.
func PaperLT(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaper), v))
	})
}

// PaperLTE applies the LTE predicate on the "paper" field.
func PaperLTE(v schema.BigInt) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaper), v))
	})
}

// HasDopes applies the HasEdge predicate on the "dopes" edge.
func HasDopes() predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DopesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DopesTable, DopesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDopesWith applies the HasEdge predicate on the "dopes" edge with a given conditions (other predicates).
func HasDopesWith(preds ...predicate.Dope) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DopesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DopesTable, DopesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.WalletItems) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHustlers applies the HasEdge predicate on the "hustlers" edge.
func HasHustlers() predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HustlersTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HustlersTable, HustlersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHustlersWith applies the HasEdge predicate on the "hustlers" edge with a given conditions (other predicates).
func HasHustlersWith(preds ...predicate.Hustler) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HustlersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HustlersTable, HustlersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(func(s *sql.Selector) {
		p(s.Not())
	})
}
