package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invitation holds the schema definition for the Invitation entity.
type Invitation struct {
	ent.Schema
}

// Fields of the Invitation.
func (Invitation) Fields() []ent.Field {
	return []ent.Field{
		field.String("invitee_email"),
		field.String("status"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Invitation.
func (Invitation) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("user", User.Type).
            Ref("invitations").
            Unique(),
    }
}
