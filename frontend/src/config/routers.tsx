import { createBrowserRouter } from "react-router-dom";
import ProtectedRouter from "@/components/router/protected-router";
import MainLayout from "@/layouts/main-layout";

import Home from "@/pages/home";
import { Login } from "@/pages/login";

import TemplateList from "@/pages/template";
import TemplateCreate from "@/pages/template/new";
import TemplateEdit from "@/pages/template/slug";

import SmtpList from "@/pages/smtp";
import SmtpCreate from "@/pages/smtp/new";
import SmtpEdit from "@/pages/smtp/slug";

import EmailList from "@/pages/email";
import ApiKeyList from "@/pages/api-key";
import EmailTemplateGenerator from "@/pages/template/slug/generator";
import EmailSend from "@/pages/email/send";
import UserHome from "@/pages/user";

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
            element: <TemplateList />,
          },
          {
            path: "/template/new",
            element: <TemplateCreate />,
          },
          {
            path: "/template/:slug/generator",
            element: <EmailTemplateGenerator />,
          },
          {
            path: "/template/:slug",
            element: <TemplateEdit />,
          },
          {
            path: "/smtp",
            element: <SmtpList />,
          },
          {
            path: "/smtp/new",
            element: <SmtpCreate />,
          },
          {
            path: "/smtp/:slug",
            element: <SmtpEdit />,
          },
          {
            path: "/email",
            element: <EmailList />,
          },
          {
            path: "/email/send",
            element: <EmailSend />,
          },
          {
            path: "/api-key",
            element: <ApiKeyList />,
          },
          {
            path: "/user",
            element: <UserHome />,
          }
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
