<script lang="ts">
  import type { Piece } from "$lib/server/pieces";
  import Container from "$lib/ui/components/Container.svelte";
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";

  let confirmDialog: HTMLDialogElement | null = $state(null);

  function cancel() {
    history.back();
  }

  function confirm() {
    confirmDialog?.showModal();
  }

  let { data } = $props();
</script>

{#snippet selectPieceCard(piece: Piece)}
  <label class="select-piece-card link-shadow">
    <h2>{piece.title}</h2>
    <input name="piece_ids" type="checkbox" value="{piece.id}">
  </label>
{/snippet}

<div class="container narrow">
  <h1>Pieces in {data.event.name}</h1>
  <FormField
    method="POST"
    action="?/createEventPieces"
    class="select-piece-form"
    id="select-event-pieces"
  >
    <Container class="select-piece-container" title="Select pieces to add">
      <div class="inner-pieces">
        {#each data.pieces as piece}{@render selectPieceCard(piece)}{/each}
      </div>
    </Container>
    <FormButtonSet {cancel} {confirm} />
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="select-event-pieces" />

<style>
  .container.narrow {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  :global(.select-piece-container) {
    display: flex;
    flex-direction: column;
    padding: var(--default-padding);
    gap: 12px;
  }

  .inner-pieces {
    display: flex;
    overflow: visible;
    flex-direction: column;
    gap: 12px;
  }

  .select-piece-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    padding: var(--default-padding);
  }
</style>
