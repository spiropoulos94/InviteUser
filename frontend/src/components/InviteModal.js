import { IconButton, TextField } from "@mui/material";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { useState } from "react";
import CloseIcon from "@mui/icons-material/Close";

export default function InviteModal() {
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

  const handleSubmit = () => {
    console.log({ email });
    handleClose();
  };

  return (
    <>
      <Button variant="outlined" color="inherit" onClick={handleClickOpen}>
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
