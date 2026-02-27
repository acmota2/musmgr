<script lang="ts">
  import type { ChangeEventHandler, HTMLImgAttributes } from "svelte/elements";
  import { IS_ADMIN } from "$lib/app";

  type EditablePictureProps = {
    name: string;
    pictureId: string;
  } & HTMLImgAttributes;

  const onFileChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    const input = event.target as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) {
      return;
    }

    previewUrl = URL.createObjectURL(file);
  };

  const onFileReset: ChangeEventHandler<HTMLInputElement> = () => {
    previewUrl = "/api/composer/picture";
  };

  let { name = "", pictureId, ...imgProps }: EditablePictureProps = $props();

  let previewUrl: string = $state("/api/composer/picture");
</script>

{#if !IS_ADMIN && pictureId !== null}
  <img class="picture" {...imgProps}>
{:else if IS_ADMIN}
  <input
    {name}
    type="file"
    id="picture"
    accept="image/*"
    hidden
    onchange={onFileChange}
    onreset={onFileReset}
  >
  <label for="picture" class="image-picker">
    {#if pictureId !== null || previewUrl}
      <figure class="picture-wrapper">
        <img class="picture-preview" src={previewUrl} alt="Couldn't load preview">
      </figure>
    {:else}
      <span id="chooser-text">Choose a picture</span>
    {/if}
  </label>
{/if}

<style>
  #chooser-text {
    display: flex;
    max-width: 100%;
    text-align: center;
    font: 24px bold;
  }

  .image-picker,
  .picture-wrapper,
  #chooser-text {
    height: 100%;
    width: 100%;
  }

  .image-picker,
  #chooser-text {
    align-items: center;
    justify-content: center;
  }

  .image-picker {
    position: relative;
    display: flex;
    overflow: hidden;
    max-height: 100%;
    cursor: pointer;
    background: linear-gradient(
      to left,
      rgba(225, 226, 238, 0) 0%,
      rgba(225, 226, 238, 0.3) 30%,
      rgba(225, 226, 238, 0.6) 70%,
      rgba(225, 226, 238, 1) 100%
    );
  }

  .image-picker:has(.picture-preview) {
    border: transparent;
  }

  .picture,
  .picture-preview {
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
    margin: 0;
    max-height: 100%;
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
    font-weight: bold;
    font-size: 24px;

    opacity: 0;
    transition: opacity 150ms ease;
    pointer-events: none;
  }

  .picture-wrapper:hover::after {
    opacity: 1;
  }
</style>
