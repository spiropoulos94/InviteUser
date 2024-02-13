// Code generated by ent, DO NOT EDIT.

package invitation

import (
	"spiropoulos94/emailchaser/invite/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Invitation {
	return predicate.Invitation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Invitation {
	return predicate.Invitation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Invitation {
	return predicate.Invitation(sql.FieldLTE(FieldID, id))
}

// InviteeEmail applies equality check predicate on the "invitee_email" field. It's identical to InviteeEmailEQ.
func InviteeEmail(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldInviteeEmail, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldStatus, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldCreatedAt, v))
}

// InviteeEmailEQ applies the EQ predicate on the "invitee_email" field.
func InviteeEmailEQ(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldInviteeEmail, v))
}

// InviteeEmailNEQ applies the NEQ predicate on the "invitee_email" field.
func InviteeEmailNEQ(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldNEQ(FieldInviteeEmail, v))
}

// InviteeEmailIn applies the In predicate on the "invitee_email" field.
func InviteeEmailIn(vs ...string) predicate.Invitation {
	return predicate.Invitation(sql.FieldIn(FieldInviteeEmail, vs...))
}

// InviteeEmailNotIn applies the NotIn predicate on the "invitee_email" field.
func InviteeEmailNotIn(vs ...string) predicate.Invitation {
	return predicate.Invitation(sql.FieldNotIn(FieldInviteeEmail, vs...))
}

// InviteeEmailGT applies the GT predicate on the "invitee_email" field.
func InviteeEmailGT(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldGT(FieldInviteeEmail, v))
}

// InviteeEmailGTE applies the GTE predicate on the "invitee_email" field.
func InviteeEmailGTE(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldGTE(FieldInviteeEmail, v))
}

// InviteeEmailLT applies the LT predicate on the "invitee_email" field.
func InviteeEmailLT(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldLT(FieldInviteeEmail, v))
}

// InviteeEmailLTE applies the LTE predicate on the "invitee_email" field.
func InviteeEmailLTE(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldLTE(FieldInviteeEmail, v))
}

// InviteeEmailContains applies the Contains predicate on the "invitee_email" field.
func InviteeEmailContains(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldContains(FieldInviteeEmail, v))
}

// InviteeEmailHasPrefix applies the HasPrefix predicate on the "invitee_email" field.
func InviteeEmailHasPrefix(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldHasPrefix(FieldInviteeEmail, v))
}

// InviteeEmailHasSuffix applies the HasSuffix predicate on the "invitee_email" field.
func InviteeEmailHasSuffix(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldHasSuffix(FieldInviteeEmail, v))
}

// InviteeEmailEqualFold applies the EqualFold predicate on the "invitee_email" field.
func InviteeEmailEqualFold(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEqualFold(FieldInviteeEmail, v))
}

// InviteeEmailContainsFold applies the ContainsFold predicate on the "invitee_email" field.
func InviteeEmailContainsFold(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldContainsFold(FieldInviteeEmail, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Invitation {
	return predicate.Invitation(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Invitation {
	return predicate.Invitation(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Invitation {
	return predicate.Invitation(sql.FieldContainsFold(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Invitation {
	return predicate.Invitation(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Invitation {
	return predicate.Invitation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Invitation {
	return predicate.Invitation(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Invitation) predicate.Invitation {
	return predicate.Invitation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Invitation) predicate.Invitation {
	return predicate.Invitation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Invitation) predicate.Invitation {
	return predicate.Invitation(sql.NotPredicates(p))
}
