export interface Hermes {
  TextDirection: "ltr" | "rtl" | undefined;
  Product: Product;
  DisableCSSInlining: boolean;
}

export interface Product {
  Name: string;
  Link?: string;
  Logo?: string;
  Copyright: string;
  TroubleText: string;
}

export interface HermesBody {
  Name: string;
  Intros: string[];
  Dictionary: Entry[];
  // Table: unknown;
  Actions: Action[];
  Outros: string[];
  Greeting: string;
  Signature: string;
  Title: string;
  FreeMarkdown?: string;
}

export interface Entry {
  Key: string;
  Value: string;
}

export interface Action {
  Instructions: string;
  Button: Button;
  InviteCode: string;
}

export interface Button {
  Color: string;
  TextColor: string;
  Text: string;
  Link?: string;
}
