<script lang="ts">
  const { class: className = "", ...rest } = $props();
</script>

<input class="text-input will-focus {className}" {...rest}>

<style>
  .text-input {
    border: solid 1px var(--border-color);
    border-radius: var(--default-radius);
    padding: var(--default-padding);
    font-size: 20px;
    width: 100%;
  }
</style>
