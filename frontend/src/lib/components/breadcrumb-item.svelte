<script lang="ts">
  import type { Snippet } from "svelte";

  let { href, children } = $props() satisfies { href: string, children: Snippet};
</script>

<a {href}>
  {@render children()}
</a>
