<script lang="ts">
  import type { Snippet } from "svelte";

  interface Breadcrumb {
    class: string;
    children: Snippet;
  }

  let { class: className, children } = $props() satisfies Breadcrumb;
</script>

<div class={className}>
  {@render children()}
</div>

