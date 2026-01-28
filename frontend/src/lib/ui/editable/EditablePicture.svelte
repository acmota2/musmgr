<script lang="ts">
  import type { ChangeEventHandler } from "svelte/elements";
  import { IS_ADMIN } from "$lib/app";

  let previewUrl: string | null = $state(null);

  const onFileChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    const input = event.target as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) {
      return;
    }

    previewUrl = URL.createObjectURL(file);
  };

  let { pictureId, alt = "", wrapperClass = "", ...imgProps } = $props();
</script>

{#if pictureId !== null}
  <figure>
    <img class="" {...imgProps}>
  </figure>
{:else if IS_ADMIN && pictureId === null}
  <article class={wrapperClass}>
    <input type="file" id="picture" accept="image/*" hidden onchange={onFileChange}>
    <label for="picture" class="image-picker">
      {#if previewUrl}
        <figure class="picture-wrapper">
          <img class="picture-preview" src={previewUrl} alt="Couldn't load preview">
        </figure>
      {:else}
        <span id="chooser-text">Choose image</span>
      {/if}
    </label>
  </article>
{/if}

<style>
  #chooser-text {
    border-radius: var(--default-radius);
    max-width: 100%;
  }

  .image-picker {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 700px;
    max-height: 100%;
    border: solid 2px rgba(100, 100, 100, 0.6);
    border-radius: var(--default-radius);
    cursor: pointer;
  }

  .image-picker:has(.picture-preview) {
    border: transparent;
  }

  .picture-preview {
    border-radius: var(--default-radius);
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center;
    display: block;
  }

  .picture-wrapper {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .picture-wrapper::after {
    content: "Replace";
    color: #fafafa;
    position: absolute;
    inset: 0;

    display: flex;
    align-items: center;
    justify-content: center;

    background-color: rgba(150, 150, 150, 0.35);
    border-radius: var(--default-radius);
    font-weight: bold;

    opacity: 0;
    transition: opacity 150ms ease;
    pointer-events: none;
  }

  .picture-wrapper:hover::after {
    opacity: 1;
  }
</style>
