<script lang="ts">
  import { capitalize } from "$lib/utils";

  let { class: className = "", pieces } = $props();
</script>

<div class="pieces-container {className}">
  {#each pieces as piece}
    <a href="/pieces/{piece.id}" class="piece-card will-focus on-hover link-shadow">
      <h2 class="title">{piece.title}</h2>
      <div class="piece-details">
        <time class="year">{piece.composedAt}</time>
        <p class="instrumentation">{capitalize(piece.instrumentation)}</p>
      </div>
    </a>
  {/each}
</div>

<style>
  .piece-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    padding: var(--default-padding);
  }

  .title {
    font-size: 40px;
  }

  .year {
    font-size: 20px;
    font-weight: bold;
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

  .pieces-container {
    display: flex;
    flex-direction: column;
    border-radius: var(--default-radius);
    gap: 20px;
    height: 100%;
    width: 100%;
  }

  @media (max-width: 480px) {
    .title {
      font-size: var(--title-small);
    }
  }
</style>
