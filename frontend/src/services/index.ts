import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
  CreateTemplateDto,
  ListView,
  LoginDto,
  TemplateDto,
  TokenDto,
} from "./dto";

export const api = createApi({
  baseQuery: fetchBaseQuery({
    // baseUrl: 'http://localhost:5555',
    prepareHeaders: (headers) => {
      const token = localStorage.getItem("access_token");
      if (token) {
        headers.set("authorization", `Bearer ${token}`);
      }
      return headers;
    },
  }),
  endpoints: (builder) => ({
    login: builder.mutation<TokenDto, LoginDto>({
      query: (login) => ({
        url: "/api/auth/token",
        method: "POST",
        body: login,
      }),
    }),

    listTemplates: builder.query<ListView<TemplateDto>, void>({
      query: () => ({
        url: "/api/template",
        method: "GET",
      }),
    }),
    postTemplate: builder.mutation<TemplateDto, CreateTemplateDto>({
      query: (template) => ({
        url: "/api/template",
        method: "POST",
        body: template,
      }),
    }),
  }),
});

export const {
  useLoginMutation,
  useListTemplatesQuery,
  usePostTemplateMutation,
} = api;
