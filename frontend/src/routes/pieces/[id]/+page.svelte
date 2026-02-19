<script lang="ts">
  import type { MusmgrEventTimetable } from "$lib/server/events.js";

  // import { IS_ADMIN } from "$lib/app";
  import type { Piece, PieceFile } from "$lib/server/pieces";
  import Events from "$lib/ui/pages/Events.svelte";

  let { data } = $props();
  let currentImageIndex = $state(0);
  let currentScoreIndex = $state(0);
  let currentImage: PieceFile | undefined = $derived(data.images[currentImageIndex]);
  let currentScore: PieceFile | undefined = $derived(data.scores[currentScoreIndex]);
  let events: MusmgrEventTimetable = $derived(data.events);
</script>

{#snippet filesContent(piece: Piece, audios: PieceFile[], events: MusmgrEventTimetable, currentImage: PieceFile, currentScore: PieceFile)}
  <div class="left scores-view">
    {#if currentScore}
      <embed
        src="/api/pieces/{piece.id}/files/{currentScore.id}"
        type="application/pdf"
        width="100%"
        height="100%"
      >
    {/if}
  </div>
  <div class="right">
    <div class="images-view">
      {#if currentImage}
        <img src="/api/pieces/{piece.id}/files/{currentImage.id}" alt="{currentImage.name}">
      {/if}
    </div>
    <div class="events-view">
      <Events descriptionCap={40} {events} />
    </div>
  </div>
  <div class="bottom audio-view">
    <audio>
      {#each audios as audio}
        <source src="/api/pieces/{piece.id}/files/{audio.id}" type="{audio.contentType}">
      {/each}
    </audio>
  </div>
{/snippet}

<div class="container wide">
  {@render filesContent(data.piece, data.audios, events, currentImage, currentScore)}
</div>

<style>
  .container.wide {
    height: 100%;
  }

  .left {
    width: 60%;
  }

  .right {
    width: 40%;
  }

  .left.scores-view {
    height: 100%;
  }
</style>
