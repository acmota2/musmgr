<script lang="ts">
  import Tile from "$lib/components/tile.svelte";
  import type { CategoryType } from "$lib/util/category";
  import type { Posts } from "$lib/util/posts";

  let { data }: Posts<CategoryType[]> = $props();
</script>

{#each data.posts as categoryType}
  <Tile href={`/category/${categoryType.id}`} text={categoryType.title} />
{/each}
