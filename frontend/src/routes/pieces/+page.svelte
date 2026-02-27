<script lang="ts">
  import { IS_ADMIN } from "$lib/app";
  import BottomControlsContainer from "$lib/ui/components/BottomControlsContainer.svelte";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";
  import NoContent from "$lib/ui/components/NoContent.svelte";
  import Pieces from "$lib/ui/pages/Pieces.svelte";

  let { data } = $props();
</script>

<div class="view">
  <div class="container narrow">
    {#if data.pieces.length === 0}
      <NoContent />
    {:else}
      <div class="pieces-container">
        <Pieces pieces={data.pieces} />
      </div>
    {/if}
  </div>
  {#if IS_ADMIN}
    <BottomControlsContainer>
      <ButtonAnchor href="/pieces/new">Create piece</ButtonAnchor>
    </BottomControlsContainer>
  {/if}
</div>

<style>
  .view {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .container.narrow {
    height: 100%;
    width: 100%;
  }

  .pieces-container {
    padding: 25px 0;
    height: 100%;
    width: 100%;
  }
</style>
