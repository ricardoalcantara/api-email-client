export interface LoginDto {
  email: string;
  password: string;
}

export interface TokenDto {
  access_token: string;
}

export interface ListView<T> {
  page: number;
  list: T[];
}

export interface TemplateDto {
  id: string;
  name: string;
  slug: string;
  json_schema: string;
  description: string;
  subject: string;
  template_html: string;
  template_text: string;
}

export type CreateTemplateDto = Omit<TemplateDto, "id">;
export type UpdateTemplateDto = Partial<TemplateDto>;

export interface SmtpDto {
  id: string;
  slug: string;
  name: string;
  server: string;
  port: number;
  email: string;
  user: string;
  password: string;
  default: boolean;
}

export type CreateSmtpDto = Omit<SmtpDto, "id">;
export type UpdateSmtpDto = Partial<SmtpDto>;

export interface EmailView {
  id: number;
  smtp_name: string;
  from: string;
  to: string;
  subject: string;
  sent_at: string | null;
  html_body?: string;
  text_body?: string;
}