import { Box, Button, Card } from "@mui/material";
import Navbar from "./components/Navbar";
import { useEffect, useState } from "react";
import { getUserByEmail, getUsers } from "./services/backend";

function App() {
  const [loggedUser, setLoggedUser] = useState(null);
  const [userTeams, setUserTeams] = useState(null);
  const [currentUserData, setCurrentUserData] = useState(null);

  const handleSignInA = async () => {
    setLoggedUser("usera@emailchaser.com");
  };

  const handleSignInB = async () => {
    setLoggedUser("userb@emailchaser.com");
  };

  const fetchTeamData = async () => {
    const data = await getUserByEmail(loggedUser);
    const currentUserData = data.users[0];

    if (!currentUserData) {
      alert("this user has not signed up yet");
    }

    setCurrentUserData(currentUserData);
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
            {currentUserData && (
              <Box>
                <h5>Teams:</h5>
                {currentUserData?.edges?.teams.map((t) => (
                  <Card key={t} sx={{ maxWidth: 120, p: 2 }}>
                    {t.name}
                  </Card>
                ))}
              </Box>
            )}
          </Box>
        )}
      </Box>
    </>
  );
}

export default App;
