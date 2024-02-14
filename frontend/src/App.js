import { Box, Button } from "@mui/material";
import Navbar from "./components/Navbar";
import { useEffect, useState } from "react";
import { getUserByEmail, getUsers } from "./services/backend";

function App() {
  const [user, setUser] = useState(null);
  const [userTeams, setUserTeams] = useState(null);
  const handleSignInA = () => {
    setUser({
      email: "usera@emailchaser.com",
    });
  };
  const handleSignInB = async () => {
    setUser({
      email: "userb@emailchaser.com",
    });
  };

  useEffect(() => {
    if (user && user.email) {
      getUserByEmail(user, "usera@emailchaser.com");
      // getUsers(user);
    }
  }, [user]);

  return (
    <>
      <Navbar user={user} />
      <Box m={2}>
        <Button onClick={handleSignInA} variant="contained">
          Login As User A
        </Button>
        <Button
          sx={{ ml: 1 }}
          onClick={() => handleSignInB()}
          variant="contained"
        >
          Login As User B
        </Button>
        {user && (
          <Box>
            <h5>Signed in as:</h5>
            <code>{JSON.stringify(user)}</code>
          </Box>
        )}
        {userTeams && (
          <Box>
            <h5>UserTeams:</h5>
            <code>{JSON.stringify(user)}</code>
          </Box>
        )}
      </Box>
    </>
  );
}

export default App;
