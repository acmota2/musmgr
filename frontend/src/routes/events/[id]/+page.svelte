<script lang="ts">
  import { IS_ADMIN } from "$lib/app";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";
  import Container from "$lib/ui/components/Container.svelte";
  import Pieces from "$lib/ui/pages/Pieces.svelte";

  let { data } = $props();
  let pieces = $derived(data.pieces);
  let currentEvent = $derived(data.event);

  function transformDate(date: string) {
    return new Date(date).toLocaleDateString("en-UK", {
      year: "numeric",
      month: "long",
      day: "2-digit",
    });
  }
</script>

<div class="container narrow">
  <div class="title">
    <h1>{currentEvent.name}</h1>
    <h3>
      <time>{transformDate(currentEvent.happenedAt)}</time>
    </h3>
    {#if IS_ADMIN}
      <ButtonAnchor href={`/events/${currentEvent.id}/edit`}>Edit</ButtonAnchor>
    {/if}
  </div>
  <div class="event-details">
    {#if currentEvent.description}
      <Container title="Description" class="event-description-container">
        <p class="description">{currentEvent.description}</p>
      </Container>
    {/if}
    <Container title="Pieces">
      <Pieces {pieces} class="event-pieces-container" />
    </Container>
  </div>
</div>

<style>
  .container.narrow {
    height: 100%;
    width: 100%;
  }

  .title {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    padding: 25px 0;
    & h1 {
      font-size: var(--title);
    }
    & time {
      font-size: var(--date);
    }
  }

  .event-details {
    display: flex;
    flex-direction: column;
    gap: 25px;
  }

  :global(.event-pieces-container) {
    padding: 25px;
  }

  .description {
    padding-bottom: 12px;
  }

  @media (max-width: 450px) {
    .title {
      flex-direction: column;
      justify-content: start;
      align-items: start;
      & h1 {
        font-size: var(--title-small);
      }
      & time {
        font-size: var(--date-small);
      }
    }
  }
</style>
