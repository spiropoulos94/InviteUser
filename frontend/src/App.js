import { Box, Button } from "@mui/material";
import Navbar from "./components/Navbar";
import { useState } from "react";

const getUsers = async () => {
  try {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    // headers.append("user-email", "usera@emailchaser.com");
    // headers.append("user-team", "Sales");

    const response = await fetch("http://localhost:8080/api/users/", {
      headers: headers,
    }); // Adjusted endpoint
    const data = await response.json();

    if (data.error) {
      console.log({ response, data });
      alert(data.error);
      throw new Error(data.error);
    }
    console.log(data); // Do something with the users data, like updating state in React component
  } catch (error) {
    console.error("Error fetching users:", error);
  }
};

function App() {
  const [user, setUser] = useState(null);
  const handleSignIn = () => {
    setUser({
      name: "Nikos Spiropoulos",
      email: "test@emailchaser.com",
    });
    getUsers();
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
