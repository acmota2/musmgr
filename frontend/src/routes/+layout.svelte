<script lang="ts">
  import { page } from "$app/state";
  import "$lib/styles/app.css";
  import GitHub from "$lib/assets/GitHub.svelte";
  import Mail from "$lib/assets/Mail.svelte";

  const { children } = $props();
  const pathname = $derived(page.url.pathname);

  function navLink(href: string) {
    return `nav-link ${(href === "/") !== pathname.startsWith(href) ? "active" : ""}`;
  }
</script>

<div class="page">
  <nav>
    <div class="nav-in">
      <a href="/" class={navLink("/")}> MusMGR </a>
      <a href="/events" class={navLink("/events")}> Events </a>
      <a href="/pieces" class={navLink("/pieces")}> Pieces </a>
    </div>
  </nav>
  <main class="frame">{@render children()}</main>
</div>

<footer>
  <p>Contacts and references:</p>
  <div class="link-section">
    <a href="mailto:acmota2@gmail.com" aria-label="Email">
      <Mail width={60} height={60} />
    </a>
    <a href="https://github.com/acmota2" aria-label="GitHub">
      <GitHub width={51} height={51} />
    </a>
  </div>
  <p class="copyright">André Mota © 2026</p>
</footer>

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

  footer {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 10px;
    background: white;
  }

  .link-section {
    display: flex;
    flex-direction: row;
    gap: 20px;
    align-items: center;
    justify-content: center;
  }

  .copyright {
    font-size: 18px;
  }

  .frame {
    max-width: var(--frame-max);
    margin-inline: auto;
  }

  .page {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    flex: 1 0 auto;
  }

  main.frame {
    min-height: 0;
    width: 100%;
    flex: 1 0 auto;
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

  .nav-link:hover::after,
  .nav-link.active::after {
    background: var(--accent);
  }

  .nav-in {
    display: flex;
    gap: 16px;
    align-items: center;
    justify-content: center;
  }
</style>
