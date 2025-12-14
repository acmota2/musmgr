<script lang="ts">
  import Tile from "$lib/components/tile.svelte";
  import type { Event } from "$lib/util";
  import type { Posts } from "$lib/util";

  let { data }: { data: Posts<Event[]> } = $props();
</script>

<h1 class="pb-4 text-3xl font-bold">Choose which event type you wish to search by</h1>

{#each data.posts as event}
  <Tile href={`/category/${event.id}`} text={event.type}></Tile>
{/each}
