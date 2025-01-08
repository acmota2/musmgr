<script lang="ts">
  import Songs from "$lib/pages/songs.svelte";
  import type { Posts, Song } from "$lib/util";

  let { data }: Posts<Song[]> = $props();
</script>

<Songs {data} />
