<script lang="ts">
  let { children, class: className = "", title = "" } = $props();
</script>

<div class="general {className}">
  {#if title !== ""}
    <h3 class="container-title">{title}</h3>
  {/if}
  <div class="container">{@render children()}</div>
</div>

<style>
  .general {
    display: flex;
    flex-direction: column;
    background: white;
    border-radius: var(--default-radius);
    box-shadow: rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;
  }

  .container-title {
    padding: 8px;
    position: sticky;
    top: 0;
    left: 0;
  }

  .container {
    padding: 0 8px;
    width: 100%;
    height: 100%;
    overflow-x: visible;
    overflow-y: auto;
  }
</style>
