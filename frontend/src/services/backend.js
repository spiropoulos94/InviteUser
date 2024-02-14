import { setupApiInstance } from "./api";

const BackendClient = setupApiInstance({
  baseURL: `http://localhost:8080`, // this is a proxy to the narrowin server
  headers: {
    "Content-Type": "application/json",
  },
});

export const makeUserGetRequest = async (
  user = null,
  endpoint,
  queryparams = {}
) => {
  try {
    const response = await BackendClient.request({
      url: endpoint,
      params: queryparams,
      headers: {
        "user-email": user?.email,
      },
    });

    console.log(12);
    console.log(response.data);
  } catch (error) {
    alert(error?.response?.data?.error);
    // alert(error); // Propagate the error for handling in the calling code
  }
};

export const getUsers = async (user) => {
  const endpoint = "api/users/";
  return await makeUserGetRequest(user, endpoint);
};

export const getUserByEmail = async (user, email) => {
  const endpoint = `api/users/`;
  return await makeUserGetRequest(user, endpoint, { email });
};
