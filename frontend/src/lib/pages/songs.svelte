<svelte:options customElement="song-list" />

<script lang="ts">
  import type { Posts, Song } from "$lib/util";

  let { data }: Posts<Song[]> = $props();
</script>

{#each data as song}
  <a href="/songs/{song.id}" class="hover:bg-gray">
    <div
      class="trasition b-2 flex flex-col border-2 border-black pb-4 pl-4 pt-4 duration-200 hover:bg-slate-400 sm:rounded-sm"
    >
      <p class="text-lg font-bold">{song.title}</p>
      <div class="flex flex-row">
        <p>Key: {song.key}</p>
      </div>
    </div>
  </a>
{/each}
