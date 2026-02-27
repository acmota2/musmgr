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

<div class="view">
  <div class="events container narrow">
    {#if !data.hasEvents}
      <NoContent />
    {:else}
      <Events {events} {descriptionCap} />
    {/if}
  </div>
  {#if IS_ADMIN}
    <BottomControlsContainer class="create-event-bar">
      <ButtonAnchor href="/events/new">Create event</ButtonAnchor>
    </BottomControlsContainer>
  {/if}
</div>

<style>
  .view {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .events {
    width: 100%;
    height: 100%;
  }

  .events.container.narrow {
    display: flex;
    flex-direction: column;
  }
</style>
