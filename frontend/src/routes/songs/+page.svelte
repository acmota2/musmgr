<script lang="ts">
  import Songs from "$lib/pages/songs.svelte";
  import type { Posts } from "$lib/util/posts";
  import type { Song } from "$lib/util/song";

  let { data }: Posts<Song[]> = $props();
</script>

<Songs {data} />
