import { RouterProvider } from "react-router-dom";
import Routers from "@/config/routers";

export default function App() {
  return <RouterProvider router={Routers} />;
}
