### POST /api/invites:

This endpoint allows User A to invite User B to join their team.
Request Body: { "invitee_email": "marcelo@emailchaser.com" }
Logic:
Validate that the inviter (User A) is authenticated and has permission to invite team members.
Validate the invitee's email address to ensure it's on the same company domain.
If validation passes, create an invitation record in the database.
Send an email notification to the invitee.
Response: { "message": "Invitation sent successfully" }

### GET /api/invites/{invitation_id}:

This endpoint allows User B to view the invitation details and accept/reject it.
Logic:
Retrieve the invitation details based on the provided invitation ID.
Check if the invitation is still pending.
If valid, present the invitation details to User B.
Response: Invitation details (e.g., inviter's email, invitation creation timestamp).

### POST /api/invites/{invitation_id}/accept:

This endpoint allows User B to accept the invitation.
Logic:
Validate the invitation token or parameters to ensure the invitation is valid.
Update the status of the invitation to "accepted" in the database.
Create entries in the UserTeam table to associate User B with the same team as User A.
Response: { "message": "Invitation accepted successfully" }

### POST /api/invites/{invitation_id}/reject:

This endpoint allows User B to reject the invitation.
Logic:
Validate the invitation token or parameters to ensure the invitation is valid.
Update the status of the invitation to "rejected" in the database.
Response: { "message": "Invitation rejected successfully" }
