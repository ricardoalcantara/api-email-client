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
  description: string;
  subject: string;
  html: string;
  text: string;
}

export type CreateTemplateDto = Omit<TemplateDto, "id">;
