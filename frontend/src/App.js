import { Box, Button } from "@mui/material";
import Navbar from "./components/Navbar";
import { useState } from "react";

function App() {
  const [user, setUser] = useState(null);
  const handleSignInA = () => {
    setUser({
      email: "usera@emailchaser.com",
    });
  };
  const handleSignInB = () => {
    setUser({
      email: "userb@emailchaser.com",
    });
  };
  return (
    <>
      <Navbar user={user} />
      <Box m={2}>
        <Button onClick={handleSignInA} variant="contained">
          Login As User A
        </Button>
        <Button sx={{ ml: 1 }} onClick={handleSignInB} variant="contained">
          Login As User B
        </Button>
        {user && (
          <Box>
            <h5>Signed in as:</h5>
            <code>{JSON.stringify(user)}</code>
          </Box>
        )}
      </Box>
    </>
  );
}

export default App;
