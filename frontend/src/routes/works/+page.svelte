<script lang="ts">
  import Works from "$lib/pages/works.svelte";
  import type { Posts } from "$lib/util";
  import type { Song } from "$lib/util";

  let { data }: Posts<Song[]> = $props();
</script>

<Works {data} />
