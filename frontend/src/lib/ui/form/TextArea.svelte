<script lang="ts">
  const { class: className = "", ...rest } = $props();
</script>

<textarea class={`textarea ${className}`} {...rest}></textarea>

<style>
  .textarea {
    min-height: 500px;
    border: solid 1px var(--border-color);
    border-radius: var(--default-radius);
    padding: var(--default-padding);
    resize: vertical;
    width: 100%;
  }

  .textarea:focus {
    outline: none;
    border-color: var(--accent);
    box-shadow: 0 0 5px var(--accent);
  }
</style>
