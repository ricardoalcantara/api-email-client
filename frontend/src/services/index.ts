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
  ApiKeyDto,
  CreateApiKeyDto,
  TemplateGeneratorDto,
  RequestTemplateGeneratorDto,
  SendEmail,
  DashboardDto,
  UpdatePasswordDto,
} from "./dto";

export const api = createApi({
  baseQuery: fetchBaseQuery({
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
      { slug: string; template: CreateTemplateDto }
    >({
      query: ({ slug, template }) => ({
        url: `/api/template/${slug}`,
        method: "PUT",
        body: template,
      }),
    }),
    patchTemplate: builder.mutation<
      TemplateDto,
      { slug: string; template: UpdateTemplateDto }
    >({
      query: ({ slug, template }) => ({
        url: `/api/template/${slug}`,
        method: "PATCH",
        body: template,
      }),
    }),
    deleteTemplate: builder.mutation<void, string>({
      query: (slug) => ({
        url: `/api/template/${slug}`,
        method: "DELETE",
      }),
    }),
    generateTemplate: builder.mutation<
      TemplateGeneratorDto,
      RequestTemplateGeneratorDto
    >({
      query: (body) => ({
        url: `/api/template/generator`,
        method: "POST",
        body,
      }),
    }),
    cloneTemplate: builder.mutation<TemplateDto, string>({
      query: (slug) => ({
        url: `/api/template/${slug}/clone`,
        method: "POST",
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
    putSmtp: builder.mutation<SmtpDto, { slug: string; smtp: CreateSmtpDto }>({
      query: ({ slug, smtp }) => ({
        url: `/api/smtp/${slug}`,
        method: "PUT",
        body: smtp,
      }),
    }),
    patchSmtp: builder.mutation<SmtpDto, { slug: string; smtp: UpdateSmtpDto }>(
      {
        query: ({ slug, smtp }) => ({
          url: `/api/smtp/${slug}`,
          method: "PATCH",
          body: smtp,
        }),
      }
    ),
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
    sendEmail: builder.mutation<void, SendEmail>({
      query: (email) => ({
        url: "/api/email",
        method: "POST",
        body: email,
      }),
    }),
    resendEmail: builder.mutation<void, number>({
      query: (id) => ({
        url: `/api/email/${id}/send`,
        method: "PATCH",
      }),
    }),

    // ApiKeys
    listApiKey: builder.query<ListView<ApiKeyDto>, void>({
      query: () => ({
        url: "/api/api-key",
        method: "GET",
      }),
    }),
    postApiKey: builder.mutation<ApiKeyDto, CreateApiKeyDto>({
      query: (apiKey) => ({
        url: "/api/api-key",
        method: "POST",
        body: apiKey,
      }),
    }),
    deleteApiKey: builder.mutation<void, number>({
      query: (id) => ({
        url: `/api/api-key/${id}`,
        method: "DELETE",
      }),
    }),

    // Dashboard
    getDashboard: builder.query<DashboardDto, void>({
      query: () => ({
        url: "/api/dashboard",
        method: "GET",
      }),
    }),

    // User
    patchPassword: builder.mutation<void, UpdatePasswordDto>({
      query: (password) => ({
        url: "/api/user/password",
        method: "PATCH",
        body: password,
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
  usePatchTemplateMutation,
  useDeleteTemplateMutation,
  useGenerateTemplateMutation,
  useCloneTemplateMutation,
  // SMTP hooks
  useListSmtpQuery,
  usePostSmtpMutation,
  useGetSmtpQuery,
  usePutSmtpMutation,
  usePatchSmtpMutation,
  useDeleteSmtpMutation,
  // Email hooks
  useListEmailQuery,
  useSendEmailMutation,
  useResendEmailMutation,
  // ApiKey hooks
  useListApiKeyQuery,
  usePostApiKeyMutation,
  useDeleteApiKeyMutation,
  // Dashboard hooks
  useGetDashboardQuery,
  // User hooks
  usePatchPasswordMutation,
} = api;
