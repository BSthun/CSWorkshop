import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyAyPliGjd793q7Q_kOul-l9wUaT0WWLLgY",
  authDomain: "anione-diffuse-dev.firebaseapp.com",
  projectId: "anione-diffuse-dev",
  storageBucket: "anione-diffuse-dev.appspot.com",
  messagingSenderId: "657426422173",
  appId: "1:657426422173:web:298fdb7002901397507774",
};

export const firebaseApp = initializeApp(firebaseConfig);

export const firebaseAuth = getAuth(firebaseApp);
