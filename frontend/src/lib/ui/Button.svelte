<script lang="ts">
  const { class: className = "", children, ...rest } = $props();
</script>

<button class="button ${className}" {...rest}>{@render children()}</button>

<style>
  .button {
    background: var(--button-plain);
    font-size: 20px;
    border: solid 1px rgba(238, 238, 238, 0.3);
    border-radius: var(--default-radius);
    padding: var(--default-padding);
  }

  .button:hover,
  .button:focus {
    background: #f8f8f8;
  }

  .button:hover {
    transition: background 200ms ease;
    cursor: pointer;
  }

  .button:focus {
    border-color: rgba(123, 123, 123, 0.8);
    transition: border-color 150ms ease;
  }
</style>
