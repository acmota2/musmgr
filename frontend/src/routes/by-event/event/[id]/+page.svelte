<script lang="ts">
  import Songs from "$lib/pages/works.svelte";
  import type { Posts, Song } from "$lib/util";

  let { data }: Posts<Song[]> = $props();
</script>

<Songs {data} />
