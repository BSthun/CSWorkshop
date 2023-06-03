import React from "react";
import { Outlet, createBrowserRouter } from "react-router-dom";
import WSRealtime from "./components/ws-realtime";
import MainView from "./views/Main";
import Login from "./views/Login";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <RouterWrapper />,
    children: [
      {
        path: "home",
        element: <MainView />,
      },
      {
        path: "",
        element: <Login />,
      },
    ],
  },
]);

export default function RouterWrapper() {
  return (
    <React.Fragment>
      <Outlet />
      <WSRealtime />
    </React.Fragment>
  );
}
