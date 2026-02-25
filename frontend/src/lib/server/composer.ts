import { error } from "@sveltejs/kit";
import { generalServerError } from "./utils";

export interface CreateComposerPayload {
  full_name: string;
  biography: string;
}

export interface GetComposerResponse {
  id: boolean;
  biography: string;
  full_name: string;
  picture_content_type: string;
  picture: string;
  created_at: Date;
  updated_at: Date;
}

export interface UpdateComposerPayload {
  biography?: string;
}

export interface Composer {
  biography: string;
  fullName: string;
  pictureContentType: string;
  pictureId: string;
}

export async function createComposer(
  fetch: typeof globalThis.fetch,
  data: CreateComposerPayload,
): Promise<string> {
  const res = await fetch("/api/composer", {
    method: "POST",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    throw error(500, generalServerError);
  }

  const location = res.headers.get("Location");
  if (location === null) {
    throw error(500, {
      code: "INVALID_LOCATION_HEADER",
      message: "It's not you, it's us",
      status: 500,
    });
  }

  return `/api${location}`;
}

export async function getComposer(fetch: typeof globalThis.fetch): Promise<Composer | null> {
  const res = await fetch("/api/composer");

  if (res.status === 404) {
    return null;
  }

  if (!res.ok) {
    throw error(500, generalServerError);
  }

  const composerData: GetComposerResponse = await res.json();

  return {
    biography: composerData.biography,
    fullName: composerData.full_name,
    pictureContentType: composerData.picture_content_type,
    pictureId: composerData.picture,
  };
}

export async function updateComposer(
  fetch: typeof globalThis.fetch,
  data: UpdateComposerPayload,
): Promise<void> {
  const res = await fetch("/api/composer", {
    method: "PATCH",
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    throw error(500, generalServerError);
  }
}

export async function updateComposerPicture(
  fetch: typeof globalThis.fetch,
  blob: File,
): Promise<void> {
  const formData = new FormData();
  formData.append("file", blob);

  const res = await fetch("/api/composer/picture", {
    method: "PUT",
    body: formData,
  });

  if (!res.ok) {
    throw error(500, generalServerError);
  }
}
