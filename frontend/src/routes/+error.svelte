<script lang="ts">
  import { page } from "$app/state";
</script>

<div class="error">
  <h1>{page.error?.status ?? 500}</h1>
  <h2>{page.error?.message ?? "It's not you, it's us"}</h2>
  <p>Error code: {page.error?.code ?? "Unknown error"}</p>
</div>

<style>
  .error {
    max-height: 100%;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
</style>
