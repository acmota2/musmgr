<script lang="ts">
  import type { MusmgrEvent, MusmgrEventTimetable } from "$lib/server/events.js";

  interface EventProps {
    descriptionCap: number;
    events: MusmgrEventTimetable;
  }

  let { descriptionCap = 40, events }: EventProps = $props();
</script>

{#snippet eventElement(event: MusmgrEvent)}
  <div class="event-item"><a href={`/events/${event.id}`}>
    <div class="event-item-headline">
      <h3>{event.name}</h3>
      <p>
        <time datetime={event.happenedAt}>
          {(new Date(event.happenedAt)).toLocaleDateString("en-UK", {
            year: "numeric", 
            month: "short", 
            day: "numeric" 
          })}
        </time>
      </p>
    </div>
    {#if event.description}
      <p>
        {event.description.slice(0, descriptionCap)}
        {#if event.description.length > descriptionCap}
          ...
        {/if}
      </p>
    {/if}
  </a></div>
{/snippet}

<ul class="timeline">
  {#each Object.entries(events) as [year, monthEvents]}
    <li>
      <h1 class="year">{year}</h1>
    </li>
    {#each Object.entries(monthEvents) as [month, events]}
      <li>
        <h2 class="month">{month}</h2>
      </li>
      <div class="month-events">
        {#each events as event}
          <li class="event-card">{@render eventElement(event)}</li>
        {/each}
      </div>
    {/each}
  {/each}
</ul>

<style>
  .year {
    font-size: 48px;
  }

  .month {
    font-size: 32px;
    padding-bottom: 10px;
  }

  .timeline li {
    list-style: none;
  }

  .timeline {
    position: relative;
    margin: 0;
    padding-left: 100px;
    min-height: 100%;
  }

  .timeline > * {
    display: block;
    width: 100%;
    position: relative;
  }

  .timeline::before {
    content: "";
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    min-height: 100%;
    height: 100%;
    max-height: none;

    width: 7px;
    background: black;

    transform: translateX(-50%);
  }

  .month-events {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .event-card {
    position: relative;
  }

  .event-card::after {
    content: "";
    position: absolute;

    /* same as padding */
    left: -100px;
    top: 50%;

    width: 20px;
    height: 20px;

    background: black;
    border-radius: 50%;

    transform: translate(-50%, -50%);
  }

  .event-item {
    border-radius: var(--default-radius);
    box-shadow: rgba(50, 50, 93, 0.1) 15px 1px 15px 0px;
    display: flex;
    flex-direction: column;
    gap: 4px;
    padding: 8px 25px;
    position: relative;
    transition:
      ease 200ms transform,
      ease 300ms box-shadow;
    width: 100%;
  }

  .event-item:hover {
    transform: scale(1.04);
    transition:
      ease 200ms transform,
      ease 300ms box-shadow;
    cursor: pointer;
    box-shadow:
      rgba(50, 50, 93, 0.25) 0px 13px 27px -5px,
      rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
  }

  .event-item > a {
    width: 100%;
    height: 100%;
  }

  .event-item-headline {
    display: flex;
    flex-direction: row;
    min-width: 0;
    width: 100%;
    justify-content: space-between;
    align-items: center;
    & h3 {
      font-size: 48px;
    }
  }

  @container (max-width: 767px) {
    .year {
      font-size: 28px;
    }

    .month {
      font-size: 24px;
    }

    .event-item-headline {
      flex-direction: column-reverse;
      justify-content: left;
      align-items: start;
    }

    .event-item-headline h3 {
      font-size: 32px;
    }

    .event-item-headline p,
    .event-item p {
      font-size: 16px;
    }

    .event-card {
      justify-content: left;
      align-items: left;
    }
  }

  @container (max-width: 999px) {
    .timeline {
      padding-left: 50px;
    }

    .timeline::before {
      left: 25px;
    }

    .event-card::after {
      left: -25px;
    }
  }
</style>
