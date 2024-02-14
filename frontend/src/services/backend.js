import { setupApiInstance } from "./api";

const BackendClient = setupApiInstance({
  baseURL: `http://localhost:8080`, // this is a proxy to the narrowin server
  headers: {
    "Content-Type": "application/json",
  },
});

export const makeUserRequest = async (
  user = null,
  method,
  endpoint,
  bodyParams = {},
  additionalParams = {}
) => {
  try {
    const response = await BackendClient.request({
      method: method,
      url: endpoint,
      data: bodyParams,
      headers: {
        "user-email": user?.email,
      },
      ...additionalParams,
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
  return await makeUserRequest(user, "GET", endpoint);
};

export const getUserByEmail = async (user, email) => {
  const endpoint = `api/users?email=${email}`;
  return await makeUserRequest(user, "GET", endpoint);
};
