import { FetchBaseQueryError } from "@reduxjs/toolkit/query";

export function IsFetchBaseQueryError(
  error: any
): error is FetchBaseQueryError {
  return "status" in error;
}

export function IsError(data: any): data is { error: string } {
  return "error" in data;
}

export function getError(error: any): string {
  if (IsFetchBaseQueryError(error)) {
    if (error.data && IsError(error.data)) {
      return error.data.error;
    }
  }

  return "Unknown error";
}
