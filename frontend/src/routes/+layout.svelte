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
  <div class="nav-in">
    <a href="/" class={navLink("/")}> MusMGR </a>
    <a href="/events" class={navLink("/events")}> Events </a>
    <a href="/pieces" class={navLink("/pieces")}> Pieces </a>
  </div>
</nav>

<main class="frame">{@render children()}</main>

<style>
  nav {
    flex: 0 0 auto;
    padding-top: 5px;
    position: sticky;
    top: 0;
    z-index: 1000;
    background: rgba(255, 255, 255, 0.6);
    backdrop-filter: blur(10px);
  }

  .frame {
    max-width: var(--frame-max);
    margin-inline: auto;
  }

  main.frame {
    min-height: 0;
    width: 100%;
    height: 100%;
    max-height: none;
  }

  .nav-link {
    padding: 2px 8px;
    position: relative;
    color: black;
    font-size: 24px;
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
    justify-content: center;
  }
</style>
