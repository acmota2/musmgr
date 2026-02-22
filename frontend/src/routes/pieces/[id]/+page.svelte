<script lang="ts">
  import Container from "$lib/ui/components/Container.svelte";
  import Events from "$lib/ui/pages/Events.svelte";
  import type { MusmgrEventTimetable } from "$lib/server/events.js";
  import type { Piece, PieceFile } from "$lib/server/pieces";
  import { IS_ADMIN } from "$lib/app/index.js";
  import ButtonAnchor from "$lib/ui/components/ButtonAnchor.svelte";

  let { data } = $props();
  let currentScoreIndex = $state(0);
  let currentAudioIndex = $state(0);
  let currentScore: PieceFile | undefined = $derived(data.scores[currentScoreIndex]);
  let events: MusmgrEventTimetable = $derived(data.events);
</script>

{#snippet filesContent(piece: Piece, audios: PieceFile[], images: PieceFile[], events: MusmgrEventTimetable, currentAudioIndex: number, currentScore: PieceFile)}
  <div class="top">
    <h1>{data.piece.title}</h1>
    {#if data.audios[currentAudioIndex]}
      <audio
        class="audio-player"
        src="/api/pieces/{piece.id}/files/{audios[currentAudioIndex].id}"
        controls
      >
        {#each audios as audio}
          <source src="/api/pieces/{piece.id}/files/{audio.id}" type="{audio.contentType}">
        {/each}
      </audio>
    {/if}
    {#if IS_ADMIN}
      <ButtonAnchor href="/pieces/{piece.id}/edit">Edit</ButtonAnchor>
    {/if}
  </div>
  <div class="bottom">
    <div class="left scores-view">
      {#if currentScore}
        <embed
          class="right-container"
          src="/api/pieces/{piece.id}/files/{currentScore.id}"
          type="application/pdf"
          width="100%"
          height="100%"
        >
      {/if}
    </div>
    <div class="right">
      <Container class="piece-description-container" title="Description">
        <p class="description">{data.piece.description}</p>
      </Container>
      <Container class="piece-images-view">
        {#if images.length > 0}
          {#each images as image}
            <img src="/api/pieces/{piece.id}/files/{image.id}" alt="{image.name}">
          {/each}
        {:else}
          <div class="no-images">
            <p>No images to show</p>
          </div>
        {/if}
      </Container>
      <Container class="piece-events-view" title="Events">
        <Events descriptionCap={40} {events} />
      </Container>
    </div>
  </div>
{/snippet}

<div class="container very-wide">
  <div class="piece-content">
    {@render filesContent(data.piece, data.audios, data.images, events, currentAudioIndex, currentScore)}
  </div>
</div>

<style>
  .right-container {
    border-radius: var(--default-radius);
    box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
    background-color: white;
  }

  .container.very-wide {
    justify-self: center;
    flex-direction: column;
    height: 100%;
    width: 100%;
    margin: 0 25px;
    display: flex;
  }

  .piece-content {
    flex-direction: column;
    height: 100%;
    display: flex;
  }

  .top {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    gap: 12px;
    & h1 {
      margin: 10px 0;
      flex-shrink: 0;
    }
  }

  .audio-player {
    width: 100%;
  }

  .audio-player::-webkit-media-controls-panel {
    background: var(--background-light);
  }

  .bottom {
    display: flex;
    flex-direction: row;
    flex: 1 1 0;
    justify-content: space-between;
    gap: 25px;
    margin-bottom: 25px;
  }

  .left {
    width: 65%;
  }

  .right {
    display: flex;
    flex-direction: column;
    width: 35%;
    gap: 25px;
  }

  .right > * {
    min-height: 0;
  }

  :global(.piece-images-view) {
    display: flex;
    flex-direction: column;
    flex: 2 1 0;
    min-height: 0;
  }

  :global(.piece-images-view img),
  :global(.piece-images-view .no-images) {
    flex: 1;
  }

  :global(.piece-images-view) .container {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .no-images {
    display: flex;
    flex-direction: row;
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
    & p {
      font-weight: bold;
    }
  }

  :global(.piece-events-view),
  :global(.piece-description-container) {
    display: flex;
    position: relative;
    flex: 1 1 0;
    min-height: 0;
    overflow: auto;
    & h3 {
      position: sticky;
      z-index: 2;
      background: inherit;
      top: 0;
      left: 0;
    }
  }

  .left.scores-view {
    position: relative;
    height: 100%;
  }

  .scores-view > embed {
    display: block;
    position: absolute;
  }

  @media (max-width: 1700px) {
    .piece-content {
      padding: 0 10px;
    }
  }

  @media (max-width: 1024px) {
    .bottom {
      display: flex;
      flex-direction: row;
      overflow-x: auto;
      overflow-y: hidden;
      scroll-snap-type: x mandatory;
      -webkit-overflow-scrolling: touch;
    }

    .left::after {
      content: "";
      position: absolute;
      top: 0;
      right: 0;
      width: 32px;
      height: 100%;
      pointer-events: none;

      /* fade from transparent to background */
      background: linear-gradient(to right, rgba(128, 128, 128, 0), rgba(128, 128, 128, 0.8));
    }

    .left,
    .right {
      flex: 0 0 100%;
      width: 100%;
      scroll-snap-align: start;
    }
  }
</style>
