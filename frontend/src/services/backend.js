import { setupApiInstance } from "./api";
import qs from "qs";

const BackendClient = setupApiInstance({
  baseURL: `http://localhost:8080`, // this is a proxy to the narrowin server
  headers: {
    "Content-Type": "application/json",
  },
});

export const makeUserGetRequest = async (
  loggedInUserMail = null,
  endpoint,
  queryparams = {}
) => {
  try {
    const response = await BackendClient.request({
      url: endpoint,
      params: queryparams,
      paramsSerializer: function (params) {
        return qs.stringify(params, { encode: false });
      },
      headers: {
        "user-email": loggedInUserMail,
      },
    });

    console.log(12);
    console.log(response.data);
  } catch (error) {
    alert(error?.response?.data?.error);
    // alert(error); // Propagate the error for handling in the calling code
  }
};

export const getUsers = async (loggedInUserMail) => {
  const endpoint = "api/users/";
  return await makeUserGetRequest(loggedInUserMail, endpoint);
};

export const getUserByEmail = async (loggedInUserMail, email) => {
  const endpoint = `api/users/`;
  return await makeUserGetRequest(loggedInUserMail, endpoint, { email });
};
