import { type Actions, fail, redirect } from "@sveltejs/kit";
import { checkStringFormField } from "$lib/utils";
import type { PageServerLoad } from "./$types";
import { createEvent, getEventTypes, type CreateEventPayload } from "$lib/server/events";

export const load: PageServerLoad = async ({ fetch }) => {
  return { eventTypes: (await getEventTypes(fetch)).sort() };
};

export const actions: Actions = {
  create: async ({ fetch, request }) => {
    const rawData = await request.formData();

    let data: CreateEventPayload;
    try {
      data = {
        happened_at: checkStringFormField(rawData.get("happened_at"), "happened_at"),
        description: checkStringFormField(rawData.get("description"), "description"),
        event_type: checkStringFormField(rawData.get("event_type"), "event_type"),
        name: checkStringFormField(rawData.get("name"), "name"),
      };
    } catch (e) {
      return fail(400, e);
    }

    const newLocation = await createEvent(fetch, data);

    throw redirect(303, newLocation);
  },
};
