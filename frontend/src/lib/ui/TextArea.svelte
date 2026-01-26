<script lang="ts">
  const { class: className = "", id = "", required = false } = $props();
</script>

<textarea class={`textarea ${className}`} {id} {required}></textarea>

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
