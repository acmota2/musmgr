<script lang="ts">
  import type { MusmgrEvent } from "$lib/server/events.js";

  let { data } = $props();
  let events: MusmgrEvent[] = $derived(data.events);
</script>

<div class="events container narrow">
  {#if events.length === 0}
    <h1 class="no-events">No events to show</h1>
  {:else}
    {#each events as event}
      <p>{event.eventType}</p>
      <p>{event.happenedAt}</p>
      <p>{event.description}</p>
    {/each}
  {/if}
</div>

<style>
  .events {
    display: flex;
    justify-content: center;
    height: 100%;
    width: 100%;
  }

  .no-events {
    justify-self: center;
    align-self: center;
  }
</style>
