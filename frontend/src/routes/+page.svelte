<script lang="ts">
  import Header from "$lib/components/header1.svelte";
  import Tile from "$lib/components/tile.svelte";
</script>

<Header>Welcome to SongMGR, your song search engine</Header>

<Tile href="/event-type" text="By event"></Tile>
<Tile href="/category-type" text="By category"></Tile>
<Tile href="/songs" text="All songs"></Tile>
