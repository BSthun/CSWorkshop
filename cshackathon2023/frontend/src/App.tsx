import { Toaster } from "react-hot-toast";
import { RouterProvider } from "react-router-dom";
import { SWRConfig } from "swr";
import { router } from "./router";
import { fetcher } from "./utils/api";

function App() {
  return (
    <SWRConfig value={{ fetcher }}>
      <RouterProvider router={router} />
      <Toaster />
    </SWRConfig>
  );
}

export default App;
