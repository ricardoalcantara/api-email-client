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
