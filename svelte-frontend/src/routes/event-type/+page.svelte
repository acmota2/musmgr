<script lang="ts">
  import Tile from "$lib/components/tile.svelte";
  import type { EventType } from "$lib/util/event";
  import type { Posts } from "$lib/util/posts";

  let { data }: { data: Posts<EventType[]> } = $props();
</script>

<h1 class="pb-4 text-3xl font-bold">Choose which event type you wish to search by</h1>

{#each data.posts as event}
  <Tile href={`/category/${event.id}`} text={event.type}></Tile>
{/each}
