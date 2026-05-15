<script lang="ts">
  import type { MusmgrEvent, MusmgrEventTimetable } from "$lib/server/events.js";

  interface EventProps {
    descriptionCap: number;
    events: MusmgrEventTimetable;
  }

  function reverseSort([s1]: [string, unknown], [s2]: [string, unknown]): number {
    return s2.localeCompare(s1);
  }

  let { descriptionCap = 40, events }: EventProps = $props();
</script>

{#snippet eventElement(event: MusmgrEvent)}
  <a href={`/events/${event.id}`}>
    <div class="event-item link-shadow">
      <article class="event-item-headline">
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
      </article>
      {#if event.description}
        <p>
          {event.description.slice(0, descriptionCap)}
          {#if event.description.length > descriptionCap}
            ...
          {/if}
        </p>
      {/if}
    </div>
  </a>
{/snippet}

<ul class="timeline">
  {#each Object.entries(events).sort(reverseSort) as [ year, monthEvents ]}
    <li>
      <h1 class="year">{year}</h1>
    </li>
    {#each Object.entries(monthEvents).sort(reverseSort) as [ month, events ]}
      <li>
        <h2 class="month">{month}</h2>
      </li>
      <div class="month-events">
        {#each events.sort((e1, e2) => e2.happenedAt.localeCompare(e1.happenedAt)) as event}
          <li class="event-card on-hover">{@render eventElement(event)}</li>
        {/each}
      </div>
    {/each}
  {/each}
</ul>

<style>
  .year {
    font-size: 48px;
    margin-bottom: 25px;
  }

  .month {
    font-size: 32px;
    padding-bottom: 10px;
  }

  .timeline li {
    list-style: none;
  }

  .timeline {
    container-type: inline-size;
    position: relative;
    margin: 0;
    min-height: 100%;
    padding: 0 0 0 50px;
    overflow-y: visible;
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
    margin-bottom: 25px;
  }

  .event-card {
    position: relative;
  }

  .event-card::after {
    content: "";
    position: absolute;

    /* same as padding */
    left: -50px;
    top: 50%;

    width: 20px;
    height: 20px;

    background: black;
    border-radius: 50%;

    transform: translate(-50%, -50%);
  }

  .event-item {
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

  .event-item-headline {
    display: flex;
    flex-direction: row;
    min-width: 0;
    width: 100%;
    justify-content: space-between;
    align-items: center;
    & h3 {
      font-size: var(--title);
    }
  }

  @media (max-width: 767px) {
    .year {
      font-size: 28px;
    }

    .month {
      font-size: 24px;
    }

    .event-item-headline {
      flex-direction: column-reverse;
      justify-content: start;
      align-items: start;
    }

    .event-item-headline h3 {
      font-size: var(--title-small);
    }

    .event-item-headline p,
    .event-item p {
      font-size: var(--date);
    }

    .event-card {
      justify-content: left;
      align-items: left;
    }
  }

  @container (max-width: 767px) {
    .timeline::before {
      left: 25px;
    }

    .event-card::after {
      left: -25px;
    }

    .year {
      font-size: 32px;
    }

    .month {
      font-size: 28px;
    }

    .event-item-headline {
      flex-direction: column-reverse;
      justify-content: start;
      align-items: start;
    }

    .event-item-headline h3 {
      font-size: 32px;
    }

    .event-item-headline p,
    .event-item p {
      font-size: var(--date-small);
    }

    .event-item {
      border-radius: 0;
    }
    .event-card {
      justify-content: left;
      align-items: left;
      border-radius: 0;
    }
  }

  @media (max-width: 999px) {
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
