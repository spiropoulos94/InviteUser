import { Box, Button } from "@mui/material";
import Navbar from "./components/Navbar";
import { useState } from "react";

function App() {
  const [user, setUser] = useState(null);
  const handleSignIn = () => {
    setUser({
      name: "Nikos Spiropoulos",
      email: "test@emailchaser.com",
    });
  };
  return (
    <>
      <Navbar />
      <Box m={2}>
        <Button onClick={handleSignIn} variant="contained">
          Login As User A
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
