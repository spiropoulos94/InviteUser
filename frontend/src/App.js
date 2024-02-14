import { Box, Button } from "@mui/material";
import Navbar from "./components/Navbar";
import { useEffect, useState } from "react";
import { getUserByEmail, getUsers } from "./services/backend";

function App() {
  const [loggedUser, setLoggedUser] = useState(null);
  const [userTeams, setUserTeams] = useState(null);

  const handleSignInA = async () => {
    setLoggedUser("usera@emailchaser.com");
  };

  const handleSignInB = async () => {
    setLoggedUser("userb@emailchaser.com");
  };

  const fetchTeamData = async () => {
    const data = await getUserByEmail(loggedUser);
    const currentUser = data.users[0];

    if (!currentUser) {
      alert("this user has not signed up yet");
    }

    console.log({ currentUser });
  };

  return (
    <>
      <Navbar user={loggedUser} />
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
        {loggedUser && (
          <Box>
            <h5>Signed in as:</h5>
            <code>{JSON.stringify(loggedUser)}</code>
          </Box>
        )}
        {loggedUser && (
          <Box sx={{ mt: 2 }}>
            <Button onClick={fetchTeamData} variant="outlined">
              Get User Team
            </Button>
          </Box>
        )}
      </Box>
    </>
  );
}

export default App;
