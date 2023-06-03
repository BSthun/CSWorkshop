import { Box, Button, Typography } from "@mui/material";

import { FaGoogle } from "react-icons/fa";
import logo from "@/assets/logo.png";
import { firebaseAuth } from "../utils/firebase";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { GoogleAuthProvider, signInWithPopup } from "firebase/auth";

const Login = () => {
  const navigate = useNavigate();

  const processCredential = (result: any) => {
    // No need to do anything with credential. Just want you to see what credential looks like.
    const credential = GoogleAuthProvider.credentialFromResult(result);
    console.log(credential);

    // Get IdToken for verification with backend.
    result.user.getIdToken().then((token: any) => {
      axios
        .post("/api/account/callback", {
          idToken: token,
        })
        .then((response) => {
          if (response.data.success) {
            // Set token from backend to cookie
            document.cookie = "user=" + response.data.data.token;
            navigate("/home");
          } else {
            alert(response.data.message);
          }
        })
        .catch((err) => {
          if (err.response.status === 500) {
            alert(err.response.data);
          } else {
            alert(err.message);
          }
        });
    });
  };

  const login = () => {
    const provider = new GoogleAuthProvider();
    signInWithPopup(firebaseAuth, provider)
      .then((result) => {
        processCredential(result);
      })
      .catch((error) => {
        alert(error.message);
      });
  };

  return (
    <Box
      sx={{
        display: "flex",
        width: "100vw",
        height: "100vh",
        justifyContent: "center",
        alignItems: "center",
      }}
    >
      <Box
        sx={{
          width: "100%",
          position: "fixed",
          top: 0,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          ml: 2,
          mt: 2,
        }}
      >
        <Box
          sx={{
            display: "flex",
            backgroundColor: "#024F97",
            alignItems: "center",
            borderRadius: 8,
            mr: 2,
          }}
        >
          <img src={logo} width="40px" height="40px" />
        </Box>
        <Typography fontSize="24px" fontWeight="400" align="center" pr={2}>
          CS Hackathon 2023
        </Typography>
      </Box>
      <Box>
        <Button variant="contained" startIcon={<FaGoogle />} onClick={login}>
          Sign in with Google
        </Button>
      </Box>
    </Box>
  );
};

export default Login;
