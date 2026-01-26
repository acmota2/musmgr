<script lang="ts">
  const { class: className = "", children, disabled = false } = $props();
</script>

<button class={`button ${className}`} {disabled}>{@render children()}</button>

<style>
  .button {
    border-radius: var(--default-radius);
    padding: var(--default-padding);
  }
</style>
