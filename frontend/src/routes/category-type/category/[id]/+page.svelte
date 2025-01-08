<script lang="ts">
  import Songs from "$lib/pages/songs.svelte";

  let { data } = $props();
</script>

<Songs {data} />
