import { createBrowserRouter } from "react-router-dom";
import { Home } from "@/pages/home";
import { Login } from "@/pages/login";
import ProtectedRouter from "@/components/router/protected-router";
import MainLayout from "@/layouts/main-layout";

const Routers = createBrowserRouter([
  {
    element: <ProtectedRouter />,
    children: [
      {
        element: <MainLayout />,
        children: [
          {
            path: "/",
            element: <Home />,
          },
          {
            path: "/template",
            element: <h1>template</h1>,
          },
          {
            path: "/template/:id",
            element: <h1>template id</h1>,
          },
          {
            path: "/smtp",
            element: <h1>smtp</h1>,
          },
          {
            path: "/smtp/:id",
            element: <h1>smtp id</h1>,
          },
          {
            path: "/email",
            element: <h1>email</h1>,
          },
          {
            path: "/api-key",
            element: <h1>Api Keys</h1>,
          },
        ],
      },
    ],
  },
  {
    path: "/login",
    element: <Login />,
  },
]);

export default Routers;
