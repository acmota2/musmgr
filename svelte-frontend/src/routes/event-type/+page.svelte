<script lang="ts">
  import Tile from "$lib/components/tile.svelte";
  import type { EventType } from "$lib/util/event";
  import type { Posts } from "$lib/util/posts";

  let { data }: { data: Posts<EventType[]> } = $props();
</script>

{#each data.posts as event}
  <Tile href={`/category/${event.id}`} text={event.type}></Tile>
{/each}
