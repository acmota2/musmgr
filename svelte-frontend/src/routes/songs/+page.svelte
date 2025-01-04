<script lang="ts">
  import type { Posts } from "$lib/util/posts";
  import type { Song } from "$lib/util/song";

  let { data }: { data: Posts<Song[]> } = $props();
</script>

{#each data.posts as song}
  <!-- change with params for better site navigation -->
  <a href="/songs/{song.id}" class="hover:bg-gray">
    <div
      class="trasition b-2 flex flex-col border-2 border-black pb-4 pt-2 duration-200 hover:bg-slate-400 sm:rounded-sm"
    >
      <p>{song.title}</p>
      <div class="flex flex-row">
        <p>Key:</p>
        <p>{song.key}</p>
      </div>
    </div>
  </a>
{/each}
