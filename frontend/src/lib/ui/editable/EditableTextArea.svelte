<script lang="ts">
  import type { HTMLTextareaAttributes } from "svelte/elements";
  import { IS_ADMIN } from "$lib/app";

  type EditableTextareaProps = {
    text: string;
  } & HTMLTextareaAttributes;

  let {
    text = $bindable(),
    class: className = "",
    id = "",
    ...props
  }: EditableTextareaProps = $props();

  const defaultText = `${text}`;
  const onReset = () => {
    text = defaultText;
  };
</script>

{#if IS_ADMIN}
  <textarea
    class="viewable-area-text editable-text {className}"
    bind:value={text}
    onreset={onReset}
    {id}
    {...props}
  ></textarea>
{:else}
  <pre class="viewable-area-text" {id}>{text}</pre>
{/if}

<style>
  .viewable-area-text {
    display: flex;
    height: 100%;
    width: 100%;
    overflow-y: auto;
    scrollbar-color: auto transparent;
    white-space: -o-pre-wrap;
    word-wrap: break-word;
  }

  .editable-text {
    background: var(--background-light);
    border: outset 1px transparent;
    border-radius: var(--default-radius);
    font-size: var(--paragraph-font);
    resize: none;
  }

  .editable-text:hover {
    border: outset 1px var(--border-color);
  }
</style>
