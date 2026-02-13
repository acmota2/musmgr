<script lang="ts">
  import { IS_ADMIN } from "$lib/app";

  let { data } = $props();
</script>

<div class="container wide">
  <div class="files-view">
    {#if IS_ADMIN}
    {:else}
    {/if}
  </div>
  <div class="data-view"></div>
</div>
