<script lang="ts">
  import { IS_ADMIN } from "$lib/app/index.js";
  import type { MusmgrEventTimetable } from "$lib/server/events.js";
  import BottomControlsContainer from "$lib/ui/components/BottomControlsContainer.svelte";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";
  import NoContent from "$lib/ui/components/NoContent.svelte";
  import Events from "$lib/ui/pages/Events.svelte";

  let { data } = $props();
  let events: MusmgrEventTimetable = $derived(data.events);
  const descriptionCap = 40;
</script>

<div class="events container narrow">
  {#if !data.hasEvents}
    <NoContent />
  {:else}
    <Events {events} {descriptionCap} />
  {/if}
  {#if IS_ADMIN}
    <BottomControlsContainer>
      <ButtonAnchor href="/events/new">Create event</ButtonAnchor>
    </BottomControlsContainer>
  {/if}
</div>

<style>
  .events {
    width: 100%;
    height: 100%;
  }
</style>
