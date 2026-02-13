<script lang="ts">
  let { class: className = "", value = $bindable(), ...rest } = $props();
</script>

<textarea bind:value class="textarea will-focus {className}" {...rest}>{value}</textarea>

<style>
  .textarea {
    font-size: 20px;
    min-height: 500px;
    border: solid 1px var(--border-color);
    border-radius: var(--default-radius);
    padding: var(--default-padding);
    resize: vertical;
    width: 100%;
  }
</style>
