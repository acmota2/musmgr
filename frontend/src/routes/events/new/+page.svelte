<script lang="ts">
  import ConfirmDialog from "$lib/ui/form/ConfirmDialog.svelte";
  import FormButtonSet from "$lib/ui/form/FormButtonSet.svelte";
  import FormField from "$lib/ui/form/FormField.svelte";
  import Select from "$lib/ui/form/Select.svelte";
  import TextArea from "$lib/ui/form/TextArea.svelte";
  import TextInput from "$lib/ui/form/TextInput.svelte";
  import { capitalize } from "$lib/utils/index.js";

  function confirm() {
    confirmDialog?.showModal();
  }

  function cancel() {
    history.back();
  }

  let confirmDialog: HTMLDialogElement | null = $state(null);
  let { data } = $props();
</script>

<div class="container narrow">
  <FormField method="POST" action="?/create" id="create-event">
    <label for="name">
      <h2>Name</h2>
    </label>
    <TextInput name="name" required />
    <label for="happened_at">
      <h2>Happened at</h2>
    </label>
    <!-- consider selects -->
    <TextInput name="happened_at" required placeholder="YYYY-MM-DD" />
    <label for="event_type">
      <h2>Type</h2>
    </label>
    <Select name="event_type">
      {#each data.eventTypes as eventType}
        <option value="{eventType}">{capitalize(eventType)}</option>
      {/each}
    </Select>
    <label for="description">
      <h2>Description</h2>
    </label>
    <TextArea name="description" required />

    <FormButtonSet {cancel} {confirm} />
  </FormField>
</div>

<ConfirmDialog bind:confirmDialog formId="create-event" />

<style>
  h2 {
    font-size: 28px;
  }

  label > * {
    margin: 6px 0;
  }
</style>
