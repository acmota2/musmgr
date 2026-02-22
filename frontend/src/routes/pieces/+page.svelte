<script lang="ts">
  import { IS_ADMIN } from "$lib/app";
  import BottomControlsContainer from "$lib/ui/components/BottomControlsContainer.svelte";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";
  import NoContent from "$lib/ui/components/NoContent.svelte";
  import Pieces from "$lib/ui/pages/Pieces.svelte";

  let { data } = $props();
</script>

{#if data.pieces.length === 0}
  <NoContent />
{:else}
  <div class="container narrow">
    <div class="pieces-container">
      <Pieces pieces={data.pieces} />
    </div>
    {#if IS_ADMIN}
      <BottomControlsContainer>
        <ButtonAnchor href="/pieces/new">Create piece</ButtonAnchor>
      </BottomControlsContainer>
    {/if}
  </div>
{/if}

<style>
  .pieces-container {
    padding: 25px 0;
  }
</style>
