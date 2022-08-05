// Code generated by entc, DO NOT EDIT.

package dope

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Claimed applies equality check predicate on the "claimed" field. It's identical to ClaimedEQ.
func Claimed(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClaimed), v))
	})
}

// Opened applies equality check predicate on the "opened" field. It's identical to OpenedEQ.
func Opened(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpened), v))
	})
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// SalePrice applies equality check predicate on the "salePrice" field. It's identical to SalePriceEQ.
func SalePrice(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalePrice), v))
	})
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// ClaimedEQ applies the EQ predicate on the "claimed" field.
func ClaimedEQ(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClaimed), v))
	})
}

// ClaimedNEQ applies the NEQ predicate on the "claimed" field.
func ClaimedNEQ(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClaimed), v))
	})
}

// OpenedEQ applies the EQ predicate on the "opened" field.
func OpenedEQ(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpened), v))
	})
}

// OpenedNEQ applies the NEQ predicate on the "opened" field.
func OpenedNEQ(v bool) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOpened), v))
	})
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScore), v))
	})
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScore), v...))
	})
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScore), v...))
	})
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScore), v))
	})
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScore), v))
	})
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScore), v))
	})
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScore), v))
	})
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScore)))
	})
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScore)))
	})
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// RankIsNil applies the IsNil predicate on the "rank" field.
func RankIsNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRank)))
	})
}

// RankNotNil applies the NotNil predicate on the "rank" field.
func RankNotNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRank)))
	})
}

// SalePriceEQ applies the EQ predicate on the "salePrice" field.
func SalePriceEQ(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSalePrice), v))
	})
}

// SalePriceNEQ applies the NEQ predicate on the "salePrice" field.
func SalePriceNEQ(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSalePrice), v))
	})
}

// SalePriceIn applies the In predicate on the "salePrice" field.
func SalePriceIn(vs ...float64) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSalePrice), v...))
	})
}

// SalePriceNotIn applies the NotIn predicate on the "salePrice" field.
func SalePriceNotIn(vs ...float64) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSalePrice), v...))
	})
}

// SalePriceGT applies the GT predicate on the "salePrice" field.
func SalePriceGT(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSalePrice), v))
	})
}

// SalePriceGTE applies the GTE predicate on the "salePrice" field.
func SalePriceGTE(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSalePrice), v))
	})
}

// SalePriceLT applies the LT predicate on the "salePrice" field.
func SalePriceLT(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSalePrice), v))
	})
}

// SalePriceLTE applies the LTE predicate on the "salePrice" field.
func SalePriceLTE(v float64) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSalePrice), v))
	})
}

// SalePriceIsNil applies the IsNil predicate on the "salePrice" field.
func SalePriceIsNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSalePrice)))
	})
}

// SalePriceNotNil applies the NotNil predicate on the "salePrice" field.
func SalePriceNotNil() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSalePrice)))
	})
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrder), v))
	})
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrder), v))
	})
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrder), v...))
	})
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Dope {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Dope(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrder), v...))
	})
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrder), v))
	})
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrder), v))
	})
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrder), v))
	})
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrder), v))
	})
}

// HasWallet applies the HasEdge predicate on the "wallet" edge.
func HasWallet() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WalletTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WalletTable, WalletColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWalletWith applies the HasEdge predicate on the "wallet" edge with a given conditions (other predicates).
func HasWalletWith(preds ...predicate.Wallet) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WalletInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WalletTable, WalletColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLastSale applies the HasEdge predicate on the "last_sale" edge.
func HasLastSale() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LastSaleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, LastSaleTable, LastSaleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLastSaleWith applies the HasEdge predicate on the "last_sale" edge with a given conditions (other predicates).
func HasLastSaleWith(preds ...predicate.Listing) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LastSaleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, LastSaleTable, LastSaleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasListings applies the HasEdge predicate on the "listings" edge.
func HasListings() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ListingsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ListingsTable, ListingsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasListingsWith applies the HasEdge predicate on the "listings" edge with a given conditions (other predicates).
func HasListingsWith(preds ...predicate.Listing) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ListingsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ListingsTable, ListingsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ItemsTable, ItemsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.Item) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ItemsTable, ItemsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIndex applies the HasEdge predicate on the "index" edge.
func HasIndex() predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IndexTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, IndexTable, IndexColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIndexWith applies the HasEdge predicate on the "index" edge with a given conditions (other predicates).
func HasIndexWith(preds ...predicate.Search) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IndexInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, IndexTable, IndexColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Dope) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Dope) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
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
func Not(p predicate.Dope) predicate.Dope {
	return predicate.Dope(func(s *sql.Selector) {
		p(s.Not())
	})
}
