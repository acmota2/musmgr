import { error } from "@sveltejs/kit";

export interface GetComposerResponse {
  id: boolean;
  biography: string;
  full_name: string;
  picture_content_type: string;
  picture: string;
  created_at: Date;
  updated_at: Date;
}

export interface Composer {
  biography: string;
  fullName: string;
  pictureContentType: string;
  pictureId: string;
}

export async function getComposer(fetch: typeof globalThis.fetch): Promise<Composer | null> {
  const res = await fetch("/api/composer");

  if (res.status === 404) {
    return null;
  }

  if (!res.ok) {
    throw error(500);
  }

  const composerData: GetComposerResponse = await res.json();

  return {
    biography: composerData.biography,
    fullName: composerData.full_name,
    pictureContentType: composerData.picture_content_type,
    pictureId: composerData.picture,
  };
}
