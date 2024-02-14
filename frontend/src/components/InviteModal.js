import { IconButton, TextField } from "@mui/material";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { useState } from "react";
import CloseIcon from "@mui/icons-material/Close";

export default function InviteModal({ user }) {
  const [email, setEmail] = useState("");
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = (event, reason) => {
    if (reason !== "backdropClick") {
      setOpen(false);
      setEmail("");
    }
  };

  const handleSubmit = async () => {
    console.log({ email });

    const data = await inviteUser(email);
    console.log({ data }, "apo to invite user");
    if (data.error) {
      alert(data.error);
    } else {
      alert("User invited successfully!");
      handleClose();
    }

    // handleClose();
  };

  const inviteUser = async (userEmail) => {
    try {
      const headers = new Headers();
      headers.append("Content-Type", "application/json");
      headers.append("user-email", user);

      const response = await fetch("http://localhost:8080/api/invites/", {
        method: "POST",
        headers: headers,
        body: JSON.stringify({
          invitee_email: userEmail,
        }),
      }); // Adjusted endpoint
      const data = await response.json();

      if (data.error) {
        console.log({ response, data });
        throw new Error(data.error);
      }

      return data;
    } catch (error) {
      return { error: "Error inviting user:", error };
    }
  };

  return (
    <>
      <Button
        disabled={Boolean(!user)}
        variant="outlined"
        color="inherit"
        onClick={handleClickOpen}
      >
        Invite your Team
      </Button>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
        fullWidth
      >
        <DialogTitle id="alert-dialog-title">
          Invite your team
          <IconButton
            aria-label="close"
            onClick={handleClose}
            sx={{
              position: "absolute",
              right: 8,
              top: 8,
              color: (theme) => theme.palette.grey[500],
            }}
          >
            <CloseIcon />
          </IconButton>
        </DialogTitle>
        <DialogContent>
          <TextField
            type="email"
            onChange={(e) => setEmail(e.target.value)}
            fullWidth
            size="medium"
          />
        </DialogContent>
        <DialogActions>
          <Button variant="contained" onClick={handleSubmit} autoFocus>
            Send
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
}
