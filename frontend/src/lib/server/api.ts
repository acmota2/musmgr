import { env } from "$env/dynamic/private";

export function getApiUrl(): string {
  const url = env.API_URL;
  if (!url) {
    throw new Error("API_URL is not set");
  }
  return url;
}
