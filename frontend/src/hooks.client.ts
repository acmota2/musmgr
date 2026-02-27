import type { HandleClientError } from "@sveltejs/kit";

export const handleError: HandleClientError = async ({ status, message }) => ({
  message,
  status,
  code: status === 404 ? "NOT_FOUND" : "",
});
