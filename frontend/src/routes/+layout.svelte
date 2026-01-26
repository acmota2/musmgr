<script lang="ts">
  import { page } from "$app/state";
  import "$lib/styles/app.css";

  const { children } = $props();
  const pathname = $derived(page.url.pathname);

  function navLink(href: string) {
    return `nav-link ${pathname.startsWith(href) ? "active" : ""}`;
  }
</script>

<nav>
  <div class="frame nav-in">
    <a href="/" class={navLink("/")}> MusMGR </a>
    <a href="/events" class={navLink("/events")}> Events </a>
    <a href="/pieces" class={navLink("/pieces")}> Pieces </a>
  </div>
</nav>

<div class="frame">{@render children()}</div>

<style>
  .frame {
    max-width: var(--frame-max);
    margin-inline: auto;
    padding-inline: var(--frame-padding);
  }

  nav {
    padding: var(--default-padding);
  }

  .nav-link {
    border: var(--default-border) solid transparent;
    position: relative;
    color: black;
    font-weight: bold;
    font-size: 24px;
    padding: var(--default-padding);
    text-decoration: none;
  }

  .nav-link::after {
    content: "";
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    height: 2px;
    background: transparent;
    transition: background 200ms ease;
  }

  .nav-link:hover::after {
    background: var(--accent);
  }

  .nav-in {
    display: flex;
    gap: 16px;
    align-items: center;
    justify-content: flex-start;
  }
</style>
