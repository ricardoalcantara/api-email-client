import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
  CreateTemplateDto,
  ListView,
  LoginDto,
  TemplateDto,
  TokenDto,
  UpdateTemplateDto,
  SmtpDto,
  CreateSmtpDto,
  UpdateSmtpDto,
  EmailView,
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

    // Templates endpoints
    listTemplate: builder.query<ListView<TemplateDto>, void>({
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
    getTemplate: builder.query<TemplateDto, string>({
      query: (slug) => ({
        url: `/api/template/${slug}`,
        method: "GET",
      }),
    }),
    putTemplate: builder.mutation<
      TemplateDto,
      { slug: string; template: UpdateTemplateDto }
    >({
      query: ({ slug, template }) => ({
        url: `/api/template/${slug}`,
        method: "PUT",
        body: template,
      }),
    }),

    // SMTP endpoints
    listSmtp: builder.query<ListView<SmtpDto>, void>({
      query: () => ({
        url: "/api/smtp",
        method: "GET",
      }),
    }),
    postSmtp: builder.mutation<SmtpDto, CreateSmtpDto>({
      query: (smtp) => ({
        url: "/api/smtp",
        method: "POST",
        body: smtp,
      }),
    }),
    getSmtp: builder.query<SmtpDto, string>({
      query: (slug) => ({
        url: `/api/smtp/${slug}`,
        method: "GET",
      }),
    }),
    putSmtp: builder.mutation<SmtpDto, { slug: string; smtp: UpdateSmtpDto }>({
      query: ({ slug, smtp }) => ({
        url: `/api/smtp/${slug}`,
        method: "PUT",
        body: smtp,
      }),
    }),
    deleteSmtp: builder.mutation<void, string>({
      query: (slug) => ({
        url: `/api/smtp/${slug}`,
        method: "DELETE",
      }),
    }),

    // Email endpoints
    listEmail: builder.query<ListView<EmailView>, void>({
      query: () => ({
        url: "/api/email",
        method: "GET",
      }),
    }),
  }),
});

export const {
  useLoginMutation,
  // Templates hooks
  useListTemplateQuery,
  usePostTemplateMutation,
  useGetTemplateQuery,
  usePutTemplateMutation,
  // SMTP hooks
  useListSmtpQuery,
  usePostSmtpMutation,
  useGetSmtpQuery,
  usePutSmtpMutation,
  useDeleteSmtpMutation,
  // Email hooks
  useListEmailQuery,
} = api;
