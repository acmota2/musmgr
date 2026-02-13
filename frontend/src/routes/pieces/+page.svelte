<script lang="ts">
  import { IS_ADMIN } from "$lib/app";
  import BottomControlsContainer from "$lib/ui/components/BottomControlsContainer.svelte";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";
  import NoContent from "$lib/ui/components/NoContent.svelte";
  import { capitalize } from "$lib/utils";

  let { data } = $props();
</script>

{#if data.pieces.length === 0}
  <NoContent />
{:else}
  <div class="pieces container narrow">
    {#each data.pieces as piece}<a href="/piece/{piece.id}" class="piece-card will-focus">
      <h2 class="title">{piece.title}</h2>
      <div class="piece-details">
        <time class="year">{piece.composedAt}</time>
        <p class="instrumentation">{capitalize(piece.instrumentation)}</p>
      </div>
    </a>{/each}
    {#if IS_ADMIN}
      <BottomControlsContainer>
        <ButtonAnchor href="/pieces/new">Create piece</ButtonAnchor>
      </BottomControlsContainer>
    {/if}
  </div>
{/if}

<style>
  .pieces {
    display: flex;
    flex-direction: column;
    gap: 20px;
    height: 100%;
    padding: 35px 0;
  }

  .piece-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    border: solid #4a4a4a 2px;
    border-radius: var(--default-radius);
    padding: var(--default-padding);
  }

  .title {
    font-size: 40px;
  }

  .year {
    font-size: 20px;
    font-weight: 400;
  }

  .instrumentation {
    font-size: 20px;
  }

  .piece-details {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-end;
    gap: 10px;
  }

  @media (max-width: 900px) {
    .pieces {
      padding: 35px 10px;
    }
  }

  @media (max-width: 480px) {
    .pieces {
      padding: 20px 10px;
    }

    .title {
      font-size: 34px;
    }
  }
</style>
