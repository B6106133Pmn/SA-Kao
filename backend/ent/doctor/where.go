// Code generated by entc, DO NOT EDIT.

package doctor

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/pmn-kao/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Password applies equality check predicate on the "Password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Tel applies equality check predicate on the "Tel" field. It's identical to TelEQ.
func Tel(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel), v))
	})
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "Password" field.
func PasswordEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "Password" field.
func PasswordNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "Password" field.
func PasswordIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "Password" field.
func PasswordNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "Password" field.
func PasswordGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "Password" field.
func PasswordGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "Password" field.
func PasswordLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "Password" field.
func PasswordLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "Password" field.
func PasswordContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "Password" field.
func PasswordHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "Password" field.
func PasswordHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "Password" field.
func PasswordEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "Password" field.
func PasswordContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TelEQ applies the EQ predicate on the "Tel" field.
func TelEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTel), v))
	})
}

// TelNEQ applies the NEQ predicate on the "Tel" field.
func TelNEQ(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTel), v))
	})
}

// TelIn applies the In predicate on the "Tel" field.
func TelIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTel), v...))
	})
}

// TelNotIn applies the NotIn predicate on the "Tel" field.
func TelNotIn(vs ...string) predicate.Doctor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTel), v...))
	})
}

// TelGT applies the GT predicate on the "Tel" field.
func TelGT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTel), v))
	})
}

// TelGTE applies the GTE predicate on the "Tel" field.
func TelGTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTel), v))
	})
}

// TelLT applies the LT predicate on the "Tel" field.
func TelLT(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTel), v))
	})
}

// TelLTE applies the LTE predicate on the "Tel" field.
func TelLTE(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTel), v))
	})
}

// TelContains applies the Contains predicate on the "Tel" field.
func TelContains(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTel), v))
	})
}

// TelHasPrefix applies the HasPrefix predicate on the "Tel" field.
func TelHasPrefix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTel), v))
	})
}

// TelHasSuffix applies the HasSuffix predicate on the "Tel" field.
func TelHasSuffix(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTel), v))
	})
}

// TelEqualFold applies the EqualFold predicate on the "Tel" field.
func TelEqualFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTel), v))
	})
}

// TelContainsFold applies the ContainsFold predicate on the "Tel" field.
func TelContainsFold(v string) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTel), v))
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDegree applies the HasEdge predicate on the "degree" edge.
func HasDegree() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DegreeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DegreeTable, DegreeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDegreeWith applies the HasEdge predicate on the "degree" edge with a given conditions (other predicates).
func HasDegreeWith(preds ...predicate.Degree) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DegreeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DegreeTable, DegreeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNametitle applies the HasEdge predicate on the "nametitle" edge.
func HasNametitle() predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NametitleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NametitleTable, NametitleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNametitleWith applies the HasEdge predicate on the "nametitle" edge with a given conditions (other predicates).
func HasNametitleWith(preds ...predicate.Nametitle) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NametitleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NametitleTable, NametitleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
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
func Not(p predicate.Doctor) predicate.Doctor {
	return predicate.Doctor(func(s *sql.Selector) {
		p(s.Not())
	})
}
