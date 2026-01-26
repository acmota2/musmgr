export function isAdmin(): boolean {
  return import.meta.env.VITE_FRONTEND_BUILD_MODE === "admin";
}
